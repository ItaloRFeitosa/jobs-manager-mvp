package job

import (
	"fmt"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/italorfeitosa/jobs-manager-mvp/common/uuid"
	"github.com/italorfeitosa/jobs-manager-mvp/manager/core"
)

type gocronJob struct {
	*gocron.Job
	id         string
	schema     core.JobSchema
	dispatcher core.JobsDispatcher
	scheduler  *gocron.Scheduler
}

func New(schema core.JobSchema, dispatcher core.JobsDispatcher) core.Job {
	return new(schema, dispatcher)
}

func new(schema core.JobSchema, dispatcher core.JobsDispatcher) *gocronJob {
	var err error

	id := uuid.New()

	scheduler := gocron.NewScheduler(time.Local)
	scheduler.TagsUnique()

	job := &gocronJob{id: id, scheduler: scheduler, schema: schema, dispatcher: dispatcher}

	scheduler.Tag(schema.Name)

	if schema.Every != "" {
		scheduler.Every(schema.Every)
	}

	if schema.Cron != "" {
		scheduler.CronWithSeconds(schema.Cron)
	}

	job.Job, err = scheduler.Do(dispatcher.Dispatch, job)

	if err != nil {
		panic(err)
	}

	return job
}

func (j *gocronJob) ID() string {
	return j.id
}

func (j *gocronJob) Name() string {
	return j.schema.Name
}

func (j *gocronJob) Schema() core.JobSchema {
	return j.schema
}
func (j *gocronJob) Run() error {
	if !j.IsActive() {
		return fmt.Errorf("job with id %s is inactive", j.ID())
	}

	return j.scheduler.RunByTag(j.schema.Name)
}

func (j *gocronJob) Start() {
	j.scheduler.StartAsync()
}

func (j *gocronJob) Stop() {
	go func() {
		j.scheduler.Stop()
		newJob := new(j.schema, j.dispatcher)
		newJob.id = j.id
		j.scheduler = newJob.scheduler
		j.Job = newJob.Job
	}()
}

func (j *gocronJob) IsActive() bool {
	return j.scheduler.IsRunning()
}
