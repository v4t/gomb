package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"

	"github.com/v4t/gomb/pkg/hardware"
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
	if len(rom) > math.MaxUint16 {
		log.Fatalf("Rom doesn't fit in memory")
	}

	cpu := hardware.InitializeCPU()
	cpu.MMU.LoadRom(rom)

	fmt.Println("STARTING")
	for i := 0; i < 1000000; i++ {
		fmt.Printf("PC %x\n", cpu.PC)
		cpu.Execute()
		if cpu.PC > 0x237 {
			panic("foo")
		}
	}

	os.Exit(0)
}

func loadRom(fname string) ([]byte, error) {
	data, err := ioutil.ReadFile(fname)
	if err != nil {
		return nil, err
	}
	return data, nil
}
