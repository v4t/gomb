package hardware

import (
	"log"
)

// Instructions represents basic cpu instructions
var Instructions = initInstr()

// Init instructions
func initInstr() []func(*CPU) {
	var instr = make([]func(*CPU), 0x100)

	instr[0x00] = nop            // NOP
	instr[0x01] = loadNNToBC     // LD BC,nn
	instr[0x02] = nop            //
	instr[0x03] = incBC          // INC BC
	instr[0x04] = incB           // INC B
	instr[0x05] = decB           // DEC B
	instr[0x06] = load8BitValToB // LD B,n
	instr[0x07] = nop
	instr[0x08] = loadSPToAddressNN // LD (nn),SP
	instr[0x09] = nop
	instr[0x0a] = nop
	instr[0x0b] = decBC          // DEC BC
	instr[0x0c] = incC           // INC C
	instr[0x0d] = decC           // DEC C
	instr[0x0e] = load8BitValToC // LD C,n
	instr[0x0f] = nop

	instr[0x10] = nop
	instr[0x11] = loadNNToDE // LD DE,nn
	instr[0x12] = nop
	instr[0x13] = incDE          // INC DE
	instr[0x14] = incD           // INC D
	instr[0x15] = decD           // DEC D
	instr[0x16] = load8BitValToD // LD D,n
	instr[0x17] = nop
	instr[0x18] = nop
	instr[0x19] = nop
	instr[0x1a] = nop
	instr[0x1b] = decDE          // DEC DE
	instr[0x1c] = incE           // INC E
	instr[0x1d] = decE           // DEC E
	instr[0x1e] = load8BitValToE // LD E,n
	instr[0x1f] = nop

	instr[0x20] = nop
	instr[0x21] = loadNNToHL // LD HL,nn
	instr[0x22] = nop
	instr[0x23] = incHL          // INC HL
	instr[0x24] = incH           // INC H
	instr[0x25] = decH           // DEC H
	instr[0x26] = load8BitValToH // LD H,n
	instr[0x27] = nop
	instr[0x28] = nop
	instr[0x29] = nop
	instr[0x2a] = nop
	instr[0x2b] = decHL          // DEC HL
	instr[0x2c] = incL           // INC L
	instr[0x2d] = decL           // DEC L
	instr[0x2e] = load8BitValToL // LD L,n
	instr[0x2f] = nop

	instr[0x30] = nop
	instr[0x31] = loadNNToSP // LD SP,nn
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

	instr[0x40] = func(cpu *CPU) { loadR2ToR1(cpu, &cpu.Registers.B, &cpu.Registers.B) } // LD B,B
	instr[0x41] = func(cpu *CPU) { loadR2ToR1(cpu, &cpu.Registers.B, &cpu.Registers.C) } // LD B,C
	instr[0x42] = func(cpu *CPU) { loadR2ToR1(cpu, &cpu.Registers.B, &cpu.Registers.D) } // LD B,D
	instr[0x43] = func(cpu *CPU) { loadR2ToR1(cpu, &cpu.Registers.B, &cpu.Registers.E) } // LD B,E
	instr[0x44] = func(cpu *CPU) { loadR2ToR1(cpu, &cpu.Registers.B, &cpu.Registers.H) } // LD B,H
	instr[0x45] = func(cpu *CPU) { loadR2ToR1(cpu, &cpu.Registers.B, &cpu.Registers.L) } // LD B,L
	instr[0x46] = nop                                                                    // LD B,(HL)
	instr[0x47] = func(cpu *CPU) { loadR2ToR1(cpu, &cpu.Registers.B, &cpu.Registers.A) } // LD B,A
	instr[0x48] = func(cpu *CPU) { loadR2ToR1(cpu, &cpu.Registers.C, &cpu.Registers.B) } // LD C,B
	instr[0x49] = func(cpu *CPU) { loadR2ToR1(cpu, &cpu.Registers.C, &cpu.Registers.C) } // LD C,C
	instr[0x4a] = func(cpu *CPU) { loadR2ToR1(cpu, &cpu.Registers.C, &cpu.Registers.D) } // LD C,D
	instr[0x4b] = func(cpu *CPU) { loadR2ToR1(cpu, &cpu.Registers.C, &cpu.Registers.E) } // LD C,E
	instr[0x4c] = func(cpu *CPU) { loadR2ToR1(cpu, &cpu.Registers.C, &cpu.Registers.H) } // LD C,H
	instr[0x4d] = func(cpu *CPU) { loadR2ToR1(cpu, &cpu.Registers.C, &cpu.Registers.L) } // LD C,L
	instr[0x4e] = nop                                                                    // LD C,(HL)
	instr[0x4f] = func(cpu *CPU) { loadR2ToR1(cpu, &cpu.Registers.C, &cpu.Registers.A) } // LD C,A

	instr[0x50] = func(cpu *CPU) { loadR2ToR1(cpu, &cpu.Registers.D, &cpu.Registers.B) } // LD D,B
	instr[0x51] = func(cpu *CPU) { loadR2ToR1(cpu, &cpu.Registers.D, &cpu.Registers.C) } // LD D,C
	instr[0x52] = func(cpu *CPU) { loadR2ToR1(cpu, &cpu.Registers.D, &cpu.Registers.D) } // LD D,D
	instr[0x53] = func(cpu *CPU) { loadR2ToR1(cpu, &cpu.Registers.D, &cpu.Registers.E) } // LD D,E
	instr[0x54] = func(cpu *CPU) { loadR2ToR1(cpu, &cpu.Registers.D, &cpu.Registers.H) } // LD D,H
	instr[0x55] = func(cpu *CPU) { loadR2ToR1(cpu, &cpu.Registers.D, &cpu.Registers.L) } // LD D,L
	instr[0x56] = nop                                                                    // LD D,(HL)
	instr[0x57] = func(cpu *CPU) { loadR2ToR1(cpu, &cpu.Registers.D, &cpu.Registers.A) } // LD D,A
	instr[0x58] = func(cpu *CPU) { loadR2ToR1(cpu, &cpu.Registers.E, &cpu.Registers.B) } // LD E,B
	instr[0x59] = func(cpu *CPU) { loadR2ToR1(cpu, &cpu.Registers.E, &cpu.Registers.C) } // LD E,C
	instr[0x5a] = func(cpu *CPU) { loadR2ToR1(cpu, &cpu.Registers.E, &cpu.Registers.D) } // LD E,D
	instr[0x5b] = func(cpu *CPU) { loadR2ToR1(cpu, &cpu.Registers.E, &cpu.Registers.E) } // LD E,E
	instr[0x5c] = func(cpu *CPU) { loadR2ToR1(cpu, &cpu.Registers.E, &cpu.Registers.H) } // LD E,H
	instr[0x5d] = func(cpu *CPU) { loadR2ToR1(cpu, &cpu.Registers.E, &cpu.Registers.L) } // LD E,L
	instr[0x5e] = nop                                                                    // LD E,(HL)
	instr[0x5f] = func(cpu *CPU) { loadR2ToR1(cpu, &cpu.Registers.E, &cpu.Registers.A) } // LD E,A

	instr[0x60] = func(cpu *CPU) { loadR2ToR1(cpu, &cpu.Registers.H, &cpu.Registers.B) } // LD H,B
	instr[0x61] = func(cpu *CPU) { loadR2ToR1(cpu, &cpu.Registers.H, &cpu.Registers.C) } // LD H,C
	instr[0x62] = func(cpu *CPU) { loadR2ToR1(cpu, &cpu.Registers.H, &cpu.Registers.D) } // LD H,D
	instr[0x63] = func(cpu *CPU) { loadR2ToR1(cpu, &cpu.Registers.H, &cpu.Registers.E) } // LD H,E
	instr[0x64] = func(cpu *CPU) { loadR2ToR1(cpu, &cpu.Registers.H, &cpu.Registers.H) } // LD H,H
	instr[0x65] = func(cpu *CPU) { loadR2ToR1(cpu, &cpu.Registers.H, &cpu.Registers.L) } // LD H,L
	instr[0x66] = nop                                                                    // LD H,(HL)
	instr[0x67] = func(cpu *CPU) { loadR2ToR1(cpu, &cpu.Registers.H, &cpu.Registers.A) } // LD H,A
	instr[0x68] = func(cpu *CPU) { loadR2ToR1(cpu, &cpu.Registers.L, &cpu.Registers.B) } // LD L,B
	instr[0x69] = func(cpu *CPU) { loadR2ToR1(cpu, &cpu.Registers.L, &cpu.Registers.C) } // LD L,C
	instr[0x6a] = func(cpu *CPU) { loadR2ToR1(cpu, &cpu.Registers.L, &cpu.Registers.D) } // LD L,D
	instr[0x6b] = func(cpu *CPU) { loadR2ToR1(cpu, &cpu.Registers.L, &cpu.Registers.E) } // LD L,E
	instr[0x6c] = func(cpu *CPU) { loadR2ToR1(cpu, &cpu.Registers.L, &cpu.Registers.H) } // LD L,H
	instr[0x6d] = func(cpu *CPU) { loadR2ToR1(cpu, &cpu.Registers.L, &cpu.Registers.L) } // LD L,L
	instr[0x6e] = nop                                                                    // LD L,(HL)
	instr[0x6f] = func(cpu *CPU) { loadR2ToR1(cpu, &cpu.Registers.L, &cpu.Registers.A) } // LD L,A

	instr[0x70] = nop
	instr[0x71] = nop
	instr[0x72] = nop
	instr[0x73] = nop
	instr[0x74] = nop
	instr[0x75] = nop
	instr[0x76] = nop
	instr[0x77] = nop
	instr[0x78] = func(cpu *CPU) { loadR2ToR1(cpu, &cpu.Registers.A, &cpu.Registers.B) } // LD A,B
	instr[0x79] = func(cpu *CPU) { loadR2ToR1(cpu, &cpu.Registers.A, &cpu.Registers.C) } // LD A,C
	instr[0x7a] = func(cpu *CPU) { loadR2ToR1(cpu, &cpu.Registers.A, &cpu.Registers.D) } // LD A,D
	instr[0x7b] = func(cpu *CPU) { loadR2ToR1(cpu, &cpu.Registers.A, &cpu.Registers.E) } // LD A,E
	instr[0x7c] = func(cpu *CPU) { loadR2ToR1(cpu, &cpu.Registers.A, &cpu.Registers.H) } // LD A,H
	instr[0x7d] = func(cpu *CPU) { loadR2ToR1(cpu, &cpu.Registers.A, &cpu.Registers.L) } // LD A,L
	instr[0x7e] = nop                                                                    // LD A,(HL)
	instr[0x7f] = func(cpu *CPU) { loadR2ToR1(cpu, &cpu.Registers.A, &cpu.Registers.A) } // LD A,A

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
	instr[0xc1] = popBC
	instr[0xc2] = nop
	instr[0xc3] = nop
	instr[0xc4] = nop
	instr[0xc5] = pushBC
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
	instr[0xd1] = popDE
	instr[0xd2] = nop
	instr[0xd3] = nop
	instr[0xd4] = nop
	instr[0xd5] = pushDE
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
	instr[0xe1] = popHL
	instr[0xe2] = nop
	instr[0xe3] = nop
	instr[0xe4] = nop
	instr[0xe5] = pushHL
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
	instr[0xf1] = popAF
	instr[0xf2] = nop
	instr[0xf3] = nop
	instr[0xf4] = nop
	instr[0xf5] = pushAF
	instr[0xf6] = nop
	instr[0xf7] = nop
	instr[0xf8] = loadSPPlusNToHL // LD HL,SP+n
	instr[0xf9] = loadHLToSP      // LD SP,HL
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

/* 8-bit ALU */
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

/* 8-bit loads */
func load8BitValToB(cpu *CPU) { load8BitValueIntoN(cpu, &cpu.Registers.B) } // LD B,n
func load8BitValToC(cpu *CPU) { load8BitValueIntoN(cpu, &cpu.Registers.C) } // LD C,n
func load8BitValToD(cpu *CPU) { load8BitValueIntoN(cpu, &cpu.Registers.D) } // LD D,n
func load8BitValToE(cpu *CPU) { load8BitValueIntoN(cpu, &cpu.Registers.E) } // LD E,n
func load8BitValToH(cpu *CPU) { load8BitValueIntoN(cpu, &cpu.Registers.H) } // LD H,n
func load8BitValToL(cpu *CPU) { load8BitValueIntoN(cpu, &cpu.Registers.L) } // LD L,n

func load8BitValueIntoN(cpu *CPU, n *byte) {
	*n = cpu.Fetch()
}

func loadR2ToR1(cpu *CPU, r1, r2 *byte) {
	*r1 = *r2
}

/* 16-bit loads */
func loadNNToBC(cpu *CPU) { loadNIntoNN(cpu, cpu.Registers.SetBC) } // LD BC,nn
func loadNNToDE(cpu *CPU) { loadNIntoNN(cpu, cpu.Registers.SetDE) } // LD DE,nn
func loadNNToHL(cpu *CPU) { loadNIntoNN(cpu, cpu.Registers.SetHL) } // LD HL,nn
func loadNNToSP(cpu *CPU) {
	// LD SP,nn
	i := uint16(cpu.Fetch())
	j := uint16(cpu.Fetch())
	var n uint16 = i<<8 | j
	cpu.SP = n
}

func loadHLToSP(cpu *CPU) {
	// LD SP,HL
	cpu.SP = cpu.Registers.HL()
}

func loadSPPlusNToHL(cpu *CPU) {
	// LD HL,SP+n
	log.Fatalln("Not implemented: LD HL,SP+n")
}

func loadSPToAddressNN(cpu *CPU) {
	// LD (nn),SP
	log.Fatalln("Not implemented: LD (nn),SP")
}

func pushAF(cpu *CPU) { pushNN(cpu, cpu.Registers.AF()) } // PUSH AF
func pushBC(cpu *CPU) { pushNN(cpu, cpu.Registers.BC()) } // PUSH BC
func pushDE(cpu *CPU) { pushNN(cpu, cpu.Registers.DE()) } // PUSH DE
func pushHL(cpu *CPU) { pushNN(cpu, cpu.Registers.HL()) } // PUSH HL

func popAF(cpu *CPU) { popNN(cpu, cpu.Registers.AF()) } // POP AF
func popBC(cpu *CPU) { popNN(cpu, cpu.Registers.BC()) } // POP BC
func popDE(cpu *CPU) { popNN(cpu, cpu.Registers.DE()) } // POP DE
func popHL(cpu *CPU) { popNN(cpu, cpu.Registers.HL()) } // POP HL

func pushNN(cpu *CPU, address uint16) {
	log.Fatalln("Not implemented: PUSH nn")
}

func popNN(cpu *CPU, address uint16) {
	log.Fatalln("Not implemented: POP nn")
}

// LD n,nn -- Put value nn into n.
func loadNIntoNN(cpu *CPU, setNN func(uint16)) {
	i := uint16(cpu.Fetch())
	j := uint16(cpu.Fetch())
	var n uint16 = i<<8 | j
	setNN(n)
}
