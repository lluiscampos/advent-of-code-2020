package day5

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMakeRange(t *testing.T) {
	myRange := makeRange(0, 127)
	assert.Equal(t, 128, len(myRange))
	assert.Equal(t, 0, myRange[0])
	assert.Equal(t, 127, myRange[127])
}

func TestDecodeRow(t *testing.T) {
	row, err := decodeRow("FBFBBFF")
	assert.NoError(t, err)
	assert.Equal(t, 44, row)
}

func TestDecodeColumn(t *testing.T) {
	column, err := decodeColumn("RLR")
	assert.NoError(t, err)
	assert.Equal(t, 5, column)
}
