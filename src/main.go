package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	winHeight      = 720
	winWidth       = 1280
	gamePadId      = int32(0)
	helpFontSize   = int32(20)
	helpWidth      = int32(360)
	helpHeight     = int32(120)
	helpPadding    = int32(10)
	helpLineHeight = int32(20)
)

func main() {

	//rl.SetConfigFlags(rl.FlagMsaa4xHint)
	rl.SetConfigFlags(rl.FlagVsyncHint) //should be set before window initialization!
	//rl.SetTargetFPS(60)

	rl.InitWindow(winWidth, winHeight, "TrimUI Rubik")
	rl.SetWindowMonitor(0)
	rl.InitAudioDevice()

	prepareTextures()

	var (
		cubeSize = 3 //currently only 3 is supported :)
		camera   = rl.Camera3D{}
		cube     = NewCube(cubeSize)
	)

	rl.SetClipPlanes(0.5, 100) //see https://github.com/raysan5/raylib/issues/4917
	rl.DisableBackfaceCulling()

	camera.Position = rl.NewVector3(10, 10, 10)
	camera.Target = rl.NewVector3(0.0, 0.0, 0.0)
	camera.Up = rl.NewVector3(0.0, 1.0, 0.0)
	camera.Fovy = 40.0
	camera.Projection = rl.CameraPerspective

	for !rl.WindowShouldClose() && !shouldExit() {
		//rl.UpdateCamera(&camera, rl.CameraOrbital)

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		rl.Color4f(1, 1, 1, 1)

		rl.BeginMode3D(camera)
		cube.update()
		cube.draw()
		//rl.DrawGrid(10, 1)
		rl.EndMode3D()

		rl.DrawText("The Breathing Cube", helpPadding*2, helpPadding*2, helpFontSize*2, rl.Blue)
		rl.DrawText("It's breathing, so it's correct.", helpPadding*2+helpPadding/2, helpPadding*8, helpFontSize, rl.DarkGreen)

		drawHelp()

		//rl.DrawFPS(5, 5)
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
