package gopa_test

import (
	"context"
	"testing"

	"github.com/cycloidio/gopa"
	"github.com/open-policy-agent/opa/server/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPolicyService(t *testing.T) {
	var (
		policyID = "example1"
	)

	t.Run("CreateOrUpdate", func(t *testing.T) {
		t.Run("Create", func(t *testing.T) {
			c, err := gopa.NewClient()
			require.NoError(t, err)

			ctx := context.Background()
			policy := []byte(`
package opa.examples

import data.servers
import data.networks
import data.ports
`)

			res, err := c.PolicyCreateOrUpdate(ctx, policyID, policy)
			require.NoError(t, err)
			assert.NotNil(t, res)
		})

		t.Run("Update", func(t *testing.T) {
			c, err := gopa.NewClient()
			require.NoError(t, err)

			ctx := context.Background()
			policy := []byte(`
package opa.examples

import data.servers
import data.networks
`)

			res, err := c.PolicyCreateOrUpdate(ctx, policyID, policy)
			require.NoError(t, err)
			assert.NotNil(t, res)
		})

		t.Run("Error", func(t *testing.T) {
			c, err := gopa.NewClient()
			require.NoError(t, err)

			ctx := context.Background()
			policy := []byte(`
potato
`)

			_, err = c.PolicyCreateOrUpdate(ctx, policyID, policy)
			assert.Contains(t, err.Error(), types.CodeInvalidParameter)
		})
	})

	t.Run("List", func(t *testing.T) {
		t.Run("Success", func(t *testing.T) {
			c, err := gopa.NewClient()
			require.NoError(t, err)

			ctx := context.Background()

			res, err := c.PolicyList(ctx)
			require.NoError(t, err)

			require.Len(t, res.Result, 1)
			assert.Equal(t, policyID, res.Result[0].ID)
			assert.Equal(t, `
package opa.examples

import data.servers
import data.networks
`, res.Result[0].Raw)
		})
	})

	t.Run("Get", func(t *testing.T) {
		t.Run("Success", func(t *testing.T) {
			c, err := gopa.NewClient()
			require.NoError(t, err)

			ctx := context.Background()

			res, err := c.PolicyGet(ctx, policyID)
			require.NoError(t, err)

			assert.Equal(t, policyID, res.Result.ID)
			assert.Equal(t, `
package opa.examples

import data.servers
import data.networks
`, res.Result.Raw)
		})
		t.Run("Error", func(t *testing.T) {
			c, err := gopa.NewClient()
			require.NoError(t, err)

			ctx := context.Background()

			_, err = c.PolicyGet(ctx, "invalidID")
			assert.Contains(t, err.Error(), types.CodeResourceNotFound)
		})
	})

	t.Run("Delete", func(t *testing.T) {
		t.Run("Success", func(t *testing.T) {
			c, err := gopa.NewClient()
			require.NoError(t, err)

			ctx := context.Background()

			res, err := c.PolicyDelete(ctx, policyID)
			require.NoError(t, err)
			assert.NotNil(t, res)
		})
		t.Run("Error", func(t *testing.T) {
			c, err := gopa.NewClient()
			require.NoError(t, err)

			ctx := context.Background()

			_, err = c.PolicyDelete(ctx, "invalidID")
			assert.Contains(t, err.Error(), types.CodeResourceNotFound)
		})
	})
}
