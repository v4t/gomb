package hardware

import (
	"log"
)

var instructions []func()

// ExecuteInstruction executes given operation and returns the amount of cycles taken.
func ExecuteInstruction(cpu *CPU, opCode byte) int {
	if len(instructions) != 0x100 {
		initInstructionList(cpu)
	}
	instructions[opCode]()
	return InstructionCycles[opCode]
}

// InstructionCycles contains the amount of cpu cycles taken per instruction
var InstructionCycles = [0x100]int{
	1, 3, 2, 2, 1, 1, 2, 1, 5, 2, 2, 2, 1, 1, 2, 1,
	0, 3, 2, 2, 1, 1, 2, 1, 3, 2, 2, 2, 1, 1, 2, 1,
	2, 3, 2, 2, 1, 1, 2, 1, 2, 2, 2, 2, 1, 1, 2, 1,
	2, 3, 2, 2, 3, 3, 3, 1, 2, 2, 2, 2, 1, 1, 2, 1,
	1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1, 1, 1, 2, 1,
	1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1, 1, 1, 2, 1,
	1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1, 1, 1, 2, 1,
	2, 2, 2, 2, 2, 2, 0, 2, 1, 1, 1, 1, 1, 1, 2, 1,
	1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1, 1, 1, 2, 1,
	1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1, 1, 1, 2, 1,
	1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1, 1, 1, 2, 1,
	1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1, 1, 1, 2, 1,
	2, 3, 3, 4, 3, 4, 2, 4, 2, 4, 3, 0, 3, 6, 2, 4,
	2, 3, 3, 0, 3, 4, 2, 4, 2, 4, 3, 0, 3, 0, 2, 4,
	3, 3, 2, 0, 0, 4, 2, 4, 4, 1, 4, 0, 0, 0, 2, 4,
	3, 3, 2, 1, 0, 4, 2, 4, 3, 2, 4, 1, 0, 0, 2, 4,
}

