package day3

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCheckTreesExample(t *testing.T) {
	stuff, err := parseFile("day3.example")
	assert.NoError(t, err)

	trees := checkTrees(stuff, 3, 1)
	assert.Equal(t, 7, trees)
}
