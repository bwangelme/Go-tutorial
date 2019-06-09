package testify_exam

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type Client struct {
	http.Client
}

func (c *Client) DoRequest(method, url, contentType string, body io.Reader) (*http.Response, error) {
	method = strings.ToUpper(method)
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", contentType)
	return c.Do(req)
}

func NewClient() *Client {
	return &Client{
		Client: http.Client{},
	}
}

func UpdateFile(filename string, client interface{}) error {
	raw_client, ok := client.(*Client)
	if !ok {
		return fmt.Errorf("INVALID client")
	}

	body, err := os.Open(filename)
	url := fmt.Sprintf("%s/%s", "https://www.bwangel.me", filename)
	if err != nil {
		return err
	}

	resp, err := raw_client.DoRequest("PUT", url, "application/octet-stream", body)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf("INVALID REQUEST")
	}

	return nil
}
