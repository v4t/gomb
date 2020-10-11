package processor

import (
	"github.com/v4t/gomb/pkg/memory"
)

// Clock represents cpu timer.
type Clock struct {
	machine int
	cpu     int
}

// Flag is a type for CPU flags.
type Flag byte

const (
	bitflagC = 0x10
	bitflagH = 0x20
	bitflagN = 0x40
	bitflagZ = 0x80
)

// CPU represents CPU state.
type CPU struct {
	Interrupts *Interrupts
	MMU       *memory.MMU
	Registers Registers
	Clock     Clock
	PC        uint16
	SP        uint16
	Halted    bool
	Stopped   bool

	enablingInterrupts  bool
	disablingInterrupts bool
}

// InitializeCPU initializes cpu values
func InitializeCPU() *CPU {
	cpu := &CPU{
		MMU: memory.InitializeMMU(),
		Interrupts: NewInterrupts(),
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
func (cpu *CPU) Execute() int {
	if cpu.Halted || cpu.Stopped {
		return 0
	}
	enableIrq := cpu.enablingInterrupts
	disableIrq := cpu.disablingInterrupts
	op := cpu.Fetch()
	cycles := 0
	if op == 0xcb {
		op = cpu.Fetch()
		cycles = ExecuteCBInstruction(cpu, op)
		// fmt.Println("CB", fmt.Sprintf("PC:%04x OP: %04x", cpu.PC, op))
	} else {
		cycles = ExecuteInstruction(cpu, op)
		// fmt.Println(fmt.Sprintf("PC:%04x OP: %02x", cpu.PC, op))
	}
	if enableIrq {
		cpu.Interrupts.Enable()
		cpu.enablingInterrupts = false
	}
	if disableIrq {
		cpu.Interrupts.Disable()
		cpu.disablingInterrupts = false
	}
	return cycles
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
	return cpu.Registers.F&bitflagC != 0
}

// SetCarry sets carry flag.
func (cpu *CPU) SetCarry(value bool) {
	if value {
		cpu.Registers.F |= bitflagC
	} else {
		cpu.Registers.F &^= bitflagC
	}
}

// HalfCarry retrieves half-carry flag.
func (cpu *CPU) HalfCarry() bool {
	return cpu.Registers.F&bitflagH != 0
}

// SetHalfCarry sets half-carry flag.
func (cpu *CPU) SetHalfCarry(value bool) {
	if value {
		cpu.Registers.F |= bitflagH
	} else {
		cpu.Registers.F &^= bitflagH
	}
}

// Negative retrieves negative/subtract flag.
func (cpu *CPU) Negative() bool {
	return cpu.Registers.F&bitflagN != 0
}

// SetNegative sets negative/subtract flag.
func (cpu *CPU) SetNegative(value bool) {
	if value {
		cpu.Registers.F |= bitflagN
	} else {
		cpu.Registers.F &^= bitflagN
	}
}

// Zero retrieves zero flag.
func (cpu *CPU) Zero() bool {
	return cpu.Registers.F&bitflagZ != 0
}

// SetZero sets zero flag.
func (cpu *CPU) SetZero(value bool) {
	if value {
		cpu.Registers.F |= bitflagZ
	} else {
		cpu.Registers.F &^= bitflagZ
	}
}

func (cpu *CPU) pushPC() {
	cpu.MMU.Write(cpu.SP-1, byte(uint16(cpu.PC&0xff00)>>8))
	cpu.MMU.Write(cpu.SP-2, byte(cpu.PC&0xff))
	cpu.SP -= 2
}
