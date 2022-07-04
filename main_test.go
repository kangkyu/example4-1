package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindAll_WithIAMRole(t *testing.T) {
	response, err := findAll()
	assert.IsType(t, nil, err)
	assert.NotNil(t, response.Body)
}