// Initialize list of basic CPU instructions.
func initInstructionList(cpu *CPU) {
	instructions = make([]func(), 0x100)

	instructions[0x00] = nop                                                           // NOP
	instructions[0x01] = func() { cpu.Registers.SetBC(cpu.Fetch16()) }                 // LD BC,nn
	instructions[0x02] = func() { cpu.MMU.Write(cpu.Registers.BC(), cpu.Registers.A) } // LD (BC),A
	instructions[0x03] = func() { incNN(cpu, cpu.Registers.BC, cpu.Registers.SetBC) }  // INC BC
	instructions[0x04] = func() { incN(cpu, &cpu.Registers.B) }                        // INC B
	instructions[0x05] = func() { decN(cpu, &cpu.Registers.B) }                        // DEC B
	instructions[0x06] = func() { cpu.Registers.B = cpu.Fetch() }                      // LD B,n
	instructions[0x07] = func() { rlca(cpu) }                                          // RLCA
	instructions[0x08] = func() { ldNNSP(cpu) }                                        // LD (nn),SP
	instructions[0x09] = func() { addHL(cpu, cpu.Registers.BC()) }                     // ADD HL,BC
	instructions[0x0a] = func() { cpu.Registers.A = cpu.MMU.Read(cpu.Registers.BC()) } // LD A,(BC)
	instructions[0x0b] = func() { decNN(cpu, cpu.Registers.BC, cpu.Registers.SetBC) }  // DEC BC
	instructions[0x0c] = func() { incN(cpu, &cpu.Registers.C) }                        // INC C
	instructions[0x0d] = func() { decN(cpu, &cpu.Registers.C) }                        // DEC C
	instructions[0x0e] = func() { cpu.Registers.C = cpu.Fetch() }                      // LD C,n
	instructions[0x0f] = func() { rrca(cpu) }                                          // RRCA

	instructions[0x10] = func() { stop(cpu) }                                          // STOP
	instructions[0x11] = func() { cpu.Registers.SetDE(cpu.Fetch16()) }                 // LD DE,nn
	instructions[0x12] = func() { cpu.MMU.Write(cpu.Registers.DE(), cpu.Registers.A) } // LD (DE),A
	instructions[0x13] = func() { incNN(cpu, cpu.Registers.DE, cpu.Registers.SetDE) }  // INC DE
	instructions[0x14] = func() { incN(cpu, &cpu.Registers.D) }                        // INC D
	instructions[0x15] = func() { decN(cpu, &cpu.Registers.D) }                        // DEC D
	instructions[0x16] = func() { cpu.Registers.D = cpu.Fetch() }                      // LD D,n
	instructions[0x17] = func() { rla(cpu) }                                           // RLA
	instructions[0x18] = func() { jr(cpu, int8(cpu.Fetch())) }                         // JR n
	instructions[0x19] = func() { addHL(cpu, cpu.Registers.DE()) }                     // ADD HL,DE
	instructions[0x1a] = func() { cpu.Registers.A = cpu.MMU.Read(cpu.Registers.DE()) } // LD A,(DE)
	instructions[0x1b] = func() { decNN(cpu, cpu.Registers.DE, cpu.Registers.SetDE) }  // DEC DE
	instructions[0x1c] = func() { incN(cpu, &cpu.Registers.E) }                        // INC E
	instructions[0x1d] = func() { decN(cpu, &cpu.Registers.E) }                        // DEC E
	instructions[0x1e] = func() { cpu.Registers.E = cpu.Fetch() }                      // LD E,n
	instructions[0x1f] = func() { rra(cpu) }                                           // RRA

	instructions[0x20] = func() { jrCC(cpu, !cpu.Zero(), int8(cpu.Fetch())) }         // JP NZ,*
	instructions[0x21] = func() { cpu.Registers.SetHL(cpu.Fetch16()) }                // LD HL,nn
	instructions[0x22] = func() { ldiHLA(cpu) }                                       // LDI (HL),A
	instructions[0x23] = func() { incNN(cpu, cpu.Registers.HL, cpu.Registers.SetHL) } // INC HL
	instructions[0x24] = func() { incN(cpu, &cpu.Registers.H) }                       // INC H
	instructions[0x25] = func() { decN(cpu, &cpu.Registers.H) }                       // DEC H
	instructions[0x26] = func() { cpu.Registers.H = cpu.Fetch() }                     // LD H,n
	instructions[0x27] = func() { daa(cpu) }                                          // DAA
	instructions[0x28] = func() { jrCC(cpu, cpu.Zero(), int8(cpu.Fetch())) }          // JP Z,*
	instructions[0x29] = func() { addHL(cpu, cpu.Registers.HL()) }                    // ADD HL,HL
	instructions[0x2a] = func() { ldiAHL(cpu) }                                       // LDI A,(HL)
	instructions[0x2b] = func() { decNN(cpu, cpu.Registers.HL, cpu.Registers.SetHL) } // DEC HL
	instructions[0x2c] = func() { incN(cpu, &cpu.Registers.L) }                       // INC L
	instructions[0x2d] = func() { decN(cpu, &cpu.Registers.L) }                       // DEC L
	instructions[0x2e] = func() { cpu.Registers.L = cpu.Fetch() }                     // LD L,n
	instructions[0x2f] = func() { cpl(cpu) }                                          // CPL

	instructions[0x30] = func() { jrCC(cpu, !cpu.Carry(), int8(cpu.Fetch())) }     // JP NC,*
	instructions[0x31] = func() { cpu.SP = cpu.Fetch16() }                         // LD SP,nn
	instructions[0x32] = func() { lddHLA(cpu) }                                    // LDD (HL),A
	instructions[0x33] = func() { cpu.SP++ }                                       // INC SP
	instructions[0x34] = func() { incHL(cpu, cpu.Registers.HL()) }                 // INC (HL)
	instructions[0x35] = func() { decHL(cpu, cpu.Registers.HL()) }                 // DEC (HL)
	instructions[0x36] = func() { cpu.MMU.Write(cpu.Registers.HL(), cpu.Fetch()) } // LD (HL),n
	instructions[0x37] = func() { scf(cpu) }                                       // SCF
	instructions[0x38] = func() { jrCC(cpu, cpu.Carry(), int8(cpu.Fetch())) }      // JP C,*
	instructions[0x39] = func() { addHL(cpu, cpu.SP) }                             // ADD HL,SP
	instructions[0x3a] = func() { lddAHL(cpu) }                                    // LDD A,(HL)
	instructions[0x3b] = func() { cpu.SP-- }                                       // DEC SP
	instructions[0x3c] = func() { incN(cpu, &cpu.Registers.A) }                    // INC A
	instructions[0x3d] = func() { decN(cpu, &cpu.Registers.A) }                    // DEC A
	instructions[0x3e] = func() { cpu.Registers.A = cpu.Fetch() }                  // LD A,#
	instructions[0x3f] = func() { ccf(cpu) }                                       // CCF

	instructions[0x40] = func() {}                                                     // LD B,B
	instructions[0x41] = func() { cpu.Registers.B = cpu.Registers.C }                  // LD B,C
	instructions[0x42] = func() { cpu.Registers.B = cpu.Registers.D }                  // LD B,D
	instructions[0x43] = func() { cpu.Registers.B = cpu.Registers.E }                  // LD B,E
	instructions[0x44] = func() { cpu.Registers.B = cpu.Registers.H }                  // LD B,H
	instructions[0x45] = func() { cpu.Registers.B = cpu.Registers.L }                  // LD B,L
	instructions[0x46] = func() { cpu.Registers.B = cpu.MMU.Read(cpu.Registers.HL()) } // LD B,(HL)
	instructions[0x47] = func() { cpu.Registers.B = cpu.Registers.A }                  // LD B,A
	instructions[0x48] = func() { cpu.Registers.C = cpu.Registers.B }                  // LD C,B
	instructions[0x49] = func() {}                                                     // LD C,C
	instructions[0x4a] = func() { cpu.Registers.C = cpu.Registers.D }                  // LD C,D
	instructions[0x4b] = func() { cpu.Registers.C = cpu.Registers.E }                  // LD C,E
	instructions[0x4c] = func() { cpu.Registers.C = cpu.Registers.H }                  // LD C,H
	instructions[0x4d] = func() { cpu.Registers.C = cpu.Registers.L }                  // LD C,L
	instructions[0x4e] = func() { cpu.Registers.C = cpu.MMU.Read(cpu.Registers.HL()) } // LD C,(HL)
	instructions[0x4f] = func() { cpu.Registers.C = cpu.Registers.A }                  // LD C,A

	instructions[0x50] = func() { cpu.Registers.D = cpu.Registers.B }                  // LD D,B
	instructions[0x51] = func() { cpu.Registers.D = cpu.Registers.C }                  // LD D,C
	instructions[0x52] = func() {}                                                     // LD D,D
	instructions[0x53] = func() { cpu.Registers.D = cpu.Registers.E }                  // LD D,E
	instructions[0x54] = func() { cpu.Registers.D = cpu.Registers.H }                  // LD D,H
	instructions[0x55] = func() { cpu.Registers.D = cpu.Registers.L }                  // LD D,L
	instructions[0x56] = func() { cpu.Registers.D = cpu.MMU.Read(cpu.Registers.HL()) } // LD D,(HL)
	instructions[0x57] = func() { cpu.Registers.D = cpu.Registers.A }                  // LD D,A
	instructions[0x58] = func() { cpu.Registers.E = cpu.Registers.B }                  // LD E,B
	instructions[0x59] = func() { cpu.Registers.E = cpu.Registers.C }                  // LD E,C
	instructions[0x5a] = func() { cpu.Registers.E = cpu.Registers.D }                  // LD E,D
	instructions[0x5b] = func() {}                                                     // LD E,E
	instructions[0x5c] = func() { cpu.Registers.E = cpu.Registers.H }                  // LD E,H
	instructions[0x5d] = func() { cpu.Registers.E = cpu.Registers.L }                  // LD E,L
	instructions[0x5e] = func() { cpu.Registers.E = cpu.MMU.Read(cpu.Registers.HL()) } // LD E,(HL)
	instructions[0x5f] = func() { cpu.Registers.E = cpu.Registers.A }                  // LD E,A

	instructions[0x60] = func() { cpu.Registers.H = cpu.Registers.B }                  // LD H,B
	instructions[0x61] = func() { cpu.Registers.H = cpu.Registers.C }                  // LD H,C
	instructions[0x62] = func() { cpu.Registers.H = cpu.Registers.D }                  // LD H,D
	instructions[0x63] = func() { cpu.Registers.H = cpu.Registers.E }                  // LD H,E
	instructions[0x64] = func() {}                                                     // LD H,H
	instructions[0x65] = func() { cpu.Registers.H = cpu.Registers.L }                  // LD H,L
	instructions[0x66] = func() { cpu.Registers.H = cpu.MMU.Read(cpu.Registers.HL()) } // LD H,(HL)
	instructions[0x67] = func() { cpu.Registers.H = cpu.Registers.A }                  // LD H,A
	instructions[0x68] = func() { cpu.Registers.L = cpu.Registers.B }                  // LD L,B
	instructions[0x69] = func() { cpu.Registers.L = cpu.Registers.C }                  // LD L,C
	instructions[0x6a] = func() { cpu.Registers.L = cpu.Registers.D }                  // LD L,D
	instructions[0x6b] = func() { cpu.Registers.L = cpu.Registers.E }                  // LD L,E
	instructions[0x6c] = func() { cpu.Registers.L = cpu.Registers.H }                  // LD L,H
	instructions[0x6d] = func() {}                                                     // LD L,L
	instructions[0x6e] = func() { cpu.Registers.L = cpu.MMU.Read(cpu.Registers.HL()) } // LD L,(HL)
	instructions[0x6f] = func() { cpu.Registers.L = cpu.Registers.A }                  // LD L,A

	instructions[0x70] = func() { cpu.MMU.Write(cpu.Registers.HL(), cpu.Registers.B) } // LD (HL),B
	instructions[0x71] = func() { cpu.MMU.Write(cpu.Registers.HL(), cpu.Registers.C) } // LD (HL),C
	instructions[0x72] = func() { cpu.MMU.Write(cpu.Registers.HL(), cpu.Registers.D) } // LD (HL),D
	instructions[0x73] = func() { cpu.MMU.Write(cpu.Registers.HL(), cpu.Registers.E) } // LD (HL),E
	instructions[0x74] = func() { cpu.MMU.Write(cpu.Registers.HL(), cpu.Registers.H) } // LD (HL),H
	instructions[0x75] = func() { cpu.MMU.Write(cpu.Registers.HL(), cpu.Registers.L) } // LD (HL),L
	instructions[0x76] = func() { halt(cpu) }                                          // HALT
	instructions[0x77] = func() { cpu.MMU.Write(cpu.Registers.HL(), cpu.Registers.A) } // LD (HL),A
	instructions[0x78] = func() { cpu.Registers.A = cpu.Registers.B }                  // LD A,B
	instructions[0x79] = func() { cpu.Registers.A = cpu.Registers.C }                  // LD A,C
	instructions[0x7a] = func() { cpu.Registers.A = cpu.Registers.D }                  // LD A,D
	instructions[0x7b] = func() { cpu.Registers.A = cpu.Registers.E }                  // LD A,E
	instructions[0x7c] = func() { cpu.Registers.A = cpu.Registers.H }                  // LD A,H
	instructions[0x7d] = func() { cpu.Registers.A = cpu.Registers.L }                  // LD A,L
	instructions[0x7e] = func() { cpu.Registers.A = cpu.MMU.Read(cpu.Registers.HL()) } // LD A,(HL)
	instructions[0x7f] = func() {}                                                     // LD A,A

	instructions[0x80] = func() { add(cpu, cpu.Registers.B) }                  // ADD A,B
	instructions[0x81] = func() { add(cpu, cpu.Registers.C) }                  // ADD A,C
	instructions[0x82] = func() { add(cpu, cpu.Registers.D) }                  // ADD A,D
	instructions[0x83] = func() { add(cpu, cpu.Registers.E) }                  // ADD A,E
	instructions[0x84] = func() { add(cpu, cpu.Registers.H) }                  // ADD A,H
	instructions[0x85] = func() { add(cpu, cpu.Registers.L) }                  // ADD A,L
	instructions[0x86] = func() { add(cpu, cpu.MMU.Read(cpu.Registers.HL())) } // ADD A,(HL)
	instructions[0x87] = func() { add(cpu, cpu.Registers.A) }                  // ADD A,A
	instructions[0x88] = func() { adc(cpu, cpu.Registers.B) }                  // ADC A,B
	instructions[0x89] = func() { adc(cpu, cpu.Registers.C) }                  // ADC A,C
	instructions[0x8a] = func() { adc(cpu, cpu.Registers.D) }                  // ADC A,D
	instructions[0x8b] = func() { adc(cpu, cpu.Registers.E) }                  // ADC A,E
	instructions[0x8c] = func() { adc(cpu, cpu.Registers.H) }                  // ADC A,H
	instructions[0x8d] = func() { adc(cpu, cpu.Registers.L) }                  // ADC A,L
	instructions[0x8e] = func() { adc(cpu, cpu.MMU.Read(cpu.Registers.HL())) } // ADC A,(HL)
	instructions[0x8f] = func() { adc(cpu, cpu.Registers.A) }                  // ADC A,A

	instructions[0x90] = func() { sub(cpu, cpu.Registers.B) }                  // SUB A,B
	instructions[0x91] = func() { sub(cpu, cpu.Registers.C) }                  // SUB A,C
	instructions[0x92] = func() { sub(cpu, cpu.Registers.D) }                  // SUB A,D
	instructions[0x93] = func() { sub(cpu, cpu.Registers.E) }                  // SUB A,E
	instructions[0x94] = func() { sub(cpu, cpu.Registers.H) }                  // SUB A,H
	instructions[0x95] = func() { sub(cpu, cpu.Registers.L) }                  // SUB A,L
	instructions[0x96] = func() { sub(cpu, cpu.MMU.Read(cpu.Registers.HL())) } // SUB A,(HL)
	instructions[0x97] = func() { sub(cpu, cpu.Registers.A) }                  // SUB A,A
	instructions[0x98] = func() { sbc(cpu, cpu.Registers.B) }                  // SBC A,B
	instructions[0x99] = func() { sbc(cpu, cpu.Registers.C) }                  // SBC A,C
	instructions[0x9a] = func() { sbc(cpu, cpu.Registers.D) }                  // SBC A,D
	instructions[0x9b] = func() { sbc(cpu, cpu.Registers.E) }                  // SBC A,E
	instructions[0x9c] = func() { sbc(cpu, cpu.Registers.H) }                  // SBC A,H
	instructions[0x9d] = func() { sbc(cpu, cpu.Registers.L) }                  // SBC A,L
	instructions[0x9e] = func() { sbc(cpu, cpu.MMU.Read(cpu.Registers.HL())) } // SBC A,(HL)
	instructions[0x9f] = func() { sbc(cpu, cpu.Registers.D) }                  // SBC A,A

	instructions[0xa0] = func() { and(cpu, cpu.Registers.B) }                  // AND B
	instructions[0xa1] = func() { and(cpu, cpu.Registers.C) }                  // AND C
	instructions[0xa2] = func() { and(cpu, cpu.Registers.D) }                  // AND D
	instructions[0xa3] = func() { and(cpu, cpu.Registers.E) }                  // AND E
	instructions[0xa4] = func() { and(cpu, cpu.Registers.H) }                  // AND H
	instructions[0xa5] = func() { and(cpu, cpu.Registers.L) }                  // AND L
	instructions[0xa6] = func() { and(cpu, cpu.MMU.Read(cpu.Registers.HL())) } // AND (HL)
	instructions[0xa7] = func() { and(cpu, cpu.Registers.A) }                  // AND A
	instructions[0xa8] = func() { xor(cpu, cpu.Registers.B) }                  // XOR B
	instructions[0xa9] = func() { xor(cpu, cpu.Registers.C) }                  // XOR C
	instructions[0xaa] = func() { xor(cpu, cpu.Registers.D) }                  // XOR D
	instructions[0xab] = func() { xor(cpu, cpu.Registers.E) }                  // XOR E
	instructions[0xac] = func() { xor(cpu, cpu.Registers.H) }                  // XOR H
	instructions[0xad] = func() { xor(cpu, cpu.Registers.L) }                  // XOR L
	instructions[0xae] = func() { xor(cpu, cpu.MMU.Read(cpu.Registers.HL())) } // XOR (HL)
	instructions[0xaf] = func() { xor(cpu, cpu.Registers.A) }                  // XOR A

	instructions[0xb0] = func() { or(cpu, cpu.Registers.B) }                  // OR B
	instructions[0xb1] = func() { or(cpu, cpu.Registers.C) }                  // OR C
	instructions[0xb2] = func() { or(cpu, cpu.Registers.D) }                  // OR D
	instructions[0xb3] = func() { or(cpu, cpu.Registers.E) }                  // OR E
	instructions[0xb4] = func() { or(cpu, cpu.Registers.H) }                  // OR H
	instructions[0xb5] = func() { or(cpu, cpu.Registers.L) }                  // OR L
	instructions[0xb6] = func() { or(cpu, cpu.MMU.Read(cpu.Registers.HL())) } // OR (HL)
	instructions[0xb7] = func() { or(cpu, cpu.Registers.A) }                  // OR A
	instructions[0xb8] = func() { cp(cpu, cpu.Registers.B) }                  // CP B
	instructions[0xb9] = func() { cp(cpu, cpu.Registers.C) }                  // CP C
	instructions[0xba] = func() { cp(cpu, cpu.Registers.D) }                  // CP D
	instructions[0xbb] = func() { cp(cpu, cpu.Registers.E) }                  // CP E
	instructions[0xbc] = func() { cp(cpu, cpu.Registers.H) }                  // CP H
	instructions[0xbd] = func() { cp(cpu, cpu.Registers.L) }                  // CP L
	instructions[0xbe] = func() { cp(cpu, cpu.MMU.Read(cpu.Registers.HL())) } // CP (HL)
	instructions[0xbf] = func() { cp(cpu, cpu.Registers.A) }                  // CP A

	instructions[0xc0] = func() { retCC(cpu, !cpu.Zero()) }                 // RET NZ
	instructions[0xc1] = func() { cpu.Registers.SetBC(popNN(cpu)) }         // POP BC
	instructions[0xc2] = func() { jpCC(cpu, !cpu.Zero(), cpu.Fetch16()) }   // JP NZ,nn
	instructions[0xc3] = func() { cpu.PC = cpu.Fetch16() }                  // JP nn
	instructions[0xc4] = func() { callCC(cpu, !cpu.Zero(), cpu.Fetch16()) } // CALL NZ,nn
	instructions[0xc5] = func() { pushNN(cpu, cpu.Registers.BC()) }         // PUSH BC
	instructions[0xc6] = func() { add(cpu, cpu.Fetch()) }                   // ADD A,#
	instructions[0xc7] = func() { rst(cpu, 0x00) }                          // RST 00H
	instructions[0xc8] = func() { retCC(cpu, cpu.Zero()) }                  // RET Z
	instructions[0xc9] = func() { cpu.PC = popNN(cpu) }                     // RET
	instructions[0xca] = func() { jpCC(cpu, cpu.Zero(), cpu.Fetch16()) }    // JP Z,nn
	instructions[0xcb] = nop                                                // Extended instructions
	instructions[0xcc] = func() { callCC(cpu, cpu.Zero(), cpu.Fetch16()) }  // CALL Z,nn
	instructions[0xcd] = func() { call(cpu, cpu.Fetch16()) }                // CALL nn
	instructions[0xce] = func() { adc(cpu, cpu.Fetch()) }                   // ADC A,#
	instructions[0xcf] = func() { rst(cpu, 0x08) }                          // RST 08H

	instructions[0xd0] = func() { retCC(cpu, !cpu.Carry()) }                 // RET NC
	instructions[0xd1] = func() { cpu.Registers.SetDE(popNN(cpu)) }          // POP DE
	instructions[0xd2] = func() { jpCC(cpu, !cpu.Carry(), cpu.Fetch16()) }   // JP NC,nn
	instructions[0xd3] = xx                                                  // XX
	instructions[0xd4] = func() { callCC(cpu, !cpu.Carry(), cpu.Fetch16()) } // CALL NC,nn
	instructions[0xd5] = func() { pushNN(cpu, cpu.Registers.DE()) }          // PUSH DE
	instructions[0xd6] = func() { sub(cpu, cpu.Fetch()) }                    // SUB A,#
	instructions[0xd7] = func() { rst(cpu, 0x10) }                           // RST 10H
	instructions[0xd8] = func() { retCC(cpu, cpu.Carry()) }                  // RET C
	instructions[0xd9] = func() { reti(cpu) }                                // RETI
	instructions[0xda] = func() { jpCC(cpu, cpu.Carry(), cpu.Fetch16()) }    // JP C,nn
	instructions[0xdb] = xx                                                  // XX
	instructions[0xdc] = func() { callCC(cpu, cpu.Carry(), cpu.Fetch16()) }  // CALL C,nn
	instructions[0xdd] = xx                                                  // XX
	instructions[0xde] = func() { sbc(cpu, cpu.Fetch()) }                    // SBC A,#
	instructions[0xdf] = func() { rst(cpu, 0x18) }                           // RST 18H

	instructions[0xe0] = func() { cpu.MMU.Write(0xff00+uint16(cpu.Fetch()), cpu.Registers.A) }     // LDH (n),A
	instructions[0xe1] = func() { cpu.Registers.SetHL(popNN(cpu)) }                                // POP HL
	instructions[0xe2] = func() { cpu.MMU.Write(0xff00+uint16(cpu.Registers.C), cpu.Registers.A) } // LD (C),A
	instructions[0xe3] = xx                                                                        // XX
	instructions[0xe4] = xx                                                                        // XX
	instructions[0xe5] = func() { pushNN(cpu, cpu.Registers.HL()) }                                // PUSH HL
	instructions[0xe6] = func() { and(cpu, cpu.Fetch()) }                                          // AND #
	instructions[0xe7] = func() { rst(cpu, 0x20) }                                                 // RST 20H
	instructions[0xe8] = func() { addSP(cpu, int8(cpu.Fetch())) }                                  // ADD SP,n
	instructions[0xe9] = func() { cpu.PC = cpu.Registers.HL() }                                    // JP (HL)
	instructions[0xea] = func() { cpu.MMU.Write(cpu.Fetch16(), cpu.Registers.A) }                  // LD (nn),A
	instructions[0xeb] = xx                                                                        // XX
	instructions[0xec] = xx                                                                        // XX
	instructions[0xed] = xx                                                                        // XX
	instructions[0xee] = func() { xor(cpu, cpu.Fetch()) }                                          // XOR #
	instructions[0xef] = func() { rst(cpu, 0x28) }                                                 // RST 28H

	instructions[0xf0] = func() { cpu.Registers.A = cpu.MMU.Read(0xff00 + uint16(cpu.Fetch())) }     // LDH A,(n)
	instructions[0xf1] = func() { cpu.Registers.SetAF(popNN(cpu)) }                                  // POP AF
	instructions[0xf2] = func() { cpu.Registers.A = cpu.MMU.Read(0xff00 + uint16(cpu.Registers.C)) } // LD A,(C)
	instructions[0xf3] = func() { DisableInterrupts() }                                              // DI
	instructions[0xf4] = xx                                                                          // XX
	instructions[0xf5] = func() { pushNN(cpu, cpu.Registers.AF()) }                                  // PUSH AF
	instructions[0xf6] = func() { or(cpu, cpu.Fetch()) }                                             // OR #
	instructions[0xf7] = func() { rst(cpu, 0x30) }                                                   // RST 30H
	instructions[0xf8] = func() { ldHLSPPlusN(cpu, int8(cpu.Fetch())) }                              // LD HL,SP+n
	instructions[0xf9] = func() { cpu.SP = cpu.Registers.HL() }                                      // LD SP,HL
	instructions[0xfa] = func() { cpu.Registers.A = cpu.MMU.Read(cpu.Fetch16()) }                    // LD A,(nn)
	instructions[0xfb] = func() { EnableInterrupts() }                                               // EI
	instructions[0xfc] = xx                                                                          // XX
	instructions[0xfd] = xx                                                                          // XX
	instructions[0xfe] = func() { cp(cpu, cpu.Fetch()) }                                             // CP #
	instructions[0xff] = func() { rst(cpu, 0x38) }                                                   // RST 38H
}

