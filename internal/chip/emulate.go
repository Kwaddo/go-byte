package chip8

import (
	"fmt"
	"os"
	"time"
)

func (c *Chip8) LoadROM(filename string) error {
	rom, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	if len(rom) > 4096-0x200 {
		return fmt.Errorf("ROM too large")
	}

	fmt.Printf("Loading ROM: %s (size: %d bytes)\n", filename, len(rom))
	for i := 0; i < len(rom); i++ {
		c.Memory[i+0x200] = rom[i]
	}
	return nil
}

func (c *Chip8) EmulateCycle() {
	opcode := uint16(c.Memory[c.Pc])<<8 | uint16(c.Memory[c.Pc+1])
	fmt.Printf("Executing opcode: 0x%04X at PC: 0x%04X\n", opcode, c.Pc)

	// Decode opcode
	switch opcode & 0xF000 {
	case 0x0000:
		switch opcode & 0x00FF {
		case 0x00E0: // Clear screen
			for i := range c.Gfx {
				c.Gfx[i] = 0
			}
			c.DrawFlag = true
			c.Pc += 2
		case 0x00EE: // Return from subroutine
			c.Sp--
			c.Pc = c.Stack[c.Sp]
			c.Pc += 2
		}
	case 0x1000: // Jump to address NNN
		c.Pc = opcode & 0x0FFF
	case 0x2000: // Call subroutine at NNN
		c.Stack[c.Sp] = c.Pc
		c.Sp++
		c.Pc = opcode & 0x0FFF
	case 0x3000: // Skip if VX == NN
		x := (opcode & 0x0F00) >> 8
		nn := byte(opcode & 0x00FF)
		if c.Vreg[x] == nn {
			c.Pc += 4
		} else {
			c.Pc += 2
		}
	case 0x4000: // Skip if VX != NN
		x := (opcode & 0x0F00) >> 8
		nn := byte(opcode & 0x00FF)
		if c.Vreg[x] != nn {
			c.Pc += 4
		} else {
			c.Pc += 2
		}
	case 0x5000: // Skip if VX == VY
		x := (opcode & 0x0F00) >> 8
		y := (opcode & 0x00F0) >> 4
		if c.Vreg[x] == c.Vreg[y] {
			c.Pc += 4
		} else {
			c.Pc += 2
		}
	case 0x6000: // Set VX = NN
		x := (opcode & 0x0F00) >> 8
		c.Vreg[x] = byte(opcode & 0x00FF)
		c.Pc += 2
	case 0x7000: // Set VX += NN
		x := (opcode & 0x0F00) >> 8
		c.Vreg[x] += byte(opcode & 0x00FF)
		c.Pc += 2
	case 0x8000:
		x := (opcode & 0x0F00) >> 8
		y := (opcode & 0x00F0) >> 4
		switch opcode & 0x000F {
		case 0x0000: // Set VX = VY
			c.Vreg[x] = c.Vreg[y]
		case 0x0001: // Set VX = VX OR VY
			c.Vreg[x] |= c.Vreg[y]
		case 0x0002: // Set VX = VX AND VY
			c.Vreg[x] &= c.Vreg[y]
		case 0x0003: // Set VX = VX XOR VY
			c.Vreg[x] ^= c.Vreg[y]
		case 0x0004: // Set VX += VY, VF = carry
			if c.Vreg[y] > 0xFF-c.Vreg[x] {
				c.Vreg[0xF] = 1
			} else {
				c.Vreg[0xF] = 0
			}
			c.Vreg[x] += c.Vreg[y]
		case 0x0005: // Set VX -= VY, VF = !borrow
			if c.Vreg[y] > c.Vreg[x] {
				c.Vreg[0xF] = 0
			} else {
				c.Vreg[0xF] = 1
			}
			c.Vreg[x] -= c.Vreg[y]
		case 0x0006: // Set VX = VY >> 1, VF = LSB
			c.Vreg[0xF] = c.Vreg[x] & 0x1
			c.Vreg[x] >>= 1
		case 0x0007: // Set VX = VY - VX, VF = !borrow
			if c.Vreg[x] > c.Vreg[y] {
				c.Vreg[0xF] = 0
			} else {
				c.Vreg[0xF] = 1
			}
			c.Vreg[x] = c.Vreg[y] - c.Vreg[x]
		case 0x000E: // Set VX = VX << 1, VF = MSB
			c.Vreg[0xF] = (c.Vreg[x] & 0x80) >> 7
			c.Vreg[x] <<= 1
		}
		c.Pc += 2
	case 0x9000: // Skip if VX != VY
		x := (opcode & 0x0F00) >> 8
		y := (opcode & 0x00F0) >> 4
		if c.Vreg[x] != c.Vreg[y] {
			c.Pc += 4
		} else {
			c.Pc += 2
		}
	case 0xA000: // Set I = NNN
		c.Index = opcode & 0x0FFF
		c.Pc += 2
	case 0xB000: // Jump to NNN + V0
		c.Pc = (opcode & 0x0FFF) + uint16(c.Vreg[0])
	case 0xC000: // Set VX = random & NN
		x := (opcode & 0x0F00) >> 8
		c.Vreg[x] = byte(time.Now().UnixNano()) & byte(opcode&0x00FF)
		c.Pc += 2
	case 0xD000: // Draw sprite
		x := uint16(c.Vreg[(opcode&0x0F00)>>8])
		y := uint16(c.Vreg[(opcode&0x00F0)>>4])
		height := opcode & 0x000F
		c.Vreg[0xF] = 0

		for yline := uint16(0); yline < height; yline++ {
			pixel := c.Memory[c.Index+yline]
			for xline := uint16(0); xline < 8; xline++ {
				if (pixel & (0x80 >> xline)) != 0 {
					if x+xline < 64 && y+yline < 32 {
						pos := (y+yline)*64 + (x + xline)
						if c.Gfx[pos] == 1 {
							c.Vreg[0xF] = 1
						}
						c.Gfx[pos] ^= 1
					}
				}
			}
		}
		c.DrawFlag = true
		c.Pc += 2
	case 0xE000:
		x := (opcode & 0x0F00) >> 8
		switch opcode & 0x00FF {
		case 0x009E: // Skip if key VX pressed
			if c.Keys[c.Vreg[x]] != 0 {
				c.Pc += 4
			} else {
				c.Pc += 2
			}
		case 0x00A1: // Skip if key VX not pressed
			if c.Keys[c.Vreg[x]] == 0 {
				c.Pc += 4
			} else {
				c.Pc += 2
			}
		}
	case 0xF000:
		x := (opcode & 0x0F00) >> 8
		switch opcode & 0x00FF {
		case 0x0007: // Set VX = delay timer
			c.Vreg[x] = c.DelayTimer
			c.Pc += 2
		case 0x000A: // Wait for key press
			keyPress := false
			for i := byte(0); i < 16; i++ {
				if c.Keys[i] != 0 {
					c.Vreg[x] = i
					keyPress = true
				}
			}
			if !keyPress {
				return
			}
			c.Pc += 2
		case 0x0015: // Set delay timer = VX
			c.DelayTimer = c.Vreg[x]
			c.Pc += 2
		case 0x0018: // Set sound timer = VX
			c.SoundTimer = c.Vreg[x]
			c.Pc += 2
		case 0x001E: // Set I += VX
			c.Index += uint16(c.Vreg[x])
			c.Pc += 2
		case 0x0029: // Set I = sprite location for digit VX
			c.Index = uint16(c.Vreg[x]) * 5
			c.Pc += 2
		case 0x0033: // Store BCD of VX
			c.Memory[c.Index] = c.Vreg[x] / 100
			c.Memory[c.Index+1] = (c.Vreg[x] / 10) % 10
			c.Memory[c.Index+2] = c.Vreg[x] % 10
			c.Pc += 2
		case 0x0055: // Store V0-VX in memory starting at I
			for i := byte(0); i <= byte(x); i++ {
				c.Memory[c.Index+uint16(i)] = c.Vreg[i]
			}
			c.Pc += 2
		case 0x0065: // Load V0-VX from memory starting at I
			for i := byte(0); i <= byte(x); i++ {
				c.Vreg[i] = c.Memory[c.Index+uint16(i)]
			}
			c.Pc += 2
		}
	default:
		fmt.Printf("Unknown opcode: 0x%X\n", opcode)
		c.Pc += 2
	}
}
