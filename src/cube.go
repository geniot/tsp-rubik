package main

import (
	"math"
	"math/rand"
)

const (
	rotationSpeedNormal = float32(3)
	rotationSpeedInc    = float32(0.5)
	rotationSpeedMax    = float32(90)
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

type Cube struct {
	size                  int
	cubies                [3][3][3]*Cubie
	angle                 float32
	rotationSpeed         float32
	isForward             bool
	selectedRotation      int
	scaleFrom             float64
	scaleTo               float64
	scaleDirection        bool
	scaleFactor           float64
	isCorrect             bool
	isShuffling           bool
	isFaceSelectionModeOn bool
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
		isCorrect:             true,
		isShuffling:           false,
		isFaceSelectionModeOn: false,
		rotationSpeed:         rotationSpeedNormal,
		scaleFrom:             scaleMax,
		scaleTo:               scaleMax,
		scaleDirection:        false,
		scaleFactor:           scaleMax,
		selectedRotation:      R_NONE,
		cubies:                cubies}
}

func (c *Cube) updateCorrect() {
	c.isCorrect = c.isFaceCorrect(RIGHT) &&
		c.isFaceCorrect(FRONT) &&
		c.isFaceCorrect(BACK) &&
		c.isFaceCorrect(LEFT) &&
		c.isFaceCorrect(TOP) &&
		c.isFaceCorrect(BOTTOM)
}

func (c *Cube) updateShuffling() {
	//shuffle mode
	if c.isShuffling {
		c.rotationSpeed += rotationSpeedInc
		if c.rotationSpeed > rotationSpeedMax {
			c.rotationSpeed = rotationSpeedMax
		}
		newSelectedRotation := int(rand.Int31n(9)) + 1
		for newSelectedRotation == c.selectedRotation {
			newSelectedRotation = int(rand.Int31n(9)) + 1
		}
		c.selectedRotation = newSelectedRotation
		c.angle = 90 //float32(rand.Int31n(3)) * 90
		c.isForward = If(rand.Int31n(2) == 0, true, false)
	} else {
		c.rotationSpeed = rotationSpeedNormal
	}
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
	c.iterateCubies(func(cubie *Cubie) {
		if cubie.isInFace(face) {
			cubies = append(cubies, cubie)
		}
	})
	return cubies
}

type applyToEachCubie func(cubie *Cubie)

func (c *Cube) iterateCubies(fn applyToEachCubie) {
	for xIterator := 0; xIterator < c.size; xIterator++ {
		for yIterator := 0; yIterator < c.size; yIterator++ {
			for zIterator := 0; zIterator < c.size; zIterator++ {
				cubie := c.cubies[xIterator][yIterator][zIterator]
				fn(cubie)
			}
		}
	}
}

func (c *Cube) update() {
	c.iterateCubies(func(cubie *Cubie) {
		if cubie.shouldSelect(c.selectedRotation) && c.isRotating() {
			cubie.update(c.selectedRotation, c.isForward, c.rotationSpeed, c.angle)
		}
	})

	isRotationJustFinished := false //used to trigger cubie's color model update
	if c.isRotating() {
		c.angle -= c.rotationSpeed
		if c.angle <= 0 {
			c.angle = 0
			isRotationJustFinished = true
		}
	}
	c.updateScale()

	c.iterateCubies(func(cubie *Cubie) {
		if isRotationJustFinished {
			cubie.updateGlobalColors(c.selectedRotation, c.isForward)
		}
	})

	if isRotationJustFinished {
		c.updateCorrect()
	}

	if !c.isRotating() {
		handleUserEvents(c)
		c.updateShuffling()
	}
}

func (c *Cube) draw() {
	c.iterateCubies(func(cubie *Cubie) {
		isSelected := If(c.isFaceSelectionModeOn, If(cubie.shouldSelect(c.selectedRotation), true, false), false)
		cubie.draw(isSelected, float32(c.scaleFactor))
	})
}

func (c *Cube) updateScale() {
	if (c.scaleFactor <= c.scaleFrom && !c.scaleDirection) || (c.scaleFactor >= c.scaleTo && c.scaleDirection) {
		c.scaleDirection = !c.scaleDirection
	} else {
		speed := If(c.scaleDirection, scaleSpeed, -scaleSpeed)
		speed += speed * c.scaleFactor
		c.scaleFactor += math.Sqrt(c.scaleFactor) * speed
	}
	//scaling is based on isCorrect
	if c.isCorrect {
		c.scaleFrom = scaleMin
		c.scaleTo = scaleAvg
	} else {
		c.scaleFrom = scaleMin
		c.scaleTo = scaleMax
		c.scaleDirection = true
	}
}

func (c *Cube) isRotating() bool {
	return c.angle != 0
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}
