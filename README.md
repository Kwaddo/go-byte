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

| CHIP-8 Key | Keyboard |
|------------|----------|
| 1 | 1 |
| 2 | 2 |
| 3 | 3 |
| C | 4 |
| 4 | Q |
| 5 | W |
| 6 | E |
| D | R |
| 7 | A |
| 8 | S |
| 9 | D |
| E | F |
| A | Z |
| 0 | X |
| B | C |
| F | V |

### Installing SDL2

- **Ubuntu/Debian**: `sudo apt-get install libsdl2-dev`  
- **macOS**: `brew install sdl2`
- **Windows**: Download SDL2 development libraries from [libsdl.org](https://libsdl.org/)

## Building

```sh
go build -o chip8 cmd/main.go
```
