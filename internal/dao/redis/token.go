package redis

import (
	"context"
	"github.com/0RAJA/Road/internal/global"
)

func SetRefreshToken(ctx context.Context, refreshToken, username string) error {
	return rdb.Set(ctx, getRedisKey(keyRefreshToken)+refreshToken, username, global.AllSetting.Token.RefreshTokenDuration).Err()
}

func GetRefreshToken(ctx context.Context, refreshToken string) (string, error) {
	return rdb.Get(ctx, getRedisKey(keyRefreshToken)+refreshToken).Result()
}
