package cpu

import (
	"github.com/v4t/gomb/pkg/memory"
	"github.com/v4t/gomb/pkg/utils"
)

// Interrupt represents individual interrupt flag.
type Interrupt = int

// Available interrupt flags.
const (
	VBlank  Interrupt = 0
	LCDStat Interrupt = 1
	Timer   Interrupt = 2
	Serial  Interrupt = 3
	Joypad  Interrupt = 4
)

var irqAddresses = []uint16{
	VBlank:  0x40,
	LCDStat: 0x48,
	Timer:   0x50,
	Serial:  0x58,
	Joypad:  0x60,
}

// EnableInterrupts enables cpu interrupts.
func EnableInterrupts(mmu *memory.MMU) {
	mmu.Write(0xffff, 255)
}

// DisableInterrupts disables cpu interrupts.
func DisableInterrupts(mmu *memory.MMU) {
	mmu.Write(0xffff, 0)
}

// SetInterrupt sets given interrupt flag.
func SetInterrupt(interrupt Interrupt, mmu *memory.MMU) {
	register := mmu.Read(0xff0f)
	register = utils.SetBit(register, interrupt)
	mmu.Write(0xff0f, register)
}

// HandleInterrupts processes raised interrupts accordingly.
func HandleInterrupts(mmu *memory.MMU) {
	flags := mmu.Read(0xff0f)
	if flags == 0 {
		return
	}
	enabled := mmu.Read(0xffff)
	for f := 0; f < 5; f++ {
		if utils.TestBit(enabled, f) && utils.TestBit(flags, f) {
			// resolveInterrupt(f)
		}
	}
}

func resolveInterrupt(interrupt int, cpu *CPU) {
	// Skip processing interrupts if cpu is halted without interrupts.
	interruptsOn := cpu.MMU.Read(0xffff) != 0
	if !interruptsOn && cpu.Halted {
		cpu.Halted = false
		return
	}

	DisableInterrupts(cpu.MMU)
	cpu.Halted = false

	// Reset interrupt
	flags := cpu.MMU.Read(0xff0f)
	flags = utils.ResetBit(flags, interrupt)
	cpu.MMU.Write(0xff0f, flags)

	// Push pc to stack and jump to interrupt address
	cpu.pushPC()
	cpu.PC = irqAddresses[interrupt]
}
