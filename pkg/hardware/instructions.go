package hardware

import (
	"log"
)

// Instructions represents basic cpu instructions
var Instructions = initInstr()

// Init instructions
func initInstr() []func(*CPU) {
	var instr = make([]func(*CPU), 0x100)

	for i := 0; i < len(instr); i++ {
		instr[i] = func(cpu *CPU) {
			log.Fatalf("Op code 0x%x not implemented yet", i)
		}
	}
	instr[0x00] = nop   // NOP
	instr[0x01] = nop   //
	instr[0x02] = nop   //
	instr[0x03] = incBC // INC BC
	instr[0x04] = incB  // INC B
	instr[0x05] = decB  // DEC B
	instr[0x06] = nop
	instr[0x07] = nop
	instr[0x08] = nop
	instr[0x09] = nop
	instr[0x0a] = nop
	instr[0x0b] = decBC // DEC BC
	instr[0x0c] = incC  // INC C
	instr[0x0d] = decC  // DEC C
	instr[0x0e] = nop
	instr[0x0f] = nop

	instr[0x10] = nop
	instr[0x11] = nop
	instr[0x12] = nop
	instr[0x13] = incDE // INC DE
	instr[0x14] = incD  // INC D
	instr[0x15] = decD  // DEC D
	instr[0x16] = nop
	instr[0x17] = nop
	instr[0x18] = nop
	instr[0x19] = nop
	instr[0x1a] = nop
	instr[0x1b] = decDE // DEC DE
	instr[0x1c] = incE  // INC E
	instr[0x1d] = decE  // DEC E
	instr[0x1e] = nop
	instr[0x1f] = nop

	instr[0x20] = nop
	instr[0x21] = nop
	instr[0x22] = nop
	instr[0x23] = incHL // INC HL
	instr[0x24] = incH  // INC H
	instr[0x25] = decH  // DEC H
	instr[0x26] = nop
	instr[0x27] = nop
	instr[0x28] = nop
	instr[0x29] = nop
	instr[0x2a] = nop
	instr[0x2b] = decHL // DEC HL
	instr[0x2c] = incL  // INC L
	instr[0x2d] = decL  // DEC L
	instr[0x2e] = nop
	instr[0x2f] = nop

	instr[0x30] = nop
	instr[0x31] = nop
	instr[0x32] = nop
	instr[0x33] = incSP // INC SP
	instr[0x34] = nop
	instr[0x35] = nop
	instr[0x36] = nop
	instr[0x37] = nop
	instr[0x38] = nop
	instr[0x39] = nop
	instr[0x3a] = nop
	instr[0x3b] = decSP // DEC SP
	instr[0x3c] = incA  // INC A
	instr[0x3d] = decA  // DEC A
	instr[0x3e] = nop
	instr[0x3f] = nop

	instr[0x40] = nop
	instr[0x41] = nop
	instr[0x42] = nop
	instr[0x43] = nop
	instr[0x44] = nop
	instr[0x45] = nop
	instr[0x46] = nop
	instr[0x47] = nop
	instr[0x48] = nop
	instr[0x49] = nop
	instr[0x4a] = nop
	instr[0x4b] = nop
	instr[0x4c] = nop
	instr[0x4d] = nop
	instr[0x4e] = nop
	instr[0x4f] = nop

	instr[0x50] = nop
	instr[0x51] = nop
	instr[0x52] = nop
	instr[0x53] = nop
	instr[0x54] = nop
	instr[0x55] = nop
	instr[0x56] = nop
	instr[0x57] = nop
	instr[0x58] = nop
	instr[0x59] = nop
	instr[0x5a] = nop
	instr[0x5b] = nop
	instr[0x5c] = nop
	instr[0x5d] = nop
	instr[0x5e] = nop
	instr[0x5f] = nop

	instr[0x60] = nop
	instr[0x61] = nop
	instr[0x62] = nop
	instr[0x63] = nop
	instr[0x64] = nop
	instr[0x65] = nop
	instr[0x66] = nop
	instr[0x67] = nop
	instr[0x68] = nop
	instr[0x69] = nop
	instr[0x6a] = nop
	instr[0x6b] = nop
	instr[0x6c] = nop
	instr[0x6d] = nop
	instr[0x6e] = nop
	instr[0x6f] = nop

	instr[0x70] = nop
	instr[0x71] = nop
	instr[0x72] = nop
	instr[0x73] = nop
	instr[0x74] = nop
	instr[0x75] = nop
	instr[0x76] = nop
	instr[0x77] = nop
	instr[0x78] = nop
	instr[0x79] = nop
	instr[0x7a] = nop
	instr[0x7b] = nop
	instr[0x7c] = nop
	instr[0x7d] = nop
	instr[0x7e] = nop
	instr[0x7f] = nop

	instr[0x80] = nop
	instr[0x81] = nop
	instr[0x82] = nop
	instr[0x83] = nop
	instr[0x84] = nop
	instr[0x85] = nop
	instr[0x86] = nop
	instr[0x87] = nop
	instr[0x88] = nop
	instr[0x89] = nop
	instr[0x8a] = nop
	instr[0x8b] = nop
	instr[0x8c] = nop
	instr[0x8d] = nop
	instr[0x8e] = nop
	instr[0x8f] = nop

	instr[0x90] = nop
	instr[0x91] = nop
	instr[0x92] = nop
	instr[0x93] = nop
	instr[0x94] = nop
	instr[0x95] = nop
	instr[0x96] = nop
	instr[0x97] = nop
	instr[0x98] = nop
	instr[0x99] = nop
	instr[0x9a] = nop
	instr[0x9b] = nop
	instr[0x9c] = nop
	instr[0x9d] = nop
	instr[0x9e] = nop
	instr[0x9f] = nop

	instr[0xa0] = nop
	instr[0xa1] = nop
	instr[0xa2] = nop
	instr[0xa3] = nop
	instr[0xa4] = nop
	instr[0xa5] = nop
	instr[0xa6] = nop
	instr[0xa7] = nop
	instr[0xa8] = nop
	instr[0xa9] = nop
	instr[0xaa] = nop
	instr[0xab] = nop
	instr[0xac] = nop
	instr[0xad] = nop
	instr[0xae] = nop
	instr[0xaf] = nop

	instr[0xb0] = nop
	instr[0xb1] = nop
	instr[0xb2] = nop
	instr[0xb3] = nop
	instr[0xb4] = nop
	instr[0xb5] = nop
	instr[0xb6] = nop
	instr[0xb7] = nop
	instr[0xb8] = nop
	instr[0xb9] = nop
	instr[0xba] = nop
	instr[0xbb] = nop
	instr[0xbc] = nop
	instr[0xbd] = nop
	instr[0xbe] = nop
	instr[0xbf] = nop

	instr[0xc0] = nop
	instr[0xc1] = nop
	instr[0xc2] = nop
	instr[0xc3] = nop
	instr[0xc4] = nop
	instr[0xc5] = nop
	instr[0xc6] = nop
	instr[0xc7] = nop
	instr[0xc8] = nop
	instr[0xc9] = nop
	instr[0xca] = nop
	instr[0xcb] = nop
	instr[0xcc] = nop
	instr[0xcd] = nop
	instr[0xce] = nop
	instr[0xcf] = nop

	instr[0xd0] = nop
	instr[0xd1] = nop
	instr[0xd2] = nop
	instr[0xd3] = nop
	instr[0xd4] = nop
	instr[0xd5] = nop
	instr[0xd6] = nop
	instr[0xd7] = nop
	instr[0xd8] = nop
	instr[0xd9] = nop
	instr[0xda] = nop
	instr[0xdb] = nop
	instr[0xdc] = nop
	instr[0xdd] = nop
	instr[0xde] = nop
	instr[0xdf] = nop

	instr[0xe0] = nop
	instr[0xe1] = nop
	instr[0xe2] = nop
	instr[0xe3] = nop
	instr[0xe4] = nop
	instr[0xe5] = nop
	instr[0xe6] = nop
	instr[0xe7] = nop
	instr[0xe8] = nop
	instr[0xe9] = nop
	instr[0xea] = nop
	instr[0xeb] = nop
	instr[0xec] = nop
	instr[0xed] = nop
	instr[0xee] = nop
	instr[0xef] = nop

	instr[0xf0] = nop
	instr[0xf1] = nop
	instr[0xf2] = nop
	instr[0xf3] = nop
	instr[0xf4] = nop
	instr[0xf5] = nop
	instr[0xf6] = nop
	instr[0xf7] = nop
	instr[0xf8] = nop
	instr[0xf9] = nop
	instr[0xfa] = nop
	instr[0xfb] = nop
	instr[0xfc] = nop
	instr[0xfd] = nop
	instr[0xfe] = nop
	instr[0xff] = nop

	return instr
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

func incBC(cpu *CPU) { incNN(cpu, cpu.Registers.BC, cpu.Registers.SetBC) } // INC BC
func incDE(cpu *CPU) { incNN(cpu, cpu.Registers.DE, cpu.Registers.SetDE) } // INC DE
func incHL(cpu *CPU) { incNN(cpu, cpu.Registers.HL, cpu.Registers.SetHL) } // INC HL
func incSP(cpu *CPU) { cpu.SP++ }                                          // INC SP
// INC (HL)

func decA(cpu *CPU) { decN(cpu, &cpu.Registers.A) } // DEC A
func decB(cpu *CPU) { decN(cpu, &cpu.Registers.B) } // DEC B
func decC(cpu *CPU) { decN(cpu, &cpu.Registers.C) } // DEC C
func decD(cpu *CPU) { decN(cpu, &cpu.Registers.D) } // DEC D
func decE(cpu *CPU) { decN(cpu, &cpu.Registers.E) } // DEC E
func decH(cpu *CPU) { decN(cpu, &cpu.Registers.H) } // DEC H
func decL(cpu *CPU) { decN(cpu, &cpu.Registers.L) } // DEC L

func decBC(cpu *CPU) { decNN(cpu, cpu.Registers.BC, cpu.Registers.SetBC) } // DEC BC
func decDE(cpu *CPU) { decNN(cpu, cpu.Registers.DE, cpu.Registers.SetDE) } // DEC DE
func decHL(cpu *CPU) { decNN(cpu, cpu.Registers.HL, cpu.Registers.SetHL) } // DEC HL
func decSP(cpu *CPU) { cpu.SP-- }                                          // DEC SP
// DEC (HL)

// INC n -- Increment register n
func incN(cpu *CPU, n *byte) {
	log.Println("incN")
	*n++
}

// INC nn -- Increment register nn
func incNN(cpu *CPU, getNN func() uint16, setNN func(uint16)) {
	log.Println("incNN")
	setNN(getNN() + 1)
}

// DEC n -- Decrement register n
func decN(cpu *CPU, n *byte) {
	log.Println("decN")
	*n--
}

// DEC nn -- Decrement register nn
func decNN(cpu *CPU, getNN func() uint16, setNN func(uint16)) {
	setNN(getNN() - 1)
}
