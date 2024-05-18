package utils

import (
	"github.com/redis/go-redis/v9"
)

func getFromRedis(link string) (string, error) {
	val, err := Client.Get(ctx, link).Result()

	if err == redis.Nil {
		return "", &notFoundError{"Not Found"}
	}
	if err == nil {
		return val, nil
	}
	return "", &internalError{"Internal Error"}
}

func setIntoRedis(link, surl string) error {
	err := Client.Set(ctx, surl, link, redis.KeepTTL).Err()
	if err == nil {
		return nil
	}
	return &internalError{"Internal Error"}
}
