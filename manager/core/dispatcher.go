package core

type JobsDispatcher interface {
	Dispatch(job Job) error
}
