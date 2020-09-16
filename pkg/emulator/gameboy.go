package emulator

import (
	"github.com/v4t/gomb/pkg/cpu"
	"github.com/v4t/gomb/pkg/graphics"
	"github.com/v4t/gomb/pkg/memory"
)

// Gameboy emulator.
type Gameboy struct {
	CPU     *cpu.CPU
	MMU     *memory.MMU
	PPU     *graphics.PPU
	Display *graphics.Display
}

// Create is constructor for gameboy emulator.
func Create() *Gameboy {
	display := graphics.Init()
	cpu := cpu.InitializeCPU()
	ppu := &graphics.PPU{MMU: cpu.MMU, Display: display}
	return &Gameboy{
		CPU:     cpu,
		PPU:     ppu,
		MMU:     cpu.MMU,
		Display: display,
	}
}

// Start gameboy emulator.
func (gb *Gameboy) Start(rom []byte) {
	gb.MMU.LoadRom(rom)

	gb.Display.Run(func() {
		gb.Display.Init()

		for !gb.Display.Closed() {
			gb.Update()
		}
	})
}

// Update gameboy state.
func (gb *Gameboy) Update() {
	maxCycles := 69905
	currentCycles := 0
	for currentCycles < maxCycles {
		cycles := gb.CPU.Execute()
		currentCycles += cycles
		gb.PPU.Execute(cycles)
	}
	gb.Display.RenderImage()
}

// func (gb *Gameboy) UpdateTimers(cycles int) {
// 	//    DoDividerRegister(cycles);

// 	// the clock must be enabled to update the clock
// 	if IsClockEnabled() {
// 		m_TimerCounter -= cycles

// 		// enough cpu clock cycles have happened to update the timer
// 		if m_TimerCounter <= 0 {
// 			// reset m_TimerTracer to the correct value
// 			SetClockFreq()

// 			// timer about to overflow
// 			if ReadMemory(TIMA) == 255 {
// 				WriteMemory(TIMA, ReadMemory(TMA))
// 				RequestInterupt(2)
// 			} else {
// 				WriteMemory(TIMA, ReadMemory(TIMA)+1)
// 			}
// 		}
// 	}
// }

// func (gb *Gameboy) DoDividerRegister(cycles int) {
// 	m_DividerRegister += cycles
// 	if m_DividerCounter >= 255 {
// 		m_DividerCounter = 0
// 		m_Rom[0xFF04]++
// 	}
// }

// func (gb *Gameboy) IsClockEnabled() bool {
// 	return utils.TestBit(gb.MMU.Read(TMC), 2)
// }
