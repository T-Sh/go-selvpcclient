package quotas

import (
	"bytes"
	"context"
	"encoding/json"
	"strings"

	"github.com/selectel/go-selvpcclient/selvpc"
)

const resourceURL = "quotas"

// GetAll returns the total amount of resources available to be allocated to projects.
func GetAll(ctx context.Context, client *selvpc.ServiceClient) ([]*Quota, *selvpc.ResponseResult, error) {
	url := strings.Join([]string{client.Endpoint, resourceURL}, "/")
	responseResult, err := client.DoRequest(ctx, "GET", url, nil)
	if err != nil {
		return nil, nil, err
	}
	if responseResult.Err != nil {
		return nil, responseResult, responseResult.Err
	}

	// Extract quotas from the response body.
	var result ResourcesQuotas
	err = responseResult.ExtractResult(&result)
	if err != nil {
		return nil, responseResult, err
	}

	return result.Quotas, responseResult, nil
}

// GetFree returns the current amount of resources available to be allocated to projects.
func GetFree(ctx context.Context, client *selvpc.ServiceClient) ([]*Quota, *selvpc.ResponseResult, error) {
	url := strings.Join([]string{client.Endpoint, resourceURL, "free"}, "/")
	responseResult, err := client.DoRequest(ctx, "GET", url, nil)
	if err != nil {
		return nil, nil, err
	}
	if responseResult.Err != nil {
		return nil, responseResult, responseResult.Err
	}

	// Extract quotas from the response body.
	var result ResourcesQuotas
	err = responseResult.ExtractResult(&result)
	if err != nil {
		return nil, responseResult, err
	}

	return result.Quotas, responseResult, nil
}

// GetProjectsQuotas returns the quotas info for all domain projects.
func GetProjectsQuotas(ctx context.Context, client *selvpc.ServiceClient) ([]*ProjectQuota, *selvpc.ResponseResult, error) {
	url := strings.Join([]string{client.Endpoint, resourceURL, "projects"}, "/")
	responseResult, err := client.DoRequest(ctx, "GET", url, nil)
	if err != nil {
		return nil, nil, err
	}
	if responseResult.Err != nil {
		return nil, responseResult, responseResult.Err
	}

	// Extract quotas from the response body.
	var result ProjectsQuotas
	err = responseResult.ExtractResult(&result)
	if err != nil {
		return nil, responseResult, err
	}

	return result.ProjectQuotas, responseResult, nil
}

// GetProjectQuotas returns the quotas info for a single project referenced by id.
func GetProjectQuotas(ctx context.Context, client *selvpc.ServiceClient, id string) ([]*Quota, *selvpc.ResponseResult, error) {
	url := strings.Join([]string{client.Endpoint, resourceURL, "projects", id}, "/")
	responseResult, err := client.DoRequest(ctx, "GET", url, nil)
	if err != nil {
		return nil, nil, err
	}
	if responseResult.Err != nil {
		return nil, responseResult, responseResult.Err
	}

	// Extract quotas from the response body.
	var result ResourcesQuotas
	err = responseResult.ExtractResult(&result)
	if err != nil {
		return nil, responseResult, err
	}

	return result.Quotas, responseResult, nil
}

// UpdateProjectQuotas updates the quotas info for a single project referenced by id.
func UpdateProjectQuotas(ctx context.Context, client *selvpc.ServiceClient, id string, updateOpts UpdateProjectQuotasOpts) ([]*Quota, *selvpc.ResponseResult, error) {
	requestBody, err := json.Marshal(&updateOpts)
	if err != nil {
		return nil, nil, err
	}

	url := strings.Join([]string{client.Endpoint, resourceURL, "projects", id}, "/")
	responseResult, err := client.DoRequest(ctx, "PATCH", url, bytes.NewReader(requestBody))
	if err != nil {
		return nil, nil, err
	}
	if responseResult.Err != nil {
		return nil, responseResult, responseResult.Err
	}

	// Extract quotas from the response body.
	var result ResourcesQuotas
	err = responseResult.ExtractResult(&result)
	if err != nil {
		return nil, responseResult, err
	}

	return result.Quotas, responseResult, nil
}
