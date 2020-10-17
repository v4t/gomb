package cartridge

import (
	"strings"
)

// MBC is an interface for various memory bank controllers.
type MBC interface {
	WriteMemory(address uint16, value byte)
	ReadMemory(address uint16) byte
}

// Cartridge manages gameboy cartridge related functionality.
type Cartridge struct {
	Title string
	mbc   MBC
}

// NewCartridge initializes cartridge based on cartridge header.
func NewCartridge(rom []byte) *Cartridge {
	cart := &Cartridge{}
	cart.Title = strings.Trim(string(rom[0x0134:0x0142]), "\x00")

	mbcType := rom[0x0147]
	if mbcType >= 0x00 && mbcType <= 0x03 {
		cart.mbc = NewMBC1(rom)
	} else if mbcType >= 0x05 && mbcType <= 0x06 {
		cart.mbc = NewMBC2(rom)
	} else {
		panic("MBC type not implemented for this cartridge.")
	}
	return cart
}

// Read from ROM or RAM using memory banking controller.
func (cart *Cartridge) Read(address uint16) byte {
	return cart.mbc.ReadMemory(address)
}

// Write to ROM or RAM memory banking controller.
func (cart *Cartridge) Write(address uint16, value byte) {
	cart.mbc.WriteMemory(address, value)
}
