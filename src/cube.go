package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
)

const (
	rotationSpeedNormal = float32(3)
	rotationSpeedInc    = float32(0.5)
	rotationSpeedMax    = float32(90)
)

// rotations
const (
	RNone = iota
	RFront
	RFbMiddle
	RBack
	RLeft
	RLrMiddle
	RRight
	RTop
	RTbMiddle
	RBottom
	RAllLeft
	RAllRight
	RAllFront
	RAllBack
	RAllTop
	RAllBottom
)

// rotations by first letters, used in tutorials
var rotationLetters = map[int]string{
	RFront:  "F",
	RBack:   "B",
	RTop:    "U",
	RBottom: "D",
	RRight:  "R",
	RLeft:   "L",
}

const (
	scaleMax     = float64(300)
	scaleAvg     = float64(30)
	scaleMin     = float64(5)
	scaleSpeed   = 0.005
	shuffleCount = 10000
)

type Cube struct {
	application           *Application
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
	hintPointer           int
}

// NewCube front-green, back-blue, left-orange, right-red, top-yellow, bottom-white
func NewCube(size int, colors [3][3][3][6]int, a *Application) *Cube {
	cubies := [3][3][3]*Cubie{}
	for xIterator := 0; xIterator < size; xIterator++ {
		for yIterator := 0; yIterator < size; yIterator++ {
			for zIterator := 0; zIterator < size; zIterator++ {

				frontColor := colors[xIterator][yIterator][zIterator][0]
				leftColor := colors[xIterator][yIterator][zIterator][1]
				backColor := colors[xIterator][yIterator][zIterator][2]
				rightColor := colors[xIterator][yIterator][zIterator][3]
				topColor := colors[xIterator][yIterator][zIterator][4]
				bottomColor := colors[xIterator][yIterator][zIterator][5]

				cubies[xIterator][yIterator][zIterator] = NewCubie(
					[6]int{frontColor, leftColor, backColor, rightColor, topColor, bottomColor},
					xIterator-1, yIterator-1, zIterator-1, a)
			}
		}
	}
	cube := Cube{
		application:           a,
		size:                  size,
		isCorrect:             true,
		isShuffling:           false,
		isFaceSelectionModeOn: false,
		rotationSpeed:         rotationSpeedNormal,
		scaleFrom:             scaleMax,
		scaleTo:               scaleMax,
		scaleDirection:        false,
		scaleFactor:           scaleMax,
		selectedRotation:      RNone,
		cubies:                cubies,
		hintPointer:           -1,
	}

	//cube.debug()
	cube.updateCorrect()
	return &cube
}

func (c *Cube) updateCorrect() {
	c.isCorrect = c.isFaceCorrect(RIGHT) &&
		c.isFaceCorrect(FRONT) &&
		c.isFaceCorrect(BACK) &&
		c.isFaceCorrect(LEFT) &&
		c.isFaceCorrect(TOP) &&
		c.isFaceCorrect(BOTTOM)
}

func (c *Cube) Shuffle(count int) {
	c.rotationSpeed = rotationSpeedMax
	for i := 0; i < count; i++ {
		c.randomize()
		c.iterateCubies(func(cubie *Cubie) {
			if cubie.shouldSelect(c.selectedRotation) && c.isRotating() {
				cubie.update(c.selectedRotation, c.isForward, rotationSpeedMax, c.angle)
			}
		})
		c.angle = 0
	}
	c.updateCorrect()
}

func (c *Cube) randomize() {
	newSelectedRotation := int(rand.Int31n(9)) + 1
	for newSelectedRotation == c.selectedRotation {
		newSelectedRotation = int(rand.Int31n(9)) + 1
	}
	c.selectedRotation = newSelectedRotation
	c.angle = 90 //float32(rand.Int31n(3)) * 90
	c.isForward = If(rand.Int31n(2) == 0, true, false)
}

func (c *Cube) updateShuffling() {
	//shuffle mode
	if c.isShuffling {
		c.rotationSpeed += rotationSpeedInc
		if c.rotationSpeed > rotationSpeedMax {
			c.rotationSpeed = rotationSpeedMax
		}
		c.randomize()
	} else {
		c.rotationSpeed = rotationSpeedNormal
	}
}

func (c *Cube) isFaceCorrect(face int) bool {
	cubies := c.getCubiesByFace(face)
	faceColor := cubies[0].getFaceColor(face)
	for _, cubie := range cubies {
		if cubie.getFaceColor(face) != faceColor {
			return false
		}
	}
	return true
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

	isRotationJustFinished := false
	if c.isRotating() {
		c.angle -= c.rotationSpeed
		if c.angle <= 0 {
			c.angle = 0
			isRotationJustFinished = true
		}
	}
	c.updateScale()

	if isRotationJustFinished {
		c.updateCorrect()
	}

	if !c.isRotating() {
		c.handleUserEvents()
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

var debugColors = map[int]string{
	GREEN:       "\033[0;42m",
	RED:         "\033[0;41m",
	BLUE:        "\033[0;44m",
	ORANGE:      "\033[0;45m",
	WHITE:       "\033[0;47m",
	YELLOW:      "\033[0;43m",
	LIGHT_BLACK: "\033[0;40m",
}

func (c *Cube) debug() {
	sides := [6]int{RIGHT, FRONT, BACK, LEFT, TOP, BOTTOM}
	for _, side := range sides {
		cubies := c.getCubiesByFace(side)
		for _, cubie := range cubies {
			fmt.Fprint(os.Stdout, debugColors[cubie.getFaceColor(side)], " ")
		}
	}
	println("\n===========================================================")
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func round32(x float32) float32 {
	return float32(math.Round(float64(x*10)) / 10)
}
