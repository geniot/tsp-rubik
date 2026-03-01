package main

import (
	"strconv"

	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Hint struct {
	rotation  int
	isForward bool
}

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

	solutions = [2]string{}

	hints = [2][]Hint{
		{
			{RTop, true},
			{RRight, true},
			{RTop, false},
			{RRight, false},
			{RTop, false},
			{RFront, false},
			{RTop, true},
			{RFront, true},
		},
		{
			{RTop, false},
			{RFront, false},
			{RTop, true},
			{RFront, true},
			{RTop, false},
			{RFront, false},
			{RTop, true},
			{RFront, true},
		},
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
	tutorialScene.cubes = make([]*Cube, len(hints))
	for i, _ := range hints {
		tutorialScene.cubes[i] = NewCube(3, split(tutorials[i]), a)
	}
	for i, _ := range hints {
		solutions[i] = genSolution(hints[i])
	}
	return &tutorialScene
}

func genSolution(hs []Hint) string {
	var s string
	for _, h := range hs {
		s += rotationLetters[h.rotation] + If(h.isForward, " ", "'")
	}
	return s
}

func (ts *TutorialScene) ShouldExit() bool {
	return rl.IsKeyPressed(rl.KeyEscape) || (rl.IsGamepadButtonDown(gamePadId, menuCode) && rl.IsGamepadButtonDown(gamePadId, startCode))
}

func (ts *TutorialScene) NextPrev(inc int) {
	ts.docPointer += inc
	if ts.docPointer < 0 {
		ts.docPointer = len(tutorials) - 1
	}
	if ts.docPointer >= len(tutorials) {
		ts.docPointer = 0
	}
}

func (ts *TutorialScene) NextHint() {
	ts.cubes[ts.docPointer].hintPointer += 1
	if ts.cubes[ts.docPointer].hintPointer >= len(hints[ts.docPointer]) {
		ts.cubes[ts.docPointer].hintPointer = 0
		ts.Reset()
	} else {
		cube := ts.cubes[ts.docPointer]
		ts.cubes[ts.docPointer].RotateAny(hints[ts.docPointer][cube.hintPointer].rotation, hints[ts.docPointer][cube.hintPointer].isForward, true)
	}
}

func (ts *TutorialScene) Reset() {
	ts.cubes[ts.docPointer] = NewCube(3, split(tutorials[ts.docPointer]), ts.a)
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

	//menu
	gui.SetState(gui.STATE_NORMAL)
	isButtonClicked = gui.Button(rl.NewRectangle(buttonHeight/2, buttonHeight/2, buttonHeight, buttonHeight), "M")
	if isButtonClicked || rl.IsGamepadButtonPressed(gamePadId, menuCode) {
		ts.a.currentSceneIndex = menuSceneKey
	}
	//reset
	gui.SetState(gui.STATE_NORMAL)
	isButtonClicked = gui.Button(rl.NewRectangle(buttonHeight/2, buttonHeight/2*4, buttonHeight, buttonHeight), "R")
	if isButtonClicked || rl.IsGamepadButtonPressed(gamePadId, selectCode) {
		ts.Reset()
	}
	//play
	gui.SetState(gui.STATE_NORMAL)
	isButtonClicked = gui.Button(rl.NewRectangle(buttonHeight/2, buttonHeight/2*7, buttonHeight, buttonHeight), "P")
	if isButtonClicked || rl.IsGamepadButtonPressed(gamePadId, startCode) {
		ts.NextHint()
	}

	setTextStyle(20, 0, int64(gui.TEXT_ALIGN_CENTER), 0)
	gui.SetState(gui.STATE_NORMAL)
	isButtonClicked = gui.Button(rl.NewRectangle(winWidth-buttonHeight/2*4.7, winHeight-buttonHeight/2*1.5, buttonHeight/2, buttonHeight/2), "<")
	if isButtonClicked || rl.IsGamepadButtonPressed(gamePadId, l1Code) || rl.IsGamepadButtonPressed(gamePadId, l2Code) {
		ts.NextPrev(-1)
	}
	gui.SetState(gui.STATE_NORMAL)
	isButtonClicked = gui.Button(rl.NewRectangle(winWidth-buttonHeight/2*1.5, winHeight-buttonHeight/2*1.5, buttonHeight/2, buttonHeight/2), ">")
	if isButtonClicked || rl.IsGamepadButtonPressed(gamePadId, l1Code) || rl.IsGamepadButtonPressed(gamePadId, l2Code) {
		ts.NextPrev(1)
	}
	setDefaultTextStyle()

	rl.DrawText(solutions[ts.docPointer], 15, winHeight-70, subTitleTextFontSize, rl.Blue)
	rl.DrawText(strconv.Itoa(ts.docPointer+1)+"/"+strconv.Itoa(len(hints)), winWidth-120, winHeight-48, subTitleTextFontSize, rl.Black)

	//rl.DrawFPS(5, 5)
	rl.EndDrawing()
}
