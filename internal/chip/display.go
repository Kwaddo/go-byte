package chip8

import (
	"github.com/veandco/go-sdl2/sdl"
	inf "chip8/internal/constants"
)

func (c *Chip8) InitDisplay() error {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		return err
	}

	window, err := sdl.CreateWindow("CHIP-8 Emulator",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		64*10, 32*10,
		sdl.WINDOW_SHOWN)
	if err != nil {
		return err
	}

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		return err
	}

	c.Window = window
	c.Renderer = renderer
	return nil
}


func (c *Chip8) DrawScreen() {
	if !c.DrawFlag {
		return
	}

	c.Renderer.SetDrawColor(0, 0, 0, 255)
	c.Renderer.Clear()

	c.Renderer.SetDrawColor(255, 255, 255, 255)
	for y := 0; y < inf.DisplayHeight; y++ {
		for x := 0; x < inf.DisplayWidth; x++ {
			if c.Gfx[y*inf.DisplayWidth+x] != 0 {
				rect := sdl.Rect{
					X: int32(x * inf.Scale),
					Y: int32(y * inf.Scale),
					W: inf.Scale,
					H: inf.Scale,
				}
				c.Renderer.FillRect(&rect)
			}
		}
	}
	c.Renderer.Present()
	c.DrawFlag = false
}
