package hardware

// CBInstructions represents extended cpu instructions.
func CBInstructions(cpu *CPU) []func() {
	var instr = make([]func(), 0x100)

	instr[0x00] = func() { cpu.Registers.B = rlc(cpu, cpu.Registers.B) }                         // RLC B
	instr[0x01] = func() { cpu.Registers.C = rlc(cpu, cpu.Registers.C) }                         // RLC C
	instr[0x02] = func() { cpu.Registers.D = rlc(cpu, cpu.Registers.D) }                         // RLC D
	instr[0x03] = func() { cpu.Registers.E = rlc(cpu, cpu.Registers.E) }                         // RLC E
	instr[0x04] = func() { cpu.Registers.H = rlc(cpu, cpu.Registers.H) }                         // RLC H
	instr[0x05] = func() { cpu.Registers.L = rlc(cpu, cpu.Registers.L) }                         // RLC L
	instr[0x06] = func() { MemWrite(cpu.Registers.HL(), rlc(cpu, MemRead(cpu.Registers.HL()))) } // RLC (HL)
	instr[0x07] = func() { cpu.Registers.A = rlc(cpu, cpu.Registers.A) }                         // RLC A
	instr[0x08] = func() { cpu.Registers.B = rrc(cpu, cpu.Registers.B) }                         // RRC B
	instr[0x09] = func() { cpu.Registers.C = rrc(cpu, cpu.Registers.C) }                         // RRC C
	instr[0x0a] = func() { cpu.Registers.D = rrc(cpu, cpu.Registers.D) }                         // RRC D
	instr[0x0b] = func() { cpu.Registers.E = rrc(cpu, cpu.Registers.E) }                         // RRC E
	instr[0x0c] = func() { cpu.Registers.H = rrc(cpu, cpu.Registers.H) }                         // RRC H
	instr[0x0d] = func() { cpu.Registers.L = rrc(cpu, cpu.Registers.L) }                         // RRC L
	instr[0x0e] = func() { MemWrite(cpu.Registers.HL(), rrc(cpu, MemRead(cpu.Registers.HL()))) } // RRC (HL)
	instr[0x0f] = func() { cpu.Registers.A = rrc(cpu, cpu.Registers.A) }                         // RRC A

	instr[0x10] = func() { cpu.Registers.B = rl(cpu, cpu.Registers.B) }                         // RL B
	instr[0x11] = func() { cpu.Registers.C = rl(cpu, cpu.Registers.C) }                         // RL C
	instr[0x12] = func() { cpu.Registers.D = rl(cpu, cpu.Registers.D) }                         // RL D
	instr[0x13] = func() { cpu.Registers.E = rl(cpu, cpu.Registers.E) }                         // RL E
	instr[0x14] = func() { cpu.Registers.H = rl(cpu, cpu.Registers.H) }                         // RL H
	instr[0x15] = func() { cpu.Registers.L = rl(cpu, cpu.Registers.L) }                         // RL L
	instr[0x16] = func() { MemWrite(cpu.Registers.HL(), rl(cpu, MemRead(cpu.Registers.HL()))) } // RL (HL)
	instr[0x17] = func() { cpu.Registers.A = rl(cpu, cpu.Registers.A) }                         // RL A
	instr[0x18] = func() { cpu.Registers.B = rr(cpu, cpu.Registers.B) }                         // RR B
	instr[0x19] = func() { cpu.Registers.C = rr(cpu, cpu.Registers.C) }                         // RR C
	instr[0x1a] = func() { cpu.Registers.D = rr(cpu, cpu.Registers.D) }                         // RR D
	instr[0x1b] = func() { cpu.Registers.E = rr(cpu, cpu.Registers.E) }                         // RR E
	instr[0x1c] = func() { cpu.Registers.H = rr(cpu, cpu.Registers.H) }                         // RR H
	instr[0x1d] = func() { cpu.Registers.L = rr(cpu, cpu.Registers.L) }                         // RR L
	instr[0x1e] = func() { MemWrite(cpu.Registers.HL(), rr(cpu, MemRead(cpu.Registers.HL()))) } // RR (HL)
	instr[0x1f] = func() { cpu.Registers.A = rr(cpu, cpu.Registers.A) }                         // RR A

	instr[0x20] = func() { cpu.Registers.B = sla(cpu, cpu.Registers.B) }                         // SLA B
	instr[0x21] = func() { cpu.Registers.C = sla(cpu, cpu.Registers.C) }                         // SLA C
	instr[0x22] = func() { cpu.Registers.D = sla(cpu, cpu.Registers.D) }                         // SLA D
	instr[0x23] = func() { cpu.Registers.E = sla(cpu, cpu.Registers.E) }                         // SLA E
	instr[0x24] = func() { cpu.Registers.H = sla(cpu, cpu.Registers.H) }                         // SLA H
	instr[0x25] = func() { cpu.Registers.L = sla(cpu, cpu.Registers.L) }                         // SLA L
	instr[0x26] = func() { MemWrite(cpu.Registers.HL(), sla(cpu, MemRead(cpu.Registers.HL()))) } // SLA (HL)
	instr[0x27] = func() { cpu.Registers.A = sla(cpu, cpu.Registers.A) }                         // SLA A
	instr[0x28] = func() { cpu.Registers.B = sra(cpu, cpu.Registers.B) }                         // SRA B
	instr[0x29] = func() { cpu.Registers.C = sra(cpu, cpu.Registers.C) }                         // SRA C
	instr[0x2a] = func() { cpu.Registers.D = sra(cpu, cpu.Registers.D) }                         // SRA D
	instr[0x2b] = func() { cpu.Registers.E = sra(cpu, cpu.Registers.E) }                         // SRA E
	instr[0x2c] = func() { cpu.Registers.H = sra(cpu, cpu.Registers.H) }                         // SRA H
	instr[0x2d] = func() { cpu.Registers.L = sra(cpu, cpu.Registers.L) }                         // SRA L
	instr[0x2e] = func() { MemWrite(cpu.Registers.HL(), sra(cpu, MemRead(cpu.Registers.HL()))) } // SRA (HL)
	instr[0x2f] = func() { cpu.Registers.A = sra(cpu, cpu.Registers.A) }                         // SRA A

	instr[0x30] = func() { cpu.Registers.B = swap(cpu, cpu.Registers.B) }                         // SWAP B
	instr[0x31] = func() { cpu.Registers.C = swap(cpu, cpu.Registers.C) }                         // SWAP C
	instr[0x32] = func() { cpu.Registers.D = swap(cpu, cpu.Registers.D) }                         // SWAP D
	instr[0x33] = func() { cpu.Registers.E = swap(cpu, cpu.Registers.E) }                         // SWAP E
	instr[0x34] = func() { cpu.Registers.H = swap(cpu, cpu.Registers.H) }                         // SWAP H
	instr[0x35] = func() { cpu.Registers.L = swap(cpu, cpu.Registers.L) }                         // SWAP L
	instr[0x36] = func() { MemWrite(cpu.Registers.HL(), swap(cpu, MemRead(cpu.Registers.HL()))) } // SWAP (HL)
	instr[0x37] = func() { cpu.Registers.A = swap(cpu, cpu.Registers.A) }                         // SWAP A
	instr[0x38] = func() { cpu.Registers.B = srl(cpu, cpu.Registers.B) }                          // SRL B
	instr[0x39] = func() { cpu.Registers.C = srl(cpu, cpu.Registers.C) }                          // SRL C
	instr[0x3a] = func() { cpu.Registers.D = srl(cpu, cpu.Registers.D) }                          // SRL D
	instr[0x3b] = func() { cpu.Registers.E = srl(cpu, cpu.Registers.E) }                          // SRL E
	instr[0x3c] = func() { cpu.Registers.H = srl(cpu, cpu.Registers.H) }                          // SRL H
	instr[0x3d] = func() { cpu.Registers.L = srl(cpu, cpu.Registers.L) }                          // SRL L
	instr[0x3e] = func() { MemWrite(cpu.Registers.HL(), srl(cpu, MemRead(cpu.Registers.HL()))) }  // SRL (HL)
	instr[0x3f] = func() { cpu.Registers.A = srl(cpu, cpu.Registers.A) }                          // SRL A

	instr[0x40] = func() { bit(cpu, cpu.Registers.B, 0) }             // BIT 0,B
	instr[0x41] = func() { bit(cpu, cpu.Registers.C, 0) }             // BIT 0,C
	instr[0x42] = func() { bit(cpu, cpu.Registers.D, 0) }             // BIT 0,D
	instr[0x43] = func() { bit(cpu, cpu.Registers.E, 0) }             // BIT 0,E
	instr[0x44] = func() { bit(cpu, cpu.Registers.H, 0) }             // BIT 0,H
	instr[0x45] = func() { bit(cpu, cpu.Registers.L, 0) }             // BIT 0,L
	instr[0x46] = func() { bit(cpu, MemRead(cpu.Registers.HL()), 0) } // BIT 0,(HL)
	instr[0x47] = func() { bit(cpu, cpu.Registers.A, 0) }             // BIT 0,A
	instr[0x48] = func() { bit(cpu, cpu.Registers.B, 1) }             // BIT 1,B
	instr[0x49] = func() { bit(cpu, cpu.Registers.C, 1) }             // BIT 1,C
	instr[0x4a] = func() { bit(cpu, cpu.Registers.D, 1) }             // BIT 1,D
	instr[0x4b] = func() { bit(cpu, cpu.Registers.E, 1) }             // BIT 1,E
	instr[0x4c] = func() { bit(cpu, cpu.Registers.H, 1) }             // BIT 1,H
	instr[0x4d] = func() { bit(cpu, cpu.Registers.L, 1) }             // BIT 1,L
	instr[0x4e] = func() { bit(cpu, MemRead(cpu.Registers.HL()), 1) } // BIT 1,(HL)
	instr[0x4f] = func() { bit(cpu, cpu.Registers.A, 1) }             // BIT 1,A

	instr[0x50] = func() { bit(cpu, cpu.Registers.B, 2) }             // BIT 2,B
	instr[0x51] = func() { bit(cpu, cpu.Registers.C, 2) }             // BIT 2,C
	instr[0x52] = func() { bit(cpu, cpu.Registers.D, 2) }             // BIT 2,D
	instr[0x53] = func() { bit(cpu, cpu.Registers.E, 2) }             // BIT 2,E
	instr[0x54] = func() { bit(cpu, cpu.Registers.H, 2) }             // BIT 2,H
	instr[0x55] = func() { bit(cpu, cpu.Registers.L, 2) }             // BIT 2,L
	instr[0x56] = func() { bit(cpu, MemRead(cpu.Registers.HL()), 2) } // BIT 2,(HL)
	instr[0x57] = func() { bit(cpu, cpu.Registers.A, 2) }             // BIT 2,A
	instr[0x58] = func() { bit(cpu, cpu.Registers.B, 3) }             // BIT 3,B
	instr[0x59] = func() { bit(cpu, cpu.Registers.C, 3) }             // BIT 3,C
	instr[0x5a] = func() { bit(cpu, cpu.Registers.D, 3) }             // BIT 3,D
	instr[0x5b] = func() { bit(cpu, cpu.Registers.E, 3) }             // BIT 3,E
	instr[0x5c] = func() { bit(cpu, cpu.Registers.H, 3) }             // BIT 3,H
	instr[0x5d] = func() { bit(cpu, cpu.Registers.L, 3) }             // BIT 3,L
	instr[0x5e] = func() { bit(cpu, MemRead(cpu.Registers.HL()), 3) } // BIT 3,(HL)
	instr[0x5f] = func() { bit(cpu, cpu.Registers.A, 3) }             // BIT 3,A

	instr[0x60] = func() { bit(cpu, cpu.Registers.B, 4) }             // BIT 4,B
	instr[0x61] = func() { bit(cpu, cpu.Registers.C, 4) }             // BIT 4,C
	instr[0x62] = func() { bit(cpu, cpu.Registers.D, 4) }             // BIT 4,D
	instr[0x63] = func() { bit(cpu, cpu.Registers.E, 4) }             // BIT 4,E
	instr[0x64] = func() { bit(cpu, cpu.Registers.H, 4) }             // BIT 4,H
	instr[0x65] = func() { bit(cpu, cpu.Registers.L, 4) }             // BIT 4,L
	instr[0x66] = func() { bit(cpu, MemRead(cpu.Registers.HL()), 4) } // BIT 4,(HL)
	instr[0x67] = func() { bit(cpu, cpu.Registers.A, 4) }             // BIT 4,A
	instr[0x68] = func() { bit(cpu, cpu.Registers.B, 5) }             // BIT 5,B
	instr[0x69] = func() { bit(cpu, cpu.Registers.C, 5) }             // BIT 5,C
	instr[0x6a] = func() { bit(cpu, cpu.Registers.D, 5) }             // BIT 5,D
	instr[0x6b] = func() { bit(cpu, cpu.Registers.E, 5) }             // BIT 5,E
	instr[0x6c] = func() { bit(cpu, cpu.Registers.H, 5) }             // BIT 5,H
	instr[0x6d] = func() { bit(cpu, cpu.Registers.L, 5) }             // BIT 5,L
	instr[0x6e] = func() { bit(cpu, MemRead(cpu.Registers.HL()), 5) } // BIT 5,(HL)
	instr[0x6f] = func() { bit(cpu, cpu.Registers.A, 5) }             // BIT 5,A

	instr[0x70] = func() { bit(cpu, cpu.Registers.B, 6) }             // BIT 6,B
	instr[0x71] = func() { bit(cpu, cpu.Registers.C, 6) }             // BIT 6,C
	instr[0x72] = func() { bit(cpu, cpu.Registers.D, 6) }             // BIT 6,D
	instr[0x73] = func() { bit(cpu, cpu.Registers.E, 6) }             // BIT 6,E
	instr[0x74] = func() { bit(cpu, cpu.Registers.H, 6) }             // BIT 6,H
	instr[0x75] = func() { bit(cpu, cpu.Registers.L, 6) }             // BIT 6,L
	instr[0x76] = func() { bit(cpu, MemRead(cpu.Registers.HL()), 6) } // BIT 6,(HL)
	instr[0x77] = func() { bit(cpu, cpu.Registers.A, 6) }             // BIT 6,A
	instr[0x78] = func() { bit(cpu, cpu.Registers.B, 7) }             // BIT 7,B
	instr[0x79] = func() { bit(cpu, cpu.Registers.C, 7) }             // BIT 7,C
	instr[0x7a] = func() { bit(cpu, cpu.Registers.D, 7) }             // BIT 7,D
	instr[0x7b] = func() { bit(cpu, cpu.Registers.E, 7) }             // BIT 7,E
	instr[0x7c] = func() { bit(cpu, cpu.Registers.H, 7) }             // BIT 7,H
	instr[0x7d] = func() { bit(cpu, cpu.Registers.L, 7) }             // BIT 7,L
	instr[0x7e] = func() { bit(cpu, MemRead(cpu.Registers.HL()), 7) } // BIT 7,(HL)
	instr[0x7f] = func() { bit(cpu, cpu.Registers.A, 7) }             // BIT 7,A

	instr[0x80] = func() { cpu.Registers.B = res(cpu, cpu.Registers.B, 0) }                         // RES 0,B
	instr[0x81] = func() { cpu.Registers.C = res(cpu, cpu.Registers.C, 0) }                         // RES 0,C
	instr[0x82] = func() { cpu.Registers.D = res(cpu, cpu.Registers.D, 0) }                         // RES 0,D
	instr[0x83] = func() { cpu.Registers.E = res(cpu, cpu.Registers.E, 0) }                         // RES 0,E
	instr[0x84] = func() { cpu.Registers.H = res(cpu, cpu.Registers.H, 0) }                         // RES 0,H
	instr[0x85] = func() { cpu.Registers.L = res(cpu, cpu.Registers.L, 0) }                         // RES 0,L
	instr[0x86] = func() { MemWrite(cpu.Registers.HL(), res(cpu, MemRead(cpu.Registers.HL()), 0)) } // RES 0,(HL)
	instr[0x87] = func() { cpu.Registers.A = res(cpu, cpu.Registers.A, 0) }                         // RES 0,A
	instr[0x88] = func() { cpu.Registers.B = res(cpu, cpu.Registers.B, 1) }                         // RES 1,B
	instr[0x89] = func() { cpu.Registers.C = res(cpu, cpu.Registers.C, 1) }                         // RES 1,C
	instr[0x8a] = func() { cpu.Registers.D = res(cpu, cpu.Registers.D, 1) }                         // RES 1,D
	instr[0x8b] = func() { cpu.Registers.E = res(cpu, cpu.Registers.E, 1) }                         // RES 1,E
	instr[0x8c] = func() { cpu.Registers.H = res(cpu, cpu.Registers.H, 1) }                         // RES 1,H
	instr[0x8d] = func() { cpu.Registers.L = res(cpu, cpu.Registers.L, 1) }                         // RES 1,L
	instr[0x8e] = func() { MemWrite(cpu.Registers.HL(), res(cpu, MemRead(cpu.Registers.HL()), 1)) } // RES 1,(HL)
	instr[0x8f] = func() { cpu.Registers.A = res(cpu, cpu.Registers.A, 1) }                         // RES 1,A

	instr[0x90] = func() { cpu.Registers.B = res(cpu, cpu.Registers.B, 2) }                         // RES 2,B
	instr[0x91] = func() { cpu.Registers.C = res(cpu, cpu.Registers.C, 2) }                         // RES 2,C
	instr[0x92] = func() { cpu.Registers.D = res(cpu, cpu.Registers.D, 2) }                         // RES 2,D
	instr[0x93] = func() { cpu.Registers.E = res(cpu, cpu.Registers.E, 2) }                         // RES 2,E
	instr[0x94] = func() { cpu.Registers.H = res(cpu, cpu.Registers.H, 2) }                         // RES 2,H
	instr[0x95] = func() { cpu.Registers.L = res(cpu, cpu.Registers.L, 2) }                         // RES 2,L
	instr[0x96] = func() { MemWrite(cpu.Registers.HL(), res(cpu, MemRead(cpu.Registers.HL()), 2)) } // RES 2,(HL)
	instr[0x97] = func() { cpu.Registers.A = res(cpu, cpu.Registers.A, 2) }                         // RES 2,A
	instr[0x98] = func() { cpu.Registers.B = res(cpu, cpu.Registers.B, 3) }                         // RES 3,B
	instr[0x99] = func() { cpu.Registers.C = res(cpu, cpu.Registers.C, 3) }                         // RES 3,C
	instr[0x9a] = func() { cpu.Registers.D = res(cpu, cpu.Registers.D, 3) }                         // RES 3,D
	instr[0x9b] = func() { cpu.Registers.E = res(cpu, cpu.Registers.E, 3) }                         // RES 3,E
	instr[0x9c] = func() { cpu.Registers.H = res(cpu, cpu.Registers.H, 3) }                         // RES 3,H
	instr[0x9d] = func() { cpu.Registers.L = res(cpu, cpu.Registers.L, 3) }                         // RES 3,L
	instr[0x9e] = func() { MemWrite(cpu.Registers.HL(), res(cpu, MemRead(cpu.Registers.HL()), 3)) } // RES 3,(HL)
	instr[0x9f] = func() { cpu.Registers.A = res(cpu, cpu.Registers.A, 3) }                         // RES 3,A

	instr[0xa0] = func() { cpu.Registers.B = res(cpu, cpu.Registers.B, 4) }                         // RES 4,B
	instr[0xa1] = func() { cpu.Registers.C = res(cpu, cpu.Registers.C, 4) }                         // RES 4,C
	instr[0xa2] = func() { cpu.Registers.D = res(cpu, cpu.Registers.D, 4) }                         // RES 4,D
	instr[0xa3] = func() { cpu.Registers.E = res(cpu, cpu.Registers.E, 4) }                         // RES 4,E
	instr[0xa4] = func() { cpu.Registers.H = res(cpu, cpu.Registers.H, 4) }                         // RES 4,H
	instr[0xa5] = func() { cpu.Registers.L = res(cpu, cpu.Registers.L, 4) }                         // RES 4,L
	instr[0xa6] = func() { MemWrite(cpu.Registers.HL(), res(cpu, MemRead(cpu.Registers.HL()), 4)) } // RES 4,(HL)
	instr[0xa7] = func() { cpu.Registers.A = res(cpu, cpu.Registers.A, 4) }                         // RES 4,A
	instr[0xa8] = func() { cpu.Registers.B = res(cpu, cpu.Registers.B, 5) }                         // RES 5,B
	instr[0xa9] = func() { cpu.Registers.C = res(cpu, cpu.Registers.C, 5) }                         // RES 5,C
	instr[0xaa] = func() { cpu.Registers.D = res(cpu, cpu.Registers.D, 5) }                         // RES 5,D
	instr[0xab] = func() { cpu.Registers.E = res(cpu, cpu.Registers.E, 5) }                         // RES 5,E
	instr[0xac] = func() { cpu.Registers.H = res(cpu, cpu.Registers.H, 5) }                         // RES 5,H
	instr[0xad] = func() { cpu.Registers.L = res(cpu, cpu.Registers.L, 5) }                         // RES 5,L
	instr[0xae] = func() { MemWrite(cpu.Registers.HL(), res(cpu, MemRead(cpu.Registers.HL()), 5)) } // RES 5,(HL)
	instr[0xaf] = func() { cpu.Registers.A = res(cpu, cpu.Registers.A, 5) }                         // RES 5,A

	instr[0xb0] = func() { cpu.Registers.B = res(cpu, cpu.Registers.B, 6) }                         // RES 6,B
	instr[0xb1] = func() { cpu.Registers.C = res(cpu, cpu.Registers.C, 6) }                         // RES 6,C
	instr[0xb2] = func() { cpu.Registers.D = res(cpu, cpu.Registers.D, 6) }                         // RES 6,D
	instr[0xb3] = func() { cpu.Registers.E = res(cpu, cpu.Registers.E, 6) }                         // RES 6,E
	instr[0xb4] = func() { cpu.Registers.H = res(cpu, cpu.Registers.H, 6) }                         // RES 6,H
	instr[0xb5] = func() { cpu.Registers.L = res(cpu, cpu.Registers.L, 6) }                         // RES 6,L
	instr[0xb6] = func() { MemWrite(cpu.Registers.HL(), res(cpu, MemRead(cpu.Registers.HL()), 6)) } // RES 6,(HL)
	instr[0xb7] = func() { cpu.Registers.A = res(cpu, cpu.Registers.A, 6) }                         // RES 6,A
	instr[0xb8] = func() { cpu.Registers.B = res(cpu, cpu.Registers.B, 7) }                         // RES 7,B
	instr[0xb9] = func() { cpu.Registers.C = res(cpu, cpu.Registers.C, 7) }                         // RES 7,C
	instr[0xba] = func() { cpu.Registers.D = res(cpu, cpu.Registers.D, 7) }                         // RES 7,D
	instr[0xbb] = func() { cpu.Registers.E = res(cpu, cpu.Registers.E, 7) }                         // RES 7,E
	instr[0xbc] = func() { cpu.Registers.H = res(cpu, cpu.Registers.H, 7) }                         // RES 7,H
	instr[0xbd] = func() { cpu.Registers.L = res(cpu, cpu.Registers.L, 7) }                         // RES 7,L
	instr[0xbe] = func() { MemWrite(cpu.Registers.HL(), res(cpu, MemRead(cpu.Registers.HL()), 7)) } // RES 7,(HL)
	instr[0xbf] = func() { cpu.Registers.A = res(cpu, cpu.Registers.A, 7) }                         // RES 7,A

	instr[0xc0] = func() { cpu.Registers.B = set(cpu, cpu.Registers.B, 0) }                         // SET 0,B
	instr[0xc1] = func() { cpu.Registers.C = set(cpu, cpu.Registers.C, 0) }                         // SET 0,C
	instr[0xc2] = func() { cpu.Registers.D = set(cpu, cpu.Registers.D, 0) }                         // SET 0,D
	instr[0xc3] = func() { cpu.Registers.E = set(cpu, cpu.Registers.E, 0) }                         // SET 0,E
	instr[0xc4] = func() { cpu.Registers.H = set(cpu, cpu.Registers.H, 0) }                         // SET 0,H
	instr[0xc5] = func() { cpu.Registers.L = set(cpu, cpu.Registers.L, 0) }                         // SET 0,L
	instr[0xc6] = func() { MemWrite(cpu.Registers.HL(), set(cpu, MemRead(cpu.Registers.HL()), 0)) } // SET 0,(HL)
	instr[0xc7] = func() { cpu.Registers.A = set(cpu, cpu.Registers.A, 0) }                         // SET 0,A
	instr[0xc8] = func() { cpu.Registers.B = set(cpu, cpu.Registers.B, 1) }                         // SET 1,B
	instr[0xc9] = func() { cpu.Registers.C = set(cpu, cpu.Registers.C, 1) }                         // SET 1,C
	instr[0xca] = func() { cpu.Registers.D = set(cpu, cpu.Registers.D, 1) }                         // SET 1,D
	instr[0xcb] = func() { cpu.Registers.E = set(cpu, cpu.Registers.E, 1) }                         // SET 1,E
	instr[0xcc] = func() { cpu.Registers.H = set(cpu, cpu.Registers.H, 1) }                         // SET 1,H
	instr[0xcd] = func() { cpu.Registers.L = set(cpu, cpu.Registers.L, 1) }                         // SET 1,L
	instr[0xce] = func() { MemWrite(cpu.Registers.HL(), set(cpu, MemRead(cpu.Registers.HL()), 1)) } // SET 1,(HL)
	instr[0xcf] = func() { cpu.Registers.A = set(cpu, cpu.Registers.A, 1) }                         // SET 1,A

	instr[0xd0] = func() { cpu.Registers.B = set(cpu, cpu.Registers.B, 2) }                         // SET 2,B
	instr[0xd1] = func() { cpu.Registers.C = set(cpu, cpu.Registers.C, 2) }                         // SET 2,C
	instr[0xd2] = func() { cpu.Registers.D = set(cpu, cpu.Registers.D, 2) }                         // SET 2,D
	instr[0xd3] = func() { cpu.Registers.E = set(cpu, cpu.Registers.E, 2) }                         // SET 2,E
	instr[0xd4] = func() { cpu.Registers.H = set(cpu, cpu.Registers.H, 2) }                         // SET 2,H
	instr[0xd5] = func() { cpu.Registers.L = set(cpu, cpu.Registers.L, 2) }                         // SET 2,L
	instr[0xd6] = func() { MemWrite(cpu.Registers.HL(), set(cpu, MemRead(cpu.Registers.HL()), 2)) } // SET 2,(HL)
	instr[0xd7] = func() { cpu.Registers.A = set(cpu, cpu.Registers.A, 2) }                         // SET 2,A
	instr[0xd8] = func() { cpu.Registers.B = set(cpu, cpu.Registers.B, 3) }                         // SET 3,B
	instr[0xd9] = func() { cpu.Registers.C = set(cpu, cpu.Registers.C, 3) }                         // SET 3,C
	instr[0xda] = func() { cpu.Registers.D = set(cpu, cpu.Registers.D, 3) }                         // SET 3,D
	instr[0xdb] = func() { cpu.Registers.E = set(cpu, cpu.Registers.E, 3) }                         // SET 3,E
	instr[0xdc] = func() { cpu.Registers.H = set(cpu, cpu.Registers.H, 3) }                         // SET 3,H
	instr[0xdd] = func() { cpu.Registers.L = set(cpu, cpu.Registers.L, 3) }                         // SET 3,L
	instr[0xde] = func() { MemWrite(cpu.Registers.HL(), set(cpu, MemRead(cpu.Registers.HL()), 3)) } // SET 3,(HL)
	instr[0xdf] = func() { cpu.Registers.A = set(cpu, cpu.Registers.A, 3) }                         // SET 3,A

	instr[0xe0] = func() { cpu.Registers.B = set(cpu, cpu.Registers.B, 4) }                         // SET 4,B
	instr[0xe1] = func() { cpu.Registers.C = set(cpu, cpu.Registers.C, 4) }                         // SET 4,C
	instr[0xe2] = func() { cpu.Registers.D = set(cpu, cpu.Registers.D, 4) }                         // SET 4,D
	instr[0xe3] = func() { cpu.Registers.E = set(cpu, cpu.Registers.E, 4) }                         // SET 4,E
	instr[0xe4] = func() { cpu.Registers.H = set(cpu, cpu.Registers.H, 4) }                         // SET 4,H
	instr[0xe5] = func() { cpu.Registers.L = set(cpu, cpu.Registers.L, 4) }                         // SET 4,L
	instr[0xe6] = func() { MemWrite(cpu.Registers.HL(), set(cpu, MemRead(cpu.Registers.HL()), 4)) } // SET 4,(HL)
	instr[0xe7] = func() { cpu.Registers.A = set(cpu, cpu.Registers.A, 4) }                         // SET 4,A
	instr[0xe8] = func() { cpu.Registers.B = set(cpu, cpu.Registers.B, 5) }                         // SET 5,B
	instr[0xe9] = func() { cpu.Registers.C = set(cpu, cpu.Registers.C, 5) }                         // SET 5,C
	instr[0xea] = func() { cpu.Registers.D = set(cpu, cpu.Registers.D, 5) }                         // SET 5,D
	instr[0xeb] = func() { cpu.Registers.E = set(cpu, cpu.Registers.E, 5) }                         // SET 5,E
	instr[0xec] = func() { cpu.Registers.H = set(cpu, cpu.Registers.H, 5) }                         // SET 5,H
	instr[0xed] = func() { cpu.Registers.L = set(cpu, cpu.Registers.L, 5) }                         // SET 5,L
	instr[0xee] = func() { MemWrite(cpu.Registers.HL(), set(cpu, MemRead(cpu.Registers.HL()), 5)) } // SET 5,(HL)
	instr[0xef] = func() { cpu.Registers.A = set(cpu, cpu.Registers.A, 5) }                         // SET 5,A

	instr[0xf0] = func() { cpu.Registers.B = set(cpu, cpu.Registers.B, 6) }                         // SET 6,B
	instr[0xf1] = func() { cpu.Registers.C = set(cpu, cpu.Registers.C, 6) }                         // SET 6,C
	instr[0xf2] = func() { cpu.Registers.D = set(cpu, cpu.Registers.D, 6) }                         // SET 6,D
	instr[0xf3] = func() { cpu.Registers.E = set(cpu, cpu.Registers.E, 6) }                         // SET 6,E
	instr[0xf4] = func() { cpu.Registers.H = set(cpu, cpu.Registers.H, 6) }                         // SET 6,H
	instr[0xf5] = func() { cpu.Registers.L = set(cpu, cpu.Registers.L, 6) }                         // SET 6,L
	instr[0xf6] = func() { MemWrite(cpu.Registers.HL(), set(cpu, MemRead(cpu.Registers.HL()), 6)) } // SET 6,(HL)
	instr[0xf7] = func() { cpu.Registers.A = set(cpu, cpu.Registers.A, 6) }                         // SET 6,A
	instr[0xf8] = func() { cpu.Registers.B = set(cpu, cpu.Registers.B, 7) }                         // SET 7,B
	instr[0xf9] = func() { cpu.Registers.C = set(cpu, cpu.Registers.C, 7) }                         // SET 7,C
	instr[0xfa] = func() { cpu.Registers.D = set(cpu, cpu.Registers.D, 7) }                         // SET 7,D
	instr[0xfb] = func() { cpu.Registers.E = set(cpu, cpu.Registers.E, 7) }                         // SET 7,E
	instr[0xfc] = func() { cpu.Registers.H = set(cpu, cpu.Registers.H, 7) }                         // SET 7,H
	instr[0xfd] = func() { cpu.Registers.L = set(cpu, cpu.Registers.L, 7) }                         // SET 7,L
	instr[0xfe] = func() { MemWrite(cpu.Registers.HL(), set(cpu, MemRead(cpu.Registers.HL()), 7)) } // SET 7,(HL)
	instr[0xff] = func() { cpu.Registers.A = set(cpu, cpu.Registers.A, 7) }                         // SET 7,A

	return instr
}

