package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/v4t/gomb/pkg/cartridge"
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

	cart := cartridge.NewCartridge(rom)
	fmt.Println(cart)

	gb := emulator.NewGameboy()
	gb.Start(cart)
	os.Exit(0)
}

func loadRom(fname string) ([]byte, error) {
	data, err := ioutil.ReadFile(fname)
	if err != nil {
		return nil, err
	}
	return data, nil
}
