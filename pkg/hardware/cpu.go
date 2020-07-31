package hardware

import (
	"fmt"
	"math"
)

var memory = make([]byte, math.MaxUint16)

// CPURegisters is collection of basic registers for cpu
type CPURegisters struct {
	A byte // accumulator
	B byte
	C byte
	D byte
	E byte
	H byte
	L byte
	F byte // flags
}

// CPUClock represents cpu timer
type CPUClock struct {
	machine int
	cpu     int
}

// CPU contains CPU state
type CPU struct {
	Registers CPURegisters
	Clock     CPUClock
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
