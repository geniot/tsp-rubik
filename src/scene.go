package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	menuSceneKey = iota
	gameSceneKey
	tutorialSceneKey
)

type Scene interface {
	Draw(camera *rl.Camera)
	ShouldExit() bool
}

var (
	currentSceneIndex = tutorialSceneKey
)
