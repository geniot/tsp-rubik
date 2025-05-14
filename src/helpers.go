package main

import (
	"bufio"
	"bytes"
	"github.com/fogleman/gg"
	rl "github.com/gen2brain/raylib-go/raylib"
	"math"
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

func orPanic(err interface{}) {
	switch v := err.(type) {
	case error:
		if v != nil {
			panic(err)
		}
	case bool:
		if !v {
			panic("condition failed: != true")
		}
	}
}

func orPanicRes[T any](res T, err interface{}) T {
	orPanic(err)
	return res
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}

func If[T any](cond bool, vTrue, vFalse T) T {
	if cond {
		return vTrue
	}
	return vFalse
}

func newInts(n, v int) []int {
	s := make([]int, n)
	for i := range s {
		s[i] = v
	}
	return s
}

func copyInts(fromArray []int, toArray []int, fromOffset int, toOffset int, length int) {
	for i := 0; i < length; i++ {
		toArray[toOffset+i] = fromArray[fromOffset+i]
	}
}
