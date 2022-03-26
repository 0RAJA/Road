package redis

import (
	"context"
	"errors"
	"github.com/0RAJA/Road/internal/global"
	"github.com/go-redis/redis/v8"
	"strconv"
)

var (
	Query *Queries
)

type Queries struct {
	rdb *redis.Client
}

func Init() {
	Query = &Queries{
		rdb: redis.NewClient(&redis.Options{
			Addr:     global.AllSetting.Redis.Address,  //ip:端口
			Password: global.AllSetting.Redis.Password, //密码
			PoolSize: global.AllSetting.Redis.PoolSize, //连接池
			DB:       global.AllSetting.Redis.DB,       //默认连接数据库
		}),
	}
	_, err := Query.rdb.Ping(context.Background()).Result() //测试连接
	if err != nil {
		panic(err)
	}
}

func IsNil(err error) bool {
	return errors.Is(err, redis.Nil)
}

func int64toA(n int64) string {
	return strconv.FormatInt(n, 10)
}

func atoInt64Must(a string) int64 {
	i, _ := strconv.ParseInt(a, 10, 64)
	return i
}
