package gopa

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewClient(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		u, _ := url.Parse(DefaultURL)
		ec := &Client{
			url:    u,
			client: http.DefaultClient,
		}

		ec.policysvc = NewPolicyService(ec)
		ec.datasvc = NewDataService(ec)
		ec.querysvc = NewQueryService(ec)

		c, err := NewClient()
		require.NoError(t, err)
		assert.Equal(t, ec, c)

		assert.Implements(t, (*Service)(nil), c)
	})
}

func TestSetURL(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		u, _ := url.Parse("http://google.com")
		ec := &Client{
			url:    u,
			client: http.DefaultClient,
		}

		ec.policysvc = NewPolicyService(ec)
		ec.datasvc = NewDataService(ec)
		ec.querysvc = NewQueryService(ec)

		c, err := NewClient(SetURL("http://google.com"))
		require.NoError(t, err)
		assert.Equal(t, ec, c)
	})
}

func TestSetToken(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		u, _ := url.Parse(DefaultURL)
		token := "my-token"
		ec := &Client{
			url:    u,
			client: http.DefaultClient,
			token:  token,
		}

		ec.policysvc = NewPolicyService(ec)
		ec.datasvc = NewDataService(ec)
		ec.querysvc = NewQueryService(ec)

		c, err := NewClient(SetToken(token))
		require.NoError(t, err)
		assert.Equal(t, ec, c)
	})
}

func TestSetClient(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		cl := &http.Client{}
		u, _ := url.Parse(DefaultURL)
		ec := &Client{
			url:    u,
			client: cl,
		}

		ec.policysvc = NewPolicyService(ec)
		ec.datasvc = NewDataService(ec)
		ec.querysvc = NewQueryService(ec)

		c, err := NewClient(SetClient(cl))
		require.NoError(t, err)
		assert.Equal(t, ec, c)
	})
}
