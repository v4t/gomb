package cartridge

// MBC2 represents memory bank controller for MBC2 type.
type MBC2 struct {
	ram           []byte
	rom           []byte
	romBankNumber byte
	ramEnabled    bool
}

func NewMBC2(rom []byte) *MBC2 {
	return &MBC2{
		rom:           rom,
		ram:           make([]byte, 0x2000),
		romBankNumber: 1,
	}
}

func (mbc *MBC2) WriteMemory(address uint16, value byte) {
	if address < 0x2000 {
		// Any value with 0x0a in the lower 4 bits enables RAM and other values disable it
		mbc.ramEnabled = value&0x0f == 0x0a
	} else if address < 0x4000 {
		// 5th bit must be 1 to select ROM bank which is specified by the lower 4 bits.
		if address&0x100 == 0x100 {
			mbc.selectRomBank(value & 0xf)
		}
	} else if address >= 0xa000 && address < 0xa200 {
		// Write to RAM
		if mbc.ramEnabled {
			mbc.ram[address-0xa000] = value & 0xf
		}
	}
}

func (mbc *MBC2) ReadMemory(address uint16) byte {
	if address < 0x4000 {
		// ROM bank 0
		return mbc.rom[address]
	} else if address < 0x8000 {
		// Switchable ROM bank
		return mbc.rom[mbc.mapAddressToRom(address)]
	} else if address >= 0xa000 && address < 0xa200 {
		// Switchable RAM bank
		return mbc.ram[address-0xa000] & 0xf
	}
	panic("Tried to read invalid memory address from MBC")
}

func (mbc *MBC2) selectRomBank(bank byte) {
	if bank == 0x00 || bank == 0x20 || bank == 0x40 || bank == 0x60 {
		bank++
	}
	mbc.romBankNumber = bank
}

func (mbc *MBC2) mapAddressToRom(address uint16) int {
	bank := int(mbc.romBankNumber)
	return int(address-0x4000) + (bank * 0x4000)
}
