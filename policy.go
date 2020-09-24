package gopa

import (
	"context"
	"net/http"
	"path"

	"github.com/open-policy-agent/opa/server/types"
)

// PolicyService is the service in charge of
// the Policy interactions
type PolicyService struct {
	client *Client
	path   string
}

// NewPolicyService initializes a new PolicyService
func NewPolicyService(c *Client) *PolicyService {
	return &PolicyService{
		client: c,
		path:   "/v1/policies",
	}
}

// CreateOrUpdate creates or updates the policy with the give id and the content policy
// https://www.openpolicyagent.org/docs/latest/rest-api/#create-or-update-a-policy
func (ps *PolicyService) CreateOrUpdate(ctx context.Context, id string, policy []byte) (*types.PolicyPutResponseV1, error) {
	var res types.PolicyPutResponseV1

	err := ps.client.do(ctx, http.MethodPut, path.Join(ps.path, id), policy, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// List returns all the policies
// https://www.openpolicyagent.org/docs/latest/rest-api/#list-policies
func (ps *PolicyService) List(ctx context.Context) (*types.PolicyListResponseV1, error) {
	var res types.PolicyListResponseV1

	err := ps.client.do(ctx, http.MethodGet, ps.path, noBody, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// Get returns the policy with the given id
// https://www.openpolicyagent.org/docs/latest/rest-api/#get-a-policy
func (ps *PolicyService) Get(ctx context.Context, id string) (*types.PolicyGetResponseV1, error) {
	var res types.PolicyGetResponseV1

	err := ps.client.do(ctx, http.MethodGet, path.Join(ps.path, id), noBody, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// Delete deletes the policy with the given id
// https://www.openpolicyagent.org/docs/latest/rest-api/#delete-a-policy
func (ps *PolicyService) Delete(ctx context.Context, id string) (*types.PolicyDeleteResponseV1, error) {
	var res types.PolicyDeleteResponseV1

	err := ps.client.do(ctx, http.MethodDelete, path.Join(ps.path, id), noBody, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
