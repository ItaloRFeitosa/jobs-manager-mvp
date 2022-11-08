package di

import (
	"github.com/go-redis/redis/v8"
	"github.com/italorfeitosa/jobs-manager-mvp/common/config"
	"github.com/italorfeitosa/jobs-manager-mvp/manager/core"
	"github.com/italorfeitosa/jobs-manager-mvp/manager/infrastructure/dispatcher"
	redisClient "github.com/italorfeitosa/jobs-manager-mvp/manager/infrastructure/redis"
	"github.com/italorfeitosa/jobs-manager-mvp/manager/infrastructure/store"
)

type Container struct {
	Config     *config.Config
	JobsStore  core.JobsStore
	Redis      *redis.Client
	Dispatcher dispatcher.Factory
}

func New() *Container {
	return new(
		registerConfig,
		registerJobsStore,
		registerRedis,
		registerDispatchFactory,
	)
}
func new(registerFns ...func(c *Container)) *Container {
	c := &Container{}
	for _, fn := range registerFns {
		fn(c)
	}

	return c
}

func registerDispatchFactory(c *Container) {
	c.Dispatcher = dispatcher.NewFactory(c.Redis)
}

func registerConfig(c *Container) {
	c.Config = config.Get()
}

func registerRedis(c *Container) {
	c.Redis = redisClient.Get()
}

func registerJobsStore(c *Container) {
	c.JobsStore = store.NewJobsStore()
}
