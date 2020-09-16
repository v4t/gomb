package main

import (
	"io/ioutil"
	"log"
	"math"
	"os"

	"github.com/v4t/gomb/pkg/emulator"
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

	gb := emulator.Create()
	gb.Start(rom)

	os.Exit(0)
}

func loadRom(fname string) ([]byte, error) {
	data, err := ioutil.ReadFile(fname)
	if err != nil {
		return nil, err
	}
	return data, nil
}
