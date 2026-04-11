package database

import (
	"context"
	"debian-ecommerce/pkg/utils"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func InitRedis(config utils.RedisConfig) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.Host, config.Port),
		Password: config.Password,
		DB:       config.DB,
	})

	if err := rdb.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}

	return rdb, nil
}
