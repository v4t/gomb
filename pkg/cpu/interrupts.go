package cpu

import (
	"github.com/v4t/gomb/pkg/memory"
	"github.com/v4t/gomb/pkg/utils"
)

// InterruptFlag represents individual interrupt flag.
type InterruptFlag = int

// Available interrupt flags.
const (
	VBlank  InterruptFlag = 0
	LCDStat InterruptFlag = 1
	Timer   InterruptFlag = 2
	Serial  InterruptFlag = 3
	Joypad  InterruptFlag = 4
)

var irqAddresses = []uint16{
	VBlank:  0x40,
	LCDStat: 0x48,
	Timer:   0x50,
	Serial:  0x58,
	Joypad:  0x60,
}

// Interrupt contains interrupt register states and handles interrupts.
type Interrupt struct {
	// Interrupt master enable flag
	IMEFlag bool

	// Interrupt Flag (IF) register
	InterruptFlag byte

	// Interupt Enable (IE) register
	InterruptEnable byte
}

// Enable interrupts.
func (interrupt *Interrupt) Enable() {
	interrupt.IMEFlag = true
}

// Disable interrupts.
func (interrupt *Interrupt) Disable() {
	interrupt.IMEFlag = false
}

// Read IF or IE register value.
func (interrupt *Interrupt) Read(address uint16) byte {
	return 0
}

// Write to IF or IE register.
func (interrupt *Interrupt) Write(address uint16, value byte) {

}

// SetInterrupt sets given interrupt flag.
func (interrupt *Interrupt) SetInterrupt(flag InterruptFlag) {
	// register := cpu.MMU.Read(0xff0f)
	// register = utils.SetBit(register, flag)
	// cpu.MMU.Write(0xff0f, register)
}

// HandleInterrupts processes raised interrupts accordingly.
func (interrupt *Interrupt) HandleInterrupts(cpu *CPU) {
	if !IMEFlag {
		return
	}
	flags := cpu.MMU.Read(0xff0f)
	if flags == 0 {
		return
	}
	enabled := cpu.MMU.Read(0xffff)
	for f := 0; f < 5; f++ {
		if utils.TestBit(enabled, f) && utils.TestBit(flags, f) {
			resolveInterrupt(f, cpu)
		}
	}
}

func (interrupt *Interrupt) resolveInterrupt(flag int, cpu *CPU) {
	// Skip processing interrupts if cpu is halted without interrupts.
	if !IMEFlag && cpu.Halted {
		cpu.Halted = false
		return
	}

	DisableInterrupts(cpu)
	cpu.Halted = false

	// Reset interrupt
	flags := cpu.MMU.Read(0xff0f)
	flags = utils.ResetBit(flags, flag)
	cpu.MMU.Write(0xff0f, flags)

	// Push pc to stack and jump to interrupt address
	cpu.pushPC()
	cpu.PC = irqAddresses[flag]
}







// OLD


// IMEFlag is Interrupt master enable flag.
var IMEFlag bool

// EnableInterrupts enables cpu interrupts.
func EnableInterrupts(cpu *CPU) {
	IMEFlag = true
}

// DisableInterrupts disables cpu interrupts.
func DisableInterrupts(cpu *CPU) {
	IMEFlag = false
}

// SetInterrupt sets given interrupt flag.
func SetInterrupt(flag InterruptFlag, cpu *CPU) {
	register := cpu.MMU.Read(0xff0f)
	register = utils.SetBit(register, flag)
	cpu.MMU.Write(0xff0f, register)
}

func SetPPUInterrupt(flag InterruptFlag, mmu *memory.MMU) {
	register := mmu.Read(0xff0f)
	register = utils.SetBit(register, flag)
	mmu.Write(0xff0f, register)
}

// HandleInterrupts processes raised interrupts accordingly.
func HandleInterrupts(cpu *CPU) {
	if !IMEFlag {
		return
	}
	flags := cpu.MMU.Read(0xff0f)
	if flags == 0 {
		return
	}
	enabled := cpu.MMU.Read(0xffff)
	for f := 0; f < 5; f++ {
		if utils.TestBit(enabled, f) && utils.TestBit(flags, f) {
			resolveInterrupt(f, cpu)
		}
	}
}

func resolveInterrupt(interrupt int, cpu *CPU) {
	// Skip processing interrupts if cpu is halted without interrupts.
	if !IMEFlag && cpu.Halted {
		cpu.Halted = false
		return
	}

	DisableInterrupts(cpu)
	cpu.Halted = false

	// Reset interrupt
	flags := cpu.MMU.Read(0xff0f)
	flags = utils.ResetBit(flags, interrupt)
	cpu.MMU.Write(0xff0f, flags)

	// Push pc to stack and jump to interrupt address
	cpu.pushPC()
	cpu.PC = irqAddresses[interrupt]
}
