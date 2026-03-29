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
	tutorials = [8][6][9]int{
		{ //1
			{GREEN, GREEN, LIGHT_BLACK, GREEN, GREEN, GREEN, GREEN, LIGHT_BLACK, LIGHT_BLACK},                        //front
			{ORANGE, ORANGE, ORANGE, ORANGE, ORANGE, ORANGE, LIGHT_BLACK, LIGHT_BLACK, LIGHT_BLACK},                  //left
			{BLUE, BLUE, LIGHT_BLACK, BLUE, BLUE, LIGHT_BLACK, BLUE, BLUE, LIGHT_BLACK},                              //back
			{RED, RED, RED, RED, RED, LIGHT_BLACK, LIGHT_BLACK, LIGHT_BLACK, LIGHT_BLACK},                            //right
			{LIGHT_BLACK, LIGHT_BLACK, LIGHT_BLACK, LIGHT_BLACK, YELLOW, RED, LIGHT_BLACK, LIGHT_BLACK, LIGHT_BLACK}, //top
			{WHITE, WHITE, WHITE, WHITE, WHITE, WHITE, WHITE, WHITE, WHITE},                                          //bottom
		},
		{ //2
			{GREEN, GREEN, LIGHT_BLACK, GREEN, GREEN, LIGHT_BLACK, GREEN, LIGHT_BLACK, LIGHT_BLACK},                    //front
			{ORANGE, ORANGE, ORANGE, ORANGE, ORANGE, ORANGE, LIGHT_BLACK, LIGHT_BLACK, LIGHT_BLACK},                    //left
			{BLUE, BLUE, LIGHT_BLACK, BLUE, BLUE, LIGHT_BLACK, BLUE, BLUE, LIGHT_BLACK},                                //back
			{RED, RED, RED, RED, RED, LIGHT_BLACK, LIGHT_BLACK, RED, LIGHT_BLACK},                                      //right
			{LIGHT_BLACK, LIGHT_BLACK, LIGHT_BLACK, LIGHT_BLACK, YELLOW, LIGHT_BLACK, LIGHT_BLACK, GREEN, LIGHT_BLACK}, //top
			{WHITE, WHITE, WHITE, WHITE, WHITE, WHITE, WHITE, WHITE, WHITE},                                            //bottom
		},
		{ //3
			{GREEN, GREEN, LIGHT_BLACK, GREEN, GREEN, YELLOW, GREEN, GREEN, LIGHT_BLACK},           //front
			{ORANGE, ORANGE, ORANGE, ORANGE, ORANGE, ORANGE, LIGHT_BLACK, YELLOW, LIGHT_BLACK},     //left
			{BLUE, BLUE, LIGHT_BLACK, BLUE, BLUE, YELLOW, BLUE, BLUE, LIGHT_BLACK},                 //back
			{RED, RED, RED, RED, RED, RED, LIGHT_BLACK, YELLOW, LIGHT_BLACK},                       //right
			{LIGHT_BLACK, ORANGE, LIGHT_BLACK, BLUE, YELLOW, GREEN, LIGHT_BLACK, RED, LIGHT_BLACK}, //top
			{WHITE, WHITE, WHITE, WHITE, WHITE, WHITE, WHITE, WHITE, WHITE},                        //bottom
		},
		{ //4
			{GREEN, GREEN, LIGHT_BLACK, GREEN, GREEN, BLUE, GREEN, GREEN, LIGHT_BLACK},                //front
			{ORANGE, ORANGE, ORANGE, ORANGE, ORANGE, ORANGE, LIGHT_BLACK, YELLOW, LIGHT_BLACK},        //left
			{BLUE, BLUE, LIGHT_BLACK, BLUE, BLUE, YELLOW, BLUE, BLUE, LIGHT_BLACK},                    //back
			{RED, RED, RED, RED, RED, RED, LIGHT_BLACK, GREEN, LIGHT_BLACK},                           //right
			{LIGHT_BLACK, ORANGE, LIGHT_BLACK, RED, YELLOW, YELLOW, LIGHT_BLACK, YELLOW, LIGHT_BLACK}, //top
			{WHITE, WHITE, WHITE, WHITE, WHITE, WHITE, WHITE, WHITE, WHITE},                           //bottom
		},
		{ //5
			{ORANGE, ORANGE, LIGHT_BLACK, ORANGE, ORANGE, YELLOW, ORANGE, ORANGE, LIGHT_BLACK},         //front
			{BLUE, BLUE, BLUE, BLUE, BLUE, BLUE, LIGHT_BLACK, GREEN, LIGHT_BLACK},                      //left
			{RED, RED, LIGHT_BLACK, RED, RED, YELLOW, RED, RED, LIGHT_BLACK},                           //back
			{GREEN, GREEN, GREEN, GREEN, GREEN, GREEN, LIGHT_BLACK, RED, LIGHT_BLACK},                  //right
			{LIGHT_BLACK, YELLOW, LIGHT_BLACK, BLUE, YELLOW, ORANGE, LIGHT_BLACK, YELLOW, LIGHT_BLACK}, //top
			{WHITE, WHITE, WHITE, WHITE, WHITE, WHITE, WHITE, WHITE, WHITE},                            //bottom
		},
		{ //6
			{GREEN, GREEN, LIGHT_BLACK, GREEN, GREEN, ORANGE, GREEN, GREEN, LIGHT_BLACK},                 //front
			{ORANGE, ORANGE, ORANGE, ORANGE, ORANGE, ORANGE, LIGHT_BLACK, BLUE, LIGHT_BLACK},             //left
			{BLUE, BLUE, LIGHT_BLACK, BLUE, BLUE, GREEN, BLUE, BLUE, LIGHT_BLACK},                        //back
			{RED, RED, RED, RED, RED, RED, LIGHT_BLACK, RED, LIGHT_BLACK},                                //right
			{LIGHT_BLACK, YELLOW, LIGHT_BLACK, YELLOW, YELLOW, YELLOW, LIGHT_BLACK, YELLOW, LIGHT_BLACK}, //top
			{WHITE, WHITE, WHITE, WHITE, WHITE, WHITE, WHITE, WHITE, WHITE},                              //bottom
		},
		{ //7
			{GREEN, GREEN, LIGHT_BLACK, GREEN, GREEN, GREEN, GREEN, GREEN, RED},                     //front
			{ORANGE, ORANGE, ORANGE, ORANGE, ORANGE, ORANGE, LIGHT_BLACK, ORANGE, LIGHT_BLACK},      //left
			{BLUE, BLUE, LIGHT_BLACK, BLUE, BLUE, BLUE, BLUE, BLUE, LIGHT_BLACK},                    //back
			{RED, RED, RED, RED, RED, RED, LIGHT_BLACK, RED, BLUE},                                  //right
			{LIGHT_BLACK, YELLOW, LIGHT_BLACK, YELLOW, YELLOW, YELLOW, LIGHT_BLACK, YELLOW, YELLOW}, //top
			{WHITE, WHITE, WHITE, WHITE, WHITE, WHITE, WHITE, WHITE, WHITE},                         //bottom
		},
		{ //8
			{BLUE, RED, RED, RED, RED, RED, YELLOW, RED, RED},                       //front
			{BLUE, BLUE, YELLOW, BLUE, BLUE, BLUE, BLUE, BLUE, BLUE},                //left
			{ORANGE, ORANGE, ORANGE, ORANGE, ORANGE, ORANGE, GREEN, ORANGE, ORANGE}, //back
			{YELLOW, GREEN, RED, GREEN, GREEN, GREEN, GREEN, GREEN, GREEN},          //right
			{WHITE, WHITE, WHITE, WHITE, WHITE, WHITE, WHITE, WHITE, WHITE},         //top
			{YELLOW, YELLOW, RED, YELLOW, YELLOW, YELLOW, ORANGE, YELLOW, GREEN},    //bottom
		},
	}

	solutions   = [8]string{}
	shouldReset = [8]bool{
		true,  //1
		true,  //2
		true,  //3
		true,  //4
		true,  //5
		true,  //6
		true,  //7
		false, //8
	}

	hints = [8][]Hint{
		{ //1
			{RTop, true},
			{RRight, true},
			{RTop, false},
			{RRight, false},
			{RTop, false},
			{RFront, false},
			{RTop, true},
			{RFront, true},
		},
		{ //2
			{RTop, false},
			{RFront, false},
			{RTop, true},
			{RFront, true},
			{RTop, true},
			{RRight, true},
			{RTop, false},
			{RRight, false},
		},
		{ //3
			{RFront, true},
			{RRight, true},
			{RTop, true},
			{RRight, false},
			{RTop, false},
			{RFront, false},
		},
		{ //4
			{RFront, true},
			{RRight, true},
			{RTop, true},
			{RRight, false},
			{RTop, false},
			{RFront, false},
		},
		{ //5
			{RFront, true},
			{RRight, true},
			{RTop, true},
			{RRight, false},
			{RTop, false},
			{RFront, false},
		},
		{ //6
			{RRight, false},
			{RTop, false},
			{RRight, true},
			{RTop, false},
			{RRight, false},
			{RTop, false},
			{RTop, false},
			{RRight, true},
		},
		{ //7
			{RRight, true},
			{RTop, false},
			{RLeft, true},
			{RTop, true},
			{RRight, false},
			{RTop, false},
			{RLeft, false},
			{RTop, true},
		},
		{ //8
			{RRight, true},
			{RTop, true},
			{RRight, false},
			{RTop, false},
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
	tutorialScene.docPointer = 0 //starts from 0, can be set to 1+ for debugging
	tutorialScene.cubes = make([]*Cube, len(hints))
	for i := range hints {
		tutorialScene.cubes[i] = NewCube(3, split(tutorials[i]), a)
	}
	for i := range hints {
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
	cube := ts.cubes[ts.docPointer]
	if !cube.isRotating() {
		cube.hintPointer += 1
		if cube.hintPointer >= len(hints[ts.docPointer]) {
			cube.hintPointer = 0
			if shouldReset[ts.docPointer] {
				ts.Reset()
			} else {
				ts.cubes[ts.docPointer].RotateAny(hints[ts.docPointer][cube.hintPointer].rotation, hints[ts.docPointer][cube.hintPointer].isForward, true)
			}
		} else {
			ts.cubes[ts.docPointer].RotateAny(hints[ts.docPointer][cube.hintPointer].rotation, hints[ts.docPointer][cube.hintPointer].isForward, true)
		}
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

	rl.DrawText(solutions[ts.docPointer], 20, winHeight-40, subTitleTextFontSize, rl.Blue)
	rl.DrawText(strconv.Itoa(ts.docPointer+1)+"/"+strconv.Itoa(len(hints)), winWidth-120, winHeight-48, subTitleTextFontSize, rl.Black)

	//rl.DrawFPS(5, 5)
	rl.EndDrawing()
}
