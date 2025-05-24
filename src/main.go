package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
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

const (
	winHeight = 720
	winWidth  = 1280
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

		fontSize := int32(20)
		width := int32(360)
		height := int32(120)
		padding := int32(10)
		lineHeight := int32(20)

		rl.DrawText("The Breathing Cube", padding*2, padding*2, fontSize*2, rl.Blue)
		rl.DrawText("It's breathing when it's correct.", padding*2+padding/2, padding*8, fontSize, rl.DarkGreen)

		rl.DrawRectangle(padding, winHeight-height-padding, width, height, rl.Fade(rl.SkyBlue, 0.5))
		rl.DrawRectangleLines(padding, winHeight-height-padding, width, height, rl.Blue)

		rl.DrawText("Desktop controls:", padding*2, winHeight-height-padding, fontSize, rl.Black)
		rl.DrawText("use arrow keys to rotate", padding*2, winHeight-height-padding+lineHeight*1, fontSize, rl.DarkGray)
		rl.DrawText("1-9 to (de)select faces", padding*2, winHeight-height-padding+lineHeight*2, fontSize, rl.DarkGray)
		rl.DrawText("'S' (hold) to shuffle", padding*2, winHeight-height-padding+lineHeight*3, fontSize, rl.DarkGray)
		rl.DrawText("'Left Control' (hold) + Up/Down", padding*2, winHeight-height-padding+lineHeight*4, fontSize, rl.DarkGray)
		rl.DrawText(" - rotate around the Z-axis", padding*6, winHeight-height-padding+lineHeight*5, fontSize, rl.DarkGray)

		rl.DrawRectangle(winWidth-width-padding, winHeight-height-padding, width, height, rl.Fade(rl.SkyBlue, 0.5))
		rl.DrawRectangleLines(winWidth-width-padding, winHeight-height-padding, width, height, rl.Blue)

		rl.DrawText("TrimUI Smart Pro controls:", winWidth-width-padding/2, winHeight-height-padding, fontSize, rl.Black)
		rl.DrawText("use arrow joystick to select", winWidth-width, winHeight-height-padding+lineHeight*1, fontSize, rl.DarkGray)
		rl.DrawText("use analogue joystick to rotate", winWidth-width, winHeight-height-padding+lineHeight*2, fontSize, rl.DarkGray)
		rl.DrawText("A/B to rotate faces", winWidth-width, winHeight-height-padding+lineHeight*3, fontSize, rl.DarkGray)
		rl.DrawText("MENU+START to exit", winWidth-width, winHeight-height-padding+lineHeight*5, fontSize, rl.DarkGray)

		//exit
		if rl.IsGamepadButtonDown(gamePadId, menuCode) && rl.IsGamepadButtonDown(gamePadId, startCode) {
			shouldExit = true //see WindowShouldClose, it checks if KeyEscape pressed or Close icon pressed
		}

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
