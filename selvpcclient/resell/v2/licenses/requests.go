package licenses

import (
	"context"
	"strings"

	"github.com/selectel/go-selvpcclient/selvpcclient"
)

const resourceURL = "licenses"

// Get returns a single license by its id.
func Get(ctx context.Context, client *selvpcclient.ServiceClient, id string) (*License, *selvpcclient.ResponseResult, error) {
	url := strings.Join([]string{client.Endpoint, resourceURL, id}, "/")
	responseResult, err := client.DoRequest(ctx, "GET", url, nil)
	if err != nil {
		return nil, nil, err
	}
	if responseResult.Err != nil {
		return nil, responseResult, responseResult.Err
	}

	// Extract a license from the response body.
	var result struct {
		License *License `json:"license"`
	}
	err = responseResult.ExtractResult(&result)
	if err != nil {
		return nil, responseResult, err
	}

	return result.License, responseResult, nil
}

// List gets a list of licenses in the current domain.
func List(ctx context.Context, client *selvpcclient.ServiceClient) ([]*License, *selvpcclient.ResponseResult, error) {
	url := strings.Join([]string{client.Endpoint, resourceURL}, "/")
	responseResult, err := client.DoRequest(ctx, "GET", url, nil)
	if err != nil {
		return nil, nil, err
	}
	if responseResult.Err != nil {
		return nil, responseResult, responseResult.Err
	}

	// Extract licenses from the response body.
	var result struct {
		Licenses []*License `json:"licenses"`
	}
	err = responseResult.ExtractResult(&result)
	if err != nil {
		return nil, responseResult, err
	}

	return result.Licenses, responseResult, nil
}