package redis

import (
	"context"
	"github.com/0RAJA/Road/internal/global"
	"github.com/go-redis/redis/v8"
	"strconv"
)

var rdb *redis.Client

func Init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     global.AllSetting.Redis.Address,  //ip:端口
		Password: global.AllSetting.Redis.Password, //密码
		PoolSize: global.AllSetting.Redis.PoolSize, //连接池
		DB:       global.AllSetting.Redis.DB,       //默认连接数据库
	})
	_, err := rdb.Ping(context.Background()).Result() //测试连接
	if err != nil {
		panic(err)
	}
}

func int64toA(n int64) string {
	return strconv.FormatInt(n, 10)
}

func atoInt64Must(a string) int64 {
	i, _ := strconv.ParseInt(a, 10, 64)
	return i
}
