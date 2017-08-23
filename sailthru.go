// Package sailthru is the main package for interacting with the Sailthru API
// as per http://docs.sailthru.com/api
package sailthru

import (
	"github.com/dailyburn/sailthru-go/client"
	"github.com/dailyburn/sailthru-go/params"
)

type Client struct {
	c *client.Client
}

// NewClient returns a new Sailthru Client with the underlying http.Client
// configured with the correct Sailthru settings
func NewClient(key, secret string) *Client {
	c := client.NewClient(key, secret)
	return &Client{c: c}
}

// Send remotely sends an email template to a single email address
func (s *Client) Send(send *params.Send) (client.Response, error) {
	return s.c.Post(send.GetEndpoint(), send)
}

// GetSend gets the status of a send
func (s *Client) GetSend(sendID string) (client.Response, error) {
	params := map[string]string{
		"send_id": sendID,
	}
	return s.c.Get("send", params)
}

// CancelSend cancels a send that was scheduled for a future time
func (s *Client) CancelSend(sendID string) (client.Response, error) {
	params := map[string]string{
		"send_id": sendID,
	}
	return s.c.Delete("send", params)
}

// ProcessImportJob performs an import from a given CSV containing one email per line
func (s *Client) ProcessImportJob(job *params.ImportJob) (client.Response, error) {
	job.Job = "import"
	return s.c.Post(job.GetEndpoint(), job)
}

// ProcessUpdateJob performs a bulk update of any number of user profiles
func (s *Client) ProcessUpdateJob(job *params.UpdateJob) (client.Response, error) {
	job.Job = "update"
	return s.c.Post(job.GetEndpoint(), job)
}

// GetJobStatus gets the status of a job
func (s *Client) GetJobStatus(jobID string) (client.Response, error) {
	params := map[string]string{
		"job_id": jobID,
	}
	return s.c.Get("job", params)
}

func (s *Client) GetUser(user *params.User) (client.Response, error) {
	return s.c.Get("user", user)
}

func (s *Client) SaveUser(user *params.SaveUser) (client.Response, error) {
	return s.c.Post("user", user)
}
