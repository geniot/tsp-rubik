package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"math/rand"
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

type Cube struct {
	size             int
	cubies           [3][3][3]*Cubie
	angle            float32
	isForward        bool
	selectedRotation int
}

// NewCube front-green, back-LBue, left-orange, right-red, top-yellow, bottom-white
func NewCube(size int) *Cube {
	//todo: use size to generate cubie config dynamically, also update possiLBe rotations
	return &Cube{size: size,
		isForward:        false,
		selectedRotation: R_NONE,
		cubies: [3][3][3]*Cubie{
			{
				{
					NewCubie([6]int{LB, O, B, LB, LB, W}, -1, -1, -1),
					NewCubie([6]int{LB, O, LB, LB, LB, W}, -1, -1, 0),
					NewCubie([6]int{G, O, LB, LB, LB, W}, -1, -1, 1),
				},
				{
					NewCubie([6]int{LB, O, B, LB, LB, LB}, -1, 0, -1),
					NewCubie([6]int{LB, O, LB, LB, LB, LB}, -1, 0, 0),
					NewCubie([6]int{G, O, LB, LB, LB, LB}, -1, 0, 1),
				},
				{
					NewCubie([6]int{LB, O, B, LB, Y, LB}, -1, 1, -1),
					NewCubie([6]int{LB, O, LB, LB, Y, LB}, -1, 1, 0),
					NewCubie([6]int{G, O, LB, LB, Y, LB}, -1, 1, 1),
				},
			},
			{
				{
					NewCubie([6]int{LB, LB, B, LB, LB, W}, 0, -1, -1),
					NewCubie([6]int{LB, LB, LB, LB, LB, W}, 0, -1, 0),
					NewCubie([6]int{G, LB, LB, LB, LB, W}, 0, -1, 1),
				},
				{
					NewCubie([6]int{LB, LB, B, LB, LB, LB}, 0, 0, -1),
					NewCubie([6]int{LB, LB, LB, LB, LB, LB}, 0, 0, 0),
					NewCubie([6]int{G, LB, LB, LB, LB, LB}, 0, 0, 1),
				},
				{
					NewCubie([6]int{LB, LB, B, LB, Y, LB}, 0, 1, -1),
					NewCubie([6]int{LB, LB, LB, LB, Y, LB}, 0, 1, 0),
					NewCubie([6]int{G, LB, LB, LB, Y, LB}, 0, 1, 1),
				},
			},
			{
				{
					NewCubie([6]int{LB, LB, B, R, LB, W}, 1, -1, -1),
					NewCubie([6]int{LB, LB, LB, R, LB, W}, 1, -1, 0),
					NewCubie([6]int{G, LB, LB, R, LB, W}, 1, -1, 1),
				},
				{
					NewCubie([6]int{LB, LB, B, R, LB, LB}, 1, 0, -1),
					NewCubie([6]int{LB, LB, LB, R, LB, LB}, 1, 0, 0),
					NewCubie([6]int{G, LB, LB, R, LB, LB}, 1, 0, 1),
				},
				{
					NewCubie([6]int{LB, LB, B, R, Y, LB}, 1, 1, -1),
					NewCubie([6]int{LB, LB, LB, R, Y, LB}, 1, 1, 0),
					NewCubie([6]int{G, LB, LB, R, Y, LB}, 1, 1, 1),
				},
			},
		}}
}

func (c *Cube) updateThenDraw() {
	if c.isRotating() {
		c.angle -= rotationSpeed
		if c.angle <= 0 {
			c.angle = 0
		}
	}
	if !c.isRotating() {
		for key, rotation := range keysToRotationsMap {
			if rl.IsKeyDown(key) {
				c.selectedRotation = rotation
				isShuffle = false
			}
		}
		if rl.IsKeyDown(rl.KeyUp) {
			c.angle = 90
			c.isForward = true
		}
		if rl.IsKeyDown(rl.KeyDown) {
			c.angle = 90
			c.isForward = false
		}
	}
	if rl.GetCharPressed() != 0 {
		isShuffle = false
	}
	if rl.IsKeyDown(rl.KeyS) {
		isShuffle = true
	}
	//shuffle mode
	if (isShuffle || c.selectedRotation == R_NONE) && !c.isRotating() {
		isShuffle = true

		newSelectedRotation := int(rand.Int31n(9)) + 1
		for newSelectedRotation == c.selectedRotation {
			newSelectedRotation = int(rand.Int31n(9)) + 1
		}
		c.selectedRotation = newSelectedRotation
		c.angle = float32(rand.Int31n(3)) * 90
		c.isForward = If(rand.Int31n(2) == 0, true, false)
	}

	for xIterator := 0; xIterator < c.size; xIterator++ {
		for yIterator := 0; yIterator < c.size; yIterator++ {
			for zIterator := 0; zIterator < c.size; zIterator++ {
				cubie := c.cubies[xIterator][yIterator][zIterator]
				cubie.update(c.selectedRotation, c.isRotating(), c.isForward)
				cubie.draw()
			}
		}
	}
}

func (c *Cube) isRotating() bool {
	return c.angle != 0
}
