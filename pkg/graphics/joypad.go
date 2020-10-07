package graphics

import (
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

// JoypadState represents currently pressed buttons.
var JoypadState byte = 0xcf

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
	state                 byte
	buttonKeysSelected    bool
	directionKeysSelected bool
}

// NewJoypad is constructor for Joypad.
func NewJoypad() *Joypad {
	return &Joypad{
		state:                 0xff,
		buttonKeysSelected:    false,
		directionKeysSelected: false,
	}
}

// KeyPress event handler.
func (joypad *Joypad) KeyPress(key JoypadButton) {
	joypad.state = utils.ResetBit(joypad.state, int(key))

	// if p15ButtonKeysSelected(ioRegister) && key <= 3 {
	// 	interruptNeeded = utils.TestBit(ioRegister, btnRegMap[key])
	// 	ioRegister = utils.ResetBit(ioRegister, btnRegMap[key])
	// }
	// if p14DirectionKeysSelected(ioRegister) && key > 3 {
	// 	interruptNeeded = utils.TestBit(ioRegister, btnRegMap[key])
	// 	ioRegister = utils.ResetBit(ioRegister, btnRegMap[key])
	// }
	// joypad.mmu.Memory[0xff00] = ioRegister
	// if interruptNeeded {
	// 	cpu.SetPPUInterrupt(cpu.Joypad, joypad.mmu)
	// }
	// fmt.Println("after press")
	// fmt.Println(
	// 	utils.TestBit(ioRegister, 5), utils.TestBit(ioRegister, 4),
	// 	utils.TestBit(ioRegister, 4), utils.TestBit(ioRegister, 2),
	// 	utils.TestBit(ioRegister, 1), utils.TestBit(ioRegister, 0))
}

// KeyRelease event handler.
func (joypad *Joypad) KeyRelease(key JoypadButton) {
	joypad.state = utils.SetBit(joypad.state, int(key))
	// ioRegister := joypad.mmu.Memory[0xff00]
	// if p15ButtonKeysSelected(ioRegister) && key <= 3 {
	// 	joypad.mmu.Memory[0xff00] = utils.SetBit(ioRegister, btnRegMap[key])
	// }
	// if p14DirectionKeysSelected(ioRegister) && key > 3 {
	// 	joypad.mmu.Memory[0xff00] = utils.SetBit(ioRegister, btnRegMap[key])
	// }
	// ioRegister = joypad.mmu.Memory[0xff00]
	// fmt.Println("after release")
	// fmt.Println(
	// 	utils.TestBit(ioRegister, 5), utils.TestBit(ioRegister, 4),
	// 	utils.TestBit(ioRegister, 4), utils.TestBit(ioRegister, 2),
	// 	utils.TestBit(ioRegister, 1), utils.TestBit(ioRegister, 0))
}

func (joypad *Joypad) Read(address uint16) byte {
	if address != 0xff00 {
		panic("Attempted to access joypad register with wrong memory address.")
	}
	registerValue := byte(0xff)
	if joypad.directionKeysSelected {
		registerValue = utils.ResetBit(registerValue, 4)
		registerValue &= ((joypad.state >> 4) | 0xf0)
	} else if joypad.buttonKeysSelected {
		registerValue = utils.ResetBit(registerValue, 5)
		registerValue &= (joypad.state | 0xf0)
	}
	return registerValue
}

func (joypad *Joypad) Write(address uint16, value byte) {
	// Set P14 and P15 selected state (0 = Select)
	// fmt.Printf("%x\n", value)
	joypad.directionKeysSelected = !utils.TestBit(value, 4)
	joypad.buttonKeysSelected = !utils.TestBit(value, 5)
}

func p15ButtonKeysSelected(value byte) bool {
	return !utils.TestBit(value, 5)
}

func p14DirectionKeysSelected(value byte) bool {
	return !utils.TestBit(value, 4)
}

// func (joypad *Joypad) getCurrentState() byte {
// 	//    BYTE res = m_Rom[0xFF00]
// 	res := joypad.mmu.Memory[0xff00]
// 	// flip all the bits
// 	res ^= 0xFF

// 	// are we interested in the standard buttons?
// 	if !utils.TestBit(res, 4) {
// 		// topJoypad := joypadState >> 4
// 		// topJoypad |= 0xF0 // turn the top 4 bits on
// 		// res &= topJoypad  // show what buttons are pressed
// 	} else if !utils.TestBit(res, 5) {
// 		//directional buttons
// 		// bottomJoypad := joypadState & 0xF
// 		// bottomJoypad |= 0xF0
// 		// res &= bottomJoypad
// 	}
// 	return res
// }
