package graphics

import (
	"github.com/v4t/gomb/pkg/memory"
	"github.com/v4t/gomb/pkg/utils"
)

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
	registers       *PPURegisters
	Display         *Display
	MMU             *memory.MMU

	clock int
}

// InitPPU is PPU constructor.
func InitPPU(mmu *memory.MMU, display *Display) *PPU {
	return &PPU{
		MMU:       mmu,
		Display:   display,
		registers: InitRegisters(mmu),
	}
}

// Execute GPU cycle.
func (ppu *PPU) Execute(cycles int) {
	line := ppu.registers.Scanline.Get()
	ppu.clock += cycles

	switch ppu.state {
	case OAMSearch:
		if ppu.clock >= 80 {
			ppu.state = PixelTransfer
		}
	case PixelTransfer:
		if ppu.clock >= 172 {
			ppu.clock = 0
			ppu.state = HBlank
			ppu.renderScan() // Put scanline to image buffer
		}
		break
	case HBlank:
		if ppu.clock >= 204 {
			ppu.clock = 0
			ppu.registers.Scanline.Set(line + 1)

			if line == 143 {
				ppu.state = VBlank
				// ppu._canvas.putImageData(ppu._scrn, 0, 0);
				ppu.Display.RenderImage()
			} else {
				ppu.state = OAMSearch
			}
		}
	case VBlank:
		if ppu.clock >= 456 {
			ppu.clock = 0
			ppu.registers.Scanline.Set(line + 1)

			if line > 153 {
				ppu.state = OAMSearch
				ppu.registers.Scanline.Set(0)
			}
		}
		break
	}
}

func (ppu *PPU) renderScan() {
	control := ppu.MMU.Read(0xff40)
	if utils.TestBit(control, 0) || true {
		ppu.RenderTiles()
	}
	// if utils.TestBit(control, 1) || true {
	// 	ppu.RenderSprites()
	// }
}

// RenderTiles doc
func (ppu *PPU) RenderTiles() {
	var tileData uint16 = 0
	var backgroundMemory uint16 = 0
	var unsigned bool = true

	// where to draw the visual area and the window
	scrollY := ppu.registers.ScrollY.Get()
	scrollX := ppu.registers.ScrollX.Get()
	windowY := ppu.registers.WindowY.Get()
	windowX := ppu.registers.WindowX.Get() - 7
	lcdControl := ppu.registers.LcdControl.Get()

	var usingWindow bool = false

	// is the window enabled?
	if utils.TestBit(lcdControl, 5) {
		if windowY <= ppu.registers.Scanline.Get() {
			usingWindow = true
		}
	}

	// which tile data are we using?
	if utils.TestBit(lcdControl, 4) {
		tileData = 0x8000
	} else {
		// IMPORTANT: This memory region uses signed bytes as tile identifiers
		tileData = 0x8800
		unsigned = false
	}

	// which background mem
	if !usingWindow {
		if utils.TestBit(lcdControl, 3) {
			backgroundMemory = 0x9c00
		} else {
			backgroundMemory = 0x9800
		}
	} else {
		// which window memory?
		if utils.TestBit(lcdControl, 6) {
			backgroundMemory = 0x9c00
		} else {
			backgroundMemory = 0x9800
		}
	}
	var yPos byte = 0

	// yPos is used to calculate which of 32 vertical tiles the current scanline is drawing
	if !usingWindow {
		yPos = scrollY + ppu.registers.Scanline.Get()
	} else {
		yPos = ppu.registers.Scanline.Get() - windowY
	}

	// which of the 8 vertical pixels of the current tile is the scanline on?
	var tileRow uint16 = uint16(yPos/8) * 32

	// time to start drawing the 160 horizontal pixels for this scanline
	for pixel := byte(0); pixel < 160; pixel++ {

		var xPos byte = pixel + scrollX

		// translate the current x pos to window space if necessary
		if usingWindow {
			if pixel >= windowX {
				xPos = pixel - windowX
			}
		}

		// which of the 32 horizontal tiles does this xPos fall within?
		var tileCol uint16 = uint16(xPos / 8)
		var tileNum int16

		// get the tile identity number. Remember it can be signed or unsigned
		var tileAddr uint16 = backgroundMemory + tileRow + tileCol
		if unsigned {
			tileNum = int16(ppu.MMU.Read(tileAddr))
		} else {
			tileNum = int16(int8(ppu.MMU.Read(tileAddr)))
		}
		// deduce where this tile identifier is in memory
		var tileLocation uint16 = tileData
		if unsigned {
			tileLocation += uint16(tileNum*16)
		} else {
			tileLocation += uint16((tileNum + 128) * 16)
		}

		// find the correct vertical line we're on of the
		// tile to get the tile data from memory
		var line byte = yPos % 8
		line *= 2 // each vertical line takes up two bytes of memory

		data1 := ppu.MMU.Read(tileLocation + uint16(line))
		data2 := ppu.MMU.Read(tileLocation + uint16(line) + 1)

		// pixel 0 in the tile is it 7 of data 1 and data2.
		// Pixel 1 is bit 6 etc..
		var colourBit int = int(xPos) % 8
		colourBit -= 7
		colourBit *= -1

		// combine data 2 and data 1 to get the colour id for this pixel in the tile
		var colourNum int = utils.GetBit(data2, colourBit)
		colourNum <<= 1
		colourNum |= utils.GetBit(data1, colourBit)

		finaly := ppu.registers.Scanline.Get()
		ppu.Display.Draw(pixel, finaly, byte(colourNum))
	}
}

// IsLCDEnabled doc
func (ppu *PPU) IsLCDEnabled() bool {
	return utils.TestBit(ppu.MMU.Memory[0xff40], 7)
}

// RequestInterrupt doc
func RequestInterrupt(mmu *memory.MMU, bit int) {
	requestFlag := mmu.Read(0xff0f)
	requestFlag = utils.SetBit(requestFlag, bit)
	mmu.Write(0xff0f, requestFlag)
}
