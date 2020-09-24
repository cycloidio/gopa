package gopa

import (
	"context"
	"encoding/json"
	"net/http"
	"path"

	"github.com/open-policy-agent/opa/server/types"
)

// DataService is the services in charge
// of the Data interactions
type DataService struct {
	client *Client
	path   string
}

// NewDataService initializes a new DataService
func NewDataService(c *Client) *DataService {
	return &DataService{
		client: c,
		path:   "/v1/data",
	}
}

// CreateOrOverride creates or replaces the given data on the path p
// https://www.openpolicyagent.org/docs/latest/rest-api/#create-or-overwrite-a-document
func (ds *DataService) CreateOrOverride(ctx context.Context, p string, data map[string]interface{}) error {
	var res interface{}

	b, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = ds.client.do(ctx, http.MethodPut, path.Join(ds.path, p), b, &res)
	if err != nil {
		return err
	}

	return nil
}

// Get get's the data on the given path p
// https://www.openpolicyagent.org/docs/latest/rest-api/#get-a-document
func (ds *DataService) Get(ctx context.Context, p string) (*types.DataResponseV1, error) {
	var res types.DataResponseV1

	err := ds.client.do(ctx, http.MethodGet, path.Join(ds.path, p), noBody, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// GetWithInput get's the data on the given path p with the input i
// https://www.openpolicyagent.org/docs/latest/rest-api/#get-a-document-with-input
func (ds *DataService) GetWithInput(ctx context.Context, p string, i map[string]interface{}) (*types.DataResponseV1, error) {
	var res types.DataResponseV1

	input := map[string]interface{}{
		"input": i,
	}
	b, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	err = ds.client.do(ctx, http.MethodPost, path.Join(ds.path, p), b, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// Update updates the data on the given path p. Can be used to do partial updates
// by using the path to specify the element
// https://www.openpolicyagent.org/docs/latest/rest-api/#patch-a-document
func (ds *DataService) Update(ctx context.Context, p string, data map[string]interface{}) error {
	var res interface{}

	b, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = ds.client.do(ctx, http.MethodPut, path.Join(ds.path, p), b, &res)
	if err != nil {
		return err
	}

	return nil
}

// Delete deletes the data on the given path p
// https://www.openpolicyagent.org/docs/latest/rest-api/#delete-a-document
func (ds *DataService) Delete(ctx context.Context, p string) error {
	var res interface{}

	err := ds.client.do(ctx, http.MethodDelete, path.Join(ds.path, p), noBody, &res)
	if err != nil {
		return err
	}

	return nil
}
