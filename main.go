package main

import (
	"fmt"
	"os"

	"github.com/mellotonio/go-chip8/Chip8"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("incorrect usage. Usage: `chipo path/to/rom`")
		os.Exit(1)
	}

	pathToROM := os.Args[1]

	chip_8, err := Chip8.Start(pathToROM)

	if err != nil {
		fmt.Printf("\nerror creating a new chip-8: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(chip_8)

	for {
		chip_8.MachineCycle()
	}

}
