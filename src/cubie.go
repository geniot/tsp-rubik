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
	application *Application
	faces       [6]*Face
}

var (
	textureCoords            = [4]rl.Vector2{{0.0, 0.0}, {1.0, 0.0}, {1.0, 1.0}, {0.0, 1.0}}
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

func NewCubie(colors [6]int, x, y, z int, a *Application) *Cubie {
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
		NewFace([4]rl.Vector3{v1, v2, v3, v4}, colors[FRONT]),
		NewFace([4]rl.Vector3{v5, v6, v7, v8}, colors[BACK]),
		NewFace([4]rl.Vector3{v1, v2, v6, v5}, colors[BOTTOM]),
		NewFace([4]rl.Vector3{v3, v4, v8, v7}, colors[TOP]),
		NewFace([4]rl.Vector3{v1, v5, v8, v4}, colors[LEFT]),
		NewFace([4]rl.Vector3{v2, v6, v7, v3}, colors[RIGHT]),
	}
	cubie := &Cubie{
		application: a,
		faces:       faces,
	}
	return cubie
}

func (c *Cubie) shouldSelect(rotation int) bool {
	center := c.getCenter(6.0)
	x, y, z := int(math.Round(float64(center.X))), int(math.Round(float64(center.Y))), int(math.Round(float64(center.Z)))

	if rotation == RAllLeft || rotation == RAllRight ||
		rotation == RAllFront || rotation == RAllBack ||
		rotation == RAllTop || rotation == RAllBottom {
		return true
	}

	return (rotation == RLeft && x == -int(cWidth)) ||
		(rotation == RBottom && y == -int(cHeight)) ||
		(rotation == RBack && z == -int(cLength)) ||
		(rotation == RLrMiddle && x == 0) ||
		(rotation == RTbMiddle && y == 0) ||
		(rotation == RFbMiddle && z == 0) ||
		(rotation == RRight && x == int(cWidth)) ||
		rotation == RTop && y == int(cHeight) ||
		(rotation == RFront && z == int(cLength))
}

func (c *Cubie) isInFace(face int) bool {
	center := c.getCenter(6.0)
	x, y, z := int(math.Round(float64(center.X))), int(math.Round(float64(center.Y))), int(math.Round(float64(center.Z)))

	return (face == LEFT && x == -int(cWidth)) ||
		(face == RIGHT && x == int(cWidth)) ||
		(face == BOTTOM && y == -int(cHeight)) ||
		face == TOP && y == int(cHeight) ||
		(face == BACK && z == -int(cLength)) ||
		(face == FRONT && z == int(cLength))
}

func (c *Cubie) getCenter(scaleFactor float32) rl.Vector3 {
	vec := rl.Vector3{}
	for _, face := range c.faces {
		faceCenter := face.getCenter()
		vec = rl.Vector3Add(vec, faceCenter)
	}
	return rl.Vector3Scale(vec, 1.0/scaleFactor)
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

func (c *Cubie) getFacesByRotation(selectedRotation int) [4]int {
	if value, ok := rotationsMap[selectedRotation]; ok {
		return value
	}
	rl.TraceLog(rl.LogFatal, "Invalid rotation: %d", selectedRotation)
	return ROTATION1
}

func (c *Cubie) draw(isSelected bool, scaleFactor float32) {
	center := c.getCenter(scaleFactor * 2.5)
	x, y, z := center.X, center.Y, center.Z

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

func (c *Cubie) getFaceColor(side int) int {
	for _, face := range c.faces {
		if face.getSide() == side {
			return face.color
		}
	}
	panic("face not found")
}
