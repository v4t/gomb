package graphics

import "github.com/v4t/gomb/pkg/memory"

// PPURegisters contains available ppu registers.
type PPURegisters struct {
	LcdControl PPURegister
	ScrollX    PPURegister
	ScrollY    PPURegister
	Scanline   PPURegister
	BgPalette  PPURegister
	WindowX    PPURegister
	WindowY    PPURegister
}

// InitRegisters is constructor for PPURegisters.
func InitRegisters() *PPURegisters {
	return &PPURegisters{
		LcdControl: PPURegister{address: 0xff40},
		ScrollY:    PPURegister{address: 0xff42},
		ScrollX:    PPURegister{address: 0xff43},
		Scanline:   PPURegister{address: 0xff44},
		BgPalette:  PPURegister{address: 0xff47},
		WindowX:    PPURegister{address: 0xff4a},
		WindowY:    PPURegister{address: 0xff4b},
	}
}

// PPURegister contains ppu register's memory address and current value.
type PPURegister struct {
	address uint16
}

// Get register value.
func (reg *PPURegister) Get() byte {
	return memory.Memory[reg.address]
}


// Set register value.
func (reg *PPURegister) Set(value byte) {
	memory.Memory[reg.address] = value
}
