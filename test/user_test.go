//go:build e2e
// +build e2e

package test

import (
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

func TestPostUser(t *testing.T) {
	client := resty.New()
	res, err := client.R().
	SetBody(
		`{
			"name": "test", 
			"email": "testmail"
			"password": "test1234",
			"role": "admin",
			"corporation_id": "test"
			`).
	Post("http://localhost:8080/api/v1/users")
	assert.NoError(t, err)
	assert.Equal(t, 400, res.StatusCode())
}

