package gopa

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path"
)

// Client is the main struct to connect and use OPA
type Client struct {
	client *http.Client
	url    *url.URL
	token  string

	policysvc *PolicyService
	datasvc   *DataService
	querysvc  *QueryService
}

// Defaults
const (
	DefaultURL = "http://localhost:8181"
)

var (
	noBody []byte = nil
)

// ClientOptionFunc is a type used to configure the Client
// on initialization time
type ClientOptionFunc func(*Client) error

// SetURL sets the u as URL
func SetURL(u string) ClientOptionFunc {
	return func(c *Client) error {
		pu, err := url.Parse(u)
		if err != nil {
			return err
		}
		c.url = pu
		return nil
	}
}

// SetToken sets the token to use on the requests
func SetToken(token string) ClientOptionFunc {
	return func(c *Client) error {
		c.token = token
		return nil
	}
}

// SetClient sets the http client used to make the
// requests to OPA
func SetClient(client *http.Client) ClientOptionFunc {
	return func(c *Client) error {
		c.client = client
		return nil
	}
}

// NewClient initializes a new client that can be
// configured with the opts
func NewClient(opts ...ClientOptionFunc) (*Client, error) {
	c := &Client{
		client: http.DefaultClient,
	}

	err := SetURL(DefaultURL)(c)
	if err != nil {
		return nil, err
	}

	for _, o := range opts {
		if err := o(c); err != nil {
			return nil, err
		}
	}

	c.policysvc = NewPolicyService(c)
	c.datasvc = NewDataService(c)
	c.querysvc = NewQueryService(c)

	return c, nil
}

// APIError models an error response sent to the client.
// We cannot use directly the type they define as it uses
// the `Errors []error` so it cannot be marshaled back to
// the type it was before
type APIError struct {
	Code     string     `json:"code"`
	Message  string     `json:"message"`
	Errors   []APIError `json:"errors,omitempty"`
	Location Location   `json:"location,omitempty"`
	Details  []string   `json:"details,omitempty"`
}

// Location records a position in source code
type Location struct {
	Text   []byte `json:"-"`    // The original text fragment from the source.
	File   string `json:"file"` // The name of the source file (which may be empty).
	Row    int    `json:"row"`  // The line in the source.
	Col    int    `json:"col"`  // The column in the row.
	Offset int    `json:"-"`    // The byte offset for the location in the source.
}

// Error transforms the error into a string
func (e *APIError) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

// do executes the query with the parameters and returns an errors or Decodes the content to the response
func (c *Client) do(ctx context.Context, method, path string, body []byte, response interface{}) error {
	req, err := c.request(ctx, method, path, body)
	if err != nil {
		return err
	}

	res, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	// If the status is not 2XX
	if res.StatusCode < 200 || res.StatusCode > 300 {
		var apiErr APIError
		err = decoder.Decode(&apiErr)
		if err != nil {
			return err
		}

		return &apiErr
	}

	if res.StatusCode == http.StatusNoContent {
		return nil
	}
	// If the status is 2XX
	err = decoder.Decode(response)
	if err != nil {
		return err
	}

	return nil
}

// request builds a new request
func (c *Client) request(ctx context.Context, method, path string, body []byte) (*http.Request, error) {
	buff := bytes.NewBuffer(body)
	req, err := http.NewRequestWithContext(ctx, method, c.buildURL(path), buff)
	if err != nil {
		return nil, err
	}

	if c.token != "" {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.token))
	}

	return req, nil
}

// buildURL build a URL with the given path p
func (c *Client) buildURL(p string) string {
	u := *c.url
	u.Path = path.Join(c.url.Path, p)
	return u.String()

}
