package memory

import "math"

// Memory contains emulator memory.
// var Memory = make([]byte, math.MaxUint16+1)

// MMU manages RAM, ROM and cartridge data.
type MMU struct {
	Memory []byte
}

// InitializeMMU creates new MMU instance.
func InitializeMMU() *MMU {
	mmu := MMU{Memory: make([]byte, math.MaxUint16+1)}
	return &mmu
}

// Read byte from memory address.
func (mmu *MMU) Read(address uint16) byte {
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
