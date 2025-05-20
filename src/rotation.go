package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	rotationSpeed = float32(5)
)

// rotations
const (
	R_NONE = iota
	R_FRONT
	R_FB_MIDDLE
	R_BACK
	R_LEFT
	R_LR_MIDDLE
	R_RIGHT
	R_TOP
	R_TB_MIDDLE
	R_BOTTOM
)

var (
	keysToRotationsMap = map[int32]int{
		rl.KeyZero:  R_NONE,
		rl.KeyOne:   R_FRONT,
		rl.KeyTwo:   R_FB_MIDDLE,
		rl.KeyThree: R_BACK,
		rl.KeyFour:  R_LEFT,
		rl.KeyFive:  R_LR_MIDDLE,
		rl.KeySix:   R_RIGHT,
		rl.KeySeven: R_TOP,
		rl.KeyEight: R_TB_MIDDLE,
		rl.KeyNine:  R_BOTTOM,
	}
)

type Rotation struct {
	angle            float32
	isForward        bool
	selectedRotation int
}

func NewRotation(selectedRotation int, isForward bool) *Rotation {
	rotation := &Rotation{isForward: isForward, selectedRotation: selectedRotation}
	return rotation
}

func (r *Rotation) update() {
	if r.isRotating() {
		r.angle -= rotationSpeed
		if r.angle <= 0 {
			r.angle = 0
		}
	}
	if !r.isRotating() {
		for key, rotation := range keysToRotationsMap {
			if rl.IsKeyDown(key) {
				r.selectedRotation = rotation
			}
		}
		if rl.IsKeyDown(rl.KeyUp) {
			r.angle = 90
			r.isForward = true
		}
		if rl.IsKeyDown(rl.KeyDown) {
			r.angle = 90
			r.isForward = false
		}
	}
}

func (r *Rotation) isRotating() bool {
	return r.angle != 0
}
