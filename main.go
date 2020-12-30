package main

import (
	"fmt"
	"os"

	"github.com/mellotonio/go-chip8/Chip8"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("incorrect usage. Usage: `chippy path/to/rom`")
		return
	}

	pathToROM := os.Args[1]

	chip_8 := Chip8.Start()

	// Tenta colocar a ROM no Chip8
	if err := chip_8.LoadROM(pathToROM); err != nil {
		fmt.Printf("\nerror loading ROM: %v\n", err)
		os.Exit(1)
	}

}
