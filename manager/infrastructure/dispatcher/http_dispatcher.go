package dispatcher

import (
	"context"
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/italorfeitosa/jobs-manager-mvp/manager/core"
)

type httpJobDispatcher struct {
	client *resty.Client
	url    string
}

func NewHttpJobDispatcher(url string) core.JobDispatcher {
	client := resty.New()
	return &httpJobDispatcher{client, url}
}

func (h *httpJobDispatcher) Dispatch(ctx context.Context, job core.JobDispatch) error {
	response, err := h.client.R().
		SetContext(ctx).
		SetBody(map[string]any{
			"idempotencyKey": job.IdempotencyKey(),
			"jobName":        job.Name(),
			"createdAt":      job.CreatedAt(),
			"data":           job.Data(),
		}).
		Post(h.url)

	if err != nil {
		return err
	}

	if response.IsError() {
		return fmt.Errorf("error on http dispatch got status %d", response.StatusCode())
	}
	return nil
}
