package hardware

type Registers struct {
	A byte // accumulator
	B byte
	C byte
	D byte
	E byte
	H byte
	L byte
	F byte // flags
}

// AF is
func (regs *Registers) AF() uint16 {
	return uint16(regs.A)<<8 | uint16(regs.F)
}

// SetAF is
func (regs *Registers) SetAF(value uint16) {
	var masked = value & 0xfff0

	regs.A = byte(masked >> 8)
	regs.F = byte(masked & 0xff)
}

// BC is
func (regs *Registers) BC() uint16 {
	return uint16(regs.B)<<8 | uint16(regs.C)
}

// SetBC is
func (regs *Registers) SetBC(value uint16) {
	regs.B = byte(value >> 8)
	regs.C = byte(value & 0xff)
}

// DE is
func (regs *Registers) DE() uint16 {
	return uint16(regs.D)<<8 | uint16(regs.E)
}

// SetDE is
func (regs *Registers) SetDE(value uint16) {
	regs.D = byte(value >> 8)
	regs.E = byte(value & 0xff)
}

// HL is
func (regs *Registers) HL() uint16 {
	return uint16(regs.H)<<8 | uint16(regs.L)
}

// SetHL is
func (regs *Registers) SetHL(value uint16) {
	regs.H = byte(value >> 8)
	regs.L = byte(value & 0xff)
}
