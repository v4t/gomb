package memory

import (
	"math"
)

// Addressable interface provides functions to read/write bytes in a given
// 16-bit address space.
type Addressable interface {
	// Contains returns true if the Addressable can handle the address.
	Contains(addr uint16) bool

	// Read returns the value stored at the given address.
	Read(addr uint16) uint8

	// Write stores the given value at the given address (if writable).
	Write(addr uint16, value uint8)
}

// Memory contains emulator memory.
var memory = make([]byte, math.MaxUint16+1)

// MMU manages an arbitrary number of ordered address spaces. It also satisfies
// the MemoryRegion interface.
type MMU struct {
	PPU Addressable
}

// LoadBoot loads
func (m *MMU) LoadBoot(rom []byte) {
	copy(memory[:], rom)
}

// Contains no
func (m *MMU) Contains(addr uint16) bool {
	// return m.space(addr) != nil
	return true
}

// Read byte from memory
func (m *MMU) Read(addr uint16) uint8 {
	// if space := m.space(addr); space != nil {
	// 	return space.Read(addr)
	// }
	// return 0xff
	if m.PPU.Contains(addr) {
		return m.PPU.Read(addr)
	}
	return memory[addr]
}

// Write byte to memory
func (m *MMU) Write(addr uint16, value uint8) {
	// if space := m.space(addr); space != nil {
	// 	space.Write(addr, value)
	// }
	if m.PPU.Contains(addr) {
		m.PPU.Write(addr, value)
	} else {
		memory[addr] = value
	}
}
