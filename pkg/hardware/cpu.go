package hardware

import (
	"fmt"
	"math"
)

var memory = make([]byte, math.MaxUint16)

// Clock represents cpu timer
type Clock struct {
	machine int
	cpu     int
}

// Flag is a type for CPU flags
type Flag byte

const (
	bitflagC = 0x10
	bitflagH = 0x20
	bitflagN = 0x40
	bitflagZ = 0x80
)

// CPU represents CPU state
type CPU struct {
	Registers Registers
	Clock     Clock
	PC        uint16
	SP        uint16
}

// Execute next CPU cycle
func (cpu *CPU) Execute() {
	op := cpu.Fetch()
	fmt.Println(op)
}

// Fetch retrieve next byte from memory
func (cpu *CPU) Fetch() byte {
	op := memory[cpu.PC]
	cpu.PC++
	return op
}

// Fetch16 retrieve next 16-bit word from memory
func (cpu *CPU) Fetch16() uint16 {
	i := uint16(cpu.Fetch())
	j := uint16(cpu.Fetch())
	return j<<8 | i
}

// Carry retrieves carry flag
func (cpu *CPU) Carry() bool {
	return cpu.Registers.F&bitflagC != 0
}

// SetCarry sets carry flag
func (cpu *CPU) SetCarry(value bool) {
	if value {
		cpu.Registers.F |= bitflagC
	} else {
		cpu.Registers.F &^= bitflagC
	}
}

// HalfCarry retrieves half carry flag
func (cpu *CPU) HalfCarry() bool {
	return cpu.Registers.F&bitflagH != 0
}

// SetHalfCarry sets half carry flag
func (cpu *CPU) SetHalfCarry(value bool) {
	if value {
		cpu.Registers.F |= bitflagH
	} else {
		cpu.Registers.F &^= bitflagH
	}
}

// Negative retrieves negative/subtract flag
func (cpu *CPU) Negative() bool {
	return cpu.Registers.F&bitflagN != 0
}

// SetNegative sets negative/subtract flag
func (cpu *CPU) SetNegative(value bool) {
	if value {
		cpu.Registers.F |= bitflagN
	} else {
		cpu.Registers.F &^= bitflagN
	}
}

// Zero retrieves zero flag
func (cpu *CPU) Zero() bool {
	return cpu.Registers.F&bitflagZ != 0
}

// SetZero sets zero flag
func (cpu *CPU) SetZero(value bool) {
	if value {
		cpu.Registers.F |= bitflagZ
	} else {
		cpu.Registers.F &^= bitflagZ
	}
}

// MemRead read byte from memory
func MemRead(address uint16) byte {
	return memory[address]
}

// MemWrite write byte to memory
func MemWrite(address uint16, value byte) {
	memory[address] = value
}
