package day8

import (
	"errors"
	"fmt"
	"github.com/lluiscampos/advent-of-code-2020/util"
	"strconv"
	"strings"
)

type bootloaderInstruction struct {
	opcode string
	arg    int
}

func (b *bootloaderInstruction) Execute(pc, acc int) (int, int, error) {
	switch b.opcode {
	case "nop":
		return pc + 1, acc, nil
	case "jmp":
		return pc + b.arg, acc, nil
	case "acc":
		return pc + 1, acc + b.arg, nil
	default:
		err := fmt.Errorf("unknown operation %q", b.opcode)
		return -1, -1, err
	}

}

func newBootloaderInstruction(line string) (*bootloaderInstruction, error) {
	bits := strings.Split(line, " ")
	if len(bits) != 2 {
		err := fmt.Errorf("Cannot parse %q", line)
		return nil, err
	}
	opcode := bits[0]
	arg, err := strconv.Atoi(bits[1])
	if err != nil {
		return nil, err
	}
	return &bootloaderInstruction{
		opcode,
		arg,
	}, nil
}

type bootloaderProgram []*bootloaderInstruction

var bootloader bootloaderProgram

func (b *bootloaderProgram) RunFirstLoop() (int, bool, error) {
	bootloaderExec := make(map[*bootloaderInstruction]bool)
	pc := 0
	acc := 0
	var err error

	for {
		if pc == len(bootloader) {
			return acc, true, nil
		}
		_, ok := bootloaderExec[bootloader[pc]]
		if ok {
			return acc, false, nil
		}

		bootloaderExec[bootloader[pc]] = true
		pc, acc, err = bootloader[pc].Execute(pc, acc)
		if err != nil {
			return -1, false, err
		}
	}
	return -1, false, errors.New("Impossible error...")

}

func loadBootloader(filename string) error {
	lines, err := util.ParseFileStrings(filename)
	if err != nil {
		return err
	}

	for _, line := range lines {
		ins, err := newBootloaderInstruction(line)
		if err != nil {
			return err
		}
		bootloader = append(bootloader, ins)
	}
	return nil
}

func SolvePart1() {
	err := loadBootloader("day8.input")
	if err != nil {
		fmt.Println(err)
		return
	}

	acc, _, err := bootloader.RunFirstLoop()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(acc)
}

type bootloaderProgramAlteration struct {
	prevInsIndex int
}

func (b *bootloaderInstruction) Alter() (bool, error) {
	switch b.opcode {
	case "nop":
		b.opcode = "jmp"
		return true, nil
	case "jmp":
		b.opcode = "nop"
		return true, nil
	case "acc":
		return false, nil
	default:
		err := fmt.Errorf("cannot alter %q", b.opcode)
		return false, err
	}

}

func (b *bootloaderProgramAlteration) Init() {
	b.prevInsIndex = -1
}

// func PrintBootloader() {
// 	for _, ins := range bootloader {
// 		fmt.Println(ins)
// 	}
// }

func (b *bootloaderProgramAlteration) NextAlteration() error {
	if b.prevInsIndex != -1 {
		bootloader[b.prevInsIndex].Alter()
	}
	for i, ins := range bootloader[b.prevInsIndex+1:] {
		ok, err := ins.Alter()
		if err != nil {
			return err
		}
		if ok {
			b.prevInsIndex = b.prevInsIndex + 1 + i
			return nil
		}
	}
	return errors.New("No more alterations")
}

func findBootloaderFix() (int, error) {
	var alteration bootloaderProgramAlteration
	alteration.Init()
	for {
		err := alteration.NextAlteration()
		if err != nil {
			return -1, err
		}

		acc, exit, err := bootloader.RunFirstLoop()
		if err != nil {
			return -1, err
		}
		if exit {
			return acc, nil
		}
	}
}

func SolvePart2() {
	acc, err := findBootloaderFix()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(acc)
}
