package redis

import (
	"github.com/go-redis/redis/v8"
	"strconv"
)

var (
	Query *Queries
)

type Queries struct {
	rdb *redis.Client
}

func QueryInit(Addr, Password string, PoolSize, DB int) {
	Query = &Queries{redisInit(Addr, Password, PoolSize, DB)}
}

func int64toA(n int64) string {
	return strconv.FormatInt(n, 10)
}

func atoInt64Must(a string) int64 {
	i, _ := strconv.ParseInt(a, 10, 64)
	return i
}
