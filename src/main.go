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
	menuSceneKey = iota
	gameSceneKey
	tutorialSceneKey
)

const (
	winHeight = 720
	winWidth  = 1280
	gamePadId = int32(0)
)

var (
	currentSceneIndex = menuSceneKey
	scenes            = make(map[int]Scene)
)

type Scene interface {
	Draw(camera *rl.Camera)
	ShouldExit() bool
}

func main() {

	//rl.SetConfigFlags(rl.FlagMsaa4xHint)
	rl.SetConfigFlags(rl.FlagVsyncHint) //should be set before window initialization!
	//rl.SetTargetFPS(60)

	rl.InitWindow(winWidth, winHeight, "TrimUI Rubik")
	rl.SetWindowMonitor(0)
	rl.InitAudioDevice()

	prepareTextures()

	var (
		camera = &rl.Camera3D{}
	)
	scenes[menuSceneKey] = NewMenuScene()
	scenes[gameSceneKey] = NewGameScene()

	rl.SetClipPlanes(0.5, 100) //see https://github.com/raysan5/raylib/issues/4917
	rl.DisableBackfaceCulling()

	camera.Position = rl.NewVector3(10, 10, 10)
	camera.Target = rl.NewVector3(0.0, 0.0, 0.0)
	camera.Up = rl.NewVector3(0.0, 1.0, 0.0)
	camera.Fovy = 40.0
	camera.Projection = rl.CameraPerspective

	for !rl.WindowShouldClose() && !scenes[currentSceneIndex].ShouldExit() {
		scenes[currentSceneIndex].Draw(camera)
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

func orPanicRes[T any](res T, err interface{}) T {
	orPanic(err)
	return res
}
