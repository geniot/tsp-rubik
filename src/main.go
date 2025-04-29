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

func main() {
	rl.SetConfigFlags(rl.FlagMsaa4xHint)
	rl.SetConfigFlags(rl.FlagVsyncHint)
	rl.InitWindow(1280, 720, "TrimUI Rubik")
	rl.InitAudioDevice()

	var (
		gamePadId  int32 = 0
		shouldExit       = false
	)

	camera := rl.Camera3D{}
	camera.Position = rl.NewVector3(10.0, 10.0, 10.0)
	camera.Target = rl.NewVector3(0.0, 0.0, 0.0)
	camera.Up = rl.NewVector3(0.0, 1.0, 0.0)
	camera.Fovy = 45.0
	camera.Projection = rl.CameraPerspective

	cubePosition := rl.NewVector3(0.0, 0.0, 0.0)

	for !rl.WindowShouldClose() && !shouldExit {
		//rl.UpdateCamera(&camera, rl.CameraThirdPerson) // Update camera with free camera mode

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		rl.BeginMode3D(camera)

		rl.DrawCube(cubePosition, 2.0, 2.0, 2.0, rl.Red)
		rl.DrawCubeWires(cubePosition, 2.0, 2.0, 2.0, rl.Maroon)

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
