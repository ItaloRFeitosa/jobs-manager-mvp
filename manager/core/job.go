package core

import (
	"time"
)

type Job interface {
	Name() string
	Dispatch() JobDispatch
	Start()
	Stop()
	Run() error
	LastRun() time.Time
	NextRun() time.Time
	RunCount() int
	IsRunning() bool
	IsActive() bool
}

type JobDispatch interface {
	Name() string
	Data() any
	IdempotencyKey() string
	CreatedAt() time.Time
}
