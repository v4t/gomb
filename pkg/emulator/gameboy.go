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
	Joypad  *graphics.Joypad
}

// Create is constructor for gameboy emulator.
func Create() *Gameboy {
	display := graphics.Init()
	cpu := cpu.InitializeCPU()
	ppu := graphics.InitPPU(cpu.MMU, display)
	joypad := graphics.NewJoypad(cpu.MMU)
	return &Gameboy{
		CPU:     cpu,
		PPU:     ppu,
		MMU:     cpu.MMU,
		Display: display,
		Joypad:  joypad,
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
		cpu.HandleInterrupts(gb.CPU)
	}
	gb.Display.RenderImage()
	gb.Display.ProcessInput(gb.Joypad)
}
