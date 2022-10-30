package core

type JobsDispatcher interface {
	Dispatch(job JobDispatch) error
}
