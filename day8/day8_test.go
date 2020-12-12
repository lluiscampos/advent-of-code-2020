package day8

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseBootloaderInstruction(t *testing.T) {
	line := "acc +1"
	ins, err := newBootloaderInstruction(line)
	assert.NoError(t, err)
	assert.Equal(t, "acc", ins.opcode)
	assert.Equal(t, 1, ins.arg)

	line = "jmp -3"
	ins, err = newBootloaderInstruction(line)
	assert.NoError(t, err)
	assert.Equal(t, "jmp", ins.opcode)
	assert.Equal(t, -3, ins.arg)
}

func TestRunFirstLoop(t *testing.T) {
	err := loadBootloader("day8.example")
	assert.NoError(t, err)
	assert.Equal(t, 9, len(bootloader))

	acc, exit, err := bootloader.RunFirstLoop()
	assert.NoError(t, err)
	assert.Equal(t, 5, acc)
	assert.False(t, exit)
}

func TestFindBootloaderFix(t *testing.T) {
	assert.Equal(t, 9, len(bootloader))

	acc, err := findBootloaderFix()
	assert.NoError(t, err)
	assert.Equal(t, 8, acc)
}
