package main

import (
	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type TutorialScene struct {
	a           *Application
	cube        *Cube
	docTextures [16]rl.Texture2D
	docPointer  int
}

func (ts *TutorialScene) ShouldExit() bool {
	return rl.IsKeyPressed(rl.KeyEscape) || (rl.IsGamepadButtonDown(gamePadId, menuCode) && rl.IsGamepadButtonDown(gamePadId, startCode))
}

func (ts *TutorialScene) Update(camera *rl.Camera) {
	//rl.UpdateCamera(&camera, rl.CameraOrbital)
	rl.BeginDrawing()
	rl.ClearBackground(rl.RayWhite)
	rl.Color4f(1, 1, 1, 1)

	rl.BeginMode3D(*camera)
	ts.cube.update()
	ts.cube.draw()
	//rl.DrawGrid(10, 1)
	rl.EndMode3D()

	//buttonCount := float32(0)
	isButtonClicked := false
	buttonHeight := float32(70)

	gui.SetState(gui.STATE_NORMAL)
	isButtonClicked = gui.Button(rl.NewRectangle(buttonHeight/2, buttonHeight/2, buttonHeight, buttonHeight), "M")
	if isButtonClicked || rl.IsGamepadButtonPressed(gamePadId, menuCode) {
		ts.a.currentSceneIndex = menuSceneKey
	}

	rl.DrawText("U R U'R'\nU'F U F'", 15, winHeight-70, subTitleTextFontSize, rl.Blue)

	//rl.DrawFPS(5, 5)
	rl.EndDrawing()
}

func NewTutorialScene(a *Application) *TutorialScene {
	return &TutorialScene{a: a, cube: NewCube(3, split(CubeTutorial1), a)}
}
