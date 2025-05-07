package main

import (
	"bufio"
	"bytes"
	"github.com/fogleman/gg"
	rl "github.com/gen2brain/raylib-go/raylib"
	"math"
)

//var (
//	//go:embed media/*
//	mediaList embed.FS
//)

// TSP button codes
const (
	upCode     = 1
	rightCode  = 2
	downCode   = 3
	leftCode   = 4
	xCode      = 5
	aCode      = 6
	bCode      = 7
	yCode      = 8
	l1Code     = 9
	l2Code     = 10
	r1Code     = 11
	r2Code     = 12
	selectCode = 13
	menuCode   = 14
	startCode  = 15
)

// https://www.schemecolor.com/rubik-cube-colors.php
var (
	black     = rl.Color{R: 0, G: 0, B: 0, A: 255}
	green     = rl.Color{R: 0, G: 155, B: 72, A: 255}
	red       = rl.Color{R: 185, G: 0, B: 0, A: 255}
	blue      = rl.Color{R: 0, G: 69, B: 173, A: 255}
	orange    = rl.Color{R: 255, G: 89, B: 0, A: 255}
	white     = rl.Color{R: 255, G: 255, B: 255, A: 255}
	yellow    = rl.Color{R: 255, G: 213, B: 0, A: 255}
	allColors = []rl.Color{black, green, red, blue, orange, white, yellow}
)

var (
	colorTextures = make(map[rl.Color]rl.Texture2D)
)

func main() {
	var (
		gamePadId  int32 = 0
		shouldExit       = false
		camera           = rl.Camera3D{}
	)

	rl.SetConfigFlags(rl.FlagMsaa4xHint)
	rl.SetConfigFlags(rl.FlagVsyncHint)
	rl.SetClipPlanes(0.5, 100)

	rl.InitWindow(1280, 720, "TrimUI Rubik")
	rl.SetWindowMonitor(1)
	rl.InitAudioDevice()

	prepareTextures()

	camera.Position = rl.NewVector3(10.0, 10.0, 10.0)
	camera.Target = rl.NewVector3(0.0, 0.0, 0.0)
	camera.Up = rl.NewVector3(0.0, 1.0, 0.0)
	camera.Fovy = 45.0
	camera.Projection = rl.CameraPerspective

	//cubePosition := rl.NewVector3(0.0, 0.0, 0.0)

	width, height, length := float32(2), float32(2), float32(2)

	for !rl.WindowShouldClose() && !shouldExit {
		rl.UpdateCamera(&camera, rl.CameraThirdPerson)
		rl.BeginDrawing()
		rl.DisableBackfaceCulling()
		rl.ClearBackground(rl.RayWhite)

		rl.BeginMode3D(camera)

		rl.PushMatrix()
		//rl.Rotatef(angle, 1, 0, 0)
		//angle += 1

		for i := range CubeDescriptors {

			cube := CubeDescriptors[i]
			x, y, z := cube.x*width, cube.y*height, cube.z*length

			rl.Begin(rl.Quads)
			{
				//front-back (z)
				rl.Color4ub(cube.frontColor.R, cube.frontColor.G, cube.frontColor.B, cube.frontColor.A)
				rl.Vertex3f(x-width/2, y-height/2, z+length/2)
				rl.Vertex3f(x+width/2, y-height/2, z+length/2)
				rl.Vertex3f(x+width/2, y+height/2, z+length/2)
				rl.Vertex3f(x-width/2, y+height/2, z+length/2)
				rl.Color4ub(cube.backColor.R, cube.backColor.G, cube.backColor.B, cube.backColor.A)
				rl.Vertex3f(x-width/2, y-height/2, z-length/2)
				rl.Vertex3f(x+width/2, y-height/2, z-length/2)
				rl.Vertex3f(x+width/2, y+height/2, z-length/2)
				rl.Vertex3f(x-width/2, y+height/2, z-length/2)
				//up-down (y)
				rl.Color4ub(cube.upColor.R, cube.upColor.G, cube.upColor.B, cube.upColor.A)
				rl.Vertex3f(x-width/2, y+height/2, z+length/2)
				rl.Vertex3f(x+width/2, y+height/2, z+length/2)
				rl.Vertex3f(x+width/2, y+height/2, z-length/2)
				rl.Vertex3f(x-width/2, y+height/2, z-length/2)
				rl.Color4ub(cube.downColor.R, cube.downColor.G, cube.downColor.B, cube.downColor.A)
				rl.Vertex3f(x-width/2, y-height/2, z+length/2)
				rl.Vertex3f(x+width/2, y-height/2, z+length/2)
				rl.Vertex3f(x+width/2, y-height/2, z-length/2)
				rl.Vertex3f(x-width/2, y-height/2, z-length/2)
				//left-right (x)
				rl.Color4ub(cube.leftColor.R, cube.leftColor.G, cube.leftColor.B, cube.leftColor.A)
				rl.Vertex3f(x-width/2, y-height/2, z+length/2)
				rl.Vertex3f(x-width/2, y-height/2, z-length/2)
				rl.Vertex3f(x-width/2, y+height/2, z-length/2)
				rl.Vertex3f(x-width/2, y+height/2, z+length/2)
				rl.Color4ub(cube.rightColor.R, cube.rightColor.G, cube.rightColor.B, cube.rightColor.A)
				rl.Vertex3f(x+width/2, y-height/2, z+length/2)
				rl.Vertex3f(x+width/2, y-height/2, z-length/2)
				rl.Vertex3f(x+width/2, y+height/2, z-length/2)
				rl.Vertex3f(x+width/2, y+height/2, z+length/2)
			}
			rl.End()

			//rl.DrawCubeWires(rl.NewVector3(x, y, z), width, height, length, rl.Black)
		}

		rl.PopMatrix()
		//rl.DrawGrid(10, 1.0)

		rl.EndMode3D()

		//exit
		if rl.IsGamepadButtonDown(gamePadId, menuCode) && rl.IsGamepadButtonDown(gamePadId, startCode) {
			shouldExit = true //see WindowShouldClose, it checks if KeyEscape pressed or Close icon pressed
		}
		rl.EndDrawing()
	}
	rl.CloseWindow()
}

func prepareTextures() {
	var (
		width  = 100
		height = 100
	)
	for _, color := range allColors {
		pngBytes := makePNG(width, height, color)
		colorTextures[color] = rl.LoadTextureFromImage(rl.LoadImageFromMemory(".png", pngBytes, int32(len(pngBytes))))
	}
}

func makePNG(width int, height int, color rl.Color) []byte {
	bytesBuffer := new(bytes.Buffer)
	dc := gg.NewContext(width, height)
	dc.DrawRectangle(0, 0, float64(width), float64(height))
	dc.SetRGBA255(int(color.R), int(color.G), int(color.B), int(color.A))
	dc.Fill()
	w := bufio.NewWriter(bytesBuffer)
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
