package main

import (
	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	CubeCorrect = [6][9]int{
		{BLUE, BLUE, BLUE, BLUE, BLUE, BLUE, BLUE, BLUE, BLUE},
		{ORANGE, ORANGE, ORANGE, ORANGE, ORANGE, ORANGE, ORANGE, ORANGE, ORANGE},
		{GREEN, GREEN, GREEN, GREEN, GREEN, GREEN, GREEN, GREEN, GREEN},
		{RED, RED, RED, RED, RED, RED, RED, RED, RED},
		{YELLOW, YELLOW, YELLOW, YELLOW, YELLOW, YELLOW, YELLOW, YELLOW, YELLOW},
		{WHITE, WHITE, WHITE, WHITE, WHITE, WHITE, WHITE, WHITE, WHITE},
	}
)

type GameScene struct {
	a                *Application
	cube             *Cube
	isStarted        bool
	buttonRectangles []rl.Rectangle
}

func NewGameScene(a *Application) *GameScene {
	gameScene := GameScene{}
	gameScene.a = a
	gameScene.cube = NewCube(3, split(CubeCorrect), a)
	gameScene.isStarted = false
	gameScene.Reset()
	gameScene.buttonRectangles = []rl.Rectangle{
		rl.NewRectangle(ButtonHeight/2, ButtonHeight/2, ButtonHeight, ButtonHeight),
	}
	return &gameScene
}

func (gs *GameScene) ShouldExit() bool {
	return false
}

func (gs *GameScene) HitsAnyButton(position rl.Vector2) bool {
	for _, button := range gs.buttonRectangles {
		if rl.CheckCollisionPointRec(position, button) {
			return true
		}
	}
	return false
}

func (gs *GameScene) Render(rt rl.RenderTexture2D, camera *rl.Camera) {
	gs.isStarted = true
	//rl.UpdateCamera(camera, rl.CameraOrbital)
	rl.BeginTextureMode(rt)
	rl.ClearBackground(rl.RayWhite)
	rl.Color4f(1, 1, 1, 1)
	rl.BeginMode3D(*camera)
	gs.cube.update()
	gs.cube.draw()
	//rl.DrawGrid(10, 1)
	rl.EndMode3D()

	//buttonCount := float32(0)
	isButtonClicked := false
	setTextStyle(singleLetterStyle)

	gui.SetState(If(rl.IsGamepadButtonDown(gamePadId, menuCode), gui.STATE_PRESSED, gui.STATE_NORMAL))
	isButtonClicked = gui.Button(gs.buttonRectangles[0], "M")
	if isButtonClicked || rl.IsGamepadButtonReleased(gamePadId, menuCode) {
		gs.a.currentSceneIndex = menuSceneKey
	}
	updateTspCamera(camera)
	rl.EndTextureMode()
}

func (gs *GameScene) Update() {

}

func (gs *GameScene) Reset() {
	gs.cube.isFaceSelectionModeOn = false
	//can be commented for testing
	gs.cube.Shuffle(shuffleCount)
}
