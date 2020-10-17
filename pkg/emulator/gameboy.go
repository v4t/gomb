package emulator

import (
	"time"

	"github.com/v4t/gomb/pkg/cartridge"
	"github.com/v4t/gomb/pkg/graphics"
	"github.com/v4t/gomb/pkg/memory"
	"github.com/v4t/gomb/pkg/processor"
)

// FPS rate for emulator. (Sped up a bit from the actual 60 fps)
const FPS = 100

// MaxCycles represents clock cycles executed for each frame.
const MaxCycles = 69905

// Gameboy emulator.
type Gameboy struct {
	CPU     *processor.CPU
	MMU     *memory.MMU
	PPU     *graphics.PPU
	Timer   *Timer
	Display *graphics.Display
	Joypad  *graphics.Joypad
}

// NewGameboy is constructor for gameboy emulator.
func NewGameboy() *Gameboy {
	display := &graphics.Display{}
	cpu := processor.NewCPU()
	ppu := graphics.NewPPU(cpu.MMU, display)
	joypad := graphics.NewJoypad()
	timer := &Timer{}

	cpu.MMU.Timer = timer
	cpu.MMU.Input = joypad
	joypad.Interrupts = cpu.Interrupts
	cpu.MMU.Interrupts = cpu.Interrupts
	ppu.Interrupts = cpu.Interrupts
	timer.Interrupts = cpu.Interrupts
	return &Gameboy{
		CPU:     cpu,
		PPU:     ppu,
		MMU:     cpu.MMU,
		Display: display,
		Joypad:  joypad,
		Timer:   timer,
	}
}

// Start gameboy emulator.
func (gb *Gameboy) Start(cart *cartridge.Cartridge) {
	gb.MMU.Cartridge = cart

	gb.Display.Run(func() {
		gb.Display.Initialize()
		t := time.NewTicker(time.Second / FPS)
		for !gb.Display.Closed() {
			select {
			case <-t.C:
				gb.Update()
			}
		}
	})
}

// Update gameboy state.
func (gb *Gameboy) Update() {
	currentCycles := 0
	for currentCycles < MaxCycles {
		cycles := 4
		if !gb.CPU.Halted {
			cycles = gb.CPU.Execute()
		}
		currentCycles += cycles
		gb.Timer.Update(cycles)
		gb.PPU.Execute(cycles)
		gb.CPU.Interrupts.Resolve(gb.CPU)
	}
	gb.Display.RenderImage()
	gb.Display.ProcessInput(gb.Joypad)
}
