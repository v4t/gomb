package graphics

import (
	"github.com/v4t/gomb/pkg/memory"
	"github.com/v4t/gomb/pkg/processor"
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
	Interrupts      *processor.Interrupts

	clock int
}

// NewPPU is PPU constructor.
func NewPPU(mmu *memory.MMU, display *Display) *PPU {
	return &PPU{
		MMU:       mmu,
		Display:   display,
		registers: NewPPURegisters(mmu),
	}
}

// Execute GPU cycle.
func (ppu *PPU) Execute(cycles int) {
	ppu.SetPPUState()
	// Is LCD enabled?
	if !utils.TestBit(ppu.registers.LcdControl.Get(), 7) {
		return
	}
	ppu.clock += cycles
	if ppu.clock >= 456 {
		currentLine := ppu.registers.Scanline.Get()
		if currentLine == 144 {
			ppu.Interrupts.SetInterrupt(processor.VBlankInterrupt)
		} else if currentLine > 153 {
			ppu.registers.Scanline.Set(0)
		} else if currentLine < 144 {
			ppu.renderScanline()
		}
		ppu.registers.Scanline.Set(currentLine + 1)
		ppu.clock = 0
	}
}

// SetPPUState sets current ppu state.
func (ppu *PPU) SetPPUState() {
	lcdStatus := ppu.registers.LcdStatus.Get()
	// Is LCD enabled?
	if !utils.TestBit(ppu.registers.LcdControl.Get(), 7) {
		ppu.resetPPU(lcdStatus)
		return
	}

	var mode PPUState
	interruptNeeded := false
	currentLine := ppu.registers.Scanline.Get()

	if currentLine >= 144 {
		mode = VBlank
		lcdStatus = utils.SetBit(lcdStatus, 0)
		lcdStatus = utils.ResetBit(lcdStatus, 1)
		interruptNeeded = utils.TestBit(lcdStatus, 4)
	} else if ppu.clock <= 80 {
		mode = OAMSearch
		lcdStatus = utils.SetBit(lcdStatus, 1)
		lcdStatus = utils.ResetBit(lcdStatus, 0)
		interruptNeeded = utils.TestBit(lcdStatus, 5)
	} else if ppu.clock <= 172 {
		mode = PixelTransfer
		lcdStatus = utils.SetBit(lcdStatus, 1)
		lcdStatus = utils.SetBit(lcdStatus, 0)
	} else {
		mode = HBlank
		lcdStatus = utils.ResetBit(lcdStatus, 1)
		lcdStatus = utils.ResetBit(lcdStatus, 0)
		interruptNeeded = utils.TestBit(lcdStatus, 3)
	}

	// Set required interrupts
	if interruptNeeded && (mode != ppu.state) {
		ppu.Interrupts.SetInterrupt(processor.LCDStatusInterrupt)
	}
	if currentLine == ppu.registers.LYC.Get() {
		lcdStatus = utils.SetBit(lcdStatus, 2)
		if utils.TestBit(lcdStatus, 6) {
			ppu.Interrupts.SetInterrupt(processor.LCDStatusInterrupt)
		}
	} else {
		lcdStatus = utils.ResetBit(lcdStatus, 2)
	}
	ppu.state = mode
	ppu.registers.LcdStatus.Set(lcdStatus)
}

// resetPPU sets ppu state to VBlank and resets scanline and clock.
func (ppu *PPU) resetPPU(lcdStatus byte) {
	ppu.clock = 0
	ppu.registers.Scanline.Set(0)
	lcdStatus &= 252
	lcdStatus = utils.SetBit(lcdStatus, 0)
	ppu.registers.LcdStatus.Set(lcdStatus)
}

func (ppu *PPU) renderScanline() {
	control := ppu.registers.LcdControl.Get()
	if utils.TestBit(control, 0) {
		ppu.renderTiles()
	}
	if utils.TestBit(control, 1) {
		ppu.renderSprites()
	}
}

