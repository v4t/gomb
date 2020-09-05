package main

import (
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/v4t/gomb/pkg/gfx"
	"github.com/v4t/gomb/pkg/hardware"
	"github.com/v4t/gomb/pkg/memory"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Program requires ROM file as parameter")
	}
	romFile := os.Args[1]

	rom, err := loadRom(romFile)
	if err != nil {
		log.Fatalf("Error when loading ROM: %v", err)
	}

	screen := &gfx.Display{}
	screen.Run(func() {
		screen.Init()
		ppu := gfx.NewPPU(screen) // Covers 0xff40, 0xff44 and 0xff47

		mmu := memory.MMU{PPU: ppu}
		mmu.LoadBoot(rom)
		ppu.Fetcher.MMU = &mmu
		cpu := hardware.InitializeCPU(&mmu)
		// cpu := CPU{mmu: mmu}

		for !screen.Closed() {
			cpu.Execute()
			ppu.Tick()
		}
		time.Sleep(3 * time.Second)
	})
}

func loadRom(fname string) ([]byte, error) {
	data, err := ioutil.ReadFile(fname)
	if err != nil {
		return nil, err
	}
	return data, nil
}
