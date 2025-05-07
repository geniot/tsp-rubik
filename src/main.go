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

const (
	blackKey = iota
	greenKey
	redKey
	blueKey
	orangeKey
	whiteKey
	yellowKey
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
	allColors = map[int]rl.Color{
		blackKey:  black,
		greenKey:  green,
		redKey:    red,
		blueKey:   blue,
		orangeKey: orange,
		whiteKey:  white,
		yellowKey: yellow,
	}
)

var (
	colorTextures = make(map[int]rl.Texture2D)
)

func main() {
	var (
		gamePadId  int32 = 0
		shouldExit       = false
		camera           = rl.Camera3D{}
		angle            = float32(0)
	)

	//rl.SetConfigFlags(rl.FlagMsaa4xHint)
	rl.SetConfigFlags(rl.FlagVsyncHint) //should be set before window initialization!

	rl.InitWindow(1280, 720, "TrimUI Rubik")
	rl.SetWindowMonitor(1)
	rl.InitAudioDevice()

	rl.SetClipPlanes(0.5, 100)
	rl.DisableBackfaceCulling()
	rl.Color4f(1, 1, 1, 1)

	prepareTextures()

	camera.Position = rl.NewVector3(10.0, 10.0, 10.0)
	camera.Target = rl.NewVector3(0.0, 0.0, 0.0)
	camera.Up = rl.NewVector3(0.0, 1.0, 0.0)
	camera.Fovy = 45.0
	camera.Projection = rl.CameraPerspective

	width, height, length := float32(2), float32(2), float32(2)

	for !rl.WindowShouldClose() && !shouldExit {
		rl.UpdateCamera(&camera, rl.CameraThirdPerson)
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		rl.BeginMode3D(camera)

		rl.PushMatrix()
		rl.Rotatef(angle, 1, 0, 0)
		angle += 1

		for i := range CubeDescriptors {

			cube := CubeDescriptors[i]
			x, y, z := cube.x*width, cube.y*height, cube.z*length

			rl.Begin(rl.Quads)
			{
				//front
				rl.SetTexture(colorTextures[greenKey].ID)
				rl.Normal3f(0.0, 0.0, 1.0)
				rl.TexCoord2f(0.0, 0.0)
				rl.Vertex3f(x-width/2, y-height/2, z+length/2)
				rl.TexCoord2f(1.0, 0.0)
				rl.Vertex3f(x+width/2, y-height/2, z+length/2)
				rl.TexCoord2f(1.0, 1.0)
				rl.Vertex3f(x+width/2, y+height/2, z+length/2)
				rl.TexCoord2f(0.0, 1.0)
				rl.Vertex3f(x-width/2, y+height/2, z+length/2)
				//back
				rl.SetTexture(colorTextures[cube.backColor].ID)
				rl.Normal3f(0.0, 0.0, -1.0)
				rl.TexCoord2f(0.0, 0.0)
				rl.Vertex3f(x-width/2, y-height/2, z-length/2)
				rl.TexCoord2f(1.0, 0.0)
				rl.Vertex3f(x+width/2, y-height/2, z-length/2)
				rl.TexCoord2f(1.0, 1.0)
				rl.Vertex3f(x+width/2, y+height/2, z-length/2)
				rl.TexCoord2f(0.0, 1.0)
				rl.Vertex3f(x-width/2, y+height/2, z-length/2)
				//up
				rl.SetTexture(colorTextures[cube.upColor].ID)
				rl.Normal3f(0.0, 1.0, 0.0)
				rl.TexCoord2f(0.0, 0.0)
				rl.Vertex3f(x-width/2, y+height/2, z+length/2)
				rl.TexCoord2f(1.0, 0.0)
				rl.Vertex3f(x+width/2, y+height/2, z+length/2)
				rl.TexCoord2f(1.0, 1.0)
				rl.Vertex3f(x+width/2, y+height/2, z-length/2)
				rl.TexCoord2f(0.0, 1.0)
				rl.Vertex3f(x-width/2, y+height/2, z-length/2)
				//down
				rl.SetTexture(colorTextures[cube.downColor].ID)
				rl.Normal3f(0.0, -1.0, 0.0)
				rl.TexCoord2f(0.0, 0.0)
				rl.Vertex3f(x-width/2, y-height/2, z+length/2)
				rl.TexCoord2f(1.0, 0.0)
				rl.Vertex3f(x+width/2, y-height/2, z+length/2)
				rl.TexCoord2f(1.0, 1.0)
				rl.Vertex3f(x+width/2, y-height/2, z-length/2)
				rl.TexCoord2f(0.0, 1.0)
				rl.Vertex3f(x-width/2, y-height/2, z-length/2)
				//right
				rl.SetTexture(colorTextures[cube.rightColor].ID)
				rl.Normal3f(1.0, 0.0, 0.0)
				rl.TexCoord2f(0.0, 0.0)
				rl.Vertex3f(x+width/2, y-height/2, z+length/2)
				rl.TexCoord2f(1.0, 0.0)
				rl.Vertex3f(x+width/2, y-height/2, z-length/2)
				rl.TexCoord2f(1.0, 1.0)
				rl.Vertex3f(x+width/2, y+height/2, z-length/2)
				rl.TexCoord2f(0.0, 1.0)
				rl.Vertex3f(x+width/2, y+height/2, z+length/2)
				//left
				rl.SetTexture(colorTextures[cube.leftColor].ID)
				rl.Normal3f(-1.0, 0.0, 0.0)
				rl.TexCoord2f(0.0, 0.0)
				rl.Vertex3f(x-width/2, y-height/2, z+length/2)
				rl.TexCoord2f(1.0, 0.0)
				rl.Vertex3f(x-width/2, y-height/2, z-length/2)
				rl.TexCoord2f(1.0, 1.0)
				rl.Vertex3f(x-width/2, y+height/2, z-length/2)
				rl.TexCoord2f(0.0, 1.0)
				rl.Vertex3f(x-width/2, y+height/2, z+length/2)
			}
			rl.End()
		}

		rl.PopMatrix()
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
