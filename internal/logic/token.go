package logic

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/0RAJA/Road/internal/dao/mysql"
	db "github.com/0RAJA/Road/internal/dao/mysql/sqlc"
	"github.com/0RAJA/Road/internal/global"
	"github.com/0RAJA/Road/internal/pkg/app/errcode"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetToken(ctx *gin.Context) {
	s := "https://github.com/login/oauth/authorize?client_id=%s&redirect_uri=%s"
	ctx.Redirect(http.StatusFound, fmt.Sprintf(s, global.AllSetting.Github.ClientID, global.AllSetting.Github.RedirectUri))
}

type githubToken struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"` // 这个字段没用到
	Scope       string `json:"scope"`      // 这个字段也没用到
}

// 向github获取 token
func getToken(url string) (*githubToken, error) {
	// 形成请求
	var req *http.Request
	var err error
	if req, err = http.NewRequest(http.MethodGet, url, nil); err != nil {
		return nil, err
	}
	req.Header.Set("accept", "application/json")

	// 发送请求并获得响应
	var httpClient = http.Client{}
	var res *http.Response
	if res, err = httpClient.Do(req); err != nil {
		return nil, err
	}

	// 将响应体解析为 token，并返回
	var token githubToken
	if err = json.NewDecoder(res.Body).Decode(&token); err != nil {
		return nil, err
	}
	return &token, nil
}

// getTokenAuthUrl 通过code获取token认证url
func getTokenAuthUrl(code string) string {
	return fmt.Sprintf(
		"https://github.com/login/oauth/access_token?client_id=%s&client_secret=%s&code=%s",
		global.AllSetting.Github.ClientID, global.AllSetting.Github.ClientSecret, code,
	)
}

// getUserInfo 获取用户信息
func getUserInfo(token *githubToken) (map[string]interface{}, error) {
	// 形成请求
	var userInfoUrl = "https://api.github.com/user" // github用户信息获取接口
	var req *http.Request
	var err error
	if req, err = http.NewRequest(http.MethodGet, userInfoUrl, nil); err != nil {
		return nil, err
	}
	req.Header.Set("accept", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("token %s", token.AccessToken))

	// 发送请求并获取响应
	var client = http.Client{}
	var res *http.Response
	if res, err = client.Do(req); err != nil {
		return nil, err
	}

	// 将响应的数据写入 userInfo 中，并返回
	var userInfo = make(map[string]interface{})
	if err = json.NewDecoder(res.Body).Decode(&userInfo); err != nil {
		return nil, err
	}
	return userInfo, nil
}

func TokenRedirect(ctx *gin.Context) (GetTokenReply, *errcode.Error) {
	var (
		err          error
		code         = ctx.Query("code") // 获取code
		tokenAuthUrl = getTokenAuthUrl(code)
		token        *githubToken
	)
	//获取token
	if token, err = getToken(tokenAuthUrl); err != nil {
		global.Logger.Info(err.Error())
		return GetTokenReply{}, errcode.ErrAuthorizationFailed
	}
	//获取用户信息
	var userInfo map[string]interface{}
	if userInfo, err = getUserInfo(token); err != nil {
		global.Logger.Info("获取用户信息失败，错误信息为:" + err.Error())
		return GetTokenReply{}, errcode.ErrAuthorizationFailed
	}
	user := User{
		Username:      userInfo["login"].(string),
		AvatarUrl:     userInfo["avatar_url"].(string),
		DepositoryUrl: userInfo["repos_url"].(string),
	}
	if ipaddr, ok := ctx.RemoteIP(); ok {
		user.Address = ipaddr.String()
	} else {
		user.Address = ctx.Request.RemoteAddr
	}
	_, err = mysql.Query.GetUserByUsername(ctx, user.Username)
	if err != nil {
		if mysql.IsNil(err) {
			err := mysql.Query.CreateUser(ctx, db.CreateUserParams{
				Username:      user.Username,
				AvatarUrl:     user.AvatarUrl,
				DepositoryUrl: user.DepositoryUrl,
				Address:       user.Address,
			})
			if err != nil {
				global.Logger.Error(err.Error())
				return GetTokenReply{}, errcode.ServerErr
			}
		}
	}
	err = mysql.Query.UpdateUser(ctx, db.UpdateUserParams{
		AvatarUrl:     user.AvatarUrl,
		DepositoryUrl: user.DepositoryUrl,
		Address:       user.Address,
		Username:      user.Username,
	})
	if err != nil {
		global.Logger.Error(err.Error())
		return GetTokenReply{}, errcode.ServerErr
	}
	token1, refreshToken1, err := generateToken(user.Username)
	if err != nil {
		return GetTokenReply{}, errcode.UnauthorizedTokenGenerateErr
	}
	return GetTokenReply{
		User: UserInfo{
			Username:      user.Username,
			AvatarUrl:     user.AvatarUrl,
			DepositoryUrl: user.DepositoryUrl,
		},
		Token:   token1,
		ReToken: refreshToken1,
	}, nil
}

func generateToken(username string) (Token, ReToken, *errcode.Error) {
	token, payload1, err := global.Maker.CreateToken(username, global.AllSetting.Token.AssessTokenDuration)
	if err != nil {
		global.Logger.Error(err.Error())
		return Token{}, ReToken{}, errcode.UnauthorizedTokenGenerateErr
	}
	refreshToken, payload2, err := global.Maker.CreateToken(username, global.AllSetting.Token.RefreshTokenDuration)
	if err != nil {
		global.Logger.Error(err.Error())
		return Token{}, ReToken{}, errcode.UnauthorizedTokenGenerateErr
	}
	return Token{
			Token:     token,
			ExpiredAt: payload1.ExpiredAt,
		}, ReToken{
			RefreshToken: refreshToken,
			ExpiredAt:    payload2.ExpiredAt,
		}, nil
}

func RefreshToken(ctx *gin.Context, params RefreshTokenReplyParams) (RefreshTokenReply, *errcode.Error) {
	payload, err := global.Maker.VerifyToken(params.ReToken)
	if err != nil {
		return RefreshTokenReply{}, errcode.UnauthorizedTokenErr
	}
	_, err = global.Maker.VerifyToken(params.Token)
	if err == nil {
		return RefreshTokenReply{}, errcode.ErrTokenNotExpired
	}
	if !errors.Is(err, errcode.UnauthorizedTokenTimeoutErr) {
		return RefreshTokenReply{}, errcode.UnauthorizedTokenErr
	}
	token, payload, err := global.Maker.CreateToken(payload.UserName, global.AllSetting.Token.AssessTokenDuration)
	if err != nil {
		global.Logger.Error(err.Error())
		return RefreshTokenReply{}, errcode.UnauthorizedTokenGenerateErr
	}
	return RefreshTokenReply{
		Token: Token{
			Token:     token,
			ExpiredAt: payload.ExpiredAt,
		},
	}, nil
}
