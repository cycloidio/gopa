package gopa

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"path"

	"github.com/open-policy-agent/opa/server/types"
)

// QueryService is the service in charge of
// making Queries
type QueryService struct {
	client *Client
	path   string
}

// NewQueryService initializes a new QueryService
func NewQueryService(c *Client) *QueryService {
	return &QueryService{
		client: c,
		path:   "/v1/query",
	}
}

// Simple makes a simple query to the path p with the give input
// https://www.openpolicyagent.org/docs/latest/rest-api/#execute-a-simple-query
func (qs *QueryService) Simple(ctx context.Context, p string, input map[string]interface{}) ([]byte, error) {
	b, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	req, err := qs.client.request(ctx, http.MethodPost, p, b)
	if err != nil {
		return nil, err
	}

	res, err := qs.client.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	b, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return b, nil
}

// QueryAdHocOptions are the options available to the AdHoc
type QueryAdHocOptions struct {
	Query    string                 `json:"query,omitempty"`
	Input    map[string]interface{} `json:"input,omitempty"`
	Unknowns []string               `json:"unknowns,omitempty"`
}

// AdHoc makes a AdHoc query to the path p with the give opt
// https://www.openpolicyagent.org/docs/latest/rest-api/#execute-an-ad-hoc-query
func (qs *QueryService) AdHoc(ctx context.Context, p string, opt QueryAdHocOptions) (*types.QueryResponseV1, error) {
	var res types.QueryResponseV1

	b, err := json.Marshal(opt)
	if err != nil {
		return nil, err
	}

	err = qs.client.do(ctx, http.MethodPost, path.Join(qs.path, p), b, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
