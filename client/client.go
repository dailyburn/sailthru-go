package sailthru_client

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"sort"
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

func (c *Client) Post(endpoint string, obj interface{}) error {
	client := &http.Client{}

	values, err := c.formValues(obj)
	if err != nil {
		return err
	}

	r, _ := http.NewRequest("POST", fmt.Sprintf("%v/%v", baseURL, endpoint), bytes.NewBufferString(values.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(values.Encode())))

	resp, err := client.Do(r)
	fmt.Println(resp)

	return err
}

func (c *Client) signature(params string) string {
	values := []string{
		c.key,
		"json",
		params,
	}
	sort.Strings(values)
	values = append([]string{c.secret}, values...)

	signature := fmt.Sprintf("%x", md5.Sum([]byte(strings.Join(values, ""))))
	return signature
}

func (c *Client) formValues(obj interface{}) (url.Values, error) {
	dd, err := json.Marshal(obj)
	if err != nil {
		return url.Values{}, err
	}

	params := string(dd)

	formData := url.Values{}
	formData.Set("api_key", c.key)
	formData.Add("sig", c.signature(params))
	formData.Add("format", "json")
	formData.Add("json", params)

	return formData, nil
}
