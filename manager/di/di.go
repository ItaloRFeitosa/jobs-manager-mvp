package di

import (
	"github.com/italorfeitosa/jobs-manager-mvp/common/config"
	"github.com/italorfeitosa/jobs-manager-mvp/manager/core"
	"github.com/italorfeitosa/jobs-manager-mvp/manager/infrastructure/dispatcher"
	"github.com/italorfeitosa/jobs-manager-mvp/manager/infrastructure/redis"
	"github.com/italorfeitosa/jobs-manager-mvp/manager/infrastructure/store"
)

type Container struct {
	Config             *config.Config
	JobsStore          core.JobsStore
	Redis              redis.Client
	DelegateDispatcher func(schema config.JobSchema) core.JobsDispatcher
}

func New() *Container {
	return &Container{
		Config:             config.Get(),
		Redis:              redis.Get(),
		JobsStore:          store.NewJobsStore(),
		DelegateDispatcher: dispatcher.Delegate,
	}
}
