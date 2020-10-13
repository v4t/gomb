package emulator

import (
	"github.com/v4t/gomb/pkg/processor"
	"github.com/v4t/gomb/pkg/utils"
)

// Addresses for timer registers.
const (
	DIV  uint16 = 0xff04
	TIMA uint16 = 0xff05
	TMA  uint16 = 0xff06
	TMC  uint16 = 0xff07
)

// Timer handles timer register updates and memory operations.
type Timer struct {
	Interrupts *processor.Interrupts

	// Internal counters
	counter int
	divider int

	// Registers
	div  byte // Divider register
	tima byte // Timer counter
	tma  byte // Timer modulo
	tmc  byte // Timer control
}

// Read value from one of the timer registers.
func (timer *Timer) Read(address uint16) byte {
	switch address {
	case DIV:
		return timer.div
	case TIMA:
		return timer.tima
	case TMA:
		return timer.tma
	case TMC:
		return timer.tmc
	default:
		panic("Attempted to read timer registers with invalid address.")
	}
}

// Write value to one of the timer registers.
func (timer *Timer) Write(address uint16, value byte) {
	switch address {
	case DIV:
		timer.div = 0
	case TIMA:
		timer.tima = value
	case TMA:
		timer.tma = value
	case TMC:
		timer.tmc = value
	default:
		panic("Attempted to write to timer registers with invalid address.")
	}
}

// Update timer registers.
func (timer *Timer) Update(cycles int) {
	timer.updateDividerRegister(cycles)
	if timer.Enabled() {
		timer.counter += cycles
		freq := timer.getClockFrequency()
		for timer.counter >= freq {
			timer.counter -= freq
			if timer.tima == 0xff {
				timer.tima = timer.tma
				timer.Interrupts.SetInterrupt(processor.TimerInterrupt)
			} else {
				timer.tima++
			}
		}
	}
}

// Enabled checks if timer is runnign from TMC register.
func (timer *Timer) Enabled() bool {
	return utils.TestBit(timer.tmc, 2)
}

// getClockFrequency returns current frequency based on TMC register.
func (timer *Timer) getClockFrequency() int {
	switch timer.tmc & 0x3 {
	case 0:
		return 1024 // 4096Hz
	case 1:
		return 16 // 262144Hz
	case 2:
		return 64 // 65536Hz
	default:
		return 256 // 16384Hz
	}
}

// updateDividerRegister performs the periodical divider register update.
// DIV register does a periodical increment at 16482Hz frequency which is equivalent to256 CPU clock cycles.
func (timer *Timer) updateDividerRegister(cycles int) {
	timer.divider += cycles
	if timer.divider >= 255 {
		timer.divider -= 255
		timer.div++
	}
}