// RLC n -- Rotate n left. Old bit 7 to Carry flag.
func rlc(cpu *CPU, n byte) byte {
	leavingBit := n >> 7
	n = n<<1 | leavingBit

	cpu.SetZero(n == 0)
	cpu.SetNegative(false)
	cpu.SetHalfCarry(false)
	cpu.SetCarry(leavingBit == 1)
	return n
}

// RL n -- Rotate n left through Carry flag.
func rl(cpu *CPU, n byte) byte {
	leavingBit := n >> 7
	var carry byte = 0
	if cpu.Carry() {
		carry = 1
	}
	n = n<<1 | carry

	cpu.SetZero(n == 0)
	cpu.SetNegative(false)
	cpu.SetHalfCarry(false)
	cpu.SetCarry(leavingBit == 1)
	return n
}

// RRC n -- Rotate n right. Old bit 0 to Carry flag.
func rrc(cpu *CPU, n byte) byte {
	leavingBit := n & 1
	n = n >> 1
	if leavingBit == 1 {
		n |= 0x80
	}

	cpu.SetZero(n == 0)
	cpu.SetNegative(false)
	cpu.SetHalfCarry(false)
	cpu.SetCarry(leavingBit == 1)
	return n
}

// RR n -- Rotate n right through Carry flag.
func rr(cpu *CPU, n byte) byte {
	leavingBit := n & 1
	var carry byte = 0
	if cpu.Carry() {
		carry = 1
	}
	n = n >> 1
	if carry == 1 {
		n |= 0x80
	}

	cpu.SetZero(n == 0)
	cpu.SetNegative(false)
	cpu.SetHalfCarry(false)
	cpu.SetCarry(leavingBit == 1)
	return n
}

