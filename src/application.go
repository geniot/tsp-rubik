package main

import (
	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Application struct {
	camera                *rl.Camera
	scenes                map[int]Scene
	currentSceneIndex     int
	colorTextures         map[int]rl.Texture2D
	selectedColorTextures map[int]rl.Texture2D
}

func (a *Application) ShouldExit() bool {
	return rl.WindowShouldClose() || a.scenes[a.currentSceneIndex].ShouldExit()
}

func (a *Application) Update() {
	a.scenes[a.currentSceneIndex].Update(a.camera)
}

func (a *Application) Exit() {
	rl.CloseWindow()
}

func NewApplication() *Application {

	app := Application{}

	// the order of these calls matters
	rl.SetTraceLogLevel(rl.LogWarning)
	rl.SetConfigFlags(rl.FlagVsyncHint) //should be set before window initialization!
	rl.InitWindow(winWidth, winHeight, "TrimUI Rubik")
	rl.SetWindowMonitor(0)
	rl.InitAudioDevice()
	rl.SetClipPlanes(0.5, 100) //see https://github.com/raysan5/raylib/issues/4917
	rl.DisableBackfaceCulling()

	gui.SetStyle(gui.DEFAULT, gui.TEXT_SIZE, 40)
	gui.SetStyle(gui.DEFAULT, gui.TEXT_SPACING, 10)
	gui.SetStyle(gui.DEFAULT, gui.TEXT_ALIGNMENT, int64(gui.TEXT_ALIGN_LEFT))
	gui.SetStyle(gui.DEFAULT, gui.TEXT_PADDING, 20)

	//camera
	app.camera = &rl.Camera3D{}
	app.scenes = make(map[int]Scene)
	app.camera.Position = rl.NewVector3(10, 10, 10)
	app.camera.Target = rl.NewVector3(0.0, 0.0, 0.0)
	app.camera.Up = rl.NewVector3(0.0, 1.0, 0.0)
	app.camera.Fovy = 40.0
	app.camera.Projection = rl.CameraPerspective

	// textures
	app.colorTextures = make(map[int]rl.Texture2D)
	app.selectedColorTextures = make(map[int]rl.Texture2D)
	prepareTextures(app.colorTextures, false)
	prepareTextures(app.selectedColorTextures, true)

	// scenes
	app.scenes[menuSceneKey] = NewMenuScene(&app)
	app.scenes[gameSceneKey] = NewGameScene(&app)
	app.scenes[tutorialSceneKey] = NewTutorialScene(&app)
	app.currentSceneIndex = menuSceneKey

	//debug
	app.scenes[gameSceneKey].(*GameScene).Reset()
	app.currentSceneIndex = gameSceneKey

	return &app
}
