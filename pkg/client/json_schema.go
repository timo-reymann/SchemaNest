package client

import (
	"context"
	"fmt"
	"github.com/timo-reymann/SchemaNest/pkg/api"
	"net/http"
	"os"
)

type Client struct {
	context   context.Context
	apiClient *api.ClientWithResponses
	apiKey    string
}

func NewClient(baseUrl string, apiKey string) (*Client, error) {
	c, err := api.NewClientWithResponses(baseUrl)
	if err != nil {
		return nil, err
	}

	return &Client{
		apiClient: c,
		apiKey:    apiKey,
		context:   context.Background(),
	}, nil
}

func (c *Client) UploadJsonSchema(identifier, version, localPath string) error {
	file, err := os.Open(localPath)
	if err != nil {
		return fmt.Errorf("failed to open file %s: %s", localPath, err)
	}

	res, err := c.apiClient.PostApiSchemaJsonSchemaIdentifierVersionVersionWithBodyWithResponse(
		c.context,
		identifier,
		version,
		"application/json",
		file,
		func(ctx context.Context, req *http.Request) error {
			if c.apiKey != "" {
				req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
			}
			return nil
		},
	)

	if err != nil {
		return fmt.Errorf("failed to communicate with SchemaNest: %s", err)
	}

	if res.StatusCode() != http.StatusCreated {
		if res.StatusCode() == http.StatusConflict {
			return fmt.Errorf("failed to upload file: %s", res.JSON409.Error)
		}
		return fmt.Errorf("failed to upload file: %s", res.Body)
	}
	return nil
}
