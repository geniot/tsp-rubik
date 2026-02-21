package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Application struct {
	camera *rl.Camera
	scenes map[int]Scene
}

func (a *Application) ShouldExit() bool {
	return rl.WindowShouldClose() || a.scenes[currentSceneIndex].ShouldExit()
}

func (a *Application) Draw() {
	a.scenes[currentSceneIndex].Draw(a.camera)
}

func (a *Application) Exit() {
	rl.CloseWindow()
}

func NewApplication() *Application {
	app := Application{}
	rl.SetConfigFlags(rl.FlagVsyncHint) //should be set before window initialization!
	rl.InitWindow(winWidth, winHeight, "TrimUI Rubik")
	rl.SetWindowMonitor(1)
	rl.InitAudioDevice()
	prepareTextures()

	app.camera = &rl.Camera3D{}
	app.scenes = make(map[int]Scene)
	app.scenes[menuSceneKey] = NewMenuScene()
	app.scenes[gameSceneKey] = NewGameScene()
	app.scenes[tutorialSceneKey] = NewTutorialScene()

	rl.SetClipPlanes(0.5, 100) //see https://github.com/raysan5/raylib/issues/4917
	rl.DisableBackfaceCulling()

	app.camera.Position = rl.NewVector3(10, 10, 10)
	app.camera.Target = rl.NewVector3(0.0, 0.0, 0.0)
	app.camera.Up = rl.NewVector3(0.0, 1.0, 0.0)
	app.camera.Fovy = 40.0
	app.camera.Projection = rl.CameraPerspective
	return &app
}
