package client_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dailyburn/sailthru-go/client"
	"github.com/stretchr/testify/assert"
)

func mockClient(url string) *client.Client {
	client.BaseURL = url
	c := client.NewClient("key", "secret")
	return c
}

// Unit Tests

func TestClientDefaultHost(t *testing.T) {
	assert.Equal(t, "https://api.sailthru.com", client.BaseURL)
}

// Functional Tests

func TestDefaultHeaders(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "application/x-www-form-urlencoded", r.Header.Get("Content-Type"))
		assert.Equal(t, "72", r.Header.Get("Content-Length"))
	}))
	defer server.Close()
	_, err := mockClient(server.URL).Post("", "")
	assert.NoError(t, err)
}

func TestPost(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
	}))
	defer server.Close()
	_, err := mockClient(server.URL).Post("", "")
	assert.NoError(t, err)
}

func Test500InternalServerError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()
	_, err := mockClient(server.URL).Post("", "")
	assert.EqualError(t, err, "500 Internal Server Error")
}

func TestInvalidJSONPayload(t *testing.T) {
	payload := make(chan int)
	_, err := mockClient("").Post("", payload)
	assert.Error(t, err)
}
