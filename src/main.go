package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"math"
)

//var (
//	//go:embed media/*
//	mediaList embed.FS
//)

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

type CubeDescriptor struct {
	OffsetX   int32
	OffsetY   int32
	ImageName string
}

var VolumeImages = []CubeDescriptor{
	{
		OffsetX:   24,
		OffsetY:   70,
		ImageName: "volume.png",
	},
}

func main() {
	var (
		gamePadId    int32 = 0
		shouldExit         = false
		camera             = rl.Camera3D{}
		angle              = float32(0)
		transparency       = uint8(255)
	)

	rl.SetConfigFlags(rl.FlagMsaa4xHint)
	rl.SetConfigFlags(rl.FlagVsyncHint)

	rl.InitWindow(1280, 720, "TrimUI Rubik")
	rl.SetWindowMonitor(1)
	rl.InitAudioDevice()

	camera.Position = rl.NewVector3(10.0, 10.0, 10.0)
	camera.Target = rl.NewVector3(0.0, 0.0, 0.0)
	camera.Up = rl.NewVector3(0.0, 1.0, 0.0)
	camera.Fovy = 45.0
	camera.Projection = rl.CameraPerspective

	//cubePosition := rl.NewVector3(0.0, 0.0, 0.0)

	x, y, z := float32(0), float32(0), float32(0)
	width, height, length := float32(2), float32(2), float32(2)

	for !rl.WindowShouldClose() && !shouldExit {
		//rl.UpdateCamera(&camera, rl.CameraFirstPerson)
		rl.BeginDrawing()
		rl.DisableBackfaceCulling()
		rl.ClearBackground(rl.RayWhite)

		rl.BeginMode3D(camera)

		rl.PushMatrix()
		rl.Rotatef(angle, 1, 0, 0)
		angle += 1

		rl.Begin(rl.Quads)
		//front-back (z)
		rl.Normal3f(0, 0, 1)
		rl.Color4ub(255, 0, 0, transparency) //red
		rl.Vertex3f(x-width/2, y-height/2, z+length/2)
		rl.Vertex3f(x+width/2, y-height/2, z+length/2)
		rl.Vertex3f(x+width/2, y+height/2, z+length/2)
		rl.Vertex3f(x-width/2, y+height/2, z+length/2)
		rl.Normal3f(0, 0, -1)
		rl.Color4ub(0, 255, 0, transparency) //green
		rl.Vertex3f(x-width/2, y-height/2, z-length/2)
		rl.Vertex3f(x+width/2, y-height/2, z-length/2)
		rl.Vertex3f(x+width/2, y+height/2, z-length/2)
		rl.Vertex3f(x-width/2, y+height/2, z-length/2)
		//left-right (x)
		rl.Normal3f(0, 1, 0)
		rl.Color4ub(0, 0, 255, transparency) //blue
		rl.Vertex3f(x-width/2, y-height/2, z+length/2)
		rl.Vertex3f(x-width/2, y-height/2, z-length/2)
		rl.Vertex3f(x-width/2, y+height/2, z-length/2)
		rl.Vertex3f(x-width/2, y+height/2, z+length/2)
		rl.Normal3f(0, -1, 0)
		rl.Color4ub(255, 255, 0, transparency) //yellow
		rl.Vertex3f(x+width/2, y-height/2, z+length/2)
		rl.Vertex3f(x+width/2, y-height/2, z-length/2)
		rl.Vertex3f(x+width/2, y+height/2, z-length/2)
		rl.Vertex3f(x+width/2, y+height/2, z+length/2)
		//up-down (y)
		rl.Normal3f(1, 0, 0)
		rl.Color4ub(255, 0, 255, transparency) //purple
		rl.Vertex3f(x-width/2, y+height/2, z+length/2)
		rl.Vertex3f(x+width/2, y+height/2, z+length/2)
		rl.Vertex3f(x+width/2, y+height/2, z-length/2)
		rl.Vertex3f(x-width/2, y+height/2, z-length/2)
		rl.Normal3f(-1, 0, 0)
		rl.Color4ub(0, 255, 255, transparency) //turquoise
		rl.Vertex3f(x-width/2, y-height/2, z+length/2)
		rl.Vertex3f(x+width/2, y-height/2, z+length/2)
		rl.Vertex3f(x+width/2, y-height/2, z-length/2)
		rl.Vertex3f(x-width/2, y-height/2, z-length/2)
		rl.End()

		rl.PopMatrix()

		//rl.DrawCube(cubePosition, 2.0, 2.0, 2.0, rl.Red)
		//rl.DrawCubeWires(cubePosition, 2.0, 2.0, 2.0, rl.Maroon)

		rl.DrawGrid(10, 1.0)

		rl.EndMode3D()

		//exit
		if rl.IsGamepadButtonDown(gamePadId, menuCode) && rl.IsGamepadButtonDown(gamePadId, startCode) {
			shouldExit = true //see WindowShouldClose, it checks if KeyEscape pressed or Close icon pressed
		}
		rl.EndDrawing()
	}
	rl.CloseWindow()
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
