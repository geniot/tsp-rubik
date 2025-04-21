package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type ButtonCode = uint8
type ResourceKey = int

const (
	APP_NAME          = "TSP Rubik"
	APP_VERSION       = "0.1"
	TSP_SCREEN_WIDTH  = 1280
	TSP_SCREEN_HEIGHT = 720
)

var (
	COLOR_RED    = sdl.Color{R: 192, G: 64, B: 64, A: 255}
	COLOR_GREEN  = sdl.Color{R: 64, G: 192, B: 64, A: 255}
	COLOR_GRAY   = sdl.Color{R: 192, G: 192, B: 192, A: 255}
	COLOR_WHITE  = sdl.Color{R: 255, G: 255, B: 255, A: 255}
	COLOR_PURPLE = sdl.Color{R: 255, G: 0, B: 255, A: 255}
	COLOR_YELLOW = sdl.Color{R: 255, G: 255, B: 0, A: 255}
	COLOR_BLUE   = sdl.Color{R: 0, G: 255, B: 255, A: 255}
	COLOR_BLACK  = sdl.Color{R: 0, G: 0, B: 0, A: 255}

	BACKGROUND_COLOR = COLOR_BLACK
)

const (
	RESOURCE_BGR_KEY           = ResourceKey(iota)
	RESOURCE_CIRCLE_YELLOW_KEY = ResourceKey(iota)
	RESOURCE_CROSS_YELLOW_KEY  = ResourceKey(iota)
)

const (
	BUTTON_CODE_MENU  = ButtonCode(5)
	BUTTON_CODE_START = ButtonCode(6)
	BUTTON_CODE_UP    = ButtonCode(11)
	BUTTON_CODE_DOWN  = ButtonCode(12)
	BUTTON_CODE_LEFT  = ButtonCode(13)
	BUTTON_CODE_RIGHT = ButtonCode(14)
)
