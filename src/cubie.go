package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// faces
const (
	FRONT = iota
	LEFT
	BACK
	RIGHT
	TOP
	BOTTOM
)

var (
	ROTATION1 = [4]int{TOP, RIGHT, BOTTOM, LEFT}
	ROTATION2 = [4]int{FRONT, TOP, BACK, BOTTOM}
	ROTATION3 = [4]int{FRONT, LEFT, BACK, RIGHT}
)

var (
	rotationsMap = map[int][4]int{
		RAllTop:    ROTATION1,
		RAllBottom: ROTATION1,
		RFront:     ROTATION1,
		RFbMiddle:  ROTATION1,
		RBack:      ROTATION1,
		RAllFront:  ROTATION2,
		RAllBack:   ROTATION2,
		RLeft:      ROTATION2,
		RLrMiddle:  ROTATION2,
		RRight:     ROTATION2,
		RAllLeft:   ROTATION3,
		RAllRight:  ROTATION3,
		RTop:       ROTATION3,
		RTbMiddle:  ROTATION3,
		RBottom:    ROTATION3,
	}
)

const (
	cubeSideLength = 2
)

type Cubie struct {
	application  *Application
	localColors  [6]int //order: front, left, back, right, top, bottom
	globalColors [6]int //colors relative to the viewer
	faces        [6]*Face
	//vertices     [8]*rl.Vector3 //order: front face, back face, starting from the bottom left corner counterclockwise, see draw()
}

var (
	textureCoords            = [4]rl.Vector2{{0.0, 0.0}, {1.0, 0.0}, {1.0, 1.0}, {0.0, 1.0}}
	frontVertIndices         = [4]int{0, 1, 2, 3}
	backVertIndices          = [4]int{4, 5, 6, 7}
	topVertIndices           = [4]int{3, 2, 6, 7}
	bottomVertIndices        = [4]int{0, 1, 5, 4}
	leftVertIndices          = [4]int{0, 4, 7, 3}
	rightVertIndices         = [4]int{1, 5, 6, 2}
	cWidth, cHeight, cLength = float32(cubeSideLength), float32(cubeSideLength), float32(cubeSideLength)
	vecX                     = rl.NewVector3(1, 0, 0)
	vecY                     = rl.NewVector3(0, 1, 0)
	vecZ                     = rl.NewVector3(0, 0, 1)
)

var (
	rotationsToVectors = map[int]*rl.Vector3{
		RLeft:      &vecX,
		RLrMiddle:  &vecX,
		RRight:     &vecX,
		RTop:       &vecY,
		RTbMiddle:  &vecY,
		RBottom:    &vecY,
		RFront:     &vecZ,
		RBack:      &vecZ,
		RFbMiddle:  &vecZ,
		RAllLeft:   &vecY,
		RAllRight:  &vecY,
		RAllFront:  &vecX,
		RAllBack:   &vecX,
		RAllTop:    &vecZ,
		RAllBottom: &vecZ,
	}
)

func NewCubie(localColors [6]int, x, y, z int, a *Application) *Cubie {
	wX := float32(x) * cWidth
	hY := float32(y) * cHeight
	lZ := float32(z) * cLength
	v1 := rl.NewVector3(wX-cWidth/2, hY-cHeight/2, lZ+cLength/2)
	v2 := rl.NewVector3(wX+cWidth/2, hY-cHeight/2, lZ+cLength/2)
	v3 := rl.NewVector3(wX+cWidth/2, hY+cHeight/2, lZ+cLength/2)
	v4 := rl.NewVector3(wX-cWidth/2, hY+cHeight/2, lZ+cLength/2)
	v5 := rl.NewVector3(wX-cWidth/2, hY-cHeight/2, lZ-cLength/2)
	v6 := rl.NewVector3(wX+cWidth/2, hY-cHeight/2, lZ-cLength/2)
	v7 := rl.NewVector3(wX+cWidth/2, hY+cHeight/2, lZ-cLength/2)
	v8 := rl.NewVector3(wX-cWidth/2, hY+cHeight/2, lZ-cLength/2)
	faces := [6]*Face{
		NewFace([4]rl.Vector3{v1, v2, v3, v4}, localColors[FRONT]),
		NewFace([4]rl.Vector3{v5, v6, v7, v8}, localColors[BACK]),
		NewFace([4]rl.Vector3{v1, v2, v6, v5}, localColors[BOTTOM]),
		NewFace([4]rl.Vector3{v3, v4, v8, v7}, localColors[TOP]),
		NewFace([4]rl.Vector3{v1, v5, v8, v4}, localColors[LEFT]),
		NewFace([4]rl.Vector3{v2, v6, v7, v3}, localColors[RIGHT]),
	}
	//To create a copy of an array in Go, we can simply assign the array to another variable using the = operator (assignment),
	//and the contents will be copied over to the new array variable.
	globalColors := localColors
	cubie := &Cubie{
		application:  a,
		faces:        faces,
		localColors:  localColors,
		globalColors: globalColors,
	}
	return cubie
}

