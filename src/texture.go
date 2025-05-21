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
	//regular
	for colorKey, color := range allColors {
		pngBytes := makePng(width, height, padding, black, color, false)
		colorTextures[colorKey] = rl.LoadTextureFromImage(rl.LoadImageFromMemory(".png", pngBytes, int32(len(pngBytes))))
	}
	//selected
	for colorKey, color := range allColors {
		pngBytes := makePng(width, height, padding, black, color, true)
		selectedColorTextures[colorKey] = rl.LoadTextureFromImage(rl.LoadImageFromMemory(".png", pngBytes, int32(len(pngBytes))))
	}
}

func makePng(width int, height int, padding int, paddingColor rl.Color, color rl.Color, isSelected bool) []byte {
	selectionColor := white
	subSelectionColor := black
	bytesBuffer := new(bytes.Buffer)
	dc := gg.NewContext(width, height)
	subPaddingFactor := 10
	if isSelected {
		drawColors(dc, width, height,
			[]int{0, padding, padding + padding/subPaddingFactor, padding + padding/subPaddingFactor + padding/subPaddingFactor},
			[]rl.Color{paddingColor, selectionColor, subSelectionColor, color})
	} else {
		drawColors(dc, width, height,
			[]int{0, padding},
			[]rl.Color{paddingColor, color})
	}
	w := bufio.NewWriter(bytesBuffer)
	//dc.SavePNG("out.png")
	orPanic(dc.EncodePNG(w))
	orPanic(w.Flush())
	return bytesBuffer.Bytes()
}

func drawColors(dc *gg.Context, width int, height int, paddings []int, colors []rl.Color) {
	for index, color := range colors {
		x := float64(index * paddings[index])
		y := float64(index * paddings[index])
		w := float64(width - index*paddings[index]*2)
		h := float64(height - index*paddings[index]*2)
		dc.DrawRectangle(x, y, w, h)
		//dc.DrawCircle(x+w/2, y+h/2, w/2)
		dc.SetRGBA255(int(color.R), int(color.G), int(color.B), int(color.A))
		dc.Fill()
	}
}