func (ppu *PPU) renderTiles() {
	// Settings for tile rendering
	var tileData uint16 = 0
	var backgroundMemory uint16 = 0
	var unsigned bool = true
	var usingWindow bool = false

	// Current register values
	scanline := ppu.registers.Scanline.Get()
	scrollY := ppu.registers.ScrollY.Get()
	scrollX := ppu.registers.ScrollX.Get()
	windowY := ppu.registers.WindowY.Get()
	windowX := ppu.registers.WindowX.Get() - 7
	lcdControl := ppu.registers.LcdControl.Get()

	// Set tile rendering settings
	if utils.TestBit(lcdControl, 5) {
		if windowY <= scanline {
			usingWindow = true
		}
	}
	if utils.TestBit(lcdControl, 4) {
		tileData = 0x8000
	} else {
		tileData = 0x8800
		unsigned = false
	}
	if !usingWindow {
		if utils.TestBit(lcdControl, 3) {
			backgroundMemory = 0x9c00
		} else {
			backgroundMemory = 0x9800
		}
	} else {
		if utils.TestBit(lcdControl, 6) {
			backgroundMemory = 0x9c00
		} else {
			backgroundMemory = 0x9800
		}
	}

	// yPos represents the current vertical line
	var yPos byte = 0
	if !usingWindow {
		yPos = scrollY + scanline
	} else {
		yPos = scanline - windowY
	}

	// Current position of scanline on tile
	var tileRow uint16 = uint16(yPos/8) * 32

	// Draw horizontal pixels for scanline
	for pixel := byte(0); pixel < 160; pixel++ {
		xPos := pixel + scrollX

		// Translate current x to window space
		if usingWindow && pixel >= windowX {
			xPos = pixel - windowX
		}

		// Get tile location
		tileAddr := backgroundMemory + tileRow + uint16(xPos/8)
		tileLocation := tileData
		if unsigned {
			tileNum := int16(ppu.MMU.Read(tileAddr))
			tileLocation += uint16(tileNum * 16)
		} else {
			tileNum := int16(int8(ppu.MMU.Read(tileAddr)))
			tileLocation += uint16((tileNum + 128) * 16)
		}

		// Fetch tile data
		var line byte = (yPos % 8) * 2
		data1 := ppu.MMU.Read(tileLocation + uint16(line))
		data2 := ppu.MMU.Read(tileLocation + uint16(line) + 1)

		// Push pixel to frame buffer
		var colourBit int = ((int(xPos) % 8) - 7) * -1
		var colourNum int = (utils.GetBit(data2, colourBit) << 1) | utils.GetBit(data1, colourBit)
		ppu.Display.Draw(pixel, scanline, byte(colourNum))
	}
}

func (ppu *PPU) renderSprites() {
	lcdControl := ppu.registers.LcdControl.Get()
	scanline := int(ppu.registers.Scanline.Get())

	ysize := 8
	if utils.TestBit(lcdControl, 2) {
		ysize = 16
	}

	for sprite := 0; sprite < 40; sprite++ {
		// Get sprite information
		index := uint16(sprite * 4)
		yPos := int(ppu.MMU.Read(0xfe00+index) - 16)
		xPos := int(ppu.MMU.Read(0xfe00+index+1) - 8)
		tileLocation := uint16(ppu.MMU.Read(0xfe00 + index + 2))
		attributes := ppu.MMU.Read(0xfe00 + index + 3)

		yFlip := utils.TestBit(attributes, 6)
		xFlip := utils.TestBit(attributes, 5)

		if scanline < yPos || scanline >= (yPos+ysize) {
			continue
		}

		// Set sprite line
		line := scanline - yPos
		if yFlip {
			line -= ysize
			line *= -1
		}

		dataAddress := (0x8000 + (tileLocation * 16)) + uint16(line*2)
		data1 := ppu.MMU.Read(dataAddress)
		data2 := ppu.MMU.Read(dataAddress + 1)

		// Draw one line of sprite
		for tilePixel := 7; tilePixel >= 0; tilePixel-- {
			colourbit := tilePixel
			if xFlip {
				colourbit -= 7
				colourbit *= -1
			}
			colourNum := (utils.GetBit(data2, colourbit) << 1) | utils.GetBit(data1, colourbit)
			if colourNum == 0 {
				continue
			}
			pixel := xPos + (7 - tilePixel)
			ppu.Display.Draw(byte(pixel), byte(scanline), byte(colourNum))
		}
	}
}
