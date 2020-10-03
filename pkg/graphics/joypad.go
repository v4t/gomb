package graphics

import (
	"fmt"

	"github.com/v4t/gomb/pkg/cpu"
	"github.com/v4t/gomb/pkg/memory"
	"github.com/v4t/gomb/pkg/utils"
)

// JoypadButton represents a single Joypad button.
type JoypadButton int

// Available Joypad buttons.
const (
	ButtonA      JoypadButton = 0
	ButtonB      JoypadButton = 1
	ButtonSelect JoypadButton = 2
	ButtonStart  JoypadButton = 3
	ButtonRight  JoypadButton = 4
	ButtonLeft   JoypadButton = 5
	ButtonUp     JoypadButton = 6
	ButtonDown   JoypadButton = 7
)

/*
 A 0 4
 B 1 4
 Select 2 4
 Start 3 4

 right 0 5
 left 1 5
 up 2 5
 down 3 5
*/

var btnRegMap = []int{
	ButtonA:      0,
	ButtonB:      1,
	ButtonSelect: 2,
	ButtonStart:  3,
	ButtonRight:  0,
	ButtonLeft:   1,
	ButtonUp:     2,
	ButtonDown:   3,
}

// Joypad represents gameboy Joypad.
type Joypad struct {
	mmu *memory.MMU
}

// NewJoypad is constructor for Joypad.
func NewJoypad(mmu *memory.MMU) *Joypad {
	return &Joypad{
		mmu: mmu,
	}
}

// KeyPress event handler.
func (joypad *Joypad) KeyPress(key JoypadButton) {
	ioRegister := joypad.mmu.Memory[0xff00]
	interruptNeeded := false
	if p15ButtonKeysSelected(ioRegister) && key <= 3 {
		interruptNeeded = utils.TestBit(ioRegister, btnRegMap[key])
		ioRegister = utils.ResetBit(ioRegister, btnRegMap[key])
		fmt.Println("p15", ioRegister)
	}
	if p14DirectionKeysSelected(ioRegister) && key > 3 {
		interruptNeeded = utils.TestBit(ioRegister, btnRegMap[key])
		ioRegister = utils.ResetBit(ioRegister, btnRegMap[key])
		fmt.Println("p14", ioRegister)
	}
	joypad.mmu.Memory[0xff00] = ioRegister
	if interruptNeeded {
		fmt.Println("irpt")
		cpu.SetPPUInterrupt(cpu.Joypad, joypad.mmu)
	}
}

// KeyRelease event handler.
func (joypad *Joypad) KeyRelease(key JoypadButton) {
	ioRegister := joypad.mmu.Memory[0xff00]
	if p15ButtonKeysSelected(ioRegister) && key <= 3 {
		joypad.mmu.Memory[0xff00] = utils.SetBit(ioRegister, btnRegMap[key])
	}
	if p14DirectionKeysSelected(ioRegister) && key > 3 {
		joypad.mmu.Memory[0xff00] = utils.SetBit(ioRegister, btnRegMap[key])
	}
}

func p15ButtonKeysSelected(value byte) bool {
	return utils.TestBit(value, 5)
}

func p14DirectionKeysSelected(value byte) bool {
	return utils.TestBit(value, 4)
}


