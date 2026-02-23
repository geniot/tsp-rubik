package main

import (
	"strconv"

	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	tutorials = [2][6][9]int{
		{
			{GREEN, GREEN, LIGHT_BLACK, GREEN, GREEN, GREEN, GREEN, LIGHT_BLACK, LIGHT_BLACK},
			{ORANGE, ORANGE, ORANGE, ORANGE, ORANGE, ORANGE, LIGHT_BLACK, LIGHT_BLACK, LIGHT_BLACK},
			{BLUE, BLUE, LIGHT_BLACK, BLUE, BLUE, LIGHT_BLACK, BLUE, BLUE, LIGHT_BLACK},
			{RED, RED, RED, RED, RED, LIGHT_BLACK, LIGHT_BLACK, LIGHT_BLACK, LIGHT_BLACK},
			{LIGHT_BLACK, LIGHT_BLACK, LIGHT_BLACK, LIGHT_BLACK, YELLOW, RED, LIGHT_BLACK, LIGHT_BLACK, LIGHT_BLACK},
			{WHITE, WHITE, WHITE, WHITE, WHITE, WHITE, WHITE, WHITE, WHITE},
		},
		{
			{GREEN, GREEN, LIGHT_BLACK, GREEN, GREEN, LIGHT_BLACK, GREEN, LIGHT_BLACK, LIGHT_BLACK},
			{ORANGE, ORANGE, ORANGE, ORANGE, ORANGE, ORANGE, LIGHT_BLACK, LIGHT_BLACK, LIGHT_BLACK},
			{BLUE, BLUE, LIGHT_BLACK, BLUE, BLUE, LIGHT_BLACK, BLUE, BLUE, LIGHT_BLACK},
			{RED, RED, RED, RED, RED, LIGHT_BLACK, LIGHT_BLACK, RED, LIGHT_BLACK},
			{LIGHT_BLACK, LIGHT_BLACK, LIGHT_BLACK, LIGHT_BLACK, YELLOW, LIGHT_BLACK, LIGHT_BLACK, GREEN, LIGHT_BLACK},
			{WHITE, WHITE, WHITE, WHITE, WHITE, WHITE, WHITE, WHITE, WHITE},
		},
	}
	solutions = [2]string{
		"U R U'R'\nU'F U F'", "U R U'R'\nU'F U F'",
	}
)

type TutorialScene struct {
	a          *Application
	cubes      []*Cube
	docPointer int
}

func NewTutorialScene(a *Application) *TutorialScene {
	tutorialScene := TutorialScene{}
	tutorialScene.a = a
	tutorialScene.docPointer = 0
	tutorialScene.cubes = make([]*Cube, len(solutions))
	for i, _ := range solutions {
		tutorialScene.cubes[i] = NewCube(3, split(tutorials[i]), a)
	}
	return &tutorialScene
}

func (ts *TutorialScene) ShouldExit() bool {
	return rl.IsKeyPressed(rl.KeyEscape) || (rl.IsGamepadButtonDown(gamePadId, menuCode) && rl.IsGamepadButtonDown(gamePadId, startCode))
}

func (ts *TutorialScene) rotate(inc int) {
	ts.docPointer += inc
	if ts.docPointer < 0 {
		ts.docPointer = len(tutorials) - 1
	}
	if ts.docPointer >= len(tutorials) {
		ts.docPointer = 0
	}
}

func (ts *TutorialScene) Update(camera *rl.Camera) {
	//rl.UpdateCamera(&camera, rl.CameraOrbital)
	rl.BeginDrawing()
	rl.ClearBackground(rl.RayWhite)
	rl.Color4f(1, 1, 1, 1)

	rl.BeginMode3D(*camera)
	ts.cubes[ts.docPointer].update()
	ts.cubes[ts.docPointer].draw()
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

	setTextStyle(20, 0, int64(gui.TEXT_ALIGN_CENTER), 0)
	gui.SetState(gui.STATE_NORMAL)
	isButtonClicked = gui.Button(rl.NewRectangle(winWidth-buttonHeight/2*4.7, winHeight-buttonHeight/2*1.5, buttonHeight/2, buttonHeight/2), "<")
	if isButtonClicked || rl.IsGamepadButtonPressed(gamePadId, l1Code) || rl.IsGamepadButtonPressed(gamePadId, l2Code) {
		ts.rotate(-1)
	}
	gui.SetState(gui.STATE_NORMAL)
	isButtonClicked = gui.Button(rl.NewRectangle(winWidth-buttonHeight/2*1.5, winHeight-buttonHeight/2*1.5, buttonHeight/2, buttonHeight/2), ">")
	if isButtonClicked || rl.IsGamepadButtonPressed(gamePadId, l1Code) || rl.IsGamepadButtonPressed(gamePadId, l2Code) {
		ts.rotate(1)
	}
	setDefaultTextStyle()

	rl.DrawText(solutions[ts.docPointer], 15, winHeight-70, subTitleTextFontSize, rl.Blue)
	rl.DrawText(strconv.Itoa(ts.docPointer+1)+"/"+strconv.Itoa(len(solutions)), winWidth-120, winHeight-48, subTitleTextFontSize, rl.Black)

	//rl.DrawFPS(5, 5)
	rl.EndDrawing()
}
