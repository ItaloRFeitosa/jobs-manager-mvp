package config

import (
	"encoding/json"
	"io/ioutil"

	"github.com/italorfeitosa/jobs-manager-mvp/manager/core"
)

type Config struct {
	JobSchemas []core.JobSchema `json:"jobSchemas"`
}

var config *Config

func Get() *Config {
	if config != nil {
		return config
	}

	var (
		err       error
		newConfig Config
	)

	configAsJson, err := ioutil.ReadFile("./config.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(configAsJson, &newConfig)
	if err != nil {
		panic(err)
	}

	config = &newConfig

	return config
}
