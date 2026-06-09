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
	return rl.IsKeyPressed(rl.KeyEscape)
}

func (cs *ControlsScene) Update(_ *rl.Camera) {
	rl.BeginDrawing()
	rl.ClearBackground(rl.RayWhite)
	rl.Color4f(1, 1, 1, 1)
	isButtonClicked := false
	buttonHeight := float32(70)

	gui.SetState(If(rl.IsGamepadButtonDown(gamePadId, menuCode), gui.STATE_PRESSED, gui.STATE_NORMAL))
	isButtonClicked = gui.Button(rl.NewRectangle(buttonHeight/2, buttonHeight/2, buttonHeight, buttonHeight), "M")
	if isButtonClicked || rl.IsGamepadButtonReleased(gamePadId, menuCode) {
		cs.a.currentSceneIndex = menuSceneKey
	}
	gui.SetState(gui.STATE_NORMAL)
	gui.SetStyle(gui.DEFAULT, gui.TEXT_ALIGNMENT_VERTICAL, int64(gui.TEXT_ALIGN_TOP))
	gui.SetStyle(gui.DEFAULT, gui.TEXT_LINE_SPACING, 50)
	padding := float32(125)
	textBoxTextLines := []string{"GAME",
		"[A,B,X,Y,L1,L2,R1,R2] Selection",
		"[Arrow keys, Left stick] Rotation",
		"TUTORIAL",
		"[MENU] Back to Menu",
		"[SELECT] Reset", "[START] Play one move", "" +
			"[DOUBLE SELECT] Next tutorial"}
	lineHeight := float32(70)
	for i := 0; i < len(textBoxTextLines); i++ {
		gui.TextBox(rl.Rectangle{padding, padding + lineHeight*float32(i), winWidth - padding*2,
			lineHeight}, &textBoxTextLines[i], 64, false)
	}
	gui.SetStyle(gui.DEFAULT, gui.TEXT_ALIGNMENT_VERTICAL, int64(gui.TEXT_ALIGN_CENTER))
	gui.SetStyle(gui.DEFAULT, gui.TEXT_LINE_SPACING, 10)
	rl.EndDrawing()
}
