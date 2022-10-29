package core

type JobsStore interface {
	GetAll() []Job
	Get(id string) (Job, error)
	Save(job Job)
	ForEach(fn func(Job))
}