// NOP -- No operation.
func nop() {
	log.Printf("NOP")
}

// XX -- Operation not supported.
func xx() {
	log.Fatalf("Operation not supported")
}

/* 8-bit ALU */

// INC n -- Increment register n.
func incN(cpu *CPU, n *byte) {
	halfCarry := (*n&0xf)+1 > 0xf
	*n++
	cpu.SetNegative(false)
	cpu.SetZero(*n == 0)
	cpu.SetHalfCarry(halfCarry)
}

// INC (HL) -- Increment value at address HL.
func incHL(cpu *CPU, address uint16) {
	value := cpu.MMU.Read(address)
	halfCarry := (value&0xf)+1 > 0xf
	cpu.MMU.Write(address, value+1)
	cpu.SetNegative(false)
	cpu.SetZero(value == 0)
	cpu.SetHalfCarry(halfCarry)
}

// DEC n -- Decrement register n.
func decN(cpu *CPU, n *byte) {
	halfCarry := *n&0x0f == 0
	*n--
	cpu.SetZero(*n == 0)
	cpu.SetNegative(true)
	cpu.SetHalfCarry(halfCarry)
}

// DEC (HL) -- Decrement value at address HL.
func decHL(cpu *CPU, address uint16) {
	value := cpu.MMU.Read(address)
	halfCarry := value&0x0f == 0
	cpu.MMU.Write(address, value+1)
	cpu.SetZero(value == 0)
	cpu.SetNegative(true)
	cpu.SetHalfCarry(halfCarry)
}

