package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

var ctx = context.Background()
var conn *redis.Client

func Setup() {
	conn = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func Set(key string, value string) {
	if err := conn.Set(ctx, key, value, 0).Err(); err != nil {
		panic(err)
	}
}

func SetEx(key string, value string, expire int) {
	if err := conn.Set(ctx, key, value, time.Duration(expire)).Err(); err != nil {
		panic(err)
	}
}

func Get(key string) string {
	value, err := conn.Get(ctx, key).Result()
	if err != nil {
		panic(err)
	}
	return value
}