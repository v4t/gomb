package hardware

import (
	"log"
)

// Instructions represents basic cpu instructions
var Instructions = [0x100]func(*CPU){
	nop, nop, nop, incBC,
	incB, nop, nop, nop,
	nop, nop, nop, nop,
	incC, nop, nop, nop,

	nop, nop, nop, incDE,
	incD, nop, nop, nop,
	nop, nop, nop, nop,
	incE, nop, nop, nop,

	nop, nop, nop, incHL,
	incH, nop, nop, nop,
	nop, nop, nop, nop,
	incL, nop, nop, nop,

	nop, nop, nop, incSP,
	nop, nop, nop, nop,
	nop, nop, nop, nop,
	nop, nop, nop, nop,

	nop, nop, nop, nop,
	nop, nop, nop, nop,
	nop, nop, nop, nop,
	nop, nop, nop, nop,

	nop, nop, nop, nop,
	nop, nop, nop, nop,
	nop, nop, nop, nop,
	nop, nop, nop, nop,

	nop, nop, nop, nop,
	nop, nop, nop, nop,
	nop, nop, nop, nop,
	nop, nop, nop, nop,

	nop, nop, nop, nop,
	nop, nop, nop, nop,
	nop, nop, nop, nop,
	nop, nop, nop, nop,

	nop, nop, nop, nop,
	nop, nop, nop, nop,
	nop, nop, nop, nop,
	nop, nop, nop, nop,

	nop, nop, nop, nop,
	nop, nop, nop, nop,
	nop, nop, nop, nop,
	nop, nop, nop, nop,

	nop, nop, nop, nop,
	nop, nop, nop, nop,
	nop, nop, nop, nop,
	nop, nop, nop, nop,

	nop, nop, nop, nop,
	nop, nop, nop, nop,
	nop, nop, nop, nop,
	nop, nop, nop, nop,

	nop, nop, nop, nop,
	nop, nop, nop, nop,
	nop, nop, nop, nop,
	nop, nop, nop, nop,

	nop, nop, nop, nop,
	nop, nop, nop, nop,
	nop, nop, nop, nop,
	nop, nop, nop, nop,

	nop, nop, nop, nop,
	nop, nop, nop, nop,
	nop, nop, nop, nop,
	nop, nop, nop, nop,

	nop, nop, nop, nop,
	nop, nop, nop, nop,
	nop, nop, nop, nop,
	nop, nop, nop, nop,
}

func nop(cpu *CPU) {
	log.Printf("NOP")
}

func incA(cpu *CPU) { incN(cpu, &cpu.Registers.A) } // INC A
func incB(cpu *CPU) { incN(cpu, &cpu.Registers.B) } // INC B
func incC(cpu *CPU) { incN(cpu, &cpu.Registers.C) } // INC C
func incD(cpu *CPU) { incN(cpu, &cpu.Registers.D) } // INC D
func incE(cpu *CPU) { incN(cpu, &cpu.Registers.E) } // INC E
func incH(cpu *CPU) { incN(cpu, &cpu.Registers.H) } // INC H
func incL(cpu *CPU) { incN(cpu, &cpu.Registers.L) } // INC L

func incBC(cpu *CPU) { incNN(cpu, &cpu.Registers.B, &cpu.Registers.A) } // INC BC
func incDE(cpu *CPU) { incNN(cpu, &cpu.Registers.D, &cpu.Registers.E) } // INC DE
func incHL(cpu *CPU) { incNN(cpu, &cpu.Registers.H, &cpu.Registers.L) } // INC HL
func incSP(cpu *CPU) { incNN(cpu, &cpu.Registers.B, &cpu.Registers.A) } // INC SP
// INC (HL)

func decA(cpu *CPU) { decN(cpu, &cpu.Registers.A) } // DEC A
func decB(cpu *CPU) { decN(cpu, &cpu.Registers.B) } // DEC B
func decC(cpu *CPU) { decN(cpu, &cpu.Registers.C) } // DEC C
func decD(cpu *CPU) { decN(cpu, &cpu.Registers.D) } // DEC D
func decE(cpu *CPU) { decN(cpu, &cpu.Registers.E) } // DEC E
func decH(cpu *CPU) { decN(cpu, &cpu.Registers.H) } // DEC H
func decL(cpu *CPU) { decN(cpu, &cpu.Registers.L) } // DEC L

func decBC(cpu *CPU) { decNN(cpu, &cpu.Registers.B, &cpu.Registers.A) } // DEC BC
func decDE(cpu *CPU) { decNN(cpu, &cpu.Registers.D, &cpu.Registers.E) } // DEC DE
func decHL(cpu *CPU) { decNN(cpu, &cpu.Registers.H, &cpu.Registers.L) } // DEC HL
func decSP(cpu *CPU) { decNN(cpu, &cpu.Registers.B, &cpu.Registers.A) } // DEC SP

// INC n -- Increment register n
func incN(cpu *CPU, n *byte) {
	log.Println("incN")
	*n++
}

// INC nn -- Increment register nn
func incNN(cpu *CPU, n1, n2 *byte) {
	log.Println("incNN")
}

func decN(cpu *CPU, n *byte) {
	log.Println("decN")
	*n--
}

func decNN(cpu *CPU, n1, n2 *byte) {
	log.Println("decNN")
}
