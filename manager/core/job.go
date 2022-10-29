package core

import (
	"time"
)

type Job interface {
	ID() string
	Name() string
	Schema() JobSchema
	Start()
	Stop()
	Run() error
	LastRun() time.Time
	NextRun() time.Time
	RunCount() int
	IsRunning() bool
	IsActive() bool
}

type JobSchema struct {
	Name          string
	Cron          string
	Every         string
	Args          string
	SingletonMode bool
}
