package memory

import (
	"math"
)

// MemoryRegion represents a specific memory region / block managed by MMU.
type MemoryRegion interface {
	Read(address uint16) byte
	Write(address uint16, value byte)
}

// MMU manages RAM, ROM and cartridge data.
type MMU struct {
	Memory     []byte
	Input      MemoryRegion
	Interrupts MemoryRegion
	Timer      MemoryRegion
}

// InitializeMMU creates new MMU instance.
func InitializeMMU() *MMU {
	mmu := MMU{Memory: make([]byte, math.MaxUint16+1)}
	mmu.Memory[0xff05] = 0x00
	mmu.Memory[0xff06] = 0x00
	mmu.Memory[0xff07] = 0x00
	mmu.Memory[0xff10] = 0x80
	mmu.Memory[0xff11] = 0xbf
	mmu.Memory[0xff12] = 0xf3
	mmu.Memory[0xff14] = 0xbf
	mmu.Memory[0xff16] = 0x3f
	mmu.Memory[0xff17] = 0x00
	mmu.Memory[0xff19] = 0xbf
	mmu.Memory[0xff1a] = 0x7f
	mmu.Memory[0xff1b] = 0xff
	mmu.Memory[0xff1c] = 0x9f
	mmu.Memory[0xff1e] = 0xbf
	mmu.Memory[0xff20] = 0xff
	mmu.Memory[0xff21] = 0x00
	mmu.Memory[0xff22] = 0x00
	mmu.Memory[0xff23] = 0xbf
	mmu.Memory[0xff24] = 0x77
	mmu.Memory[0xff25] = 0xf3
	mmu.Memory[0xff26] = 0xf1
	mmu.Memory[0xff40] = 0x91
	mmu.Memory[0xff42] = 0x00
	mmu.Memory[0xff43] = 0x00
	mmu.Memory[0xff45] = 0x00
	mmu.Memory[0xff47] = 0xfc
	mmu.Memory[0xff48] = 0xff
	mmu.Memory[0xff49] = 0xff
	mmu.Memory[0xff4a] = 0x00
	mmu.Memory[0xff4b] = 0x00
	mmu.Memory[0xffff] = 0x00

	// Joypad
	mmu.Memory[0xff00] = 0xcf
	return &mmu
}

// Read byte from memory address.
func (mmu *MMU) Read(address uint16) byte {
	if address == 0xff00 {
		return mmu.Input.Read(address)
	} else if address == 0xffff || address == 0xff0f {
		return mmu.Interrupts.Read(address)
	}else if address == 0xff04 || address == 0xff05 || address == 0xff06 || address == 0xff07 {
		return mmu.Timer.Read(address)
	}
	return mmu.Memory[address]
}

// Write byte to memory address.
func (mmu *MMU) Write(address uint16, value byte) {
	if address < 0x8000 {
		// Read only memory
		return
	} else if (address >= 0xe000) && (address < 0xfe00) {
		// Echo ram
		mmu.Memory[address] = value
		mmu.Write(address-0x2000, value)
	} else if (address >= 0xfea0) && (address < 0xfeff) {
		// Restricted area
		return
	} else if address == 0xff44 {
		mmu.Memory[address] = 0
	} else if address == 0xff46 {
		mmu.dmaTransfer(value)
	} else if address == 0xff00 {
		mmu.Input.Write(address, value)
	} else if address == 0xffff || address == 0xff0f {
		mmu.Interrupts.Write(address, value)
	} else if address == 0xff04 || address == 0xff05 || address == 0xff06 || address == 0xff07 {
		mmu.Timer.Write(address, value)
	} else {
		mmu.Memory[address] = value
	}
}

// LoadRom copies rom to memory.
func (mmu *MMU) LoadRom(rom []byte) {
	copy(mmu.Memory[:], rom)
}

func (mmu *MMU) dmaTransfer(value byte) {
	address := uint16(value) << 8
	for i := uint16(0); i < 0xa0; i++ {
		mmu.Write(0xfe00+i, mmu.Read(address+i))
	}
}
