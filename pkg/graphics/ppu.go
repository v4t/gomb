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
	registers       PPURegisters
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

// 00: H-Blank
// 01: V-Blank
// 10: Searching Sprites Atts
// 11: Transfering Data to LCD Driver

// SetPPUState is
func (ppu *PPU) SetPPUState() {
	status := ppu.MMU.Read(0xff41)
	if !ppu.IsLCDEnabled() {
		// set the mode to 1 during lcd disabled and reset scanline
		ppu.scanlineCounter = 456
		memory.Memory[0xff44] = 0
		status &= 252
		status = utils.SetBit(status, 0)
		ppu.MMU.Write(0xff41, status)
		return
	}

	currentline := ppu.registers.Scanline.Get()
	currentmode := status & 0x3

	var mode byte = 0
	reqInt := false

	// in vblank so set mode to 1
	if currentline >= 144 {
		mode = 1
		status = utils.SetBit(status, 0)
		status = utils.ResetBit(status, 1)
		reqInt = utils.TestBit(status, 4)
	} else {
		mode2bounds := 456 - 80
		mode3bounds := mode2bounds - 172

		// mode 2 -- oam search
		if ppu.scanlineCounter >= mode2bounds {
			mode = 2
			status = utils.SetBit(status, 1)
			status = utils.ResetBit(status, 0)
			reqInt = utils.TestBit(status, 5)
		} else if ppu.scanlineCounter >= mode3bounds {
			// mode 3 -- pixeltransfer
			mode = 3
			status = utils.SetBit(status, 1)
			status = utils.SetBit(status, 0)
		} else {
			// mode 0 -- h blank
			mode = 0
			status = utils.ResetBit(status, 1)
			status = utils.ResetBit(status, 0)
			reqInt = utils.TestBit(status, 3)
		}
	}

	// just entered a new mode so request interupt
	if reqInt && (mode != currentmode) {
		RequestInterrupt(ppu.MMU, 1)
	}

	// check the coincidence flag
	if currentline == ppu.MMU.Read(0xff45) {
		status = utils.SetBit(status, 2)
		if utils.TestBit(status, 6) {
			RequestInterrupt(ppu.MMU, 1)
		}
	} else {
		status = utils.ResetBit(status, 2)
	}
	ppu.MMU.Write(0xff41, status)
}

// DrawScanLine draws scanline...
func (ppu *PPU) DrawScanLine() {
	control := ppu.MMU.Read(0xff40)
	if utils.TestBit(control, 0) {
		ppu.RenderTiles()
	}
	if utils.TestBit(control, 1) {
		ppu.RenderSprites()
	}
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
		// is the current scanline we're drawing within the windows Y pos?,
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

	// which background mem?
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
	var tileRow uint16 = uint16((byte(yPos / 8)) * 32)

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
		var tileAddrss uint16 = backgroundMemory + tileRow + tileCol
		if unsigned {
			tileNum = int16(ppu.MMU.Read(tileAddrss))
		} else {
			tileNum = int16(int8(ppu.MMU.Read(tileAddrss)))
		}

		// deduce where this tile identifier is in memory
		var tileLocation uint16 = tileData

		if unsigned {
			tileLocation += uint16(tileNum * 16)
		} else {
			tileLocation += uint16((tileNum + 128) * 16)
		}

		// find the correct vertical line we're on of the
		// tile to get the tile data from memory
		var line uint16 = uint16(yPos % 8)
		line *= 2 // each vertical line takes up two bytes of memory
		data1 := ppu.MMU.Read(tileLocation + line)
		data2 := ppu.MMU.Read(tileLocation + line + 1)

		// pixel 0 in the tile is it 7 of data 1 and data2.
		// Pixel 1 is bit 6 etc..
		var colourBit int = int(xPos) % 8
		colourBit -= 7
		colourBit *= -1

		// combine data 2 and data 1 to get the colour id for this pixel
		// in the tile
		var colourNum int = utils.GetBit(data2, colourBit)
		colourNum <<= 1
		colourNum |= utils.GetBit(data1, colourBit)

		finaly := ppu.registers.Scanline.Get()
		ppu.Display.Draw(pixel, finaly, byte(colourNum))
	}
}

// RenderSprites doc
func (ppu *PPU) RenderSprites() {
	lcdControl := ppu.registers.LcdControl.Get()
	use8x16 := false
	if utils.TestBit(lcdControl, 2) {
		use8x16 = true
	}

	for sprite := byte(0); sprite < 40; sprite++ {
		// sprite occupies 4 bytes in the sprite attributes table
		var index uint16 = uint16(sprite) * 4
		var yPos byte = ppu.MMU.Read(0xfe00+index) - 16
		var xPos byte = ppu.MMU.Read(0xfe00+index+1) - 8
		var tileLocation byte = ppu.MMU.Read(0xfe00 + index + 2)
		var attributes byte = ppu.MMU.Read(0xfe00 + index + 3)

		yFlip := utils.TestBit(attributes, 6)
		xFlip := utils.TestBit(attributes, 5)

		scanline := ppu.registers.Scanline.Get()

		ysize := byte(8)
		if use8x16 {
			ysize = 16
		}

		// does this sprite intercept with the scanline?
		if (scanline >= yPos) && (scanline < (yPos + ysize)) {
			line := int16(scanline - yPos)

			// read the sprite in backwards in the y axis
			if yFlip {
				line -= int16(ysize)
				line *= -1
			}

			line *= 2 // same as for tiles
			var dataAddress uint16 = (0x8000 + uint16(tileLocation*16)) + uint16(line)
			data1 := ppu.MMU.Read(dataAddress)
			data2 := ppu.MMU.Read(dataAddress + 1)

			// its easier to read in from right to left as pixel 0 is
			// bit 7 in the colour data, pixel 1 is bit 6 etc...
			for tilePixel := 7; tilePixel >= 0; tilePixel-- {
				colourbit := tilePixel
				// read the sprite in backwards for the x axis
				if xFlip {
					colourbit -= 7
					colourbit *= -1
				}

				// the rest is the same as for tiles
				colourNum := utils.GetBit(data2, colourbit)
				colourNum <<= 1
				colourNum |= utils.GetBit(data1, colourbit)

				xPix := 0 - tilePixel
				xPix += 7

				pixel := xPos + byte(xPix)
				ppu.Display.Draw(pixel, scanline, byte(colourNum))
			}
		}
	}
}

// IsLCDEnabled doc
func (ppu *PPU) IsLCDEnabled() bool {
	return utils.TestBit(memory.Memory[0xff40], 7)
}

// RequestInterrupt doc
func RequestInterrupt(mmu *memory.MMU, bit int) {
	requestFlag := mmu.Read(0xff0f)
	requestFlag = utils.SetBit(requestFlag, bit)
	mmu.Write(0xff0f, requestFlag)
}
