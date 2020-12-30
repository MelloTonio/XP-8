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
	opcode          uint16        // Referência de instrução do processador
	memory          [4096]byte    // O Chip-8, originalmente, é capaz de acessar 4096 bytes de RAM (4KB)
	Vx              [16]byte      // Registradores de proposito geral, Vx aonde x é um hexadecimal z (0 até F)
	index           uint16        // Registrador de indice
	program_counter uint16        // Usado para guardar o endereço atual da instrução que está sendo executada (0x000 - 0 => 0xFFF - 4095)
	stack           [16]uint16    // Stack para "acumular" instruções
	stack_pointer   uint16        // Registro que guarda o ultimo endereço requisitado na pilha
	gfx             [64 * 32]byte // Pixels da tela
	key             [16]byte      // "16-key hexadecimal keypad for input"
	drawFloag       bool
}

// Inicialização do Chip8 com a fonte inicializada nos primeiros 80 bytes
func Start(pathToROM string) (*chip_8_VM, error) {

	chip8_INIT := chip_8_VM{
		memory:          [4096]byte{},
		Vx:              [16]byte{},
		program_counter: 0x200, // Começa no byte 512, já reservado para o inicio dos programas
		stack:           [16]uint16{},
		gfx:             [64 * 32]byte{},
		key:             [16]byte{},
	}

	chip8_INIT.loadFontSet()

	// Tenta iniciar a ROM
	if err := chip8_INIT.LoadROM(pathToROM); err != nil {
		return nil, err
	}

	return &chip8_INIT, nil

}

// Carrega a font nos primeiros 80 bytes de memoria
func (chip_8 *chip_8_VM) loadFontSet() {
	for i := 0; i < 80; i++ {
		chip_8.memory[i] = Display.FontSet[i]
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
		chip_8.memory[0x50+i] = rom[i] // Memoria começa 0x50 (80) + x, tirando espaço reservado para as fontes (80 bytes)
	}

	return nil
}
