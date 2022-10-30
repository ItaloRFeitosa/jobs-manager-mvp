package dispatcher

import (
	"github.com/italorfeitosa/jobs-manager-mvp/common/config"
	"github.com/italorfeitosa/jobs-manager-mvp/manager/core"
)

func Delegate(schema config.JobSchema) core.JobsDispatcher {
	if schema.HasHttp() {
		return NewHttpJobsDispatcher(schema.Http.URL)
	}

	return NewNoopJobDispatcher()
}
