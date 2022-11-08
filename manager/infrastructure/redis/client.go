package redis

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
	"github.com/italorfeitosa/jobs-manager-mvp/common/config"
)

var instance *redis.Client

func Get() *redis.Client {
	if instance != nil {
		return instance
	}

	opt, err := redis.ParseURL(config.Get().Redis.URL)
	if err != nil {
		panic(err)
	}

	instance = redis.NewClient(opt)
	cmd := instance.Ping(context.Background())
	if cmd.Err() != nil {
		panic(cmd.Err())
	}
	result, err := cmd.Result()
	if err != nil {
		panic(err)
	}
	log.Printf("result: %#+v\n", result)
	return instance
}
