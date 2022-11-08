package dispatcher

import (
	"github.com/go-redis/redis/v8"
	"github.com/italorfeitosa/jobs-manager-mvp/common/config"
	"github.com/italorfeitosa/jobs-manager-mvp/manager/core"
)

type Factory interface {
	FromJobSchema(schema config.JobSchema) core.JobDispatcher
}

type factory struct {
	redis *redis.Client
}

func NewFactory(rc *redis.Client) Factory {
	return &factory{rc}
}
func (f *factory) FromJobSchema(schema config.JobSchema) core.JobDispatcher {
	if schema.HasHttp() {
		return NewHttpJobDispatcher(schema.Http.URL)
	}

	if schema.HasRedis() {
		return NewRedisJobDispatcher(f.redis, schema.Redis.Channel)
	}

	return NewNoopJobDispatcher()
}
