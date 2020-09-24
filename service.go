package gopa

import (
	"context"

	"github.com/open-policy-agent/opa/server/types"
)

// Service is the main interface that we support of OPA
type Service interface {
	PolicyCreateOrUpdate(ctx context.Context, id string, policy []byte) (*types.PolicyPutResponseV1, error)
	PolicyList(ctx context.Context) (*types.PolicyListResponseV1, error)
	PolicyGet(ctx context.Context, id string) (*types.PolicyGetResponseV1, error)
	PolicyDelete(ctx context.Context, id string) (*types.PolicyDeleteResponseV1, error)

	DataCreateOrOverride(ctx context.Context, path string, data map[string]interface{}) error
	DataGet(ctx context.Context, path string) (*types.DataResponseV1, error)
	DataGetWithInput(ctx context.Context, path string, input map[string]interface{}) (*types.DataResponseV1, error)
	DataUpdate(ctx context.Context, path string, data map[string]interface{}) error
	DataDelete(ctx context.Context, path string) error

	QuerySimple(ctx context.Context, path string, input map[string]interface{}) ([]byte, error)
	QueryAdHoc(ctx context.Context, path string, opt QueryAdHocOptions) (*types.QueryResponseV1, error)
}

// PolicyCreateOrUpdate creates or updates the policy with the give id and the content policy
// https://www.openpolicyagent.org/docs/latest/rest-api/#create-or-update-a-policy
func (c *Client) PolicyCreateOrUpdate(ctx context.Context, id string, policy []byte) (*types.PolicyPutResponseV1, error) {
	return c.policysvc.CreateOrUpdate(ctx, id, policy)
}

// PolicyList returns all the policies
// https://www.openpolicyagent.org/docs/latest/rest-api/#list-policies
func (c *Client) PolicyList(ctx context.Context) (*types.PolicyListResponseV1, error) {
	return c.policysvc.List(ctx)
}

// PolicyGet returns the policy with the given id
// https://www.openpolicyagent.org/docs/latest/rest-api/#get-a-policy
func (c *Client) PolicyGet(ctx context.Context, id string) (*types.PolicyGetResponseV1, error) {
	return c.policysvc.Get(ctx, id)
}

// PolicyDelete deletes the policy with the given id
// https://www.openpolicyagent.org/docs/latest/rest-api/#delete-a-policy
func (c *Client) PolicyDelete(ctx context.Context, id string) (*types.PolicyDeleteResponseV1, error) {
	return c.policysvc.Delete(ctx, id)
}

// DataCreateOrOverride creates or replaces the given data on the path p
// https://www.openpolicyagent.org/docs/latest/rest-api/#create-or-overwrite-a-document
func (c *Client) DataCreateOrOverride(ctx context.Context, path string, data map[string]interface{}) error {
	return c.datasvc.CreateOrOverride(ctx, path, data)
}

// DataGet get's the data on the given path p
// https://www.openpolicyagent.org/docs/latest/rest-api/#get-a-document
func (c *Client) DataGet(ctx context.Context, path string) (*types.DataResponseV1, error) {
	return c.datasvc.Get(ctx, path)
}

// DataGetWithInput get's the data on the given path p with the input i
// https://www.openpolicyagent.org/docs/latest/rest-api/#get-a-document-with-input
func (c *Client) DataGetWithInput(ctx context.Context, path string, input map[string]interface{}) (*types.DataResponseV1, error) {
	return c.datasvc.GetWithInput(ctx, path, input)
}

// DataUpdate updates the data on the given path p. Can be used to do partial updates
// by using the path to specify the element
// https://www.openpolicyagent.org/docs/latest/rest-api/#patch-a-document
func (c *Client) DataUpdate(ctx context.Context, path string, data map[string]interface{}) error {
	return c.datasvc.Update(ctx, path, data)
}

// DataDelete deletes the data on the given path p
// https://www.openpolicyagent.org/docs/latest/rest-api/#delete-a-document
func (c *Client) DataDelete(ctx context.Context, path string) error {
	return c.datasvc.Delete(ctx, path)
}

// QuerySimple makes a simple query to the path p with the give input
// https://www.openpolicyagent.org/docs/latest/rest-api/#execute-a-simple-query
func (c *Client) QuerySimple(ctx context.Context, path string, input map[string]interface{}) ([]byte, error) {
	return c.querysvc.Simple(ctx, path, input)
}

// QueryAdHoc makes a AdHoc query to the path p with the give opt
// https://www.openpolicyagent.org/docs/latest/rest-api/#execute-an-ad-hoc-query
func (c *Client) QueryAdHoc(ctx context.Context, path string, opt QueryAdHocOptions) (*types.QueryResponseV1, error) {
	return c.querysvc.AdHoc(ctx, path, opt)
}
