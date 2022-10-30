package dispatcher

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/italorfeitosa/jobs-manager-mvp/manager/core"
)

type httpJobsDispatcher struct {
	client *resty.Client
	url    string
}

func NewHttpJobsDispatcher(url string) core.JobsDispatcher {
	client := resty.New()
	return &httpJobsDispatcher{client, url}
}

func (h *httpJobsDispatcher) Dispatch(job core.JobDispatch) error {

	response, err := h.client.R().
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
