package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// test ensures that we can generate a simple number
func TestSimpleRandomName(t *testing.T) {
	result := randomValue("foobar")
	assert.Contains(t, result, "foobar-")
	assert.Len(t, result, 14) // foobar = 6 + '-' = 1 + seven digits
}