func (c *Cubie) shouldSelect(rotation int) bool {
	x, y, z := c.xyz()

	if rotation == RAllLeft || rotation == RAllRight ||
		rotation == RAllFront || rotation == RAllBack ||
		rotation == RAllTop || rotation == RAllBottom {
		return true
	}

	return (rotation == RLeft && math.Round(x) == -float64(cWidth)) ||
		(rotation == RBottom && math.Round(y) == -float64(cHeight)) ||
		(rotation == RBack && math.Round(z) == -float64(cLength)) ||
		(rotation == RLrMiddle && math.Round(x) == 0) ||
		(rotation == RTbMiddle && math.Round(y) == 0) ||
		(rotation == RFbMiddle && math.Round(z) == 0) ||
		(rotation == RRight && math.Round(x) == float64(cWidth)) ||
		rotation == RTop && math.Round(y) == float64(cHeight) ||
		(rotation == RFront && math.Round(z) == float64(cLength))
}

func (c *Cubie) isInFace(face int) bool {
	x, y, z := c.xyz()

	return (face == LEFT && math.Round(x) == -float64(cWidth)) ||
		(face == BOTTOM && math.Round(y) == -float64(cHeight)) ||
		(face == BACK && math.Round(z) == -float64(cLength)) ||
		(face == RIGHT && math.Round(x) == float64(cWidth)) ||
		face == TOP && math.Round(y) == float64(cHeight) ||
		(face == FRONT && math.Round(z) == float64(cLength))
}

func (c *Cubie) xyz() (float64, float64, float64) {
	x := float64(c.faces[0].vertices[0].X+c.faces[1].vertices[2].X) / 2
	y := float64(c.faces[0].vertices[0].Y+c.faces[1].vertices[2].Y) / 2
	z := float64(c.faces[0].vertices[0].Z+c.faces[1].vertices[2].Z) / 2
	return x, y, z
}

func (c *Cubie) update(selectedRotation int, isForward bool, rotationSpeed float32, angle float32) {
	delta := If(rotationSpeed > angle, angle, rotationSpeed)
	delta *= If(isForward, float32(1), float32(-1))
	vec := rotationsToVectors[selectedRotation]
	for i, face := range c.faces {
		for k, vertex := range face.vertices {
			res := rl.Vector3RotateByAxisAngle(vertex, *vec, rl.Deg2rad*delta) //and that's where the magic happens
			c.faces[i].vertices[k].X = res.X
			c.faces[i].vertices[k].Y = res.Y
			c.faces[i].vertices[k].Z = res.Z
		}
	}
}

func (c *Cubie) setColors(to [4]int, from [4]int) {
	for i := 0; i < len(to); i++ {
		c.globalColors[to[i]] = from[i]
	}
}

func (c *Cubie) getFacesByRotation(selectedRotation int) [4]int {
	if value, ok := rotationsMap[selectedRotation]; ok {
		return value
	}
	rl.TraceLog(rl.LogFatal, "Invalid rotation: %d", selectedRotation)
	return ROTATION1
}

func (c *Cubie) getColorsByFaces(faces [4]int, isForward bool) [4]int {
	if isForward {
		return [4]int{c.globalColors[faces[1]], c.globalColors[faces[2]], c.globalColors[faces[3]], c.globalColors[faces[0]]}
	}
	return [4]int{c.globalColors[faces[3]], c.globalColors[faces[0]], c.globalColors[faces[1]], c.globalColors[faces[2]]}
}

func (c *Cubie) updateGlobalColors(selectedRotation int, isForward bool) {
	faces := c.getFacesByRotation(selectedRotation)
	colors := c.getColorsByFaces(faces, isForward)
	c.setColors(faces, colors)
}

func (c *Cubie) draw(isSelected bool, scaleFactor float32) {
	x := (c.faces[0].vertices[0].X + c.faces[1].vertices[2].X) / scaleFactor
	y := (c.faces[0].vertices[0].Y + c.faces[1].vertices[2].Y) / scaleFactor
	z := (c.faces[0].vertices[0].Z + c.faces[1].vertices[2].Z) / scaleFactor

	rl.PushMatrix()
	rl.Translatef(x, y, z)
	rl.Begin(rl.Quads)
	{
		for _, face := range c.faces {
			face.draw(c, isSelected, textureCoords)
		}
	}
	rl.End()
	rl.PopMatrix()
}

func (c *Cubie) getFaceColor(face int) int {
	return c.globalColors[face]
}
