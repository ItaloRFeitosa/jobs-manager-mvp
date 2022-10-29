package manager

import (
	"github.com/italorfeitosa/jobs-manager-mvp/manager/api"
	"github.com/italorfeitosa/jobs-manager-mvp/manager/background"
	"github.com/italorfeitosa/jobs-manager-mvp/manager/di"
)

func Start() {
	deps := di.New()
	go background.Start(deps)

	api.Start(deps)
}
