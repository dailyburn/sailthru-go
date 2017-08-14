package sailthru_client_test

import (
	"net/http"
	"testing"

	"github.com/dailyburn/sailthru-go/client"
	"github.com/stretchr/testify/assert"
)

func mockClient(url string) *sailthru_client.Client {
	sailthru_client.BaseURL = url
	return &sailthru_client.Client{HTTPClient: http.DefaultClient}
}

// Unit Tests

func TestClientDefaultHost(t *testing.T) {
	assert.Equal(t, "https://api.sailthru.com", sailthru_client.BaseURL)
}
