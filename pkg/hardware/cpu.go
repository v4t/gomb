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

// Run executes next CPU cycle
func (cpu *CPU) Run() {
	op := cpu.nextInstruction()
	fmt.Println(op)
}

func (cpu *CPU) nextInstruction() byte {
	op := memory[cpu.PC]
	cpu.PC++
	return op
}
