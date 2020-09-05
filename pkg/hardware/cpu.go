package hardware

import (
	// "fmt"
	"log"
	"math"

	"github.com/v4t/gomb/pkg/memory"
)

// Mem contains emulator memory
var Mem = make([]byte, math.MaxUint16+1)

// Clock represents cpu timer.
type Clock struct {
	machine int
	cpu     int
}

// Flag is a type for CPU flags.
type Flag byte

const (
	// FlagC -- Carry flag
	FlagC = 0x10
	// FlagH -- Half carry flag
	FlagH = 0x20
	// FlagN -- Negative / subtraction flag
	FlagN = 0x40
	// FlagZ -- Zero flag
	FlagZ = 0x80
)

// CPU represents CPU state.
type CPU struct {
	MMU       *memory.MMU
	Registers Registers
	Clock     Clock
	PC        uint16
	SP        uint16
}

// InitializeCPU initializes cpu values
func InitializeCPU(mmu *memory.MMU) *CPU {
	cpu := &CPU{
		MMU: mmu,
		Registers: Registers{
			A: 0x11,
			B: 0x00,
			C: 0x00,
			D: 0xff,
			E: 0x56,
			H: 0x00,
			L: 0x0d,
			F: 0x80,
		},
		Clock: Clock{machine: 0, cpu: 0},
		PC:    0x100,
		SP:    0xfffe,
	}
	return cpu
}

// Execute next CPU cycle.
func (cpu *CPU) Execute() {
	op := cpu.Fetch()
	if op == 0xcb {
		op = cpu.Fetch()
		ExecuteCBInstruction(cpu, op)
		// fmt.Println("CB", fmt.Sprintf("%02x", op))
	} else {
		ExecuteInstruction(cpu, op)
		// fmt.Println(fmt.Sprintf("%02x", op))
	}
}

func (cpu *CPU) printDebug() {

}

// Fetch retrieve next byte from memory.
func (cpu *CPU) Fetch() byte {
	op := cpu.MMU.Read(cpu.PC)
	cpu.PC++
	return op
}

// Fetch16 retrieve next 16-bit word from memory.
func (cpu *CPU) Fetch16() uint16 {
	i := uint16(cpu.Fetch())
	j := uint16(cpu.Fetch())
	return j<<8 | i
}

// Carry retrieves carry flag.
func (cpu *CPU) Carry() bool {
	return cpu.Registers.F&FlagC != 0
}

// SetCarry sets carry flag.
func (cpu *CPU) SetCarry(value bool) {
	if value {
		cpu.Registers.F |= FlagC
	} else {
		cpu.Registers.F &^= FlagC
	}
}

// HalfCarry retrieves half-carry flag.
func (cpu *CPU) HalfCarry() bool {
	return cpu.Registers.F&FlagH != 0
}

// SetHalfCarry sets half-carry flag.
func (cpu *CPU) SetHalfCarry(value bool) {
	if value {
		cpu.Registers.F |= FlagH
	} else {
		cpu.Registers.F &^= FlagH
	}
}

// Negative retrieves negative/subtract flag.
func (cpu *CPU) Negative() bool {
	return cpu.Registers.F&FlagN != 0
}

// SetNegative sets negative/subtract flag.
func (cpu *CPU) SetNegative(value bool) {
	if value {
		cpu.Registers.F |= FlagN
	} else {
		cpu.Registers.F &^= FlagN
	}
}

// Zero retrieves zero flag.
func (cpu *CPU) Zero() bool {
	return cpu.Registers.F&FlagZ != 0
}

// SetZero sets zero flag.
func (cpu *CPU) SetZero(value bool) {
	if value {
		cpu.Registers.F |= FlagZ
	} else {
		cpu.Registers.F &^= FlagZ
	}
}

// EnableInterrupts enables cpu interrupts.
func EnableInterrupts() {
	log.Println("Enable interrupts")
}

// DisableInterrupts disables cpu interrupts.
func DisableInterrupts() {
	log.Println("Enable interrupts")
}
