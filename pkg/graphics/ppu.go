package graphics

import "github.com/v4t/gomb/pkg/memory"

// PPU represents pixel processing unit.
type PPU struct {
	scanlineCounter int
	Display         *Display
	MMU             *memory.MMU
}

// Execute ppu cycle.
func (ppu *PPU) Execute(cycles int) {

	if ppu.Display.IsEnabled() {
		ppu.scanlineCounter -= cycles
	}

	if ppu.scanlineCounter <= 0 {

		var currentLine = memory.Memory[0xff44] + 1
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

// DrawScanLine draws scanline...
func (ppu *PPU) DrawScanLine() {

}
