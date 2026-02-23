package main

import (
	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	CubeCorrect = [6][9]int{
		{GREEN, GREEN, GREEN, GREEN, GREEN, GREEN, GREEN, GREEN, GREEN},
		{ORANGE, ORANGE, ORANGE, ORANGE, ORANGE, ORANGE, ORANGE, ORANGE, ORANGE},
		{BLUE, BLUE, BLUE, BLUE, BLUE, BLUE, BLUE, BLUE, BLUE},
		{RED, RED, RED, RED, RED, RED, RED, RED, RED},
		{YELLOW, YELLOW, YELLOW, YELLOW, YELLOW, YELLOW, YELLOW, YELLOW, YELLOW},
		{WHITE, WHITE, WHITE, WHITE, WHITE, WHITE, WHITE, WHITE, WHITE},
	}
)

type GameScene struct {
	a    *Application
	cube *Cube
}

func NewGameScene(a *Application) *GameScene {
	gameScene := GameScene{}
	gameScene.a = a
	gameScene.cube = NewCube(3, split(CubeCorrect), a)
	gameScene.Reset()
	return &gameScene
}

func (gs *GameScene) ShouldExit() bool {
	return rl.IsKeyPressed(rl.KeyEscape) || (rl.IsGamepadButtonDown(gamePadId, menuCode) && rl.IsGamepadButtonDown(gamePadId, startCode))
}

func (gs *GameScene) Update(camera *rl.Camera) {
	//rl.UpdateCamera(&camera, rl.CameraOrbital)
	rl.BeginDrawing()
	rl.ClearBackground(rl.RayWhite)
	rl.Color4f(1, 1, 1, 1)
	rl.BeginMode3D(*camera)
	gs.cube.update()
	gs.cube.draw()
	//rl.DrawGrid(10, 1)
	rl.EndMode3D()

	//buttonCount := float32(0)
	isButtonClicked := false
	buttonHeight := float32(70)

	gui.SetState(gui.STATE_NORMAL)
	isButtonClicked = gui.Button(rl.NewRectangle(buttonHeight/2, buttonHeight/2, buttonHeight, buttonHeight), "M")
	if isButtonClicked || rl.IsGamepadButtonPressed(gamePadId, menuCode) {
		gs.a.currentSceneIndex = menuSceneKey
	}

	rl.EndDrawing()
}

func (gs *GameScene) Reset() {
	gs.cube.isFaceSelectionModeOn = false
	gs.cube.Shuffle(shuffleCount)
}
