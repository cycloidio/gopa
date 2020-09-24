package gopa_test

import (
	"context"
	"path"
	"testing"

	"github.com/cycloidio/gopa"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDataService(t *testing.T) {
	dataRootPath := "test-data"
	dataBody := map[string]interface{}{
		"example": map[string]interface{}{
			"key": "value",
		},
	}

	t.Run("CreateOrOverride", func(t *testing.T) {
		c, err := gopa.NewClient()
		require.NoError(t, err)

		ctx := context.Background()

		err = c.DataCreateOrOverride(ctx, dataRootPath, dataBody)
		require.NoError(t, err)
	})

	t.Run("Get", func(t *testing.T) {
		c, err := gopa.NewClient()
		require.NoError(t, err)

		ctx := context.Background()

		res, err := c.DataGet(ctx, dataRootPath)
		require.NoError(t, err)
		r := *res.Result
		assert.Equal(t, dataBody, r.(map[string]interface{}))
	})

	t.Run("Get_Partially", func(t *testing.T) {
		c, err := gopa.NewClient()
		require.NoError(t, err)

		ctx := context.Background()

		res, err := c.DataGet(ctx, path.Join(dataRootPath, "example"))
		require.NoError(t, err)
		r := *res.Result
		assert.Equal(t, dataBody["example"], r.(map[string]interface{}))
	})

	t.Run("GetWithInput", func(t *testing.T) {
		c, err := gopa.NewClient()
		require.NoError(t, err)

		ctx := context.Background()
		policy := []byte(`
package opa.examples

import input.example.flag

allow_request { flag == true }
`)
		policyID := "example-data"

		_, err = c.PolicyCreateOrUpdate(ctx, policyID, policy)
		require.NoError(t, err)

		input := map[string]interface{}{
			"example": map[string]interface{}{
				"flag": true,
			},
		}

		res, err := c.DataGetWithInput(ctx, "/opa/examples/allow_request", input)
		require.NoError(t, err)
		r := *res.Result
		assert.Equal(t, true, r.(bool))

		_, err = c.PolicyDelete(ctx, policyID)
		require.NoError(t, err)
	})

	t.Run("Update", func(t *testing.T) {
		c, err := gopa.NewClient()
		require.NoError(t, err)
		newBody := map[string]interface{}{
			"example2": map[string]interface{}{
				"key2": "value",
				"map": map[string]interface{}{
					"key3": "value3",
				},
			},
		}

		ctx := context.Background()

		err = c.DataUpdate(ctx, dataRootPath, newBody)
		require.NoError(t, err)

		res, err := c.DataGet(ctx, dataRootPath)
		require.NoError(t, err)
		r := *res.Result
		assert.Equal(t, newBody, r.(map[string]interface{}), "New updated body")

		newMapBody := map[string]interface{}{
			"key4": "value4",
		}
		err = c.DataUpdate(ctx, path.Join(dataRootPath, "map"), newMapBody)
		require.NoError(t, err)

		endBody := newBody
		endBody["map"] = newMapBody

		res, err = c.DataGet(ctx, dataRootPath)
		require.NoError(t, err)
		r = *res.Result
		assert.Equal(t, endBody, r.(map[string]interface{}), "New patch updated body")
	})

	t.Run("Delete", func(t *testing.T) {
		c, err := gopa.NewClient()
		require.NoError(t, err)

		ctx := context.Background()

		err = c.DataDelete(ctx, dataRootPath)
		require.NoError(t, err)

		res, err := c.DataGet(ctx, dataRootPath)
		require.NoError(t, err)
		assert.Empty(t, res.Result)
	})
}
