package sailthru_client

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const (
	baseURL = "https://api.sailthru.com"
)

type Client struct {
	key    string
	secret string
}

func NewClient(key, secret string) *Client {
	return &Client{key: key, secret: secret}
}

func (c *Client) Post(endpoint string, params []byte) error {
	client := &http.Client{}

	values, err := c.formValues(params)
	if err != nil {
		return err
	}

	r, _ := http.NewRequest("POST", fmt.Sprintf("%v/%v", baseURL, endpoint), bytes.NewBufferString(values.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(values.Encode())))

	_, err = client.Do(r)

	return err
}

func (c *Client) signature(params []byte) (string, error) {
	sigElements := []string{
		c.secret,
		c.key,
		"json",
		string(params),
	}

	sig := fmt.Sprintf("%x", md5.Sum([]byte(strings.Join(sigElements, ""))))
	return sig, nil
}

func (c *Client) formValues(params []byte) (url.Values, error) {
	signature, err := c.signature(params)
	if err != nil {
		return url.Values{}, err
	}

	formData := url.Values{}
	formData.Set("api_key", c.key)
	formData.Add("sig", signature)
	formData.Add("format", "json")
	formData.Add("json", string(params))

	return formData, nil
}
