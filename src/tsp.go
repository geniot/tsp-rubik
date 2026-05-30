package main

import rl "github.com/gen2brain/raylib-go/raylib"

// TSP button codes
const (
	upCode = iota + 1
	rightCode
	downCode
	leftCode
	xCode
	aCode
	bCode
	yCode
	l1Code
	l2Code
	r1Code
	r2Code
	selectCode
	menuCode
	startCode
)

const (
	winHeight = 720
	winWidth  = 1280
	gamePadId = int32(0)
)

func updateTspCameraX(camera *rl.Camera) {
	rightX := rl.GetGamepadAxisMovement(gamePadId, rl.GamepadAxisRightX)
	var rotation = rl.MatrixRotate(rl.GetCameraUp(camera), -1.8*rightX)
	var view = rl.Vector3Subtract(InitialCameraPosition, camera.Target)
	view = rl.Vector3Transform(view, rotation)
	camera.Position = rl.Vector3Add(camera.Target, view)
}

func updateTspCameraY(camera *rl.Camera) {
	rightY := rl.GetGamepadAxisMovement(gamePadId, rl.GamepadAxisRightY)
	factor := If(rightY < 0, 1.5, 1)
	var rotation = rl.MatrixRotate(rl.NewVector3(-1, 0.0, 1), float32(factor)*rightY)
	var view = rl.Vector3Subtract(camera.Position, camera.Target)
	view = rl.Vector3Transform(view, rotation)
	camera.Position = rl.Vector3Add(camera.Target, view)
}

func updateTspCamera(camera *rl.Camera) {
	updateTspCameraX(camera)
	updateTspCameraY(camera)
}
