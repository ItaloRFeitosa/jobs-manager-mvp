package config

type JobSchema struct {
	Name          string
	Cron          string
	Every         string
	SingletonMode bool
	Http          HttpSchema
	Redis         RedisSchema
	Data          map[string]any
}

func (s JobSchema) HasHttp() bool {
	return s.Http.NotEmpty()
}

func (s JobSchema) HasRedis() bool {
	return s.Redis.NotEmpty()
}

type HttpSchema struct {
	URL string
}

func (s HttpSchema) NotEmpty() bool {
	return s.URL != ""
}

type RedisSchema struct {
	Channel string
}

func (s RedisSchema) NotEmpty() bool {
	return s.Channel != ""
}
