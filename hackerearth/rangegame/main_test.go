package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFn(t *testing.T) {
	assert.Equal(t, 0, solve(0))
	assert.Equal(t, 0, solve(1))
	assert.Equal(t, 1, solve(2))
	assert.Equal(t, 4, solve(3))
	assert.Equal(t, 6, solve(4))
	assert.Equal(t, 11, solve(5))
	assert.Equal(t, 17, solve(6))
	assert.Equal(t, 24, solve(7))
	assert.Equal(t, 47, solve(10))
	assert.Equal(t, 4986, solve(100))
	assert.Equal(t, 499988, solve(1000))
	assert.Equal(t, 499999975712, solve(1000000))
}
