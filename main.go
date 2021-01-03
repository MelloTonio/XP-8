package main

import (
	"fmt"
	"os"

	"github.com/faiface/pixel/pixelgl"
	"github.com/mellotonio/go-chip8/Chip8"
)

func main() {
	fmt.Printf("WELCOME to XP-8\n\nPONG COMMANDS:\nPlayer1: Keyboard [1] [Q]\nPlayer2: Keyboard [4] [R]\n\nSPACE INVADERS: Keyboard [Q] [W] [E]\n1. Pong\n2. Space Invaders\n\n")
	pixelgl.Run(mainFunc) // Pixelgl precisa do controle da função principal

}

func mainFunc() {
	var number int

	var pathToROM string

	_, err := fmt.Scanf("%d", &number)

	if number == 1 {
		pathToROM = "./Chip8/roms/pong.ch8"
	} else if number == 2 {
		pathToROM = "./Chip8/roms/Space Invaders [David Winter].ch8"
	}

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
