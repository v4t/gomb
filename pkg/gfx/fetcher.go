package gfx

import "github.com/v4t/gomb/pkg/memory"

// FetcherState is an enum-like type to define all admissible states for the PPU's fetcher.
type FetcherState uint8

// Possible Fetcher states. Values don't really matter, they start from zero.
const (
	ReadTileID FetcherState = iota
	ReadTileData0
	ReadTileData1
	PushToFIFO
)

// Fetcher reading tile data from VRAM and pushing pixels to a FIFO queue.
type Fetcher struct {
	FIFO     FIFO         // Pixel FIFO that the PPU will read.
	MMU      *memory.MMU  // Reference to the global MMU.
	ticks    int          // Clock cycle counter for timings.
	state    FetcherState // Current state of our state machine.
	mapAddr  uint16       // Start address of BG/Windows map row.
	dataAddr uint16       // Start address of Sprite/BG tile data.
	tileLine uint8        // Y offset (in pixels) in the tile.

	// Index of the tile to read in the current row of the background map.
	tileIndex uint8

	tileID   uint8    // Tile number in the tilemap.
	tileData [8]uint8 // Pixel data for one row of the fetched tile.
}

// Start fetching a line of pixels starting from the given tile address in the
// background map. Here, tileLine indicates which row of pixels to pick from
// each tile we read.
func (f *Fetcher) Start(mapAddr uint16, tileLine uint8) {
	f.tileIndex = 0
	f.mapAddr = mapAddr
	f.tileLine = tileLine
	f.state = ReadTileID

	// Clear FIFO between calls, as it may still contain leftover tile data
	// from the very end of the previous scanline.
	f.FIFO.Clear()
}

// Tick advances the fetcher's state machine one step.
func (f *Fetcher) Tick() {
	// The Fetcher runs at half the speed of the PPU (every 2 clock cycles).
	f.ticks++
	if f.ticks < 2 {
		return
	}
	f.ticks = 0 // Reset tick counter and execute next state.

	switch f.state {
	case ReadTileID:
		// Read the tile's number from the background map. This will be used
		// in the next states to find the address where the tile's actual pixel
		// data is stored in memory.
		f.tileID = f.MMU.Read(f.mapAddr + uint16(f.tileIndex))
		f.state = ReadTileData0

	case ReadTileData0:
		f.ReadTileLine(0, f.tileID, f.tileLine, &f.tileData)
		f.state = ReadTileData1

	case ReadTileData1:
		f.ReadTileLine(1, f.tileID, f.tileLine, &f.tileData)
		f.state = PushToFIFO

	case PushToFIFO:
		if f.FIFO.Size() <= 8 {
			// We stored pixel bits from least significant (rightmost) to most
			// (leftmost) in the data array, so we must push them in reverse
			// order.
			for i := 7; i >= 0; i-- {
				f.FIFO.Push(f.tileData[i])
			}
			// Advance to the next tile in the map's row.
			f.tileIndex++
			f.state = ReadTileID
		}
	}
}

// ReadTileLine updates the fetcher's internal pixel buffer with tile data
// depending on the current state. Each pixel needs 2 bits of information,
// which are read in two separate steps.
func (f *Fetcher) ReadTileLine(bitPlane uint8, tileID uint8, tileLine uint8,
	data *[8]uint8) {
	// A tile's graphical data takes 16 bytes (2 bytes per row of 8 pixels).
	// Tile data starts at address 0x8000 so we first compute an offset to
	// find out where the data for the tile we want starts.
	offset := 0x8000 + (uint16(tileID) * 16)

	// Then, from that starting offset, we compute the final address to read
	// by finding out which of the 8-pixel rows of the tile we want to display.
	addr := offset + (uint16(tileLine) * 2)

	// Finally, read the first or second byte of graphical data depending on
	// what state we're in.
	pixelData := f.MMU.Read(addr + uint16(bitPlane))
	for bitPos := uint(0); bitPos <= 7; bitPos++ {
		// Separate each bit fom the data byte we just read. Each of these bits
		// is half of a pixel's color value.
		if bitPlane == 0 {
			// Least significant bit, replace the previous value.
			data[bitPos] = (pixelData >> bitPos) & 1
		} else {
			// Most significant bit, update the previous value.
			data[bitPos] |= ((pixelData >> bitPos) & 1) << 1
		}
	}
}
