package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"math"
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
	R_ALL_LEFT
	R_ALL_RIGHT
	R_ALL_FRONT
	R_ALL_BACK
	R_ALL_TOP
	R_ALL_BOTTOM
)

const (
	scaleMax   = float64(300)
	scaleAvg   = float64(30)
	scaleMin   = float64(5)
	scaleSpeed = 0.005
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
	scaleFrom        float64
	scaleTo          float64
	scaleDirection   bool
	scaleFactor      float64
}

// NewCube front-green, back-blue, left-orange, right-red, top-yellow, bottom-white
func NewCube(size int) *Cube {
	cubies := [3][3][3]*Cubie{}
	for xIterator := 0; xIterator < size; xIterator++ {
		for yIterator := 0; yIterator < size; yIterator++ {
			for zIterator := 0; zIterator < size; zIterator++ {
				frontColor := If(zIterator == size-1, G, LB)
				leftColor := If(xIterator == 0, O, LB)
				backColor := If(zIterator == 0, B, LB)
				rightColor := If(xIterator == size-1, R, LB)
				topColor := If(yIterator == size-1, Y, LB)
				bottomColor := If(yIterator == 0, W, LB)
				cubies[xIterator][yIterator][zIterator] = NewCubie([6]int{frontColor, leftColor, backColor, rightColor, topColor, bottomColor}, xIterator-1, yIterator-1, zIterator-1)
			}
		}
	}
	return &Cube{size: size,
		scaleFrom:        scaleMax,
		scaleTo:          scaleMax,
		scaleDirection:   false,
		scaleFactor:      scaleMax,
		selectedRotation: R_NONE,
		cubies:           cubies}
}

func (c *Cube) isCorrect() bool {
	return c.isFaceCorrect(RIGHT) &&
		c.isFaceCorrect(FRONT) &&
		c.isFaceCorrect(BACK) &&
		c.isFaceCorrect(LEFT) &&
		c.isFaceCorrect(TOP) &&
		c.isFaceCorrect(BOTTOM)
}

func (c *Cube) isFaceCorrect(face int) bool {
	cubies := c.getCubiesByFace(face)
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

func (c *Cube) getCubiesByFace(face int) []*Cubie {
	var cubies = make([]*Cubie, 0)
	for xIterator := 0; xIterator < c.size; xIterator++ {
		for yIterator := 0; yIterator < c.size; yIterator++ {
			for zIterator := 0; zIterator < c.size; zIterator++ {
				cubie := c.cubies[xIterator][yIterator][zIterator]
				if cubie.isInFace(face) {
					cubies = append(cubies, cubie)
				}
			}
		}
	}
	return cubies
}

func (c *Cube) update() {
	isRotationFinished := false //used to trigger cubie's color model update
	if c.isRotating() {
		c.angle -= rotationSpeed
		if c.angle <= 0 {
			c.angle = 0
			isRotationFinished = true
		}
	}

	//scaling is based on isCorrect()
	if !c.isRotating() && c.isCorrect() {
		c.scaleFrom = scaleMin
		c.scaleTo = scaleAvg
	} else {
		c.scaleFrom = scaleMin
		c.scaleTo = scaleMax
		c.scaleDirection = true
	}
	c.updateScale()

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
		if rl.IsKeyDown(rl.KeyUp) {
			c.angle = 90
			c.selectedRotation = If(c.selectedRotation == R_NONE, If(rl.IsKeyDown(rl.KeyLeftControl), R_ALL_TOP, R_ALL_FRONT), c.selectedRotation)
			c.isForward = If(c.selectedRotation <= R_BACK, true, false)
			c.isForward = If(rl.IsKeyDown(rl.KeyLeftControl), !c.isForward, c.isForward)
		}
		if rl.IsKeyDown(rl.KeyDown) {
			c.angle = 90
			c.selectedRotation = If(c.selectedRotation == R_NONE, If(rl.IsKeyDown(rl.KeyLeftControl), R_ALL_BOTTOM, R_ALL_BACK), c.selectedRotation)
			c.isForward = If(c.selectedRotation <= R_BACK, false, true)
			c.isForward = If(rl.IsKeyDown(rl.KeyLeftControl), !c.isForward, c.isForward)
		}
		if rl.IsKeyDown(rl.KeyLeft) {
			c.angle = 90
			c.selectedRotation = If(c.selectedRotation == R_NONE, R_ALL_LEFT, c.selectedRotation)
			c.isForward = If(c.selectedRotation <= R_BACK, true, false)
		}
		if rl.IsKeyDown(rl.KeyRight) {
			c.angle = 90
			c.selectedRotation = If(c.selectedRotation == R_NONE, R_ALL_RIGHT, c.selectedRotation)
			c.isForward = If(c.selectedRotation <= R_BACK, false, true)
		}
	}
	//shuffle mode
	if isShuffling() && !c.isRotating() {
		newSelectedRotation := int(rand.Int31n(9)) + 1
		for newSelectedRotation == c.selectedRotation {
			newSelectedRotation = int(rand.Int31n(9)) + 1
		}
		c.selectedRotation = newSelectedRotation
		c.angle = 90 //float32(rand.Int31n(3)) * 90
		c.isForward = If(rand.Int31n(2) == 0, true, false)
	}

	for xIterator := 0; xIterator < c.size; xIterator++ {
		for yIterator := 0; yIterator < c.size; yIterator++ {
			for zIterator := 0; zIterator < c.size; zIterator++ {
				cubie := c.cubies[xIterator][yIterator][zIterator]
				cubie.update(c.selectedRotation, c.isRotating(), c.isForward, isRotationFinished)
			}
		}
	}

	if !c.isRotating() {
		if isRotationFinished && c.isCorrect() { //you win!
			c.selectedRotation = R_NONE
		}
		if c.selectedRotation == R_ALL_LEFT || c.selectedRotation == R_ALL_RIGHT ||
			c.selectedRotation == R_ALL_FRONT || c.selectedRotation == R_ALL_BACK ||
			c.selectedRotation == R_ALL_TOP || c.selectedRotation == R_ALL_BOTTOM {
			c.selectedRotation = R_NONE
		}
	}
}

func isShuffling() bool {
	return rl.IsKeyDown(rl.KeyS)
}

func (c *Cube) draw() {
	for xIterator := 0; xIterator < c.size; xIterator++ {
		for yIterator := 0; yIterator < c.size; yIterator++ {
			for zIterator := 0; zIterator < c.size; zIterator++ {
				cubie := c.cubies[xIterator][yIterator][zIterator]
				isSelected := If(isShuffling(), false, If(cubie.shouldSelect(c.selectedRotation, true), true, false))
				cubie.draw(isSelected, float32(c.scaleFactor))
			}
		}
	}
}

func (c *Cube) updateScale() {
	if (c.scaleFactor <= c.scaleFrom && !c.scaleDirection) || (c.scaleFactor >= c.scaleTo && c.scaleDirection) {
		c.scaleDirection = !c.scaleDirection
	} else {
		speed := If(c.scaleDirection, scaleSpeed, -scaleSpeed)
		speed += speed * c.scaleFactor
		c.scaleFactor += math.Sqrt(c.scaleFactor) * speed
	}
}

func (c *Cube) isRotating() bool {
	return c.angle != 0
}
