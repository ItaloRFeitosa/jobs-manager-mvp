package dispatcher

import (
	"log"
	"time"

	"github.com/italorfeitosa/jobs-manager-mvp/manager/core"
)

type noopJobsDispatcher struct{}

func NewNoopJobDispatcher() core.JobsDispatcher {
	return noopJobsDispatcher{}
}

func (noopJobsDispatcher) Dispatch(job core.Job) error {
	log.Printf("job_id: %s started", job.ID())
	log.Printf("job_definition: %#+v\n", job.Schema())
	time.Sleep(15 * time.Second)
	log.Printf("job_id: %s ended", job.ID())
	return nil
}
