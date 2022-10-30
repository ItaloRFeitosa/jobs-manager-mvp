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

func (noopJobsDispatcher) Dispatch(job core.JobDispatch) error {
	log.Printf("job_name: %s started", job.Name())
	time.Sleep(15 * time.Second)
	log.Printf("job_name: %s ended", job.Name())
	return nil
}
