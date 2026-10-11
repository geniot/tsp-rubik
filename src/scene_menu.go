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
	gameDrawTarget      rl.RenderTexture2D
	gameSourceRect      rl.Rectangle
	gameDestRect        rl.Rectangle
}

func NewMenuScene(app *Application) *MenuScene {
	gs := MenuScene{}
	gs.a = app

	gs.gameDrawTarget = rl.LoadRenderTexture(ScreenLogicalWidth, ScreenLogicalHeight)
	gs.gameSourceRect = rl.NewRectangle(0, 0, float32(ScreenLogicalWidth), -float32(ScreenLogicalHeight))
	gs.gameDestRect = rl.NewRectangle(0, 0, float32(ScreenLogicalWidth), float32(ScreenLogicalHeight))

	gs.yTitleOffset = 100
	gs.yButtonsOffset = 250
	gs.yButtonsSpacing = 30
	gs.buttonWidth = 400
	gs.buttonHeight = 60
	gs.titleTextSize = rl.MeasureTextEx(rl.GetFontDefault(), titleText, titleTextFontSize, 1)
	gs.subTitleTextSize = rl.MeasureTextEx(rl.GetFontDefault(), subTitleText, subTitleTextFontSize, 1)
	gs.isExitButtonClicked = false

	return &gs
}

func (ms *MenuScene) HitsAnyButton(_ rl.Vector2) bool {
	return false
}

func (ms *MenuScene) ShouldExit() bool {
	return ms.isExitButtonClicked
}

func (ms *MenuScene) Render(drawTarget rl.RenderTexture2D, _ *rl.Camera) {
	rl.BeginTextureMode(ms.gameDrawTarget)
	{
		//rl.UpdateCamera(&camera, rl.CameraOrbital)
		rl.ClearBackground(rl.RayWhite)
		rl.Color4f(1, 1, 1, 1)

		rl.DrawText(titleText, int32((ScreenLogicalWidth-ms.titleTextSize.X)/2), ms.yTitleOffset, titleTextFontSize, rl.Blue)
		rl.DrawText(subTitleText, int32((ScreenLogicalWidth-ms.subTitleTextSize.X)/2), ms.yTitleOffset+int32(ms.titleTextSize.Y), subTitleTextFontSize, rl.DarkGreen)

		buttonCount := float32(0)
		isButtonClicked := false

		setTextStyle(defaultTextStyle)

		//new game startCode
		gui.SetState(If(rl.IsGamepadButtonDown(gamePadId, startCode), gui.STATE_PRESSED, gui.STATE_NORMAL))
		isButtonClicked = gui.Button(rl.NewRectangle((ScreenLogicalWidth-ms.buttonWidth)/2,
			ms.yButtonsOffset+ms.buttonHeight*buttonCount+ms.yButtonsSpacing*buttonCount,
			ms.buttonWidth, ms.buttonHeight), "[ST] New Game")
		if isButtonClicked || rl.IsGamepadButtonReleased(gamePadId, startCode) {
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
		isButtonClicked = gui.Button(rl.NewRectangle((ScreenLogicalWidth-ms.buttonWidth)/2,
			ms.yButtonsOffset+ms.buttonHeight*buttonCount+ms.yButtonsSpacing*buttonCount,
			ms.buttonWidth, ms.buttonHeight), "[Y] Continue")
		if gui.GetState() != gui.STATE_DISABLED && (isButtonClicked || rl.IsGamepadButtonReleased(gamePadId, yCode)) {
			ms.a.currentSceneIndex = gameSceneKey
		}
		buttonCount += 1
		//controls aCode
		gui.SetState(If(rl.IsGamepadButtonDown(gamePadId, aCode), gui.STATE_PRESSED, gui.STATE_NORMAL))
		isButtonClicked = gui.Button(rl.NewRectangle((ScreenLogicalWidth-ms.buttonWidth)/2,
			ms.yButtonsOffset+ms.buttonHeight*buttonCount+ms.yButtonsSpacing*buttonCount,
			ms.buttonWidth, ms.buttonHeight), "[A] Controls")
		if isButtonClicked || rl.IsGamepadButtonReleased(gamePadId, aCode) {
			ms.a.currentSceneIndex = controlsSceneKey
		}
		buttonCount += 1
		//tutorial xCode
		gui.SetState(If(rl.IsGamepadButtonDown(gamePadId, xCode), gui.STATE_PRESSED, gui.STATE_NORMAL))
		isButtonClicked = gui.Button(rl.NewRectangle((ScreenLogicalWidth-ms.buttonWidth)/2,
			ms.yButtonsOffset+ms.buttonHeight*buttonCount+ms.yButtonsSpacing*buttonCount,
			ms.buttonWidth, ms.buttonHeight), "[X] Tutorial")
		if isButtonClicked || rl.IsGamepadButtonReleased(gamePadId, xCode) {
			ms.a.currentSceneIndex = tutorialSceneKey
		}
		if !ms.a.isWeb {
			buttonCount += 1
			//exit bCode
			gui.SetState(If(rl.IsGamepadButtonDown(gamePadId, bCode) || ms.isExitButtonClicked, gui.STATE_PRESSED, gui.STATE_NORMAL))
			isButtonClicked = gui.Button(rl.NewRectangle((ScreenLogicalWidth-ms.buttonWidth)/2,
				ms.yButtonsOffset+ms.buttonHeight*buttonCount+ms.yButtonsSpacing*buttonCount,
				ms.buttonWidth, ms.buttonHeight), "[B] Exit")
			if isButtonClicked || rl.IsGamepadButtonReleased(gamePadId, bCode) {
				ms.isExitButtonClicked = true
			}
		}
		//
		//rl.DrawFPS(5, 5)
	}
	rl.EndTextureMode()

	rl.BeginTextureMode(drawTarget)
	{
		rl.ClearBackground(rl.Black)
		rl.DrawTexturePro(ms.gameDrawTarget.Texture, ms.gameSourceRect, ms.gameDestRect, ZeroVector2, 0, rl.White)
	}
	rl.EndTextureMode()
}

func (ms *MenuScene) Update() {

}
