package client

import "encoding/json"

type Response interface {
	IsOK() bool
	GetResponse() map[string]interface{}
}

type JSONResponse struct {
	response map[string]interface{}
}

func (c *JSONResponse) IsOK() bool {
	if c.response != nil && c.response["error"] == nil {
		return true
	}
	return false
}

func (c *JSONResponse) GetResponse() map[string]interface{} {
	if c.response == nil {
		c.response = map[string]interface{}{}
	}
	return c.response
}

func (c *JSONResponse) UnmarshalJSON(b []byte) error {
	err := json.Unmarshal(b, &c.response)
	if err != nil {
		return err
	}
	return nil
}
