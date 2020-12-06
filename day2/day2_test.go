package day2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPolicyPassword(t *testing.T) {
	line := "1-3 a: abcde"
	pp, err := NewPolicyPassword(line)
	assert.NoError(t, err)
	assert.Equal(t, pp.p.minInst, 1)
	assert.Equal(t, pp.p.maxInst, 3)
	assert.Equal(t, pp.p.letter, 'a')
	assert.Equal(t, pp.pass, "abcde")

	assert.True(t, pp.Valid())

	lineInvalid := "1-3 b: cdefg"
	pp, err = NewPolicyPassword(lineInvalid)
	assert.NoError(t, err)
	assert.False(t, pp.Valid())
}

func TestPolicyPasswordAdvanced(t *testing.T) {
	line := "1-3 a: abcde"
	pp, err := NewPolicyPassword(line)
	assert.NoError(t, err)
	assert.Equal(t, pp.p.minInst, 1)
	assert.Equal(t, pp.p.maxInst, 3)
	assert.Equal(t, pp.p.letter, 'a')
	assert.Equal(t, pp.pass, "abcde")

	assert.True(t, pp.ValidAdvanced())

	lineInvalid := "2-9 c: ccccccccc"
	pp, err = NewPolicyPassword(lineInvalid)
	assert.NoError(t, err)
	assert.False(t, pp.ValidAdvanced())
}
