package main

import (
	"fmt"
	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	titleText            = "The Breathing Cube"
	titleTextFontSize    = 60
	subTitleText         = "It's breathing, when it's correct."
	subTitleTextFontSize = 30
)

type MenuScene struct {
	isButtonClicked     bool
	titleTextSize       rl.Vector2
	subTitleTextSize    rl.Vector2
	yTitleOffset        int32
	yButtonsOffset      float32
	yButtonsSpacing     float32
	buttonWidth         float32
	buttonHeight        float32
	isExitButtonClicked bool
}

func NewMenuScene() *MenuScene {
	return &MenuScene{
		yTitleOffset:        100,
		yButtonsOffset:      250,
		yButtonsSpacing:     30,
		buttonWidth:         400,
		buttonHeight:        80,
		titleTextSize:       rl.MeasureTextEx(rl.GetFontDefault(), titleText, titleTextFontSize, 1),
		subTitleTextSize:    rl.MeasureTextEx(rl.GetFontDefault(), subTitleText, subTitleTextFontSize, 1),
		isExitButtonClicked: false,
	}
}

func (ms *MenuScene) ShouldExit() bool {
	return rl.IsKeyPressed(rl.KeyY) || (rl.IsGamepadButtonDown(gamePadId, menuCode) && rl.IsGamepadButtonDown(gamePadId, startCode)) || ms.isExitButtonClicked
}

func (ms *MenuScene) Draw(camera *rl.Camera) {
	//rl.UpdateCamera(&camera, rl.CameraOrbital)
	rl.BeginDrawing()
	rl.ClearBackground(rl.RayWhite)
	rl.Color4f(1, 1, 1, 1)

	rl.DrawText(titleText, int32((winWidth-ms.titleTextSize.X)/2), ms.yTitleOffset, titleTextFontSize, rl.Blue)
	rl.DrawText(subTitleText, int32((winWidth-ms.subTitleTextSize.X)/2), ms.yTitleOffset+int32(ms.titleTextSize.Y), subTitleTextFontSize, rl.DarkGreen)

	gui.SetStyle(gui.DEFAULT, gui.TEXT_SIZE, 40)
	gui.SetStyle(gui.DEFAULT, gui.TEXT_SPACING, 10)
	gui.SetStyle(gui.DEFAULT, gui.TEXT_ALIGNMENT, int64(gui.TEXT_ALIGN_LEFT))
	gui.SetStyle(gui.DEFAULT, gui.TEXT_PADDING, 20)

	buttonCount := float32(0)

	gui.SetState(gui.STATE_NORMAL)
	ms.isButtonClicked = gui.Button(rl.NewRectangle((winWidth-ms.buttonWidth)/2, ms.yButtonsOffset+ms.buttonHeight*buttonCount+ms.yButtonsSpacing*buttonCount, ms.buttonWidth, ms.buttonHeight), "(A) New Game")
	if ms.isButtonClicked {
		scenes[gameSceneKey].(*GameScene).cube.Shuffle(shuffleCount)
		currentSceneIndex = gameSceneKey
	}
	gui.SetState(gui.STATE_NORMAL)
	if !scenes[gameSceneKey].(*GameScene).cube.isCorrect {
		buttonCount += 1
		ms.isButtonClicked = gui.Button(rl.NewRectangle((winWidth-ms.buttonWidth)/2, ms.yButtonsOffset+ms.buttonHeight*buttonCount+ms.yButtonsSpacing*buttonCount, ms.buttonWidth, ms.buttonHeight), "(B) Continue")
		if ms.isButtonClicked {
			currentSceneIndex = gameSceneKey
		}
	}
	gui.SetState(gui.STATE_NORMAL)
	buttonCount += 1
	ms.isButtonClicked = gui.Button(rl.NewRectangle((winWidth-ms.buttonWidth)/2, ms.yButtonsOffset+ms.buttonHeight*buttonCount+ms.yButtonsSpacing*buttonCount, ms.buttonWidth, ms.buttonHeight), "(X) Tutorial")
	if ms.isButtonClicked {
		fmt.Println("Clicked on button")
	}
	gui.SetState(gui.STATE_NORMAL)
	buttonCount += 1
	ms.isExitButtonClicked = gui.Button(rl.NewRectangle((winWidth-ms.buttonWidth)/2, ms.yButtonsOffset+ms.buttonHeight*buttonCount+ms.yButtonsSpacing*buttonCount, ms.buttonWidth, ms.buttonHeight), "(Y) Exit")
	//rl.DrawFPS(5, 5)
	rl.EndDrawing()
}
