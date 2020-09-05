package memory

import "math"

// Memory contains emulator memory.
var memory = make([]byte, math.MaxUint16+1)

// MMU manages RAM, ROM and cartridge data.
type MMU struct {
}

// InitializeMMU creates new MMU instance.
func InitializeMMU() *MMU {
	mmu := MMU{}
	return &mmu
}

// Read byte from memory address.
func (mmu *MMU) Read(address uint16) byte {
	return memory[address]
}

// Write byte to memory address.
func (mmu *MMU) Write(address uint16, value byte) {
	memory[address] = value
}
