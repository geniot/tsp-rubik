package main

import (
	"embed"
	rl "github.com/gen2brain/raylib-go/raylib"
	"strconv"
)

var (
	//go:embed media/*
	mediaList embed.FS
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
	winHeight      = 720
	winWidth       = 1280
	gamePadId      = int32(0)
	helpFontSize   = int32(20)
	helpWidth      = int32(360)
	helpHeight     = int32(120)
	helpPadding    = int32(10)
	helpLineHeight = int32(20)
)

var (
	docTextures = [16]rl.Texture2D{}
	docPointer  = 0
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
		cubeSize  = 3 //currently only 3 is supported :)
		camera    = rl.Camera3D{}
		cube      = NewCube(cubeSize)
		isDocMode = false
	)

	for i := 0; i < len(docTextures); i++ {
		textureBytes := orPanicRes(mediaList.ReadFile("media/doc" + strconv.Itoa(i) + ".png"))
		docTextures[i] = rl.LoadTextureFromImage(rl.LoadImageFromMemory(".png", textureBytes, int32(len(textureBytes))))
	}

	rl.SetClipPlanes(0.5, 100) //see https://github.com/raysan5/raylib/issues/4917
	rl.DisableBackfaceCulling()

	camera.Position = rl.NewVector3(10, 10, 10)
	camera.Target = rl.NewVector3(0.0, 0.0, 0.0)
	camera.Up = rl.NewVector3(0.0, 1.0, 0.0)
	camera.Fovy = 40.0
	camera.Projection = rl.CameraPerspective

	for !rl.WindowShouldClose() && !shouldExit() {
		if rl.IsKeyPressed(rl.KeyH) || rl.IsGamepadButtonPressed(gamePadId, selectCode) {
			isDocMode = !isDocMode
		}
		//rl.UpdateCamera(&camera, rl.CameraOrbital)

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		rl.Color4f(1, 1, 1, 1)

		if isDocMode {
			rl.DrawTexture(docTextures[docPointer], 0, 0, rl.White)
			rl.DrawText(strconv.Itoa(docPointer+1)+"/"+strconv.Itoa(len(docTextures)), helpPadding*2, helpPadding*2, helpFontSize*2, rl.Blue)
			handleDocUserEvents()
		} else {
			rl.BeginMode3D(camera)
			cube.update()
			cube.draw()
			//rl.DrawGrid(10, 1)
			rl.EndMode3D()

			//rl.DrawText("The Breathing Cube", helpPadding*2, helpPadding*2, helpFontSize*2, rl.Blue)
			//rl.DrawText("It's breathing, when it's correct.", helpPadding*2+helpPadding/2, helpPadding*8, helpFontSize, rl.DarkGreen)

			//rl.DrawFPS(5, 5)
		}
		rl.EndDrawing()
	}
	rl.CloseWindow()
}

func handleDocUserEvents() {
	if rl.IsKeyPressed(rl.KeyRight) || rl.IsGamepadButtonPressed(gamePadId, rightCode) {
		docPointer += 1
		if docPointer >= len(docTextures) {
			docPointer = 0
		}
	}
	if rl.IsKeyPressed(rl.KeyLeft) || rl.IsGamepadButtonPressed(gamePadId, leftCode) {
		docPointer -= 1
		if docPointer < 0 {
			docPointer = len(docTextures) - 1
		}
	}
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

func orPanicRes[T any](res T, err interface{}) T {
	orPanic(err)
	return res
}
