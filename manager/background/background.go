package background

import (
	"github.com/italorfeitosa/jobs-manager-mvp/manager/core"
	"github.com/italorfeitosa/jobs-manager-mvp/manager/di"
	"github.com/italorfeitosa/jobs-manager-mvp/manager/infrastructure/job"
)

func Start(deps *di.Container) {
	for _, schema := range deps.Config.JobSchemas {
		job := job.New(schema, deps.DelegateDispatcher(schema))
		deps.JobsStore.Save(job)
	}

	deps.JobsStore.ForEach(func(j core.Job) {
		j.Start()
	})
}
