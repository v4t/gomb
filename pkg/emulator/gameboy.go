package emulator

import (
	"time"

	"github.com/v4t/gomb/pkg/cartridge"
	"github.com/v4t/gomb/pkg/graphics"
	"github.com/v4t/gomb/pkg/memory"
	"github.com/v4t/gomb/pkg/processor"
)

// FPS rate for emulator.
const FPS = 60

// Gameboy emulator.
type Gameboy struct {
	CPU     *processor.CPU
	MMU     *memory.MMU
	PPU     *graphics.PPU
	Timer   *Timer
	Display *graphics.Display
	Joypad  *graphics.Joypad
}

// Create is constructor for gameboy emulator.
func Create() *Gameboy {
	display := graphics.Init()
	cpu := processor.InitializeCPU()
	ppu := graphics.InitPPU(cpu.MMU, display)
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
		gb.Display.Init()
		t := time.NewTicker(time.Second / FPS)
		for !gb.Display.Closed() {
			select {
			case <-t.C:
				gb.Update()
			}
		}
	})
}

// MaxCycles represents clock cycles executed for each frame.
const MaxCycles = 69905

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
