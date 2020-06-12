package problem159

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, 'b', Solution("acbbac"))
	assert.Equal(t, rune(-1), Solution("abcdef"))
	assert.Equal(t, rune(-1), Solution(""))
}
