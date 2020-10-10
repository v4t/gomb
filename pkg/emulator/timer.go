package emulator

// import "github.com/v4t/gomb/pkg/utils"

// var timer int

// const (
// 	DIV  uint16 = 0xff04
// 	TIMA uint16 = 0xff05
// 	TMA  uint16 = 0xff06
// 	TMC  uint16 = 0xff07
// )

// func UpdateTimers(cycles int) {
// 	//    DoDividerRegister(cycles)

// 	// the clock must be enabled to update the clock
// 	if IsClockEnabled() {
// 		timer -= cycles

// 		// enough cpu clock cycles have happened to update the timer
// 		if timer <= 0 {
// 			// reset m_TimerTracer to the correct value
// 			SetClockFreq()

// 			// timer about to overflow
// 			if ReadMemory(TIMA) == 255 {
// 				// WriteMemory(TIMA, ReadMemory(TMA))
// 				// RequestInterupt(2)
// 			} else {
// 				// WriteMemory(TIMA, ReadMemory(TIMA)+1)
// 			}
// 		}
// 	}
// }

// var dividerCounter byte

// func DoDividerRegister(cycles int) {
// 	m_DividerRegister += cycles
// 	if dividerCounter >= 255 {
// 		dividerCounter = 0
// 		m_Rom[0xFF04]++
// 	}
// }

// func IsClockEnabled() bool {
// 	return utils.TestBit(ReadMemory(TMC), 2)
// }

// func SetClockFreq() {
// 	freq := GetClockFreq()
// 	switch freq {
// 	case 0:
// 		timer = 1024
// 	case 1:
// 		timer = 16
// 	case 2:
// 		timer = 64
// 	case 3:
// 		timer = 256
// 	}
// }

// func GetClockFreq() byte {
// 	return ReadMemory(TMC) & 0x3
// }
