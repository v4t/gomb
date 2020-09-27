package graphics

import "github.com/v4t/gomb/pkg/memory"

// PPURegisters contains available ppu registers.
type PPURegisters struct {
	LcdControl  PPURegister
	LcdStatus   PPURegister
	ScrollX     PPURegister
	ScrollY     PPURegister
	Scanline    PPURegister
	BgPalette   PPURegister
	ObjPalette1 PPURegister
	ObjPalette2 PPURegister
	WindowX     PPURegister
	WindowY     PPURegister
}

// InitRegisters is constructor for PPURegisters.
func InitRegisters(mmu *memory.MMU) *PPURegisters {
	mmu.Memory[0xff40] = 0x91
	mmu.Memory[0xff41] = 0x85
	return &PPURegisters{
		LcdControl:  PPURegister{mmu: mmu, address: 0xff40},
		LcdStatus:   PPURegister{mmu: mmu, address: 0xff41},
		ScrollY:     PPURegister{mmu: mmu, address: 0xff42},
		ScrollX:     PPURegister{mmu: mmu, address: 0xff43},
		Scanline:    PPURegister{mmu: mmu, address: 0xff44},
		BgPalette:   PPURegister{mmu: mmu, address: 0xff47},
		ObjPalette1: PPURegister{mmu: mmu, address: 0xff48},
		ObjPalette2: PPURegister{mmu: mmu, address: 0xff49},
		WindowX:     PPURegister{mmu: mmu, address: 0xff4a},
		WindowY:     PPURegister{mmu: mmu, address: 0xff4b},
	}
}

// PPURegister contains ppu register's memory address and current value.
type PPURegister struct {
	mmu     *memory.MMU
	address uint16
}

// Get register value.
func (reg *PPURegister) Get() byte {
	return reg.mmu.Memory[reg.address]
}

// Set register value.
func (reg *PPURegister) Set(value byte) {
	reg.mmu.Memory[reg.address] = value
}
