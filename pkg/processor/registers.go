package processor

// Registers represents CPU Registers.
type Registers struct {
	A byte
	B byte
	C byte
	D byte
	E byte
	H byte
	L byte
	F byte
}

// AF returns 16-bit value of AF union register.
func (regs *Registers) AF() uint16 {
	return uint16(regs.A)<<8 | uint16(regs.F)
}

// SetAF sets 16-bit value for AF union register.
func (regs *Registers) SetAF(value uint16) {
	masked := value & 0xfff0 // F register value is masked

	regs.A = byte(masked >> 8)
	regs.F = byte(masked & 0xff)
}

// BC returns 16-bit value of BC union register.
func (regs *Registers) BC() uint16 {
	return uint16(regs.B)<<8 | uint16(regs.C)
}

// SetBC sets 16-bit value for BC union register.
func (regs *Registers) SetBC(value uint16) {
	regs.B = byte(value >> 8)
	regs.C = byte(value & 0xff)
}

// DE returns 16-bit value of DE union register.
func (regs *Registers) DE() uint16 {
	return uint16(regs.D)<<8 | uint16(regs.E)
}

// SetDE sets 16-bit value for DE union register.
func (regs *Registers) SetDE(value uint16) {
	regs.D = byte(value >> 8)
	regs.E = byte(value & 0xff)
}

// HL returns 16-bit value of HL union register.
func (regs *Registers) HL() uint16 {
	return uint16(regs.H)<<8 | uint16(regs.L)
}

// SetHL sets 16-bit value for HL union register.
func (regs *Registers) SetHL(value uint16) {
	regs.H = byte(value >> 8)
	regs.L = byte(value & 0xff)
}
