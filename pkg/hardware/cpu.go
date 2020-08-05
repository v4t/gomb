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

// CPU contains CPU state
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

// Fetch next byte from memory
func (cpu *CPU) Fetch() byte {
	op := memory[cpu.PC]
	cpu.PC++
	return op
}
