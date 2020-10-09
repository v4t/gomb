package processor

import (
	"github.com/v4t/gomb/pkg/utils"
)

// InterruptFlag represents individual interrupt flag.
type InterruptFlag = int

// Available interrupt flags.
const (
	VBlankInterrupt    InterruptFlag = 0
	LCDStatusInterrupt InterruptFlag = 1
	TimerInterrupt     InterruptFlag = 2
	SerialInterrupt    InterruptFlag = 3
	JoypadInterrupt    InterruptFlag = 4
)

var irqAddresses = []uint16{
	VBlankInterrupt:    0x40,
	LCDStatusInterrupt: 0x48,
	TimerInterrupt:     0x50,
	SerialInterrupt:    0x58,
	JoypadInterrupt:    0x60,
}

// Interrupts contains interrupt register states and handles interrupts.
type Interrupts struct {
	// IME - Interrupt master enable flag
	IME bool

	// IF - Interrupt Flag register
	IF byte

	// IE - Interrupt Enable register
	IE byte
}

// NewInterrupts is constructor for interrupt manager.
func NewInterrupts() *Interrupts {
	return &Interrupts{
		IME: false,
		IF:  0,
		IE:  0xe1,
	}
}

// Enable interrupts.
func (interrupts *Interrupts) Enable() {
	interrupts.IME = true
}

// Disable interrupts.
func (interrupts *Interrupts) Disable() {
	interrupts.IME = false
}

// Read IF or IE register value.
func (interrupts *Interrupts) Read(address uint16) byte {
	if address == 0xffff {
		return interrupts.IE
	} else if address == 0xff0f {
		return interrupts.IF
	} else {
		panic("Attempted to read interrupt registers with invalid memory address.")
	}
}

// Write to IF or IE register.
func (interrupts *Interrupts) Write(address uint16, value byte) {
	if address == 0xffff {
		interrupts.IE = value
	} else if address == 0xff0f {
		interrupts.IF = value
	} else {
		panic("Attempted writing to interrupt registers with invalid memory address.")
	}
}

// SetInterrupt sets given interrupt flag.
func (interrupts *Interrupts) SetInterrupt(flag InterruptFlag) {
	if flag == JoypadInterrupt {
	}
	interrupts.IF = utils.SetBit(interrupts.IF, flag)
}

// Resolve raised interrupts accordingly.
func (interrupts *Interrupts) Resolve(cpu *CPU) {
	if !interrupts.IME || interrupts.IF == 0 {
		return
	}
	for f := 0; f < 5; f++ {
		if utils.TestBit(interrupts.IE, f) && utils.TestBit(interrupts.IF, f) {
			interrupts.resolveInterrupt(f, cpu)
		}
	}
}

// resolveInterrupt is a helper function for resolving a specific interrupt.
func (interrupts *Interrupts) resolveInterrupt(flag int, cpu *CPU) {
	// Skip processing interrupts if cpu is halted without interrupts.
	if !interrupts.IME && cpu.Halted {
		cpu.Halted = false
		return
	}

	// Reset interrupt
	interrupts.IF = utils.ResetBit(interrupts.IF, flag)
	interrupts.Disable()
	cpu.Halted = false

	// Push pc to stack and jump to interrupt address
	cpu.pushPC()
	cpu.PC = irqAddresses[flag]
}

// OLD

// // IMEFlag is Interrupt master enable flag.
// var IMEFlag bool

// // EnableInterrupts enables cpu interrupts.
// func EnableInterrupts(cpu *CPU) {
// 	IMEFlag = true
// }

// // DisableInterrupts disables cpu interrupts.
// func DisableInterrupts(cpu *CPU) {
// 	IMEFlag = false
// }

// // SetInterrupt sets given interrupt flag.
// func SetInterrupt(flag InterruptFlag, cpu *CPU) {
// 	register := cpu.MMU.Read(0xff0f)
// 	register = utils.SetBit(register, flag)
// 	cpu.MMU.Write(0xff0f, register)
// }

// func SetPPUInterrupt(flag InterruptFlag, mmu *memory.MMU) {
// 	register := mmu.Read(0xff0f)
// 	register = utils.SetBit(register, flag)
// 	mmu.Write(0xff0f, register)
// }

// // HandleInterrupts processes raised interrupts accordingly.
// func HandleInterrupts(cpu *CPU) {
// 	if !IMEFlag {
// 		return
// 	}
// 	flags := cpu.MMU.Read(0xff0f)
// 	if flags == 0 {
// 		return
// 	}
// 	enabled := cpu.MMU.Read(0xffff)
// 	for f := 0; f < 5; f++ {
// 		if utils.TestBit(enabled, f) && utils.TestBit(flags, f) {
// 			resolveInterrupt(f, cpu)
// 		}
// 	}
// }

// func resolveInterrupt(interrupt int, cpu *CPU) {
// 	// Skip processing interrupts if cpu is halted without interrupts.
// 	if !IMEFlag && cpu.Halted {
// 		cpu.Halted = false
// 		return
// 	}

// 	DisableInterrupts(cpu)
// 	cpu.Halted = false

// 	// Reset interrupt
// 	flags := cpu.MMU.Read(0xff0f)
// 	flags = utils.ResetBit(flags, interrupt)
// 	cpu.MMU.Write(0xff0f, flags)

// 	// Push pc to stack and jump to interrupt address
// 	cpu.pushPC()
// 	cpu.PC = irqAddresses[interrupt]
// }
