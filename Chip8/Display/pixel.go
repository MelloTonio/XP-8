package Display

import (
	"fmt"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

// Fonte em hexadecimal
var FontSet = [80]byte{
	0xF0, 0x90, 0x90, 0x90, 0xF0, // 0
	0x20, 0x60, 0x20, 0x20, 0x70, // 1
	0xF0, 0x10, 0xF0, 0x80, 0xF0, // 2
	0xF0, 0x10, 0xF0, 0x10, 0xF0, // 3
	0x90, 0x90, 0xF0, 0x10, 0x10, // 4
	0xF0, 0x80, 0xF0, 0x10, 0xF0, // 5
	0xF0, 0x80, 0xF0, 0x90, 0xF0, // 6
	0xF0, 0x10, 0x20, 0x40, 0x40, // 7
	0xF0, 0x90, 0xF0, 0x90, 0xF0, // 8
	0xF0, 0x90, 0xF0, 0x10, 0xF0, // 9
	0xF0, 0x90, 0xF0, 0x90, 0x90, // A
	0xE0, 0x90, 0xe0, 0x90, 0xE0, // B
	0xF0, 0x80, 0x80, 0x80, 0x80, // C
	0xF0, 0x90, 0x90, 0x90, 0xE0, // D
	0xF0, 0x80, 0xF0, 0x80, 0xF0, // E
	0xF0, 0x80, 0xF0, 0x80, 0x80, // F
}

const windowX float64 = 64
const windowY float64 = 32

const screenWidth float64 = 1024
const screenHeight float64 = 768

type Window struct {
	*pixelgl.Window
	KeyMap   map[uint16]pixelgl.Button
	KeysDown [16]*time.Ticker
}

// https://github.com/faiface/pixel/wiki/Creating-a-Window
func NewWindow() (*Window, error) {
	cfg := pixelgl.WindowConfig{
		Title:  "XP-8",
		Bounds: pixel.R(0, 0, 1024, 768), // Rectangle minX, minY, maxX, maxY
		VSync:  true,
	}
	w, err := pixelgl.NewWindow(cfg)
	if err != nil {
		return nil, fmt.Errorf("error creating new window: %v", err)
	}
	// Teclas utilizadas no chip-8
	km := map[uint16]pixelgl.Button{
		0x1: pixelgl.Key1, 0x2: pixelgl.Key2,
		0x3: pixelgl.Key3, 0xC: pixelgl.Key4,
		0x4: pixelgl.KeyQ, 0x5: pixelgl.KeyW,
		0x6: pixelgl.KeyE, 0xD: pixelgl.KeyR,
		0x7: pixelgl.KeyA, 0x8: pixelgl.KeyS,
		0x9: pixelgl.KeyD, 0xE: pixelgl.KeyF,
		0xA: pixelgl.KeyZ, 0x0: pixelgl.KeyX,
		0xB: pixelgl.KeyC, 0xF: pixelgl.KeyV,
	}
	return &Window{
		Window:   w,
		KeyMap:   km,
		KeysDown: [16]*time.Ticker{},
	}, nil
}

func (w *Window) DrawGraphics(gfx [64 * 32]byte) {
	w.Clear(colornames.Black)
	imDraw := imdraw.New(nil)
	imDraw.Color = pixel.RGB(1, 1, 1)
	width, height := screenWidth/windowX, screenHeight/windowY

	for i := 0; i < 64; i++ {
		for j := 0; j < 32; j++ {
			// If the gfx byte in question is turned off,
			// continue and skip drawing the rectangle
			if gfx[(31-j)*64+i] == 0 {
				continue
			}
			imDraw.Push(pixel.V(width*float64(i), height*float64(j))) // Adiciona um pixel desenhado nas coordenadas x,y
			imDraw.Push(pixel.V(width*float64(i)+width, height*float64(j)+height))
			imDraw.Rectangle(0)
		}
	}

	imDraw.Draw(w)
	w.Update()
}
