package graphics

import "github.com/v4t/gomb/pkg/memory"

// PPURegisters contains available ppu registers.
type PPURegisters struct {
	LcdControl PPURegister
	ScrollX    PPURegister
	ScrollY    PPURegister
	Scanline   PPURegister
	BgPalette  PPURegister
}

// InitRegisters is constructor for PPURegisters.
func InitRegisters() *PPURegisters {
	return &PPURegisters{
		LcdControl: PPURegister{address: 0xff40, Value: 0},
		ScrollY:    PPURegister{address: 0xff42, Value: 0},
		ScrollX:    PPURegister{address: 0xff43, Value: 0},
		Scanline:   PPURegister{address: 0xff44, Value: 0},
		BgPalette:  PPURegister{address: 0xff47, Value: 0},
	}
}

// PPURegister contains ppu register's memory address and current value.
type PPURegister struct {
	address uint16
	Value   byte
}

// Set register value.
func (reg *PPURegister) Set(value byte) {
	memory.Memory[reg.address] = value
	reg.Value = value
}
