package projects

import (
	"context"
	"strings"

	"github.com/selectel/go-selvpcclient/selvpcclient"
)

const resourceURL = "projects"

// Get returns a single project by its id.
func Get(ctx context.Context, client *selvpcclient.ServiceClient, id string) (*Project, *selvpcclient.ResponseResult, error) {
	url := strings.Join([]string{client.Endpoint, resourceURL, id}, "/")
	responseResult, err := client.DoRequest(ctx, "GET", url, nil)
	if err != nil {
		return nil, nil, err
	}
	if responseResult.Err != nil {
		return nil, responseResult, responseResult.Err
	}

	// Extract a project from the response body.
	var result struct {
		Project *Project `json:"project"`
	}
	err = responseResult.ExtractResult(&result)
	if err != nil {
		return nil, responseResult, err
	}

	return result.Project, responseResult, nil
}