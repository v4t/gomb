package cartridge

// MBC1 represents memory bank controller for MBC1 type.
type MBC1 struct {
	rom           []byte
	romBankNumber byte

	ram           []byte
	ramBankNumber byte

	romBanking bool
	ramEnabled bool
}

// NewMBC1 is a constructor for MBC1 type memory banking controller.
func NewMBC1(rom []byte) *MBC1 {
	return &MBC1{
		rom:           rom,
		ram:           make([]byte, 0x8000),
		romBankNumber: 1,
		romBanking:    true,
	}
}

// WriteMemory handles writes to MBC1.
func (mbc *MBC1) WriteMemory(address uint16, value byte) {
	if address < 0x2000 {
		// Any value with 0x0a in the lower 4 bits enables RAM and other values disable it
		mbc.ramEnabled = value&0x0f == 0x0a
	} else if address < 0x4000 {
		// Set lower bits for ROM bank number
		bank := (mbc.romBankNumber & 0xe0) | (value & 0x1f)
		mbc.selectRomBank(bank)
	} else if address < 0x6000 {
		// Select RAM bank or set higher bits for ROM bank number
		if mbc.romBanking {
			bank := (mbc.romBankNumber & 0x1f) | (value & 0xe0)
			mbc.selectRomBank(bank)
		} else {
			mbc.selectRamBank(value & 0x03)
		}
	} else if address < 0x8000 {
		// Select ROM/RAM banking mode
		mbc.romBanking = value&1 == 0
		if mbc.romBanking {
			mbc.selectRamBank(0)
		} else {
			mbc.selectRomBank(mbc.romBankNumber & 0x1f)
		}
	} else if address >= 0xa000 && address < 0xc000 {
		// Write to RAM
		if mbc.ramEnabled {
			mbc.ram[mbc.mapAddressToRam(address)] = value
		}
	}
}

// ReadMemory handles reads from memory for MBC1.
func (mbc *MBC1) ReadMemory(address uint16) byte {
	if address < 0x4000 {
		// ROM bank 0
		return mbc.rom[address]
	} else if address < 0x8000 {
		// Switchable ROM bank
		return mbc.rom[mbc.mapAddressToRom(address)]
	} else if address >= 0xa000 && address < 0xc000 {
		// Switchable RAM bank
		return mbc.ram[mbc.mapAddressToRam(address)]
	}
	panic("Tried to read invalid memory address from MBC")
}

func (mbc *MBC1) selectRomBank(bank byte) {
	if bank == 0x00 || bank == 0x20 || bank == 0x40 || bank == 0x60 {
		bank++
	}
	mbc.romBankNumber = bank
}

func (mbc *MBC1) selectRamBank(bank byte) {
	mbc.ramBankNumber = bank
}

func (mbc *MBC1) mapAddressToRom(address uint16) int {
	bank := int(mbc.romBankNumber)
	return int(address-0x4000) + (bank * 0x4000)
}

func (mbc *MBC1) mapAddressToRam(address uint16) int {
	bank := int(mbc.ramBankNumber)
	return int(address-0xa000) + (bank * 0x2000)
}
