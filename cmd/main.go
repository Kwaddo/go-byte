package main

import (
	"fmt"
	"os"
	"time"

	ch "chip8/internal/chip"
	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: chip8 <ROM file>")
		return
	}

	chip8 := ch.NewChip8()

	if err := chip8.InitDisplay(); err != nil {
		fmt.Printf("Failed to initialize display: %v\n", err)
		return
	}
	defer sdl.Quit()
	defer chip8.Window.Destroy()
	defer chip8.Renderer.Destroy()

	if err := chip8.LoadROM(os.Args[1]); err != nil {
		fmt.Printf("Failed to load ROM: %v\n", err)
		return
	}

	fmt.Printf("ROM loaded successfully\n")

	lastCycle := time.Now()
	lastFrame := time.Now()
	frames := 0
	lastFPSUpdate := time.Now()

	for {
		if !chip8.HandleInput() {
			break
		}

		// CPU cycle timing
		if time.Since(lastCycle) >= chip8.CycleDelay {
			chip8.EmulateCycle()
			lastCycle = time.Now()

			if chip8.DelayTimer > 0 {
				chip8.DelayTimer--
			}
			if chip8.SoundTimer > 0 {
				chip8.SoundTimer--
			}
		}

		// Frame timing
		if time.Since(lastFrame) >= chip8.FrameDelay {
			chip8.DrawScreen()
			lastFrame = time.Now()
			frames++
		}

		// FPS counter
		if time.Since(lastFPSUpdate) >= time.Second {
			fmt.Printf("FPS: %d\n", frames)
			frames = 0
			lastFPSUpdate = time.Now()
		}

		// Prevent CPU hogging
		time.Sleep(time.Millisecond)
	}
}
