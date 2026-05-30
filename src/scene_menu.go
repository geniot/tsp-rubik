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
		buttonHeight:        60,
		titleTextSize:       rl.MeasureTextEx(rl.GetFontDefault(), titleText, titleTextFontSize, 1),
		subTitleTextSize:    rl.MeasureTextEx(rl.GetFontDefault(), subTitleText, subTitleTextFontSize, 1),
		isExitButtonClicked: false,
	}
}

func (ms *MenuScene) ShouldExit() bool {
	return ms.isExitButtonClicked
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

	//new game aCode
	gui.SetState(If(rl.IsGamepadButtonDown(gamePadId, aCode), gui.STATE_PRESSED, gui.STATE_NORMAL))
	isButtonClicked = gui.Button(rl.NewRectangle((winWidth-ms.buttonWidth)/2,
		ms.yButtonsOffset+ms.buttonHeight*buttonCount+ms.yButtonsSpacing*buttonCount,
		ms.buttonWidth, ms.buttonHeight), "[A] New Game")
	if isButtonClicked || rl.IsGamepadButtonReleased(gamePadId, aCode) {
		ms.a.scenes[gameSceneKey].(*GameScene).Reset()
		ms.a.currentSceneIndex = gameSceneKey
	}
	buttonCount += 1
	//continue yCode
	//we only enable the Continue button if the cube is not correct (shuffled) and the game has been started at least once
	//there is no point to continue a finished game
	gui.SetState(If(ms.a.scenes[gameSceneKey].(*GameScene).isStarted,
		If(ms.a.scenes[gameSceneKey].(*GameScene).cube.isCorrect, gui.STATE_DISABLED,
			If(rl.IsGamepadButtonDown(gamePadId, yCode), gui.STATE_PRESSED, gui.STATE_NORMAL)),
		gui.STATE_DISABLED))
	isButtonClicked = gui.Button(rl.NewRectangle((winWidth-ms.buttonWidth)/2,
		ms.yButtonsOffset+ms.buttonHeight*buttonCount+ms.yButtonsSpacing*buttonCount,
		ms.buttonWidth, ms.buttonHeight), "[Y] Continue")
	if isButtonClicked || rl.IsGamepadButtonReleased(gamePadId, yCode) {
		ms.a.currentSceneIndex = gameSceneKey
	}
	buttonCount += 1
	//controls selectCode
	gui.SetState(If(rl.IsGamepadButtonDown(gamePadId, selectCode), gui.STATE_PRESSED, gui.STATE_NORMAL))
	isButtonClicked = gui.Button(rl.NewRectangle((winWidth-ms.buttonWidth)/2,
		ms.yButtonsOffset+ms.buttonHeight*buttonCount+ms.yButtonsSpacing*buttonCount,
		ms.buttonWidth, ms.buttonHeight), "[SL] Controls")
	if isButtonClicked || rl.IsGamepadButtonReleased(gamePadId, selectCode) {
		ms.a.currentSceneIndex = controlsSceneKey
	}
	buttonCount += 1
	//tutorial xCode
	gui.SetState(If(rl.IsGamepadButtonDown(gamePadId, xCode), gui.STATE_PRESSED, gui.STATE_NORMAL))
	isButtonClicked = gui.Button(rl.NewRectangle((winWidth-ms.buttonWidth)/2,
		ms.yButtonsOffset+ms.buttonHeight*buttonCount+ms.yButtonsSpacing*buttonCount,
		ms.buttonWidth, ms.buttonHeight), "[X] Tutorial")
	if isButtonClicked || rl.IsGamepadButtonReleased(gamePadId, xCode) {
		ms.a.currentSceneIndex = tutorialSceneKey
	}
	buttonCount += 1
	//exit bCode
	gui.SetState(If(rl.IsGamepadButtonDown(gamePadId, bCode) || ms.isExitButtonClicked, gui.STATE_PRESSED, gui.STATE_NORMAL))
	isButtonClicked = gui.Button(rl.NewRectangle((winWidth-ms.buttonWidth)/2,
		ms.yButtonsOffset+ms.buttonHeight*buttonCount+ms.yButtonsSpacing*buttonCount,
		ms.buttonWidth, ms.buttonHeight), "[B] Exit")
	if isButtonClicked || rl.IsGamepadButtonReleased(gamePadId, bCode) {
		ms.isExitButtonClicked = true
	}
	//
	//rl.DrawFPS(5, 5)
	rl.EndDrawing()
}
