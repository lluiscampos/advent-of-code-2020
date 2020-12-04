package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseFile(t *testing.T) {
	stuff, err := parseFile("day1.example")
	assert.NoError(t, err)
	assert.Equal(t, len(stuff), 6)
	assert.Equal(t, stuff[0], 1721)
	assert.Equal(t, stuff[5], 1456)
}

func TestFindPairSumEqual(t *testing.T) {
	var dummy = []int{123, 30, 20, 70}
	a, b, err := findPairSumEqual(dummy, 100)
	assert.NoError(t, err)
	assert.Equal(t, a, 30)
	assert.Equal(t, b, 70)

	var fake = []int{50, 0, 0, 0}
	a, b, err = findPairSumEqual(fake, 100)
	assert.Error(t, err)

	example, _ := parseFile("day1.example")
	a, b, err = findPairSumEqual(example, 2020)
	assert.NoError(t, err)
	assert.Equal(t, a, 1721)
	assert.Equal(t, b, 299)
}

func TestFindTrioSumEqual(t *testing.T) {
	var dummy = []int{123, 30, 123, 20, 123, 50}
	a, b, c, err := findTrioSumEqual(dummy, 100)
	assert.NoError(t, err)
	assert.Equal(t, a, 30)
	assert.Equal(t, b, 20)
	assert.Equal(t, c, 50)

	example, _ := parseFile("day1.example")
	a, b, c, err = findTrioSumEqual(example, 2020)
	assert.NoError(t, err)
	assert.Equal(t, a, 979)
	assert.Equal(t, b, 366)
	assert.Equal(t, c, 675)

}
