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
}

func NewClient(baseUrl string) (*Client, error) {
	c, err := api.NewClientWithResponses(baseUrl)
	if err != nil {
		return nil, err
	}

	return &Client{
		apiClient: c,
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
	)

	if err != nil {
		return fmt.Errorf("failed to communicate with SchemaNest: %s", err)
	}

	if res.StatusCode() != http.StatusCreated {
		return fmt.Errorf("failed to upload file: %s", res.JSON409.Error)
	}
	return nil
}
