package main

import (
	"bufio"
	"bytes"
	"github.com/fogleman/gg"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	G = iota
	R
	B
	O
	W
	Y
	LB
	BL
)

// colors
// https://www.schemecolor.com/rubik-cube-colors.php
var (
	black      = rl.Color{R: 0, G: 0, B: 0, A: 255}
	lightBlack = rl.Color{R: 77, G: 77, B: 77, A: 255}
	lightGray  = rl.Color{R: 211, G: 211, B: 211, A: 255}
	green      = rl.Color{R: 0, G: 155, B: 72, A: 255}
	red        = rl.Color{R: 185, G: 0, B: 0, A: 255}
	blue       = rl.Color{R: 0, G: 69, B: 173, A: 255}
	orange     = rl.Color{R: 255, G: 89, B: 0, A: 255}
	white      = rl.Color{R: 255, G: 255, B: 255, A: 255}
	yellow     = rl.Color{R: 255, G: 213, B: 0, A: 255}
	allColors  = map[int]rl.Color{
		LB: lightBlack,
		BL: black,
		G:  green,
		R:  red,
		B:  blue,
		O:  orange,
		W:  white,
		Y:  yellow,
	}
)

var (
	colorTextures         = make(map[int]rl.Texture2D)
	selectedColorTextures = make(map[int]rl.Texture2D)
)

func prepareTextures() {
	var (
		width   = 100
		height  = 100
		padding = 5
	)
	for colorKey, color := range allColors {
		pngBytes := makePng(width, height, padding, black, color, false)
		colorTextures[colorKey] = rl.LoadTextureFromImage(rl.LoadImageFromMemory(".png", pngBytes, int32(len(pngBytes))))
	}
	for colorKey, color := range allColors {
		pngBytes := makePng(width, height, padding, black, color, true)
		selectedColorTextures[colorKey] = rl.LoadTextureFromImage(rl.LoadImageFromMemory(".png", pngBytes, int32(len(pngBytes))))
	}
}

func makePng(width int, height int, padding int, paddingColor rl.Color, color rl.Color, isSelected bool) []byte {
	selectionColor := white
	subSelectionColor := black
	subSelectionPadding := 2
	bytesBuffer := new(bytes.Buffer)
	dc := gg.NewContext(width, height)
	dc.DrawRectangle(0, 0, float64(width), float64(height))
	dc.SetRGBA255(int(paddingColor.R), int(paddingColor.G), int(paddingColor.B), int(paddingColor.A))
	dc.Fill()
	if isSelected {
		dc.DrawRectangle(float64(padding), float64(padding), float64(width-padding*2), float64(height-padding*2))
		dc.SetRGBA255(int(selectionColor.R), int(selectionColor.G), int(selectionColor.B), int(selectionColor.A))
		dc.Fill()
		dc.DrawRectangle(float64(padding*2-subSelectionPadding),
			float64(padding*2-subSelectionPadding),
			float64(width-padding*4+subSelectionPadding*2),
			float64(height-padding*4+subSelectionPadding*2))
		dc.SetRGBA255(int(subSelectionColor.R), int(subSelectionColor.G), int(subSelectionColor.B), int(subSelectionColor.A))
		dc.Fill()
		dc.DrawRectangle(float64(padding*2), float64(padding*2), float64(width-padding*4), float64(height-padding*4))
		dc.SetRGBA255(int(color.R), int(color.G), int(color.B), int(color.A))
		dc.Fill()
	} else {
		dc.DrawRectangle(float64(padding), float64(padding), float64(width-padding*2), float64(height-padding*2))
		dc.SetRGBA255(int(color.R), int(color.G), int(color.B), int(color.A))
		dc.Fill()
	}
	w := bufio.NewWriter(bytesBuffer)
	//dc.SavePNG("out.png")
	orPanic(dc.EncodePNG(w))
	orPanic(w.Flush())
	return bytesBuffer.Bytes()
}