// ADD A,n -- Add n to A.
func add(cpu *CPU, n byte) {
	sum := int16(cpu.Registers.A) + int16(n)
	halfCarry := (((cpu.Registers.A & 0xf) + (n & 0xf)) & 0x10) == 0x10
	cpu.Registers.A = byte(sum)

	cpu.SetZero(cpu.Registers.A == 0)
	cpu.SetNegative(false)
	cpu.SetHalfCarry(halfCarry)
	cpu.SetCarry(sum > 0xff)
}

// ADC A,n -- Add n + Carry flag to A.
func adc(cpu *CPU, n byte) {
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

// SUB n -- Subtract n from A.
func sub(cpu *CPU, n byte) {
	diff := int16(cpu.Registers.A) + int16(n)
	halfCarry := ((cpu.Registers.A & 0xf) - (n & 0xf)) < 0
	cpu.Registers.A = byte(diff)

	cpu.SetZero(cpu.Registers.A == 0)
	cpu.SetNegative(true)
	cpu.SetHalfCarry(halfCarry)
	cpu.SetCarry(diff < 0)
}

// SBC n + Carry flag from A.
func sbc(cpu *CPU, n byte) {
	diff := int16(cpu.Registers.A) + int16(n)
	halfCarry := ((cpu.Registers.A & 0xf) - (n & 0xf)) < 0
	cpu.Registers.A = byte(diff)

	cpu.SetZero(cpu.Registers.A == 0)
	cpu.SetNegative(true)
	cpu.SetHalfCarry(halfCarry)
	cpu.SetCarry(diff < 0)
}

// Logically AND n with A, result in A.
func and(cpu *CPU, n byte) {
	cpu.Registers.A &= n

	cpu.SetZero(cpu.Registers.A == 0)
	cpu.SetNegative(false)
	cpu.SetHalfCarry(true)
	cpu.SetCarry(false)
}

// Logical OR n with register A, result in A.
func or(cpu *CPU, n byte) {
	cpu.Registers.A |= n

	cpu.SetZero(cpu.Registers.A == 0)
	cpu.SetNegative(false)
	cpu.SetHalfCarry(false)
	cpu.SetCarry(false)
}

// Logical exclusive OR n with register A, result in A.
func xor(cpu *CPU, n byte) {
	cpu.Registers.A ^= n

	cpu.SetZero(cpu.Registers.A == 0)
	cpu.SetNegative(false)
	cpu.SetHalfCarry(false)
	cpu.SetCarry(false)
}

// Compare A with n. This is basically an A - n subtraction instruction but the results are thrown away.
func cp(cpu *CPU, n byte) {
	result := cpu.Registers.A - n

	cpu.SetZero(result == 0)
	cpu.SetNegative(true)
	cpu.SetHalfCarry((cpu.Registers.A & 0x0f) > (n & 0x0f))
	cpu.SetCarry(cpu.Registers.A > n)
}

/* 16-bit arithmetic */

// INC nn -- Increment register nn.
func incNN(cpu *CPU, getNN func() uint16, setNN func(uint16)) {
	setNN(getNN() + 1)
}

// DEC nn -- Decrement register nn.
func decNN(cpu *CPU, getNN func() uint16, setNN func(uint16)) {
	setNN(getNN() - 1)
}

// ADD HL,n -- Add n to HL.
func addHL(cpu *CPU, n uint16) {
	hl := cpu.Registers.HL()
	sum := int32(hl) + int32(n)
	cpu.Registers.SetHL(uint16(sum))
	cpu.SetNegative(false)
	cpu.SetHalfCarry(int32(hl&0xfff) > (sum & 0xfff))
	cpu.SetCarry(sum > 0xffff)
}

// ADD SP,n -- Add signed byte n to Stack Pointer (SP).
func addSP(cpu *CPU, n int8) {
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

// LDD (HL),A -- Put A into memory address HL. Decrement HL.
func lddHLA(cpu *CPU) {
	cpu.MMU.Write(cpu.Registers.HL(), cpu.Registers.A)
	cpu.Registers.SetHL(cpu.Registers.HL() - 1)
}

// LDD A,(HL) -- Put value at address HL into A. Decrement HL.
func lddAHL(cpu *CPU) {
	cpu.Registers.A = cpu.MMU.Read(cpu.Registers.HL())
	cpu.Registers.SetHL(cpu.Registers.HL() - 1)
}

// LDI (HL),A -- Put A into memory address HL. Increment HL.
func ldiHLA(cpu *CPU) {
	cpu.MMU.Write(cpu.Registers.HL(), cpu.Registers.A)
	cpu.Registers.SetHL(cpu.Registers.HL() + 1)
}

// LDI A,(HL) -- Put value at address HL into A. Increment HL.
func ldiAHL(cpu *CPU) {
	cpu.Registers.A = cpu.MMU.Read(cpu.Registers.HL())
	cpu.Registers.SetHL(cpu.Registers.HL() + 1)
}

/* 16-bit loads */

// LD HL,SP+n -- Put SP + n effective address into HL
func ldHLSPPlusN(cpu *CPU, n int8) {
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
func ldNNSP(cpu *CPU) {
	address := cpu.Fetch16()
	cpu.MMU.Write(address, byte(cpu.SP&0xff))
	cpu.MMU.Write(address+1, byte(cpu.SP>>8))
}

// PUSH nn -- Push register pair nn onto stack.
// Decrement Stack Pointer (SP) twice
func pushNN(cpu *CPU, address uint16) {
	cpu.MMU.Write(cpu.SP-1, byte(uint16(address&0xff00)>>8))
	cpu.MMU.Write(cpu.SP-2, byte(address&0xff))
	cpu.SP -= 2
}

// POP nn --  Pop two bytes off stack into register pair nn.
// Increment Stack Pointer (SP) twice.
func popNN(cpu *CPU) uint16 {
	byte1 := uint16(cpu.MMU.Read(cpu.SP))
	byte2 := uint16(cpu.MMU.Read(cpu.SP+1)) << 8
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

// RLA -- Rotate A left through Carry flag.
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

// RRCA -- Rotate A right. Old bit 0 to Carry flag.
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

// DAA --  Decimal adjust register A.
// This instruction adjusts register A so that the correct
// representation of Binary Coded Decimal (BCD) is obtained.
func daa(cpu *CPU) {
	a := cpu.Registers.A
	if !cpu.Negative() {
		if cpu.Carry() || a > 0x99 {
			a += 0x60
			cpu.SetCarry(true)
		}
		if cpu.HalfCarry() || (a&0x0f) > 0x09 {
			a += 0x6
		}
	} else {
		if cpu.Carry() {
			a -= 0x60
		}
		if cpu.HalfCarry() {
			a -= 0x6
		}
	}
	cpu.Registers.A = a
	cpu.SetZero(a == 0)
	cpu.SetHalfCarry(false)
}

// CPL -- Complement A register.
func cpl(cpu *CPU) {
	cpu.Registers.A = 0xff ^ cpu.Registers.A
	cpu.SetNegative(true)
	cpu.SetHalfCarry(true)
}

// CCF -- Complement carry flag.
func ccf(cpu *CPU) {
	cpu.SetNegative(false)
	cpu.SetHalfCarry(false)
	cpu.SetCarry(!cpu.Carry())
}

// SCF -- Set carry flag.
func scf(cpu *CPU) {
	cpu.SetNegative(false)
	cpu.SetHalfCarry(false)
	cpu.SetCarry(true)
}

// HALT -- Power down CPU until an interrupt occurs.
func halt(cpu *CPU) {
	log.Fatalln("TODO: HALT")
}

// STOP -- Halt CPU & LCD display until button pressed.
func stop(cpu *CPU) {
	log.Fatalln("TODO: STOP")
}

/* Jumps */

// JP cc,nn -- Jump to address n if condition is true.
func jpCC(cpu *CPU, condition bool, n uint16) {
	if condition {
		cpu.PC = n
	}
}

// JR n -- Add n to current address and jump to it.
func jr(cpu *CPU, n int8) {
	address := int32(cpu.PC) + int32(n)
	cpu.PC = uint16(address)
}

// JR cc,nn -- If  condition is true then add n to current address and jump to it.
func jrCC(cpu *CPU, condition bool, n int8) {
	if condition {
		address := int32(cpu.PC) + int32(n)
		cpu.PC = uint16(address)
	}
}

/* Calls */

// CALL nn -- Push address of next instruction onto stack and then jump to address nn.
func call(cpu *CPU, next uint16) {
	pushNN(cpu, cpu.PC)
	cpu.PC = next
}

// CALL cc,nn -- Call address n if condition is true.
func callCC(cpu *CPU, condition bool, next uint16) {
	if condition {
		call(cpu, next)
	}
}

/* Restarts */

// RST -- Push present address onto stack.
// Jump to address $0000 + n.
func rst(cpu *CPU, n byte) {
	pushNN(cpu, cpu.PC)
	cpu.PC = uint16(n)
}

/* Returns */

// RET cc -- Return if condition is true.
func retCC(cpu *CPU, condition bool) {
	if condition {
		cpu.PC = popNN(cpu)
	}
}

// RETI -- Pop two bytes from stack & jump to that address then enable interrupts.
func reti(cpu *CPU) {
	cpu.PC = popNN(cpu)
	EnableInterrupts()
}
