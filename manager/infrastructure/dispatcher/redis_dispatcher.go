package dispatcher

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/italorfeitosa/jobs-manager-mvp/manager/core"
)

type redisJobDispatcher struct {
	channel string
	redis   *redis.Client
}

func NewRedisJobDispatcher(redis *redis.Client, channel string) core.JobDispatcher {
	return &redisJobDispatcher{channel, redis}
}

func (jd *redisJobDispatcher) Dispatch(ctx context.Context, job core.JobDispatch) error {
	ackChannel := fmt.Sprintf("%s:%s:ack", jd.channel, job.IdempotencyKey())
	pubsub := jd.redis.Subscribe(context.Background(), ackChannel)
	defer pubsub.Close()

	payload, err := prepareRedisPayload(job)
	if err != nil {
		return err
	}

	err = jd.redis.Publish(ctx, jd.channel, payload).Err()

	if err != nil {
		return err
	}

	controlCh := pubsub.Channel()

	for msg := range controlCh {
		ack, err := NewAck(msg.Payload)
		if err != nil {
			return err
		}

		if ack.HasFailed() {
			return fmt.Errorf(ack.Failed)
		}

		if ack.HasSucceeded() {
			return nil
		}

		if ack.HasCanceled() {
			return nil
		}

		if ack.HasStarted() {
			continue
		}
	}

	return nil
}

func prepareRedisPayload(job core.JobDispatch) (string, error) {
	jsonAsBytes, err := json.Marshal(map[string]any{
		"idempotencyKey": job.IdempotencyKey(),
		"jobName":        job.Name(),
		"createdAt":      job.CreatedAt(),
		"data":           job.Data(),
	})

	if err != nil {
		return "", err
	}

	return string(jsonAsBytes), nil
}

func NewAck(payload string) (AckMessage, error) {
	var ackMessage AckMessage
	if err := json.Unmarshal([]byte(payload), &ackMessage); err != nil {
		return ackMessage, err
	}

	return ackMessage, nil
}

type AckMessage struct {
	Failed    string `json:"failed,omitempty"`
	Succeeded string `json:"succeeded,omitempty"`
	Started   string `json:"started,omitempty"`
	Canceled  string `json:"canceled,omitempty"`
}

func (a AckMessage) HasFailed() bool {
	return a.Failed != ""
}

func (a AckMessage) HasStarted() bool {
	return a.Started != ""
}

func (a AckMessage) HasSucceeded() bool {
	return a.Succeeded != ""
}

func (a AckMessage) HasCanceled() bool {
	return a.Canceled != ""
}
