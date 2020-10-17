package cartridge

import (
	"fmt"
	"log"
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
	switch rom[0x0147] {
	case 0x00: // ROM only
		cart.mbc = NewROM(rom)
	case 0x01: // MBC 1
		cart.mbc = NewMBC1(rom)
	case 0x02: // MBC 1 + RA
		cart.mbc = NewMBC1(rom)
	case 0x03: // MBC 1 + RAM + Battery
		cart.mbc = NewMBC1(rom)
	case 0x05: // MBC 2
		fmt.Println("Mbac2")
		cart.mbc = NewMBC2(rom)
	case 0x06: // MBC 2 + RAM + Battery
		fmt.Println("Mbac2")
		cart.mbc = NewMBC2(rom)
	default:
		log.Fatalf("MBC not implemented")
	}
	return cart
}

// Read from ROM or RAM
func (cart *Cartridge) Read(address uint16) byte {
	return cart.mbc.ReadMemory(address)
}

// Write to ROM or RAM
func (cart *Cartridge) Write(address uint16, value byte) {
	cart.mbc.WriteMemory(address, value)
}
