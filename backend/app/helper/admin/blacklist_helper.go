package authHelper

import (
	"context"
	"franchise-web/config/initializers"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func InitRedis(redisAddr string) {
	initializers.Rdb = redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})

	_, err := initializers.Rdb.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
}

func BlacklistToken(tokenID string) error {
	return initializers.Rdb.Set(ctx, tokenID, true, 24*time.Hour).Err()
}

func IsTokenBlacklisted(tokenID string) bool {
	log.Println("tokenID: ", tokenID)
	val, err := initializers.Rdb.Get(ctx, tokenID).Result()
	log.Println("val: ", val)
	log.Println(val == "1")
	return err == nil && val == "1"
}
