package client_test

import (
	"net/http"
	"testing"

	"github.com/dailyburn/sailthru-go/client"
	"github.com/stretchr/testify/assert"
)

func mockClient(url string) *client.Client {
	client.BaseURL = url
	return &client.Client{HTTPClient: http.DefaultClient}
}

// Unit Tests

func TestClientDefaultHost(t *testing.T) {
	assert.Equal(t, "https://api.sailthru.com", client.BaseURL)
}
