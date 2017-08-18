package sailthru

import (
	"github.com/dailyburn/sailthru-go/client"
	"github.com/dailyburn/sailthru-go/params"
)

type Client struct {
	c *client.Client
}

func NewClient(key, secret string) *Client {
	c := client.NewClient(key, secret)
	return &Client{c: c}
}

func (s *Client) Send(send *params.Send) (client.Response, error) {
	return s.c.Post(send.GetEndpoint(), send)
}

func (s *Client) GetSend(sendID string) (client.Response, error) {
	params := map[string]string{
		"send_id": sendID,
	}
	return s.c.Get("send", params)
}

func (s *Client) CancelSend(sendID string) (client.Response, error) {
	params := map[string]string{
		"send_id": sendID,
	}
	return s.c.Delete("send", params)
}

func (s *Client) PostJob(job *params.UpdateJob) (client.Response, error) {
	job.Job = "update"
	return s.c.Post(job.GetEndpoint(), job)
}

func (s *Client) GetJobStatus(jobID string) (client.Response, error) {
	params := map[string]string{
		"job_id": jobID,
	}
	return s.c.Get("job", params)
}
