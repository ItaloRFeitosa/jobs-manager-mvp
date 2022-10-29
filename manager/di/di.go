package di

import (
	"github.com/italorfeitosa/jobs-manager-mvp/common/config"
	"github.com/italorfeitosa/jobs-manager-mvp/manager/core"
	"github.com/italorfeitosa/jobs-manager-mvp/manager/infrastructure/dispatcher"
	"github.com/italorfeitosa/jobs-manager-mvp/manager/infrastructure/store"
)

type Container struct {
	Config         *config.Config
	JobsStore      core.JobsStore
	JobsDispatcher core.JobsDispatcher
}

func New() *Container {
	return &Container{
		Config:         config.Get(),
		JobsStore:      store.NewJobsStore(),
		JobsDispatcher: dispatcher.NewNoopJobDispatcher(),
	}
}
