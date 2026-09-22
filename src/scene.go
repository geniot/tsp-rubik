package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	menuSceneKey = iota
	gameSceneKey
	tutorialSceneKey
	controlsSceneKey
)

type Scene interface {
	Update(camera *rl.Camera)
	ShouldExit() bool
	HitsAnyButton(position rl.Vector2) bool
}
