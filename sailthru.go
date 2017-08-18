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

func (s *Client) PostJob(job *params.UpdateJob) (client.Response, error) {
	job.Job = "update"
	return s.c.Post(job.GetEndpoint(), job)
}
