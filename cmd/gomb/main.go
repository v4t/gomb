package main

import (
	"io/ioutil"
	"log"
	"math"
	"os"
)

var memory = make([]byte, math.MaxUint16)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Program requires ROM file as parameter")
	}
	romFile := os.Args[1]

	rom, err := loadRom(romFile)
	if err != nil {
		log.Fatalf("Error when loading ROM: %v", err)
	}
	if len(rom) > len(memory) {
		log.Fatalf("Rom doesn't fit in memory")
	}

	copy(memory[:], rom)
	os.Exit(0)
}

func loadRom(fname string) ([]byte, error) {
	data, err := ioutil.ReadFile(fname)
	if err != nil {
		return nil, err
	}
	return data, nil
}
