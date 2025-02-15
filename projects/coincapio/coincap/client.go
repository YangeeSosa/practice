package coincap

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Client struct {
	client *http.Client
}

func NewClient() (*Client, error) {
	return &Client{
		client: &http.Client{
			Transport: &loggingRoundTripper{
				logger: os.Stdout,
				next:   http.DefaultTransport,
			},
		},
	}, nil
}

func (c *Client) GetAssets() ([]Asset, error) {
	resp, err := c.client.Get(
		"https://api.coincap.io/v2/assets/",
	)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var r assetsResponse
	err = json.Unmarshal(body, &r)
	if err != nil {
		return nil, err
	}

	return r.Assets, nil
}

func (c *Client) GetAsset(name string) (Asset, error) {
	url := fmt.Sprintf("https://api.coincap.io/v2/assets/%s", name)
	resp, err := c.client.Get(
		url,
	)

	if err != nil {
		return Asset{}, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Asset{}, err
	}

	var r assetResponse
	err = json.Unmarshal(body, &r)
	if err != nil {
		return Asset{}, err
	}

	return r.Asset, nil
}
