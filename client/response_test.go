package client_test

import (
	"encoding/json"
	"testing"

	"github.com/dailyburn/sailthru-go/client"
	"github.com/stretchr/testify/assert"
)

func TestJSONResponseUnmarshal(t *testing.T) {
	response := &client.JSONResponse{}
	payload := "{\"send_id\":\"1\"}"
	json.Unmarshal([]byte(payload), &response)

	expected := map[string]interface{}{"send_id": "1"}
	assert.Equal(t, expected, response.GetResponse())
}

func TestJSONResponseInvalidPayloadUnmarshal(t *testing.T) {
	response := &client.JSONResponse{}
	err := response.UnmarshalJSON([]byte("*"))
	expected := map[string]interface{}{}
	assert.Error(t, err)
	assert.Equal(t, expected, response.GetResponse())
}

func TestJSONResponseIsOK(t *testing.T) {
	response := &client.JSONResponse{}
	payload := "{\"send_id\":\"1\"}"
	json.Unmarshal([]byte(payload), &response)
	assert.True(t, response.IsOK())
}

func TestJSONResponseIsOKWithErrorPayload(t *testing.T) {
	response := &client.JSONResponse{}
	payload := "{\"error\":5,\"errormsg\":\"Signature hash does not match\"}"
	json.Unmarshal([]byte(payload), &response)
	assert.False(t, response.IsOK())
}

func TestJSONResponseIsOKWithoutPayload(t *testing.T) {
	response := &client.JSONResponse{}
	assert.False(t, response.IsOK())
}
