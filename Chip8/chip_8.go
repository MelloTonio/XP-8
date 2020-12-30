package Chip8

import (
	"io/ioutil"

	"github.com/mellotonio/go-chip8/Chip8/Display"
)

// Uso de memória
// 0x000-0x1FF - Reservado para o interpretador do Chip-8 -> 0 ~ 512 (bytes)
// 0x050-0x0A0 - "Used for the built in 4x5 pixel font set (0-F)"" -> 80 ~ 160 (bytes)
// 0x200-0xFFF - Reservado para os programas e funcionalidades -> 512 ~ 4095 (bytes)

type chip_8_VM struct {
	opcode          uint16      // Referência de instrução do processador
	memory          [4096]byte  // O Chip-8, originalmente, é capaz de acessar 4096 bytes de RAM (4KB)
	Vx              [16]byte    // Registradores de proposito geral, Vx aonde x é um hexadecimal z (0 até F)
	index           uint16      // Registrador de indice
	program_counter uint16      // Usado para guardar o endereço atual da instrução que está sendo executada (0x000 - 0 => 0xFFF - 4095)
	stack           [16]uint16  // Stack para "acumular" instruções
	stack_pointer   uint16      // Registro que guarda o ultimo endereço requisitado na pilha
	gfx             [32][8]byte // Pixels da tela
	keypad          [16]byte    // "16-key hexadecimal keypad for input"
}

// Inicialização do Chip8 com a fonte inicializada nos primeiros 80 bytes
func Start() *chip_8_VM {
	var memory [4096]byte

	for i := 0; i < Display.FontOffset; i++ {
		memory[i] = Display.FontSet[i]
	}

	return &chip_8_VM{
		opcode:          0,
		memory:          memory,
		Vx:              [16]byte{},
		index:           0,
		program_counter: 0x200, // Começa no byte 512, já reservado para o inicio dos programas
		stack:           [16]uint16{},
		stack_pointer:   0,
		gfx:             [32][8]byte{},
		keypad:          [16]byte{},
	}

}

// Pega o caminho da ROM e carrega ela no Chip8
func (chip_8 *chip_8_VM) LoadROM(path string) error {
	rom, err := ioutil.ReadFile(path)

	if err != nil {
		return err
	}

	if len(rom) >= 3585 {
		panic("ERROR: ROM TOO LARGE - MAX SIZE: 3584") // Se a ROM ultrapassar o espaço dedicado para o interpretador ocorrerá um "panic"
	}

	for i := 0; i < len(rom); i++ {
		chip_8.memory[Display.FontOffset+i] = rom[i] // Memoria começa 80 + x, tirando espaço reservado para as fontes (80 bytes)
	}

	return nil
}
