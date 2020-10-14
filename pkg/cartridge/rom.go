package cartridge

// ROM handles games that don't need a MBC but are mapped to memory directly.
type ROM struct {
	data []byte
}

// NewROM is a constructor.
func NewROM(rom []byte) *ROM {
	return &ROM{
		data: rom,
	}
}

// WriteToMB handles writes to ROM.
// Since there is no memory bank controller, writes are not allowed.
func (rom *ROM) WriteToMB(address uint16, value byte) { }

// ReadFromMB handles reads from ROM.
func (rom *ROM) ReadFromMB(address uint16) byte {
	return rom.data[address]
}
