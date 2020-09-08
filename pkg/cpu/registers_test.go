package cpu

import "testing"

func initRegisters() Registers {
	var regs Registers
	regs.A = 0xaa
	regs.F = 0xbb
	regs.B = 0xaa
	regs.C = 0xbb
	regs.D = 0xaa
	regs.E = 0xbb
	regs.H = 0xaa
	regs.L = 0xbb
	return regs
}

func TestGettingRegisterAF(t *testing.T) {
	regs := initRegisters()
	const expected uint16 = 0xaabb
	AF := regs.AF()
	if AF != expected {
		t.Errorf("AF register value should be %x, got %x", expected, AF)
	}
}

func TestSettingRegisterAF(t *testing.T) {
	regs := initRegisters()
	const expectedA byte = 0xcc
	const expectedF byte = 0xd0
	regs.SetAF(0xccdd)

	if regs.A != expectedA {
		t.Errorf("A register value should be %x, got %x", expectedA, regs.A)
	}
	if regs.F != expectedF {
		t.Errorf("F register value should be %x, got %x", expectedF, regs.F)
	}
}

func TestGettingRegisterBC(t *testing.T) {
	regs := initRegisters()
	const expected uint16 = 0xaabb
	BC := regs.BC()
	if BC != expected {
		t.Errorf("BC register value should be %x, got %x", expected, BC)
	}
}

func TestSettingRegisterBC(t *testing.T) {
	regs := initRegisters()
	const expectedB byte = 0xcc
	const expectedC byte = 0xdd
	regs.SetBC(0xccdd)

	if regs.B != expectedB {
		t.Errorf("B register value should be %x, got %x", expectedB, regs.B)
	}
	if regs.C != expectedC {
		t.Errorf("C register value should be %x, got %x", expectedC, regs.C)
	}
}

func TestGettingRegisterDE(t *testing.T) {
	regs := initRegisters()
	const expected uint16 = 0xaabb
	DE := regs.DE()
	if DE != expected {
		t.Errorf("DE register value should be %x, got %x", expected, DE)
	}
}

func TestSettingRegisterDE(t *testing.T) {
	regs := initRegisters()
	const expectedD byte = 0xcc
	const expectedE byte = 0xdd
	regs.SetDE(0xccdd)

	if regs.D != expectedD {
		t.Errorf("D register value should be %x, got %x", expectedD, regs.D)
	}
	if regs.E != expectedE {
		t.Errorf("E register value should be %x, got %x", expectedE, regs.E)
	}
}

func TestGettingRegisterHL(t *testing.T) {
	regs := initRegisters()
	const expected uint16 = 0xaabb
	HL := regs.HL()
	if HL != expected {
		t.Errorf("HL register value should be %x, got %x", expected, HL)
	}
}

func TestSettingRegisterHL(t *testing.T) {
	regs := initRegisters()
	const expectedH byte = 0xcc
	const expectedL byte = 0xdd
	regs.SetHL(0xccdd)

	if regs.H != expectedH {
		t.Errorf("H register value should be %x, got %x", expectedH, regs.H)
	}
	if regs.L != expectedL {
		t.Errorf("L register value should be %x, got %x", expectedL, regs.L)
	}
}
