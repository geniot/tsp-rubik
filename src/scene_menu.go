package main

import (
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
	a                   *Application
	titleTextSize       rl.Vector2
	subTitleTextSize    rl.Vector2
	yTitleOffset        int32
	yButtonsOffset      float32
	yButtonsSpacing     float32
	buttonWidth         float32
	buttonHeight        float32
	isExitButtonClicked bool
}

func NewMenuScene(app *Application) *MenuScene {
	return &MenuScene{
		a:                   app,
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

func (ms *MenuScene) Update(_ *rl.Camera) {
	//rl.UpdateCamera(&camera, rl.CameraOrbital)
	rl.BeginDrawing()
	rl.ClearBackground(rl.RayWhite)
	rl.Color4f(1, 1, 1, 1)

	rl.DrawText(titleText, int32((winWidth-ms.titleTextSize.X)/2), ms.yTitleOffset, titleTextFontSize, rl.Blue)
	rl.DrawText(subTitleText, int32((winWidth-ms.subTitleTextSize.X)/2), ms.yTitleOffset+int32(ms.titleTextSize.Y), subTitleTextFontSize, rl.DarkGreen)

	buttonCount := float32(0)
	isButtonClicked := false

	gui.SetState(gui.STATE_NORMAL)
	isButtonClicked = gui.Button(rl.NewRectangle((winWidth-ms.buttonWidth)/2, ms.yButtonsOffset+ms.buttonHeight*buttonCount+ms.yButtonsSpacing*buttonCount, ms.buttonWidth, ms.buttonHeight), "[A] New Game")
	if isButtonClicked || rl.IsGamepadButtonPressed(gamePadId, aCode) {
		ms.a.scenes[gameSceneKey].(*GameScene).Reset()
		ms.a.currentSceneIndex = gameSceneKey
	}
	gui.SetState(gui.STATE_NORMAL)
	if !ms.a.scenes[gameSceneKey].(*GameScene).cube.isCorrect {
		buttonCount += 1
		isButtonClicked = gui.Button(rl.NewRectangle((winWidth-ms.buttonWidth)/2, ms.yButtonsOffset+ms.buttonHeight*buttonCount+ms.yButtonsSpacing*buttonCount, ms.buttonWidth, ms.buttonHeight), "[B] Continue")
		if isButtonClicked || rl.IsGamepadButtonPressed(gamePadId, bCode) {
			ms.a.currentSceneIndex = gameSceneKey
		}
	}
	gui.SetState(gui.STATE_NORMAL)
	buttonCount += 1
	isButtonClicked = gui.Button(rl.NewRectangle((winWidth-ms.buttonWidth)/2, ms.yButtonsOffset+ms.buttonHeight*buttonCount+ms.yButtonsSpacing*buttonCount, ms.buttonWidth, ms.buttonHeight), "[X] Tutorial")
	if isButtonClicked || rl.IsGamepadButtonPressed(gamePadId, xCode) {
		ms.a.currentSceneIndex = tutorialSceneKey
	}
	gui.SetState(gui.STATE_NORMAL)
	buttonCount += 1
	ms.isExitButtonClicked = gui.Button(rl.NewRectangle((winWidth-ms.buttonWidth)/2, ms.yButtonsOffset+ms.buttonHeight*buttonCount+ms.yButtonsSpacing*buttonCount, ms.buttonWidth, ms.buttonHeight), "[Y] Exit")
	//rl.DrawFPS(5, 5)
	rl.EndDrawing()
}
