package chip8

import (
	"github.com/veandco/go-sdl2/sdl"
)

func (c *Chip8) HandleInput() bool {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch t := event.(type) {
		case *sdl.QuitEvent:
			return false
		case *sdl.KeyboardEvent:
			if t.Type == sdl.KEYDOWN {
				switch t.Keysym.Sym {
				case sdl.K_1:
					c.Keys[0x1] = 1
				case sdl.K_2:
					c.Keys[0x2] = 1
				case sdl.K_3:
					c.Keys[0x3] = 1
				case sdl.K_4:
					c.Keys[0xC] = 1
				case sdl.K_q:
					c.Keys[0x4] = 1
				case sdl.K_w:
					c.Keys[0x5] = 1
				case sdl.K_e:
					c.Keys[0x6] = 1
				case sdl.K_r:
					c.Keys[0xD] = 1
				case sdl.K_a:
					c.Keys[0x7] = 1
				case sdl.K_s:
					c.Keys[0x8] = 1
				case sdl.K_d:
					c.Keys[0x9] = 1
				case sdl.K_f:
					c.Keys[0xE] = 1
				case sdl.K_z:
					c.Keys[0xA] = 1
				case sdl.K_x:
					c.Keys[0x0] = 1
				case sdl.K_c:
					c.Keys[0xB] = 1
				case sdl.K_v:
					c.Keys[0xF] = 1
				}
			} else if t.Type == sdl.KEYUP {
				switch t.Keysym.Sym {
				case sdl.K_1:
					c.Keys[0x1] = 0
				case sdl.K_2:
					c.Keys[0x2] = 0
				case sdl.K_3:
					c.Keys[0x3] = 0
				case sdl.K_4:
					c.Keys[0xC] = 0
				case sdl.K_q:
					c.Keys[0x4] = 0
				case sdl.K_w:
					c.Keys[0x5] = 0
				case sdl.K_e:
					c.Keys[0x6] = 0
				case sdl.K_r:
					c.Keys[0xD] = 0
				case sdl.K_a:
					c.Keys[0x7] = 0
				case sdl.K_s:
					c.Keys[0x8] = 0
				case sdl.K_d:
					c.Keys[0x9] = 0
				case sdl.K_f:
					c.Keys[0xE] = 0
				case sdl.K_z:
					c.Keys[0xA] = 0
				case sdl.K_x:
					c.Keys[0x0] = 0
				case sdl.K_c:
					c.Keys[0xB] = 0
				case sdl.K_v:
					c.Keys[0xF] = 0
				}
			}
		}
	}
	return true
}
