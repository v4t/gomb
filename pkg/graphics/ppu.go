package graphics

import "github.com/v4t/gomb/pkg/memory"

// PPUState defines possible PPU states.
type PPUState byte

// Possible PPU states.
const (
	HBlank        = 0
	VBlank        = 1
	OAMSearch     = 2
	PixelTransfer = 3
)

// PPU represents pixel processing unit.
type PPU struct {
	scanlineCounter int
	state           PPUState
	Display         *Display
	MMU             *memory.MMU
}

// Execute ppu cycle.
func (ppu *PPU) Execute(cycles int) {
	ppu.SetPPUState()
	if ppu.Display.IsEnabled() {
		ppu.scanlineCounter -= cycles
	}

	if ppu.scanlineCounter <= 0 {
		currentLine := memory.Memory[0xff44] + 1
		memory.Memory[0xff44] = currentLine

		ppu.scanlineCounter = 456

		if currentLine == 144 {
			// vblank
		} else if currentLine > 153 {
			memory.Memory[0xff44] = 0
		} else if currentLine < 144 {
			ppu.DrawScanLine()
		}
	}
}

/*
00: H-Blank
01: V-Blank
10: Searching Sprites Atts
11: Transfering Data to LCD Driver
*/
// SetPPUState is
func (ppu *PPU) SetPPUState() {
	status := ppu.MMU.Read(0xff41)
	if !IsLCDEnabled() {
		// set the mode to 1 during lcd disabled and reset scanline
		ppu.scanlineCounter = 456
		memory.Memory[0xff44] = 0
		status &= 252
		status = BitSet(status, 0)
		ppu.MMU.Write(0xFF41, status)
		return
	}

	currentline := ppu.MMU.Read(0xff44)
	currentmode := status & 0x3

	var mode byte = 0
	reqInt := false

	// in vblank so set mode to 1
	if currentline >= 144 {
		mode = 1
		status = BitSet(status, 0)
		status = BitReset(status, 1)
		reqInt = TestBit(status, 4)
	} else {
		mode2bounds := 456 - 80
		mode3bounds := mode2bounds - 172

		// mode 2 -- oam search
		if ppu.scanlineCounter >= mode2bounds {
			mode = 2
			status = BitSet(status, 1)
			status = BitReset(status, 0)
			reqInt = TestBit(status, 5)
		} else if ppu.scanlineCounter >= mode3bounds {
			// mode 3 -- pixeltransfer
			mode = 3
			status = BitSet(status, 1)
			status = BitSet(status, 0)
		} else {
			// mode 0 -- h blank
			mode = 0
			status = BitReset(status, 1)
			status = BitReset(status, 0)
			reqInt = TestBit(status, 3)
		}
	}

	// just entered a new mode so request interupt
	if reqInt && (mode != currentmode) {
		RequestInterrupt(ppu.MMU, 1)
	}

	// check the coincidence flag
	if currentline == ppu.MMU.Read(0xff45) {
		status = BitSet(status, 2)
		if TestBit(status, 6) {
			RequestInterrupt(ppu.MMU, 1)
		}
	} else {
		status = BitReset(status, 2)
	}
	ppu.MMU.Write(0xff41, status)
}

// DrawScanLine draws scanline...
func (ppu *PPU) DrawScanLine() {
	control := ppu.MMU.Read(0xff40)
	if TestBit(control, 0) {
		ppu.RenderTiles()
	}
	if TestBit(control, 1) {
		ppu.RenderSprites()
	}
}

// RenderTiles doc
func (ppu *PPU) RenderTiles() {
}

// RenderSprites doc
func (ppu *PPU) RenderSprites() {

}

// IsLCDEnabled doc
func IsLCDEnabled() bool {
	return TestBit(memory.Memory[0xFF40], 7)
}

// BitSet doc
func BitSet(value byte, pos int) byte {
	value |= (1 << pos)
	return value
}

// BitReset doc
func BitReset(value byte, pos int) byte {
	value &= ^(1 << pos)
	return value
}

// TestBit doc
func TestBit(value byte, pos int) bool {
	result := value & (1 << pos)
	return result != 0
}

// RequestInterrupt doc
func RequestInterrupt(mmu *memory.MMU, bit int) {
	requestFlag := mmu.Read(0xff0f)
	requestFlag = BitSet(requestFlag, bit)
	mmu.Write(0xff0f, requestFlag)
}
