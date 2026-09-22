package main

import (
	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type ControlsScene struct {
	a *Application
}

func NewControlsScene(a *Application) *ControlsScene {
	tutorialScene := ControlsScene{}
	tutorialScene.a = a
	return &tutorialScene
}

func (cs *ControlsScene) ShouldExit() bool {
	return false
}

func (cs *ControlsScene) HitsAnyButton(_ rl.Vector2) bool {
	return false
}

func (cs *ControlsScene) Update(_ *rl.Camera) {
	rl.BeginDrawing()
	rl.ClearBackground(rl.RayWhite)
	rl.Color4f(1, 1, 1, 1)
	setTextStyle(singleLetterStyle)
	gui.SetState(If(rl.IsGamepadButtonDown(gamePadId, menuCode), gui.STATE_PRESSED, gui.STATE_NORMAL))
	isButtonClicked := gui.Button(rl.NewRectangle(ButtonHeight/2, ButtonHeight/2, ButtonHeight, ButtonHeight), "M")
	if isButtonClicked || rl.IsGamepadButtonReleased(gamePadId, menuCode) {
		cs.a.currentSceneIndex = menuSceneKey
	}

	setTextStyle(defaultTextStyle)
	padding := float32(125)
	lineHeight := float32(70)
	textBoxTextLines := []string{"GAME",
		If(rl.IsGamepadAvailable(gamePadId), "[A,B,X,Y,L1,L2,R1,R2] Selection", "[1,2,3,4,5,6,7,8,9] Selection"),
		"[Arrow keys, Left stick] Rotation",
		"TUTORIAL",
		"[MENU] Back to Menu",
		"[SELECT] Reset", "[START] Play one move", "" +
			"[DOUBLE SELECT] Next tutorial"}
	for i := range textBoxTextLines {
		gui.Label(rl.Rectangle{
			X:      padding,
			Y:      padding + lineHeight*float32(i),
			Width:  winWidth - padding*2,
			Height: lineHeight},
			textBoxTextLines[i])
	}
	rl.EndDrawing()
}
