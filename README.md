# go-byte

A CHIP-8 emulator written in Go using SDL2 for graphics and input handling.

## Description

go-byte is a complete CHIP-8 emulator/interpreter that can run classic CHIP-8 ROMs. The emulator features:

- Full CHIP-8 instruction set implementation
- 64x64 pixel display with configurable scaling
- Keyboard input mapping
- Configurable CPU and display refresh rates
- Sound timer support
- Built-in debugging output

## Requirements

- Go 1.24 or higher
- SDL2 library

## Key Mappings

CHIP-8 Key   Keyboard
---------    ---------
1 2 3 C      1 2 3 4
4 5 6 D      Q W E R  
7 8 9 E      A S D F
A 0 B F      Z X C V

### Installing SDL2

- **Ubuntu/Debian**: `sudo apt-get install libsdl2-dev`  
- **macOS**: `brew install sdl2`
- **Windows**: Download SDL2 development libraries from [libsdl.org](https://libsdl.org/)

## Building

```sh
go build -o chip8 cmd/main.go
```
