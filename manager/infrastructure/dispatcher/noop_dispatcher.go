package dispatcher

import (
	"context"
	"log"
	"time"

	"github.com/italorfeitosa/jobs-manager-mvp/manager/core"
)

type noopJobDispatcher struct{}

func NewNoopJobDispatcher() core.JobDispatcher {
	return noopJobDispatcher{}
}

func (noopJobDispatcher) Dispatch(ctx context.Context, job core.JobDispatch) error {
	log.Printf("job_name: %s started", job.Name())
	time.Sleep(15 * time.Second)
	log.Printf("job_name: %s ended", job.Name())
	return nil
}
