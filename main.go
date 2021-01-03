package main

import (
	"fmt"
	"os"

	"github.com/faiface/pixel/pixelgl"
	"github.com/mellotonio/go-chip8/Chip8"
)

func main() {

	pixelgl.Run(mainFunc) // Pixelgl precisa do controle da função principal

}

func mainFunc() {

	pathToROM := "./Chip8/roms/pong.ch8"

	// chip_8, err := Chip8.Start(pathToROM)

	chip_8, err := Chip8.Start(pathToROM)

	if err != nil {
		fmt.Printf("\nerror creating a new chip-8 VM: %v\n", err)
		os.Exit(1)
	}

	go chip_8.Run()
	go chip_8.ManageAudio()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	//	ticker := time.NewTicker(time.Second / 60)
	//		defer ticker.Stop()
	<-chip_8.Shutdown
}