// SLA n -- Shift n left into Carry. LSB of n set to 0.
func sla(cpu *CPU, n byte) byte {
	leavingBit := n >> 7
	n <<= 1
	cpu.SetZero(n == 0)
	cpu.SetNegative(false)
	cpu.SetHalfCarry(false)
	cpu.SetCarry(leavingBit == 1)
	return n
}

// SRA n -- Shift n right into Carry. MSB doesn't change.
func sra(cpu *CPU, n byte) byte {
	leavingBit := n & 1
	n = (n & 0x80) | (n >> 1)
	cpu.SetZero(n == 0)
	cpu.SetNegative(false)
	cpu.SetHalfCarry(false)
	cpu.SetCarry(leavingBit == 1)
	return n
}

// SRL n -- Shift n right into Carry. MSB set to 0.
func srl(cpu *CPU, n byte) byte {
	leavingBit := n & 1
	n = (n & 0x80) | (n >> 1)
	cpu.SetZero(n == 0)
	cpu.SetNegative(false)
	cpu.SetHalfCarry(false)
	cpu.SetCarry(leavingBit == 1)
	return n
}

// SWAP n -- Swap upper & lower nibles of n.
func swap(cpu *CPU, n byte) byte {
	swapped := ((n&0x0f)<<4 | (n&0xf0)>>4)
	cpu.SetZero(swapped == 0)
	cpu.SetNegative(false)
	cpu.SetHalfCarry(false)
	cpu.SetCarry(false)
	return swapped
}

// BIT b,r -- Test bit b in register r.
func bit(cpu *CPU, b, r byte) {
	bit := r>>b&1 == 0
	cpu.SetZero(bit)
	cpu.SetNegative(false)
	cpu.SetHalfCarry(true)
}

// SET b,r -- Set bit b in register r.
func set(cpu *CPU, b, r byte) byte {
	return r | (1 << b)
}

// RES b,r -- Reset bit b in register r.
func res(cpu *CPU, b, r byte) byte {
	return r & ^(1 << b)
}
