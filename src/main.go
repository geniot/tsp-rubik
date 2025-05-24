package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	rotationSpeed  = float32(3)
	cubeSideLength = 2
)

// TSP button codes
const (
	noCode = iota
	upCode
	rightCode
	downCode
	leftCode
	xCode
	aCode
	bCode
	yCode
	l1Code
	l2Code
	r1Code
	r2Code
	selectCode
	menuCode
	startCode
)

func main() {

	//rl.SetConfigFlags(rl.FlagMsaa4xHint)
	rl.SetConfigFlags(rl.FlagVsyncHint) //should be set before window initialization!
	//rl.SetTargetFPS(60)

	rl.InitWindow(1280, 720, "TrimUI Rubik")
	rl.SetWindowMonitor(0)
	rl.InitAudioDevice()

	prepareTextures()

	var (
		cubeSize         = 3 //currently only 3 is supported :)
		gamePadId  int32 = 0
		shouldExit       = false
		camera           = rl.Camera3D{}
		cube             = NewCube(cubeSize)
	)

	rl.SetClipPlanes(0.5, 100) //see https://github.com/raysan5/raylib/issues/4917
	rl.DisableBackfaceCulling()

	camera.Position = rl.NewVector3(10, 10, 10)
	camera.Target = rl.NewVector3(0.0, 0.0, 0.0)
	camera.Up = rl.NewVector3(0.0, 1.0, 0.0)
	camera.Fovy = 40.0
	camera.Projection = rl.CameraPerspective

	for !rl.WindowShouldClose() && !shouldExit {
		//rl.UpdateCamera(&camera, rl.CameraOrbital)

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		rl.Color4f(1, 1, 1, 1)

		rl.BeginMode3D(camera)
		cube.update()
		cube.draw()
		//rl.DrawGrid(10, 1)
		rl.EndMode3D()

		//exit
		if rl.IsGamepadButtonDown(gamePadId, menuCode) && rl.IsGamepadButtonDown(gamePadId, startCode) {
			shouldExit = true //see WindowShouldClose, it checks if KeyEscape pressed or Close icon pressed
		}

		rl.DrawFPS(5, 5)
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

func If[T any](cond bool, vTrue, vFalse T) T {
	if cond {
		return vTrue
	}
	return vFalse
}
