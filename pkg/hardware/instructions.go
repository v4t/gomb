package hardware

import (
	"log"
)

// Instructions represents basic cpu instructions
var Instructions = initInstr()

// Init instructions
func initInstr() []func(*CPU) {
	var instr = make([]func(*CPU), 0x100)

	instr[0x00] = nop                                                              // NOP
	instr[0x01] = loadNNToBC                                                       // LD BC,nn
	instr[0x02] = func(cpu *CPU) { MemWrite(cpu.Registers.BC(), cpu.Registers.A) } // LD (BC),A
	instr[0x03] = incBC                                                            // INC BC
	instr[0x04] = incB                                                             // INC B
	instr[0x05] = decB                                                             // DEC B
	instr[0x06] = load8BitValToB                                                   // LD B,n
	instr[0x07] = func(cpu *CPU) { rlca(cpu) }                                     // RLCA
	instr[0x08] = loadSPToAddressNN                                                // LD (nn),SP
	instr[0x09] = nop
	instr[0x0a] = func(cpu *CPU) { cpu.Registers.A = MemRead(cpu.Registers.BC()) } // LD A,(BC)
	instr[0x0b] = decBC                                                            // DEC BC
	instr[0x0c] = incC                                                             // INC C
	instr[0x0d] = decC                                                             // DEC C
	instr[0x0e] = load8BitValToC                                                   // LD C,n
	instr[0x0f] = func(cpu *CPU) { rrca(cpu) }                                     // RRCA

	instr[0x10] = nop
	instr[0x11] = loadNNToDE                                                             // LD DE,nn
	instr[0x12] = func(cpu *CPU) { MemWrite(cpu.Registers.DE(), cpu.Registers.A) }       // LD (DE),A
	instr[0x13] = incDE                                                                  // INC DE
	instr[0x14] = incD                                                                   // INC D
	instr[0x15] = decD                                                                   // DEC D
	instr[0x16] = load8BitValToD                                                         // LD D,n
	instr[0x17] = func(cpu *CPU) { rla(cpu) }                                            // RLA
	instr[0x18] = func(cpu *CPU) { cpu.PC = uint16(int32(cpu.PC) + int32(cpu.Fetch())) } // JR n
	instr[0x19] = nop
	instr[0x1a] = func(cpu *CPU) { cpu.Registers.A = MemRead(cpu.Registers.DE()) } // LD A,(DE)
	instr[0x1b] = decDE                                                            // DEC DE
	instr[0x1c] = incE                                                             // INC E
	instr[0x1d] = decE                                                             // DEC E
	instr[0x1e] = load8BitValToE                                                   // LD E,n
	instr[0x1f] = func(cpu *CPU) { rra(cpu) }                                      // RRA

	instr[0x20] = func(cpu *CPU) { conditionalRelativeJump(cpu, !cpu.Zero(), cpu.Fetch()) } // JP NZ,*
	instr[0x21] = loadNNToHL                                                                // LD HL,nn
	instr[0x22] = func(cpu *CPU) {
		// LDI (HL),A
		MemWrite(cpu.Registers.HL(), cpu.Registers.A)
		cpu.Registers.SetHL(cpu.Registers.HL() + 1)
	}
	instr[0x23] = incHL          // INC HL
	instr[0x24] = incH           // INC H
	instr[0x25] = decH           // DEC H
	instr[0x26] = load8BitValToH // LD H,n
	instr[0x27] = nop
	instr[0x28] = func(cpu *CPU) { conditionalRelativeJump(cpu, cpu.Zero(), cpu.Fetch()) } // JP Z,*
	instr[0x29] = nop
	instr[0x2a] = func(cpu *CPU) {
		// LDI A,(HL)
		cpu.Registers.A = MemRead(cpu.Registers.HL())
		cpu.Registers.SetHL(cpu.Registers.HL() + 1)
	}
	instr[0x2b] = decHL          // DEC HL
	instr[0x2c] = incL           // INC L
	instr[0x2d] = decL           // DEC L
	instr[0x2e] = load8BitValToL // LD L,n
	instr[0x2f] = complementA    // CPL

	instr[0x30] = func(cpu *CPU) { conditionalRelativeJump(cpu, !cpu.Carry(), cpu.Fetch()) } // JP NC,*
	instr[0x31] = loadNNToSP                                                                 // LD SP,nn
	instr[0x32] = func(cpu *CPU) {
		// LDD (HL),A
		MemWrite(cpu.Registers.HL(), cpu.Registers.A)
		cpu.Registers.SetHL(cpu.Registers.HL() - 1)
	}
	instr[0x33] = incSP // INC SP
	instr[0x34] = nop
	instr[0x35] = nop
	instr[0x36] = func(cpu *CPU) { MemWrite(cpu.Registers.HL(), cpu.Fetch()) } // LD (HL),n
	instr[0x37] = setCarryFlag
	instr[0x38] = func(cpu *CPU) { conditionalRelativeJump(cpu, cpu.Carry(), cpu.Fetch()) } // JP C,*
	instr[0x39] = nop
	instr[0x3a] = func(cpu *CPU) {
		// LDD A,(HL)
		cpu.Registers.A = MemRead(cpu.Registers.HL())
		cpu.Registers.SetHL(cpu.Registers.HL() - 1)
	}
	instr[0x3b] = decSP                                            // DEC SP
	instr[0x3c] = incA                                             // INC A
	instr[0x3d] = decA                                             // DEC A
	instr[0x3e] = func(cpu *CPU) { cpu.Registers.A = cpu.Fetch() } // LD A,#
	instr[0x3f] = complementCarry                                  // CCF

	instr[0x40] = func(cpu *CPU) { loadR2ToR1(cpu, &cpu.Registers.B, &cpu.Registers.B) } // LD B,B
	instr[0x41] = func(cpu *CPU) { loadR2ToR1(cpu, &cpu.Registers.B, &cpu.Registers.C) } // LD B,C
	instr[0x42] = func(cpu *CPU) { loadR2ToR1(cpu, &cpu.Registers.B, &cpu.Registers.D) } // LD B,D
	instr[0x43] = func(cpu *CPU) { loadR2ToR1(cpu, &cpu.Registers.B, &cpu.Registers.E) } // LD B,E
	instr[0x44] = func(cpu *CPU) { loadR2ToR1(cpu, &cpu.Registers.B, &cpu.Registers.H) } // LD B,H
	instr[0x45] = func(cpu *CPU) { loadR2ToR1(cpu, &cpu.Registers.B, &cpu.Registers.L) } // LD B,L
	instr[0x46] = func(cpu *CPU) { cpu.Registers.B = MemRead(cpu.Registers.HL()) }       // LD B,(HL)
	instr[0x47] = func(cpu *CPU) { loadR2ToR1(cpu, &cpu.Registers.B, &cpu.Registers.A) } // LD B,A
	instr[0x48] = func(cpu *CPU) { loadR2ToR1(cpu, &cpu.Registers.C, &cpu.Registers.B) } // LD C,B
	instr[0x49] = func(cpu *CPU) { loadR2ToR1(cpu, &cpu.Registers.C, &cpu.Registers.C) } // LD C,C
	instr[0x4a] = func(cpu *CPU) { loadR2ToR1(cpu, &cpu.Registers.C, &cpu.Registers.D) } // LD C,D
	instr[0x4b] = func(cpu *CPU) { loadR2ToR1(cpu, &cpu.Registers.C, &cpu.Registers.E) } // LD C,E
	instr[0x4c] = func(cpu *CPU) { loadR2ToR1(cpu, &cpu.Registers.C, &cpu.Registers.H) } // LD C,H
	instr[0x4d] = func(cpu *CPU) { loadR2ToR1(cpu, &cpu.Registers.C, &cpu.Registers.L) } // LD C,L
	instr[0x4e] = func(cpu *CPU) { cpu.Registers.C = MemRead(cpu.Registers.HL()) }       // LD C,(HL)
	instr[0x4f] = func(cpu *CPU) { loadR2ToR1(cpu, &cpu.Registers.C, &cpu.Registers.A) } // LD C,A

	instr[0x50] = func(cpu *CPU) { loadR2ToR1(cpu, &cpu.Registers.D, &cpu.Registers.B) } // LD D,B
	instr[0x51] = func(cpu *CPU) { loadR2ToR1(cpu, &cpu.Registers.D, &cpu.Registers.C) } // LD D,C
	instr[0x52] = func(cpu *CPU) { loadR2ToR1(cpu, &cpu.Registers.D, &cpu.Registers.D) } // LD D,D
	instr[0x53] = func(cpu *CPU) { loadR2ToR1(cpu, &cpu.Registers.D, &cpu.Registers.E) } // LD D,E
	instr[0x54] = func(cpu *CPU) { loadR2ToR1(cpu, &cpu.Registers.D, &cpu.Registers.H) } // LD D,H
	instr[0x55] = func(cpu *CPU) { loadR2ToR1(cpu, &cpu.Registers.D, &cpu.Registers.L) } // LD D,L
	instr[0x56] = func(cpu *CPU) { cpu.Registers.D = MemRead(cpu.Registers.HL()) }       // LD D,(HL)
	instr[0x57] = func(cpu *CPU) { loadR2ToR1(cpu, &cpu.Registers.D, &cpu.Registers.A) } // LD D,A
	instr[0x58] = func(cpu *CPU) { loadR2ToR1(cpu, &cpu.Registers.E, &cpu.Registers.B) } // LD E,B
	instr[0x59] = func(cpu *CPU) { loadR2ToR1(cpu, &cpu.Registers.E, &cpu.Registers.C) } // LD E,C
	instr[0x5a] = func(cpu *CPU) { loadR2ToR1(cpu, &cpu.Registers.E, &cpu.Registers.D) } // LD E,D
	instr[0x5b] = func(cpu *CPU) { loadR2ToR1(cpu, &cpu.Registers.E, &cpu.Registers.E) } // LD E,E
	instr[0x5c] = func(cpu *CPU) { loadR2ToR1(cpu, &cpu.Registers.E, &cpu.Registers.H) } // LD E,H
	instr[0x5d] = func(cpu *CPU) { loadR2ToR1(cpu, &cpu.Registers.E, &cpu.Registers.L) } // LD E,L
	instr[0x5e] = func(cpu *CPU) { cpu.Registers.E = MemRead(cpu.Registers.HL()) }       // LD E,(HL)
	instr[0x5f] = func(cpu *CPU) { loadR2ToR1(cpu, &cpu.Registers.E, &cpu.Registers.A) } // LD E,A

	instr[0x60] = func(cpu *CPU) { loadR2ToR1(cpu, &cpu.Registers.H, &cpu.Registers.B) } // LD H,B
	instr[0x61] = func(cpu *CPU) { loadR2ToR1(cpu, &cpu.Registers.H, &cpu.Registers.C) } // LD H,C
	instr[0x62] = func(cpu *CPU) { loadR2ToR1(cpu, &cpu.Registers.H, &cpu.Registers.D) } // LD H,D
	instr[0x63] = func(cpu *CPU) { loadR2ToR1(cpu, &cpu.Registers.H, &cpu.Registers.E) } // LD H,E
	instr[0x64] = func(cpu *CPU) { loadR2ToR1(cpu, &cpu.Registers.H, &cpu.Registers.H) } // LD H,H
	instr[0x65] = func(cpu *CPU) { loadR2ToR1(cpu, &cpu.Registers.H, &cpu.Registers.L) } // LD H,L
	instr[0x66] = func(cpu *CPU) { cpu.Registers.H = MemRead(cpu.Registers.HL()) }       // LD H,(HL)
	instr[0x67] = func(cpu *CPU) { loadR2ToR1(cpu, &cpu.Registers.H, &cpu.Registers.A) } // LD H,A
	instr[0x68] = func(cpu *CPU) { loadR2ToR1(cpu, &cpu.Registers.L, &cpu.Registers.B) } // LD L,B
	instr[0x69] = func(cpu *CPU) { loadR2ToR1(cpu, &cpu.Registers.L, &cpu.Registers.C) } // LD L,C
	instr[0x6a] = func(cpu *CPU) { loadR2ToR1(cpu, &cpu.Registers.L, &cpu.Registers.D) } // LD L,D
	instr[0x6b] = func(cpu *CPU) { loadR2ToR1(cpu, &cpu.Registers.L, &cpu.Registers.E) } // LD L,E
	instr[0x6c] = func(cpu *CPU) { loadR2ToR1(cpu, &cpu.Registers.L, &cpu.Registers.H) } // LD L,H
	instr[0x6d] = func(cpu *CPU) { loadR2ToR1(cpu, &cpu.Registers.L, &cpu.Registers.L) } // LD L,L
	instr[0x6e] = func(cpu *CPU) { cpu.Registers.L = MemRead(cpu.Registers.HL()) }       // LD L,(HL)
	instr[0x6f] = func(cpu *CPU) { loadR2ToR1(cpu, &cpu.Registers.L, &cpu.Registers.A) } // LD L,A

	instr[0x70] = func(cpu *CPU) { MemWrite(cpu.Registers.HL(), cpu.Registers.B) } // LD (HL),B
	instr[0x71] = func(cpu *CPU) { MemWrite(cpu.Registers.HL(), cpu.Registers.C) } // LD (HL),C
	instr[0x72] = func(cpu *CPU) { MemWrite(cpu.Registers.HL(), cpu.Registers.D) } // LD (HL),D
	instr[0x73] = func(cpu *CPU) { MemWrite(cpu.Registers.HL(), cpu.Registers.E) } // LD (HL),E
	instr[0x74] = func(cpu *CPU) { MemWrite(cpu.Registers.HL(), cpu.Registers.H) } // LD (HL),H
	instr[0x75] = func(cpu *CPU) { MemWrite(cpu.Registers.HL(), cpu.Registers.L) } // LD (HL),L
	instr[0x76] = nop
	instr[0x77] = func(cpu *CPU) { MemWrite(cpu.Registers.HL(), cpu.Registers.A) }       // LD (HL),A
	instr[0x78] = func(cpu *CPU) { loadR2ToR1(cpu, &cpu.Registers.A, &cpu.Registers.B) } // LD A,B
	instr[0x79] = func(cpu *CPU) { loadR2ToR1(cpu, &cpu.Registers.A, &cpu.Registers.C) } // LD A,C
	instr[0x7a] = func(cpu *CPU) { loadR2ToR1(cpu, &cpu.Registers.A, &cpu.Registers.D) } // LD A,D
	instr[0x7b] = func(cpu *CPU) { loadR2ToR1(cpu, &cpu.Registers.A, &cpu.Registers.E) } // LD A,E
	instr[0x7c] = func(cpu *CPU) { loadR2ToR1(cpu, &cpu.Registers.A, &cpu.Registers.H) } // LD A,H
	instr[0x7d] = func(cpu *CPU) { loadR2ToR1(cpu, &cpu.Registers.A, &cpu.Registers.L) } // LD A,L
	instr[0x7e] = func(cpu *CPU) { cpu.Registers.A = MemRead(cpu.Registers.HL()) }       // LD A,(HL)
	instr[0x7f] = func(cpu *CPU) { loadR2ToR1(cpu, &cpu.Registers.A, &cpu.Registers.A) } // LD A,A

	instr[0x80] = func(cpu *CPU) { addToA(cpu, cpu.Registers.B) }                      // ADD A,B
	instr[0x81] = func(cpu *CPU) { addToA(cpu, cpu.Registers.C) }                      // ADD A,C
	instr[0x82] = func(cpu *CPU) { addToA(cpu, cpu.Registers.D) }                      // ADD A,D
	instr[0x83] = func(cpu *CPU) { addToA(cpu, cpu.Registers.E) }                      // ADD A,E
	instr[0x84] = func(cpu *CPU) { addToA(cpu, cpu.Registers.H) }                      // ADD A,H
	instr[0x85] = func(cpu *CPU) { addToA(cpu, cpu.Registers.L) }                      // ADD A,L
	instr[0x86] = func(cpu *CPU) { addToA(cpu, MemRead(cpu.Registers.HL())) }          // ADD A,(HL)
	instr[0x87] = func(cpu *CPU) { addToA(cpu, cpu.Registers.A) }                      // ADD A,A
	instr[0x88] = func(cpu *CPU) { addWithCarryToA(cpu, cpu.Registers.B) }             // ADC A,B
	instr[0x89] = func(cpu *CPU) { addWithCarryToA(cpu, cpu.Registers.C) }             // ADC A,C
	instr[0x8a] = func(cpu *CPU) { addWithCarryToA(cpu, cpu.Registers.D) }             // ADC A,D
	instr[0x8b] = func(cpu *CPU) { addWithCarryToA(cpu, cpu.Registers.E) }             // ADC A,E
	instr[0x8c] = func(cpu *CPU) { addWithCarryToA(cpu, cpu.Registers.H) }             // ADC A,H
	instr[0x8d] = func(cpu *CPU) { addWithCarryToA(cpu, cpu.Registers.L) }             // ADC A,L
	instr[0x8e] = func(cpu *CPU) { addWithCarryToA(cpu, MemRead(cpu.Registers.HL())) } // ADC A,(HL)
	instr[0x8f] = func(cpu *CPU) { addWithCarryToA(cpu, cpu.Registers.A) }             // ADC A,A

	instr[0x90] = func(cpu *CPU) { subFromA(cpu, cpu.Registers.B) }                      // SUB A,B
	instr[0x91] = func(cpu *CPU) { subFromA(cpu, cpu.Registers.C) }                      // SUB A,C
	instr[0x92] = func(cpu *CPU) { subFromA(cpu, cpu.Registers.D) }                      // SUB A,D
	instr[0x93] = func(cpu *CPU) { subFromA(cpu, cpu.Registers.E) }                      // SUB A,E
	instr[0x94] = func(cpu *CPU) { subFromA(cpu, cpu.Registers.H) }                      // SUB A,H
	instr[0x95] = func(cpu *CPU) { subFromA(cpu, cpu.Registers.L) }                      // SUB A,L
	instr[0x96] = func(cpu *CPU) { subFromA(cpu, MemRead(cpu.Registers.HL())) }          // SUB A,(HL)
	instr[0x97] = func(cpu *CPU) { subFromA(cpu, cpu.Registers.A) }                      // SUB A,A
	instr[0x98] = func(cpu *CPU) { subWithCarryFromA(cpu, cpu.Registers.B) }             // SBC A,B
	instr[0x99] = func(cpu *CPU) { subWithCarryFromA(cpu, cpu.Registers.C) }             // SBC A,C
	instr[0x9a] = func(cpu *CPU) { subWithCarryFromA(cpu, cpu.Registers.D) }             // SBC A,D
	instr[0x9b] = func(cpu *CPU) { subWithCarryFromA(cpu, cpu.Registers.E) }             // SBC A,E
	instr[0x9c] = func(cpu *CPU) { subWithCarryFromA(cpu, cpu.Registers.H) }             // SBC A,H
	instr[0x9d] = func(cpu *CPU) { subWithCarryFromA(cpu, cpu.Registers.L) }             // SBC A,L
	instr[0x9e] = func(cpu *CPU) { subWithCarryFromA(cpu, MemRead(cpu.Registers.HL())) } // SBC A,(HL)
	instr[0x9f] = func(cpu *CPU) { subWithCarryFromA(cpu, cpu.Registers.D) }             // SBC A,A

	instr[0xa0] = func(cpu *CPU) { andA(cpu, cpu.Registers.B) }             // AND B
	instr[0xa1] = func(cpu *CPU) { andA(cpu, cpu.Registers.C) }             // AND C
	instr[0xa2] = func(cpu *CPU) { andA(cpu, cpu.Registers.D) }             // AND D
	instr[0xa3] = func(cpu *CPU) { andA(cpu, cpu.Registers.E) }             // AND E
	instr[0xa4] = func(cpu *CPU) { andA(cpu, cpu.Registers.H) }             // AND H
	instr[0xa5] = func(cpu *CPU) { andA(cpu, cpu.Registers.L) }             // AND L
	instr[0xa6] = func(cpu *CPU) { andA(cpu, MemRead(cpu.Registers.HL())) } // AND (HL)
	instr[0xa7] = func(cpu *CPU) { andA(cpu, cpu.Registers.A) }             // AND A
	instr[0xa8] = func(cpu *CPU) { xorA(cpu, cpu.Registers.B) }             // XOR B
	instr[0xa9] = func(cpu *CPU) { xorA(cpu, cpu.Registers.C) }             // XOR C
	instr[0xaa] = func(cpu *CPU) { xorA(cpu, cpu.Registers.D) }             // XOR D
	instr[0xab] = func(cpu *CPU) { xorA(cpu, cpu.Registers.E) }             // XOR E
	instr[0xac] = func(cpu *CPU) { xorA(cpu, cpu.Registers.H) }             // XOR H
	instr[0xad] = func(cpu *CPU) { xorA(cpu, cpu.Registers.L) }             // XOR L
	instr[0xae] = func(cpu *CPU) { xorA(cpu, MemRead(cpu.Registers.HL())) } // XOR (HL)
	instr[0xaf] = func(cpu *CPU) { xorA(cpu, cpu.Registers.A) }             // XOR A

	instr[0xb0] = func(cpu *CPU) { orA(cpu, cpu.Registers.B) }             // OR B
	instr[0xb1] = func(cpu *CPU) { orA(cpu, cpu.Registers.C) }             // OR C
	instr[0xb2] = func(cpu *CPU) { orA(cpu, cpu.Registers.D) }             // OR D
	instr[0xb3] = func(cpu *CPU) { orA(cpu, cpu.Registers.E) }             // OR E
	instr[0xb4] = func(cpu *CPU) { orA(cpu, cpu.Registers.H) }             // OR H
	instr[0xb5] = func(cpu *CPU) { orA(cpu, cpu.Registers.L) }             // OR L
	instr[0xb6] = func(cpu *CPU) { orA(cpu, MemRead(cpu.Registers.HL())) } // OR (HL)
	instr[0xb7] = func(cpu *CPU) { orA(cpu, cpu.Registers.A) }             // OR A
	instr[0xb8] = func(cpu *CPU) { cpA(cpu, cpu.Registers.B) }             // CP B
	instr[0xb9] = func(cpu *CPU) { cpA(cpu, cpu.Registers.C) }             // CP C
	instr[0xba] = func(cpu *CPU) { cpA(cpu, cpu.Registers.D) }             // CP D
	instr[0xbb] = func(cpu *CPU) { cpA(cpu, cpu.Registers.E) }             // CP E
	instr[0xbc] = func(cpu *CPU) { cpA(cpu, cpu.Registers.H) }             // CP H
	instr[0xbd] = func(cpu *CPU) { cpA(cpu, cpu.Registers.L) }             // CP L
	instr[0xbe] = func(cpu *CPU) { cpA(cpu, MemRead(cpu.Registers.HL())) } // CP (HL)
	instr[0xbf] = func(cpu *CPU) { cpA(cpu, cpu.Registers.A) }             // CP A

	instr[0xc0] = nop
	instr[0xc1] = popBC
	instr[0xc2] = func(cpu *CPU) { conditionalJump(cpu, !cpu.Zero(), cpu.Fetch16()) } // JP NZ,nn
	instr[0xc3] = func(cpu *CPU) { cpu.PC = cpu.Fetch16() }                           // JP nn
	instr[0xc4] = func(cpu *CPU) { conditionalCall(cpu, !cpu.Zero(), cpu.Fetch16()) } // CALL NZ,nn
	instr[0xc5] = pushBC
	instr[0xc6] = func(cpu *CPU) { addToA(cpu, cpu.Fetch()) } // ADD A,#
	instr[0xc7] = nop
	instr[0xc8] = nop
	instr[0xc9] = nop
	instr[0xca] = func(cpu *CPU) { conditionalJump(cpu, cpu.Zero(), cpu.Fetch16()) } // JP Z,nn
	instr[0xcb] = nop
	instr[0xcc] = func(cpu *CPU) { conditionalCall(cpu, cpu.Zero(), cpu.Fetch16()) } // CALL Z,nn
	instr[0xcd] = func(cpu *CPU) { call(cpu, cpu.Fetch16()) }                        // CALL nn
	instr[0xce] = func(cpu *CPU) { addWithCarryToA(cpu, cpu.Fetch()) }               // ADC A,#
	instr[0xcf] = nop

	instr[0xd0] = nop
	instr[0xd1] = popDE
	instr[0xd2] = func(cpu *CPU) { conditionalJump(cpu, !cpu.Carry(), cpu.Fetch16()) } // JP NC,nn
	instr[0xd3] = nop
	instr[0xd4] = func(cpu *CPU) { conditionalCall(cpu, !cpu.Carry(), cpu.Fetch16()) } // CALL NC,nn
	instr[0xd5] = pushDE
	instr[0xd6] = func(cpu *CPU) { subFromA(cpu, cpu.Fetch()) } // SUB A,#
	instr[0xd7] = nop
	instr[0xd8] = nop
	instr[0xd9] = nop
	instr[0xda] = func(cpu *CPU) { conditionalJump(cpu, cpu.Carry(), cpu.Fetch16()) } // JP C,nn
	instr[0xdb] = nop
	instr[0xdc] = func(cpu *CPU) { conditionalCall(cpu, cpu.Carry(), cpu.Fetch16()) } // CALL C,nn
	instr[0xdd] = nop
	instr[0xde] = func(cpu *CPU) { subWithCarryFromA(cpu, cpu.Fetch()) } // SBC A,#
	instr[0xdf] = nop

	instr[0xe0] = func(cpu *CPU) { MemWrite(0xff00+uint16(cpu.Fetch()), cpu.Registers.A) } // LDH (n),A
	instr[0xe1] = popHL
	instr[0xe2] = func(cpu *CPU) { MemWrite(0xff00+uint16(cpu.Registers.C), cpu.Registers.A) } // LD (C),A
	instr[0xe3] = nop
	instr[0xe4] = nop
	instr[0xe5] = pushHL
	instr[0xe6] = func(cpu *CPU) { andA(cpu, cpu.Fetch()) } // AND #
	instr[0xe7] = nop
	instr[0xe8] = func(cpu *CPU) { addSignedNToSp(cpu, int8(cpu.Fetch())) }   // ADD SP,n
	instr[0xe9] = func(cpu *CPU) { cpu.PC = cpu.Registers.HL() }              // JP (HL)
	instr[0xea] = func(cpu *CPU) { MemWrite(cpu.Fetch16(), cpu.Registers.A) } // LD (nn),A
	instr[0xeb] = nop
	instr[0xec] = nop
	instr[0xed] = nop
	instr[0xee] = func(cpu *CPU) { xorA(cpu, cpu.Fetch()) } // XOR #
	instr[0xef] = nop

	instr[0xf0] = func(cpu *CPU) { cpu.Registers.A = MemRead(0xff00 + uint16(cpu.Fetch())) } // LDH A,(n)
	instr[0xf1] = popAF
	instr[0xf2] = func(cpu *CPU) { cpu.Registers.A = MemRead(0xff00 + uint16(cpu.Registers.C)) } // LD A,(C)
	instr[0xf3] = nop
	instr[0xf4] = nop
	instr[0xf5] = pushAF
	instr[0xf6] = func(cpu *CPU) { orA(cpu, cpu.Fetch()) } // OR #
	instr[0xf7] = nop
	instr[0xf8] = func(cpu *CPU) { loadSPPlusNToHL(cpu, int8(cpu.Fetch())) }  // LD HL,SP+n
	instr[0xf9] = loadHLToSP                                                  // LD SP,HL
	instr[0xfa] = func(cpu *CPU) { cpu.Registers.A = MemRead(cpu.Fetch16()) } // LD A,(nn)
	instr[0xfb] = nop
	instr[0xfc] = nop
	instr[0xfd] = nop
	instr[0xfe] = func(cpu *CPU) { cpA(cpu, cpu.Fetch()) } // CP #
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

// Add n to A
func addToA(cpu *CPU, n byte) {
	sum := int16(cpu.Registers.A) + int16(n)
	halfCarry := (((cpu.Registers.A & 0xf) + (n & 0xf)) & 0x10) == 0x10
	cpu.Registers.A = byte(sum)

	cpu.SetZero(cpu.Registers.A == 0)
	cpu.SetNegative(false)
	cpu.SetHalfCarry(halfCarry)
	cpu.SetCarry(sum > 0xff)
}

// Add n + Carry flag to A
func addWithCarryToA(cpu *CPU, n byte) {
	var carry int16 = 0
	if cpu.Carry() {
		carry = 1
	}
	sum := int16(cpu.Registers.A) + int16(n) + carry
	halfCarry := ((cpu.Registers.A & 0xf) + (n & 0xf) + byte(carry)) > 0xf
	cpu.Registers.A = byte(sum)

	cpu.SetZero(cpu.Registers.A == 0)
	cpu.SetNegative(false)
	cpu.SetHalfCarry(halfCarry)
	cpu.SetCarry(sum > 0xff)
}

// Sub n from A
func subFromA(cpu *CPU, n byte) {
	diff := int16(cpu.Registers.A) + int16(n)
	halfCarry := ((cpu.Registers.A & 0xf) - (n & 0xf)) < 0
	cpu.Registers.A = byte(diff)

	cpu.SetZero(cpu.Registers.A == 0)
	cpu.SetNegative(true)
	cpu.SetHalfCarry(halfCarry)
	cpu.SetCarry(diff < 0)
}

// Sub n + Carry flag from A
func subWithCarryFromA(cpu *CPU, n byte) {
	diff := int16(cpu.Registers.A) + int16(n)
	halfCarry := ((cpu.Registers.A & 0xf) - (n & 0xf)) < 0
	cpu.Registers.A = byte(diff)

	cpu.SetZero(cpu.Registers.A == 0)
	cpu.SetNegative(true)
	cpu.SetHalfCarry(halfCarry)
	cpu.SetCarry(diff < 0)
}

// Logically AND n with A, result in A
func andA(cpu *CPU, n byte) {
	cpu.Registers.A &= n

	cpu.SetZero(cpu.Registers.A == 0)
	cpu.SetNegative(false)
	cpu.SetHalfCarry(true)
	cpu.SetCarry(false)
}

// Logical OR n with register A, result in A
func orA(cpu *CPU, n byte) {
	cpu.Registers.A |= n

	cpu.SetZero(cpu.Registers.A == 0)
	cpu.SetNegative(false)
	cpu.SetHalfCarry(false)
	cpu.SetCarry(false)
}

// Logical exclusive OR n with register A, result in A
func xorA(cpu *CPU, n byte) {
	cpu.Registers.A ^= n

	cpu.SetZero(cpu.Registers.A == 0)
	cpu.SetNegative(false)
	cpu.SetHalfCarry(false)
	cpu.SetCarry(false)
}

// Compare A with n. This is basically an A - n subtraction instruction but the results are thrown away.
func cpA(cpu *CPU, n byte) {
	result := cpu.Registers.A - n

	cpu.SetZero(result == 0)
	cpu.SetNegative(true)
	cpu.SetHalfCarry((cpu.Registers.A & 0x0f) > (n & 0x0f))
	cpu.SetCarry(cpu.Registers.A > n)
}

/* 16-bit arithmetic */

// ADD SP,n -- Add signed byte n to Stack Pointer (SP)
func addSignedNToSp(cpu *CPU, n int8) {
	sum := int32(cpu.SP) + int32(n)
	if n >= 0 {
		cpu.SetCarry(((sum & 0xff) + int32(n)) > 0xff)
		cpu.SetHalfCarry(((sum & 0xf) + int32(n&0xf)) > 0xf)
	} else {
		cpu.SetCarry((sum & 0xff) <= int32(cpu.SP&0xff))
		cpu.SetHalfCarry((sum & 0xf) <= int32(cpu.SP&0xf))
	}
	cpu.SetZero(false)
	cpu.SetNegative(false)
	cpu.SP = uint16(sum)
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
	cpu.SP = cpu.Fetch16()
}

// LD SP,HL
func loadHLToSP(cpu *CPU) {
	cpu.SP = cpu.Registers.HL()
}

// LD HL,SP+n
func loadSPPlusNToHL(cpu *CPU, n int8) {
	sum := int32(cpu.SP) + int32(n)
	if n >= 0 {
		cpu.SetCarry(((sum & 0xff) + int32(n)) > 0xff)
		cpu.SetHalfCarry(((sum & 0xf) + int32(n&0xf)) > 0xf)
	} else {
		cpu.SetCarry((sum & 0xff) <= int32(cpu.SP&0xff))
		cpu.SetHalfCarry((sum & 0xf) <= int32(cpu.SP&0xf))
	}
	cpu.SetZero(false)
	cpu.SetNegative(false)
	cpu.Registers.SetHL(uint16(sum))
}

// LD (nn),SP -- Put Stack Pointer (SP) at address n.
func loadSPToAddressNN(cpu *CPU) {
	address := cpu.Fetch16()
	MemWrite(address, byte(cpu.SP&0xff))
	MemWrite(address+1, byte(cpu.SP>>8))
}

func pushAF(cpu *CPU) { pushNN(cpu, cpu.Registers.AF()) } // PUSH AF
func pushBC(cpu *CPU) { pushNN(cpu, cpu.Registers.BC()) } // PUSH BC
func pushDE(cpu *CPU) { pushNN(cpu, cpu.Registers.DE()) } // PUSH DE
func pushHL(cpu *CPU) { pushNN(cpu, cpu.Registers.HL()) } // PUSH HL

func popAF(cpu *CPU) { cpu.Registers.SetBC(popNN(cpu)) } // POP AF
func popBC(cpu *CPU) { cpu.Registers.SetBC(popNN(cpu)) } // POP BC
func popDE(cpu *CPU) { cpu.Registers.SetDE(popNN(cpu)) } // POP DE
func popHL(cpu *CPU) { cpu.Registers.SetHL(popNN(cpu)) } // POP HL

func pushNN(cpu *CPU, address uint16) {
	MemWrite(cpu.SP-1, byte(uint16(address&0xff00)>>8))
	MemWrite(cpu.SP-2, byte(address&0xff))
	cpu.SP -= 2
}

func popNN(cpu *CPU) uint16 {
	byte1 := uint16(MemRead(cpu.SP))
	byte2 := uint16(MemRead(cpu.SP+1)) << 8
	cpu.SP += 2
	return byte1 | byte2
}

// LD n,nn -- Put value nn into n.
func loadNIntoNN(cpu *CPU, setNN func(uint16)) {
	setNN(cpu.Fetch16())
}

/* Rotates & shifts */

// RLCA -- Rotate A left. Old bit 7 to Carry flag.
func rlca(cpu *CPU) {
	leavingBit := cpu.Registers.A >> 7
	cpu.Registers.A = cpu.Registers.A<<1 | leavingBit

	cpu.SetZero(cpu.Registers.A == 0)
	cpu.SetNegative(false)
	cpu.SetHalfCarry(false)
	cpu.SetCarry(leavingBit == 1)
}

// RLA -- Rotate A left through Carry flag
func rla(cpu *CPU) {
	leavingBit := cpu.Registers.A >> 7
	var carry byte = 0
	if cpu.Carry() {
		carry = 1
	}
	cpu.Registers.A = cpu.Registers.A<<1 | carry

	cpu.SetZero(cpu.Registers.A == 0)
	cpu.SetNegative(false)
	cpu.SetHalfCarry(false)
	cpu.SetCarry(leavingBit == 1)
}

// RRCA -- Rotate A right. Old bit 0 to Carry flag
func rrca(cpu *CPU) {
	leavingBit := cpu.Registers.A & 1
	cpu.Registers.A = cpu.Registers.A >> 1
	if leavingBit == 1 {
		cpu.Registers.A |= 0x80
	}

	cpu.SetZero(cpu.Registers.A == 0)
	cpu.SetNegative(false)
	cpu.SetHalfCarry(false)
	cpu.SetCarry(leavingBit == 1)
}

// RRA -- Rotate A right through Carry flag.
func rra(cpu *CPU) {
	leavingBit := cpu.Registers.A & 1
	var carry byte = 0
	if cpu.Carry() {
		carry = 1
	}
	cpu.Registers.A = cpu.Registers.A >> 1
	if carry == 1 {
		cpu.Registers.A |= 0x80
	}

	cpu.SetZero(cpu.Registers.A == 0)
	cpu.SetNegative(false)
	cpu.SetHalfCarry(false)
	cpu.SetCarry(leavingBit == 1)
}

/* Misc. */

// DAA -- Decimal adjust register A
func decimalAdjustA(cpu *CPU) {
	log.Fatalln("TODO")
	cpu.SetZero(cpu.Registers.A == 0)
	cpu.SetHalfCarry(false)
}

// CPL -- Complement A register
func complementA(cpu *CPU) {
	cpu.Registers.A = 0xff ^ cpu.Registers.A
	cpu.SetNegative(true)
	cpu.SetHalfCarry(true)
}

// CCF -- Complement carry flag
func complementCarry(cpu *CPU) {
	cpu.SetNegative(false)
	cpu.SetHalfCarry(false)
	cpu.SetCarry(!cpu.Carry())
}

// SCF -- Set carry flag
func setCarryFlag(cpu *CPU) {
	cpu.SetNegative(false)
	cpu.SetHalfCarry(false)
	cpu.SetCarry(true)
}

// HALT -- Power down CPU until an interrupt occurs
func halt(cpu *CPU) {
	log.Fatalln("TODO: HALT")
}

// STOP -- Halt CPU & LCD display until button pressed
func stop(cpu *CPU) {
	log.Fatalln("TODO: STOP")
}

/* Jumps */

// JP cc,nn -- Jump to address n if condition is true
func conditionalJump(cpu *CPU, condition bool, n uint16) {
	if condition {
		cpu.PC = n
	}
}

// JR cc,nn -- If  condition is true then add n to current address and jump to it:
func conditionalRelativeJump(cpu *CPU, condition bool, n byte) {
	if condition {
		address := int32(cpu.PC) + int32(n)
		cpu.PC = uint16(address)
	}
}

// CALL nn -- Push address of next instruction onto stack and then jump to address nn
func call(cpu *CPU, next uint16) {
	pushNN(cpu, cpu.PC)
	cpu.PC = next
}

// CALL cc,nn -- Call address n if condition is true
func conditionalCall(cpu *CPU, condition bool, next uint16) {
	if condition {
		call(cpu, next)
	}
}
