package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"math"
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

const (
	scaleMax   = float64(70)
	scaleAvg   = float64(40)
	scaleMin   = float64(10)
	scaleSpeed = 0.005
)

type ScaleRange struct {
	scaleFrom      float64
	scaleTo        float64
	scaleDirection bool
	scaleFactor    float64
}

func (c *ScaleRange) update() {
	if (c.scaleFactor <= c.scaleFrom && !c.scaleDirection) || (c.scaleFactor >= c.scaleTo && c.scaleDirection) {
		c.scaleDirection = !c.scaleDirection
	} else {
		speed := If(c.scaleDirection, scaleSpeed, -scaleSpeed)
		speed += speed * c.scaleFactor
		c.scaleFactor += math.Sqrt(c.scaleFactor) * speed
	}

}

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
	scaleRange       ScaleRange
}

// NewCube front-green, back-LBue, left-orange, right-red, top-yellow, bottom-white
func NewCube(size int) *Cube {
	//todo: use size to generate cubie config dynamically, also update possiLBe rotations
	return &Cube{size: size,
		scaleRange:       ScaleRange{scaleFrom: scaleMax, scaleTo: scaleMax, scaleDirection: false, scaleFactor: scaleMax},
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

func (c *Cube) isCorrect() bool {
	return c.isFaceCorrect(RIGHT, R_RIGHT) &&
		c.isFaceCorrect(FRONT, R_FRONT) &&
		c.isFaceCorrect(BACK, R_BACK) &&
		c.isFaceCorrect(LEFT, R_LEFT) &&
		c.isFaceCorrect(TOP, R_TOP) &&
		c.isFaceCorrect(BOTTOM, R_BOTTOM)
}

func (c *Cube) isFaceCorrect(face int, rotation int) bool {
	cubies := c.getCubiesByRotation(rotation)
	var faceColors = make([]int, 0)
	for _, cubie := range cubies {
		faceColors = append(faceColors, cubie.globalColors[face])
	}
	isFaceCorrect := true
	firstColor := faceColors[0]
	for _, color := range faceColors {
		if color != firstColor {
			isFaceCorrect = false
			break
		}
	}
	return isFaceCorrect
}

func (c *Cube) getCubiesByRotation(rotation int) []*Cubie {
	var cubies = make([]*Cubie, 0)
	for xIterator := 0; xIterator < c.size; xIterator++ {
		for yIterator := 0; yIterator < c.size; yIterator++ {
			for zIterator := 0; zIterator < c.size; zIterator++ {
				cubie := c.cubies[xIterator][yIterator][zIterator]
				if cubie.shouldSelect(rotation) {
					cubies = append(cubies, cubie)
				}
			}
		}
	}
	return cubies
}

func (c *Cube) updateThenDraw() {
	isRotationFinished := false
	if c.isRotating() {
		c.angle -= rotationSpeed
		if c.angle <= 0 {
			c.angle = 0
			isRotationFinished = true
		}
	}
	if c.isCorrect() {
		c.scaleRange = ScaleRange{scaleFactor: c.scaleRange.scaleFactor,
			scaleFrom: scaleMin, scaleTo: scaleAvg, scaleDirection: c.scaleRange.scaleDirection}
	} else {
		c.scaleRange = ScaleRange{scaleFactor: c.scaleRange.scaleFactor,
			scaleFrom: scaleMin, scaleTo: scaleMax, scaleDirection: true}
	}
	c.scaleRange.update()

	if !c.isRotating() {
		for key, rotation := range keysToRotationsMap {
			if rl.IsKeyPressed(key) {
				if c.selectedRotation == rotation {
					c.selectedRotation = R_NONE
				} else {
					c.selectedRotation = rotation
				}
			}
		}
		if rl.IsKeyPressed(rl.KeyUp) {
			c.angle = 90
			c.isForward = If(c.selectedRotation <= R_BACK, true, false)
		}
		if rl.IsKeyPressed(rl.KeyDown) {
			c.angle = 90
			c.isForward = If(c.selectedRotation <= R_BACK, false, true)
		}
		if rl.IsKeyPressed(rl.KeyLeft) {
			c.angle = 90
			c.isForward = If(c.selectedRotation <= R_BACK, true, false)
		}
		if rl.IsKeyPressed(rl.KeyRight) {
			c.angle = 90
			c.isForward = If(c.selectedRotation <= R_BACK, false, true)
		}
	}
	//shuffle mode
	//if (c.selectedRotation == R_NONE) && !c.isRotating() {
	//	newSelectedRotation := int(rand.Int31n(9)) + 1
	//	for newSelectedRotation == c.selectedRotation {
	//		newSelectedRotation = int(rand.Int31n(9)) + 1
	//	}
	//	c.selectedRotation = newSelectedRotation
	//	c.angle = float32(rand.Int31n(3)) * 90
	//	c.isForward = If(rand.Int31n(2) == 0, true, false)
	//}

	for xIterator := 0; xIterator < c.size; xIterator++ {
		for yIterator := 0; yIterator < c.size; yIterator++ {
			for zIterator := 0; zIterator < c.size; zIterator++ {
				cubie := c.cubies[xIterator][yIterator][zIterator]
				cubie.update(c.selectedRotation, c.isRotating(), c.isForward, isRotationFinished)
				cubie.draw(float32(c.scaleRange.scaleFactor))
			}
		}
	}
}

func (c *Cube) isRotating() bool {
	return c.angle != 0
}
