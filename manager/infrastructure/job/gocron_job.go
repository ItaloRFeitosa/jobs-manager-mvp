package job

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/italorfeitosa/jobs-manager-mvp/common/config"
	"github.com/italorfeitosa/jobs-manager-mvp/common/uuid"
	"github.com/italorfeitosa/jobs-manager-mvp/manager/core"
)

type gocronJob struct {
	*gocron.Job
	id         string
	schema     config.JobSchema
	dispatcher core.JobDispatcher
	scheduler  *gocron.Scheduler
}

func New(schema config.JobSchema, dispatcher core.JobDispatcher) core.Job {
	return new(schema, dispatcher)
}

func new(schema config.JobSchema, dispatcher core.JobDispatcher) *gocronJob {
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

	if schema.SingletonMode {
		scheduler.SingletonMode()
	}

	job.Job, err = scheduler.Do(func(job core.Job) {
		err := dispatcher.Dispatch(context.Background(), job.Dispatch())
		log.Printf("err: %#+v\n", err)
	}, job)

	if err != nil {
		panic(err)
	}

	return job
}

func (j *gocronJob) Name() string {
	return j.schema.Name
}

func (j *gocronJob) Dispatch() core.JobDispatch {

	return gocronJobDispatch{
		idempotencyKey: uuid.New(),
		data:           j.schema.Data,
		name:           j.Name(),
		createdAt:      time.Now().UTC(),
	}
}

func (j *gocronJob) Run() error {
	if !j.IsActive() {
		return fmt.Errorf("job with name %s is inactive", j.Name())
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

type gocronJobDispatch struct {
	name           string
	data           any
	idempotencyKey string
	createdAt      time.Time
}

func (j gocronJobDispatch) Name() string {
	return j.name
}

func (j gocronJobDispatch) IdempotencyKey() string {
	return j.idempotencyKey
}

func (j gocronJobDispatch) Data() any {
	return j.data
}

func (j gocronJobDispatch) CreatedAt() time.Time {
	return j.createdAt
}
