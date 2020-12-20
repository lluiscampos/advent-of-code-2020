package day9

import (
	"github.com/lluiscampos/advent-of-code-2020/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestXmasDecode(t *testing.T) {
	nums, err := util.ParseFileInts("day9.example")
	assert.NoError(t, err)

	seq := XmasSequence{
		5,
		nums,
	}

	_, n, err := seq.FirstInvalid()
	assert.NoError(t, err)
	assert.Equal(t, 127, n)
}

func TestXmasWeakness(t *testing.T) {
	nums, err := util.ParseFileInts("day9.example")
	assert.NoError(t, err)

	seq := XmasSequence{
		5,
		nums,
	}

	n, err := seq.FindWeakness()
	assert.NoError(t, err)
	assert.Equal(t, 62, n)
}
