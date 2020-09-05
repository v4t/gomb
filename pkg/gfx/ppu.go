package gfx

// PPUState is an enum-like type to define all admissible states for the PPU.
type PPUState uint8

// Possible PPU states. Values don't really matter, they start from zero.
const (
	OAMSearch PPUState = iota
	PixelTransfer
	HBlank
	VBlank
)

// LCD Control register bits (see https://golang.org/ref/spec#Iota)
const (
	// Bit 0 - BG/Window Display/Priority     (0=Off, 1=On)
	LCDCBGDisplay uint8 = 1 << iota
	// Bit 1 - OBJ (Sprite) Display Enable    (0=Off, 1=On)
	LCDCSpriteDisplayEnable
	// Bit 2 - OBJ (Sprite) Size              (0=8x8, 1=8x16)
	LCDCSpriteSize
	// Bit 3 - BG Tile Map Display Select     (0=9800-9BFF, 1=9C00-9FFF)
	LCDCBGTileMapDisplayeSelect
	// Bit 4 - BG & Window Tile Data Select   (0=8800-97FF, 1=8000-8FFF)
	LCDCBGWindowTileDataSelect
	// Bit 5 - Window Display Enable          (0=Off, 1=On)
	LCDCWindowDisplayEnable
	// Bit 6 - Window Tile Map Display Select (0=9800-9BFF, 1=9C00-9FFF)
	LCDCWindowTileMapDisplayeSelect
	// Bit 7 - LCD Display Enable             (0=Off, 1=On)
	LCDCDisplayEnable
)

// PPU that continuously scans video RAM, enqueues pixels to display in a FIFO
// and pops them out to an arbitrary display. This version only implements the
// few registers required to display the background map (no scrolling). For
// synchronicity with the CPU, the PPU is implemented as a state machine whose
// Tick() method is called every clock cycle.
// Since it holds hardware registers, the PPU also implements the Addressable
// interface.
type PPU struct {
	LCDC uint8 // LCD Control register.
	LY   uint8 // Number of the scanline currently being displayed.
	BGP  uint8 // Background map tiles palette.

	// Fetcher runs at half the PPU's speed and fetches pixel data from the
	// background map's tiles, according to the current scanline. It also holds
	// the FIFO pixel queue that we will be writing out to the display.
	Fetcher Fetcher

	// Screen object implementing our Display interface above. Just a text
	// console for now, but we can swap it for a proper window later.
	Screen *Display

	state PPUState // Current state of the state machine.
	ticks uint     // Clock ticks counter for the current line.
	x     uint8    // Current count of pixels already output in a scanline.
}

// NewPPU returns an instance of PPU using the given display object.
func NewPPU(screen *Display) *PPU {
	ppu := PPU{Screen: screen}
	return &ppu
}

// Contains return true if the requested address is LCDC, LY or BGP.
// We'll soon come up with a way to automatically map an address to a register.
func (p *PPU) Contains(addr uint16) bool {
	return addr == 0xff40 || addr == 0xff44 || addr == 0xff47
}

// Read returns the current value in the LCDC, LY or BGP register.
func (p *PPU) Read(addr uint16) uint8 {
	switch addr {
	case 0xff40:
		return p.LCDC
	case 0xff44:
		return p.LY
	case 0xff47:
		return p.BGP
	}
	panic("invalid PPU read address")
}

// Write updates the value in LCDC or BGP and ignore writes to LY.
func (p *PPU) Write(addr uint16, value uint8) {
	switch addr {
	case 0xff40:
		p.LCDC = value
	case 0xff44:
		// LY is read-only.
	case 0xff47:
		p.BGP = value
	default:
		panic("invalid PPU write address")
	}
}

// Tick advances the PPU state one step.
func (p *PPU) Tick() {
	// Check if the screen should be turned on or off depending on LCDC value.
	if !p.Screen.Enabled() {
		if p.LCDC&LCDCDisplayEnable != 0 {
			p.Screen.Enable()
			p.state = OAMSearch
		} else {
			return
		}
	} else {
		if p.LCDC&LCDCDisplayEnable == 0 {
			p.LY = 0
			p.x = 0
			p.Screen.Disable()
			return
		}
	}

	p.ticks++

	switch p.state {
	case OAMSearch:
		// In this state, the PPU would scan the OAM (Objects Attribute Memory)
		// from 0xfe00 to 0xfe9f to mix sprite pixels in the current line later.
		// This always takes 40 clock ticks.

		//
		// OAM search will happen here (when implemented).
		//

		if p.ticks == 40 {
			// Move to Pixel Transfer state. Initialize the fetcher to start
			// reading background tiles from VRAM. We don't do scrolling yet
			// and the boot ROM does nothing fancy with map addresses, so we
			// just give the fetcher the base address of the row of tiles we
			// need in video RAM:
			//
			// - The background map is 32×32 tiles big.
			// - The viewport starts at the top left of that map when LY is 0.
			// - Each tile is 8×8 pixels.
			//
			// In the present case, we only need to figure out in which row of
			// the background map our current line (at position LY) is. Then we
			// start fetching pixels from that row's address in VRAM, and for
			// each tile, we can tell which 8-pixel line to fetch by computing
			// LY modulo 8.
			p.x = 0
			tileLine := p.LY % 8
			tileMapRowAddr := 0x9800 + (uint16(p.LY/8) * 32)
			p.Fetcher.Start(tileMapRowAddr, tileLine)

			p.state = PixelTransfer
		}

	case PixelTransfer:
		// Fetch pixel data into our pixel FIFO.
		p.Fetcher.Tick()

		// Stop here if the FIFO isn't holding at least 8 pixels. This will
		// be used to mix in sprite data when we implement these. It also
		// guarantees the FIFO will always have data to Pop() later.
		if p.Fetcher.FIFO.Size() <= 8 {
			return
		}

		// Put a pixel from the FIFO on screen. We take a value between 0 and 3
		// and use it to look up an actual color (yet another value between 0
		// and 3 where 0 is the lightest color and 3 the darkest) in the BGP
		// register.
		pixelColor, _ := p.Fetcher.FIFO.Pop()

		// BGP contains four consecutive 2-bit values. We take the one whose
		// index is given by pixelColor by shifting those 2-bit values right
		// that many times and only keeping the rightmost 2 bits. I initially
		// got the order wrong and fixed it thanks to coffee-gb.
		paletteColor := (p.BGP >> (pixelColor.(uint8) * 2)) & 3
		p.Screen.Write(paletteColor)

		p.x++
		if p.x == 160 {
			p.state = HBlank
		}

	case HBlank:
		if p.ticks == 456 {
			p.ticks = 0
			p.LY++
			if p.LY == 144 {
				p.Screen.DrawDisplayImage()
				p.state = VBlank
			} else {
				p.state = OAMSearch
			}
		}

	case VBlank:
		if p.ticks == 456 {
			p.ticks = 0
			p.LY++
			if p.LY == 153 {
				p.LY = 0
				p.state = OAMSearch
			}
		}
	}
}
