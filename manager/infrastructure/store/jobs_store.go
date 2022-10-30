package store

import (
	"fmt"
	"sync"

	"github.com/italorfeitosa/jobs-manager-mvp/manager/core"
)

type Jobs struct {
	jobs sync.Map
}

func NewJobsStore() core.JobsStore {
	var jobs sync.Map
	return &Jobs{jobs}
}

func (r *Jobs) GetAll() []core.Job {
	var jobs []core.Job

	r.jobs.Range(func(key, value any) bool {
		jobs = append(jobs, value.(core.Job))
		return true
	})

	return jobs
}

func (r *Jobs) Get(name string) (core.Job, error) {
	found, ok := r.jobs.Load(name)
	if !ok {
		return nil, fmt.Errorf("job with name: %s not found", name)
	}

	return found.(core.Job), nil
}

func (r *Jobs) ForEach(fn func(core.Job)) {
	r.jobs.Range(func(key, value any) bool {
		fn(value.(core.Job))
		return true
	})
}

func (r *Jobs) Save(job core.Job) {
	r.jobs.Store(job.Name(), job)
}
