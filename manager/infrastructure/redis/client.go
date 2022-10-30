package redis

import (
	"github.com/go-redis/redis/v8"
	"github.com/italorfeitosa/jobs-manager-mvp/common/config"
)

type Client *redis.Client

var client Client

func Get() Client {
	if client != nil {
		return client
	}

	opt, err := redis.ParseURL(config.Get().Redis.URL)
	if err != nil {
		panic(err)
	}

	client = redis.NewClient(opt)

	return client
}
