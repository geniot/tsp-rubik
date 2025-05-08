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
		pngBytes := makePng(width, height, padding, black, color)
		colorTextures[colorKey] = rl.LoadTextureFromImage(rl.LoadImageFromMemory(".png", pngBytes, int32(len(pngBytes))))
	}
}

func makePng(width int, height int, padding int, paddingColor rl.Color, color rl.Color) []byte {
	bytesBuffer := new(bytes.Buffer)
	dc := gg.NewContext(width, height)
	dc.DrawRectangle(0, 0, float64(width), float64(height))
	dc.SetRGBA255(int(paddingColor.R), int(paddingColor.G), int(paddingColor.B), int(paddingColor.A))
	dc.Fill()
	dc.DrawRectangle(float64(padding), float64(padding), float64(width-padding*2), float64(height-padding*2))
	dc.SetRGBA255(int(color.R), int(color.G), int(color.B), int(color.A))
	dc.Fill()
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
