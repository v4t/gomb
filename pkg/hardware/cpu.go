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

// MemRead read byte from memory
func MemRead(address uint16) byte {
	return memory[address]
}

// MemWrite write byte to memory
func MemWrite(address uint16, value byte) {
	memory[address] = value
}
