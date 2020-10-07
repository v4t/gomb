package utils

// SetBit sets bit at given position.
func SetBit(value byte, pos int) byte {
	value |= (1 << pos)
	return value
}

// ResetBit clears bit at given position.
func ResetBit(value byte, pos int) byte {
	value &= ^(1 << pos)
	return value
}

// TestBit checks if bit at given position is set.
func TestBit(value byte, pos int) bool {
	result := value & (1 << pos)
	return result != 0
}

// GetBit return bit value at given position.
func GetBit(value byte, pos int) int {
	return int((value >> pos) & 1)
}
