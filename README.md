# Gopa 

[![PkgGoDev](https://pkg.go.dev/badge/github.com/cycloidio/gopa)](https://pkg.go.dev/github.com/cycloidio/gopa)

Gopa is a client written in Go for [OPA](https://openpolicyagent.org/).

## Usage

This example can be found on the [example_test.go](https://github.com/cycloidio/gopa/blob/master/example_test.go)

```go
package main

import (
	"context"
	"fmt"

	"github.com/cycloidio/gopa"
)

func main() {
	c, err := gopa.NewClient()
	if err != nil {
		// Handle error
	}

	ctx := context.Background()

	policyID := "my-policy-id"
	policyBody := []byte(`
package opa.examples

import input.example.flag

default allow_request = false
allow_request { flag == true }
`)

	// First we create a Policy to be used
	_, err = c.PolicyCreateOrUpdate(ctx, policyID, policyBody)
	if err != nil {
		// Handle error
	}

	input := map[string]interface{}{
		"example": map[string]interface{}{
			"flag": true,
		},
	}

	res, err := c.DataGetWithInput(ctx, "/opa/examples/allow_request", input)
	if err != nil {
		// Handle error
	}
	fmt.Println(*res.Result)

	res, err = c.DataGet(ctx, "/opa/examples/allow_request")
	if err != nil {
		// Handle error
	}
	fmt.Println(*res.Result)

	_, err = c.PolicyDelete(ctx, policyID)

	// Output:
	// true
	// false
}
```

## Implementation

The current implementation supports:

* [x] Policy API
* [x] Data API
* [ ] Query API (WIP)
* [ ] Compile API
