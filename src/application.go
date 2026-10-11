package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	InitialCameraPosition = rl.NewVector3(10, 10, 10)
	ZeroVector2           = rl.Vector2{}
)

const (
	ButtonHeight = float32(78)
)

type Application struct {
	camera                *rl.Camera
	scenes                map[int]Scene
	currentSceneIndex     int
	colorTextures         map[int]rl.Texture2D
	selectedColorTextures map[int]rl.Texture2D
	isWeb                 bool
	drawTarget            rl.RenderTexture2D
	sourceRect            rl.Rectangle
	destRect              rl.Rectangle
	resizeRatio           float32
}

func (a *Application) onResize() {
	screenWidth := float32(rl.GetScreenWidth())
	screenHeight := float32(rl.GetScreenHeight())
	ratioX := screenWidth / float32(ScreenLogicalWidth)
	ratioY := screenHeight / float32(ScreenLogicalHeight)
	a.resizeRatio = If(ratioX < ratioY, ratioX, ratioY)

	a.sourceRect = rl.NewRectangle(0, 0, float32(ScreenLogicalWidth), float32(-ScreenLogicalHeight))
	a.destRect = rl.NewRectangle(
		(screenWidth-(float32(ScreenLogicalWidth)*a.resizeRatio))*0.5,
		(screenHeight-(float32(ScreenLogicalHeight)*a.resizeRatio))*0.5,
		float32(ScreenLogicalWidth)*a.resizeRatio,
		float32(ScreenLogicalHeight)*a.resizeRatio,
	)
	rl.SetMouseOffset(-int(a.destRect.X), -int(a.destRect.Y))
	rl.SetMouseScale(1/a.resizeRatio, 1/a.resizeRatio)
}

func (a *Application) getMousePositionOriginal() rl.Vector2 {
	rl.SetMouseOffset(0, 0)
	rl.SetMouseScale(1, 1)
	mousePosition := rl.GetMousePosition()
	rl.SetMouseOffset(-int(a.destRect.X), -int(a.destRect.Y))
	rl.SetMouseScale(1/a.resizeRatio, 1/a.resizeRatio)
	return mousePosition
}

func (a *Application) ShouldExit() bool {
	return rl.WindowShouldClose() || a.scenes[a.currentSceneIndex].ShouldExit()
}

func (a *Application) Update() {
	a.scenes[a.currentSceneIndex].Update()
}

func (a *Application) ProcessInput() {
	if rl.IsWindowResized() {
		a.onResize()
	}
	//a.scenes[a.currentSceneIndex].ProcessInput()
}

func (a *Application) Render() {
	a.scenes[a.currentSceneIndex].Render(a.drawTarget, a.camera)

	rl.BeginDrawing()
	rl.ClearBackground(rl.RayWhite)
	rl.DrawTexturePro(a.drawTarget.Texture,
		a.sourceRect,
		a.destRect,
		ZeroVector2, 0, rl.White)
	rl.EndDrawing()
}

func (a *Application) Exit() {
	rl.CloseWindow()
}

func NewApplication(iw bool) *Application {

	app := Application{}
	app.isWeb = iw

	// the order of these calls matters
	rl.SetTraceLogLevel(rl.LogWarning)
	rl.SetConfigFlags(rl.FlagVsyncHint | rl.FlagWindowResizable) //should be set before window initialization!
	rl.InitWindow(ScreenLogicalWidth, ScreenLogicalHeight, "TrimUI Rubik")
	// used for testing on multiple monitors (desktop), should be set back to 0 for TrimUI and web,
	// otherwise SDL can fail with WARNING: SDL: Failed to find selected monitor
	rl.SetWindowMonitor(0)
	rl.InitAudioDevice()
	rl.SetClipPlanes(0.5, 100) //these values are found by trial and error for TrimUI, see https://github.com/raysan5/raylib/issues/4917
	rl.DisableBackfaceCulling()
	rl.SetExitKey(0) //should be set after window initialization!

	setTextStyle(defaultTextStyle)

	//camera
	app.camera = &rl.Camera3D{}
	app.camera.Position = InitialCameraPosition
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
	app.scenes = make(map[int]Scene)
	app.scenes[menuSceneKey] = NewMenuScene(&app)
	app.scenes[gameSceneKey] = NewGameScene(&app)
	app.scenes[tutorialSceneKey] = NewTutorialScene(&app)
	app.scenes[controlsSceneKey] = NewControlsScene(&app)
	app.currentSceneIndex = menuSceneKey

	app.drawTarget = rl.LoadRenderTexture(ScreenLogicalWidth, ScreenLogicalHeight)
	app.onResize()

	//debug
	//app.currentSceneIndex = gameSceneKey

	return &app
}
