package lolo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

type Client struct {
	apiKey     string
	baseUrl    string
	HttpClient *http.Client
}

func NewClient(apiKey string) (*Client, error) {
	baseUrl := os.Getenv("LO_API")

	if baseUrl == "" {
		baseUrl = "https://dev.lolo.company/api"
	}

	return &Client{
		apiKey:     apiKey,
		baseUrl:    baseUrl,
		HttpClient: http.DefaultClient,
	}, nil
}

func (client *Client) sendRequest(m, p string, b, out interface{}) error {
	req, err := client.createRequest(m, p, b)
	if err != nil {
		return err
	}

	var resp *http.Response
	resp, err = client.HttpClient.Do(req)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return fmt.Errorf("API error %s: %s", resp.Status, body)
	}

	if len(body) == 0 {
		body = []byte{'{', '}'}
	}

	if out == nil {
		return nil
	}

	return json.Unmarshal(body, &out)
}

func (client *Client) buildUrl(p string) string {
	return client.baseUrl + p
}

func (client *Client) createRequest(m, p string, b interface{}) (*http.Request, error) {
	var bReader io.Reader

	if m != "GET" && b != nil {
		bJson, err := json.Marshal(b)
		if err != nil {
			return nil, err
		}
		fmt.Println("Request body:", string(bJson))
		bReader = bytes.NewReader(bJson)
	}

	url := client.buildUrl(p)

	req, err := http.NewRequest(m, url, bReader)
	if err != nil {
		return nil, err
	}

	req.Header.Set("LO-API-KEY", client.apiKey)
	req.Header.Add("Content-Type", "application/json")

	return req, nil
}
