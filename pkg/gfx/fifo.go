package gfx

import "errors"

// FIFO structure for shifting out pixels to the display or enqueue CPU
// micro-operations. We use interface{} as a placeholder for any type.
// We also make it a fixed size that works for CPU operations as well as PPU
// pixels.
type FIFO struct {
	fifo [16]interface{} // Array of values in the FIFO.
	out  int             // Current index of the tail (output) of the FIFO.
	in   int             // Current index of the head (input) of the FIFO.
	len  int             // Current length of the FIFO.
}

// Pre-defined errors to only instantiate them once.
var errFIFOOverflow = errors.New("FIFO buffer overflow")
var errFIFOUnderrun = errors.New("FIFO buffer underrun")

// Push an item to the FIFO.
func (f *FIFO) Push(item interface{}) error {
	if f.len == len(f.fifo) {
		return errFIFOOverflow
	}
	f.fifo[f.in] = item
	f.in = (f.in + 1) % len(f.fifo)
	f.len++
	return nil
}

// Pop an item out of the FIFO.
func (f *FIFO) Pop() (item interface{}, err error) {
	if f.len == 0 {
		return 0, errFIFOUnderrun
	}
	item = f.fifo[f.out]
	f.out = (f.out + 1) % len(f.fifo)
	f.len--
	return item, nil
}

// Size returns the current amount of items in the FIFO.
func (f *FIFO) Size() int {
	return f.len
}

// Clear resets internal indexes, effectively clearing out the FIFO.
func (f *FIFO) Clear() {
	f.in, f.out, f.len = 0, 0, 0
}
