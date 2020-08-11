package hardware

// CBInstructions represents extended cpu instructions
var CBInstructions = initInstr()

// Init extended instructions
func initCBInstr() []func(*CPU) {
	var instr = make([]func(*CPU), 0x100)

	instr[0x00] = func(cpu *CPU) { cpu.Registers.B = rotateLeft(cpu, cpu.Registers.B, true) }                          // RLC B
	instr[0x01] = func(cpu *CPU) { cpu.Registers.C = rotateLeft(cpu, cpu.Registers.C, true) }                          // RLC C
	instr[0x02] = func(cpu *CPU) { cpu.Registers.D = rotateLeft(cpu, cpu.Registers.D, true) }                          // RLC D
	instr[0x03] = func(cpu *CPU) { cpu.Registers.E = rotateLeft(cpu, cpu.Registers.E, true) }                          // RLC E
	instr[0x04] = func(cpu *CPU) { cpu.Registers.H = rotateLeft(cpu, cpu.Registers.H, true) }                          // RLC H
	instr[0x05] = func(cpu *CPU) { cpu.Registers.L = rotateLeft(cpu, cpu.Registers.L, true) }                          // RLC L
	instr[0x06] = func(cpu *CPU) { MemWrite(cpu.Registers.HL(), rotateLeft(cpu, MemRead(cpu.Registers.HL()), true)) }  // RLC (HL)
	instr[0x07] = func(cpu *CPU) { cpu.Registers.A = rotateLeft(cpu, cpu.Registers.A, true) }                          // RLC A
	instr[0x08] = func(cpu *CPU) { cpu.Registers.B = rotateRight(cpu, cpu.Registers.B, true) }                         // RRC B
	instr[0x09] = func(cpu *CPU) { cpu.Registers.C = rotateRight(cpu, cpu.Registers.C, true) }                         // RRC C
	instr[0x0a] = func(cpu *CPU) { cpu.Registers.D = rotateRight(cpu, cpu.Registers.D, true) }                         // RRC D
	instr[0x0b] = func(cpu *CPU) { cpu.Registers.E = rotateRight(cpu, cpu.Registers.E, true) }                         // RRC E
	instr[0x0c] = func(cpu *CPU) { cpu.Registers.H = rotateRight(cpu, cpu.Registers.H, true) }                         // RRC H
	instr[0x0d] = func(cpu *CPU) { cpu.Registers.L = rotateRight(cpu, cpu.Registers.L, true) }                         // RRC L
	instr[0x0e] = func(cpu *CPU) { MemWrite(cpu.Registers.HL(), rotateRight(cpu, MemRead(cpu.Registers.HL()), true)) } // RRC (HL)
	instr[0x0f] = func(cpu *CPU) { cpu.Registers.A = rotateRight(cpu, cpu.Registers.A, true) }                         // RRC A

	instr[0x10] = func(cpu *CPU) { cpu.Registers.B = rotateLeft(cpu, cpu.Registers.B, false) }                          // RL B
	instr[0x11] = func(cpu *CPU) { cpu.Registers.C = rotateLeft(cpu, cpu.Registers.C, false) }                          // RL C
	instr[0x12] = func(cpu *CPU) { cpu.Registers.D = rotateLeft(cpu, cpu.Registers.D, false) }                          // RL D
	instr[0x13] = func(cpu *CPU) { cpu.Registers.E = rotateLeft(cpu, cpu.Registers.E, false) }                          // RL E
	instr[0x14] = func(cpu *CPU) { cpu.Registers.H = rotateLeft(cpu, cpu.Registers.H, false) }                          // RL H
	instr[0x15] = func(cpu *CPU) { cpu.Registers.L = rotateLeft(cpu, cpu.Registers.L, false) }                          // RL L
	instr[0x16] = func(cpu *CPU) { MemWrite(cpu.Registers.HL(), rotateLeft(cpu, MemRead(cpu.Registers.HL()), false)) }  // RL (HL)
	instr[0x17] = func(cpu *CPU) { cpu.Registers.A = rotateLeft(cpu, cpu.Registers.A, false) }                          // RL A
	instr[0x18] = func(cpu *CPU) { cpu.Registers.B = rotateRight(cpu, cpu.Registers.B, false) }                         // RR B
	instr[0x19] = func(cpu *CPU) { cpu.Registers.C = rotateRight(cpu, cpu.Registers.C, false) }                         // RR C
	instr[0x1a] = func(cpu *CPU) { cpu.Registers.D = rotateRight(cpu, cpu.Registers.D, false) }                         // RR D
	instr[0x1b] = func(cpu *CPU) { cpu.Registers.E = rotateRight(cpu, cpu.Registers.E, false) }                         // RR E
	instr[0x1c] = func(cpu *CPU) { cpu.Registers.H = rotateRight(cpu, cpu.Registers.H, false) }                         // RR H
	instr[0x1d] = func(cpu *CPU) { cpu.Registers.L = rotateRight(cpu, cpu.Registers.L, false) }                         // RR L
	instr[0x1e] = func(cpu *CPU) { MemWrite(cpu.Registers.HL(), rotateRight(cpu, MemRead(cpu.Registers.HL()), false)) } // RR (HL)
	instr[0x1f] = func(cpu *CPU) { cpu.Registers.A = rotateRight(cpu, cpu.Registers.A, false) }                         // RR A

	instr[0x20] = nop
	instr[0x21] = nop
	instr[0x22] = nop
	instr[0x23] = nop
	instr[0x24] = nop
	instr[0x25] = nop
	instr[0x26] = nop
	instr[0x27] = nop
	instr[0x28] = nop
	instr[0x29] = nop
	instr[0x2a] = nop
	instr[0x2b] = nop
	instr[0x2c] = nop
	instr[0x2d] = nop
	instr[0x2e] = nop
	instr[0x2f] = nop

	instr[0x30] = nop
	instr[0x31] = nop
	instr[0x32] = nop
	instr[0x33] = nop
	instr[0x34] = nop
	instr[0x35] = nop
	instr[0x36] = nop
	instr[0x37] = nop
	instr[0x38] = nop
	instr[0x39] = nop
	instr[0x3a] = nop
	instr[0x3b] = nop
	instr[0x3c] = nop
	instr[0x3d] = nop
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

func rotateLeft(cpu *CPU, val byte, withCarry bool) byte {
	if withCarry {
		rlc(cpu, val)
	} else {
		rl(cpu, val)
	}
	return val
}

func rotateRight(cpu *CPU, val byte, withCarry bool) byte {
	if withCarry {
		rrc(cpu, val)
	} else {
		rr(cpu, val)
	}
	return val
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

// RL n -- Rotate n left through Carry flag
func rl(cpu *CPU, n byte) {
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

// RRC n -- Rotate n right. Old bit 0 to Carry flag
func rrc(cpu *CPU, n byte) {
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

// RR n -- Rotate n right through Carry flag.
func rr(cpu *CPU, n byte) {
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
