package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
)

var client = redis.NewClient(&redis.Options{
	Addr:     "127.0.0.1:6379",
	Password: "",
	DB:       1,
})

var ctx = context.Background()

func Set(key string, value []byte) error {
	err := client.Set(ctx, key, value, 0).Err()
	return err
}

func Get(key string) []byte {
	value, err := client.Get(ctx, key).Bytes()
	if err != nil {

	}
	return value
}
