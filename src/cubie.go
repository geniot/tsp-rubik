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

const (
	cubeSideLength = 2
)

type Cubie struct {
	application  *Application
	localColors  [6]int         //order: front, left, back, right, top, bottom
	globalColors [6]int         //colors relative to the viewer
	vertices     [8]*rl.Vector3 //order: front face, back face, starting from the bottom left corner counterclockwise, see draw()
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
	vertices := [8]*rl.Vector3{&v1, &v2, &v3, &v4, &v5, &v6, &v7, &v8}
	//To create a copy of an array in Go, we can simply assign the array to another variable using the = operator (assignment),
	//and the contents will be copied over to the new array variable.
	globalColors := localColors
	cubie := &Cubie{
		application:  a,
		vertices:     vertices,
		localColors:  localColors,
		globalColors: globalColors,
	}
	return cubie
}

func (c *Cubie) shouldSelect(rotation int) bool {
	x := float64(c.vertices[0].X+c.vertices[6].X) / 2
	y := float64(c.vertices[0].Y+c.vertices[6].Y) / 2
	z := float64(c.vertices[0].Z+c.vertices[6].Z) / 2

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
	x := float64(c.vertices[0].X+c.vertices[6].X) / 2
	y := float64(c.vertices[0].Y+c.vertices[6].Y) / 2
	z := float64(c.vertices[0].Z+c.vertices[6].Z) / 2

	return (face == LEFT && math.Round(x) == -float64(cWidth)) ||
		(face == BOTTOM && math.Round(y) == -float64(cHeight)) ||
		(face == BACK && math.Round(z) == -float64(cLength)) ||
		(face == RIGHT && math.Round(x) == float64(cWidth)) ||
		face == TOP && math.Round(y) == float64(cHeight) ||
		(face == FRONT && math.Round(z) == float64(cLength))
}

func (c *Cubie) update(selectedRotation int, isForward bool, rotationSpeed float32, angle float32) {
	delta := If(rotationSpeed > angle, angle, rotationSpeed)
	delta *= If(isForward, float32(1), float32(-1))
	vec := rotationsToVectors[selectedRotation]
	for _, vertex := range c.vertices {
		res := rl.Vector3RotateByAxisAngle(*vertex, *vec, rl.Deg2rad*delta)
		vertex.X = res.X
		vertex.Y = res.Y
		vertex.Z = res.Z
	}
}

func (c *Cubie) setColors(to [4]int, from [4]int) {
	for i := 0; i < len(to); i++ {
		c.globalColors[to[i]] = from[i]
	}
}

func (c *Cubie) updateGlobalColors(selectedRotation int, isForward bool) {

	frontColor := c.globalColors[FRONT]
	leftColor := c.globalColors[LEFT]
	backColor := c.globalColors[BACK]
	rightColor := c.globalColors[RIGHT]
	topColor := c.globalColors[TOP]
	bottomColor := c.globalColors[BOTTOM]

	if selectedRotation == RAllTop || selectedRotation == RAllBottom || selectedRotation == RFront || selectedRotation == RFbMiddle || selectedRotation == RBack {
		if isForward {
			c.setColors([4]int{TOP, RIGHT, BOTTOM, LEFT}, [4]int{rightColor, bottomColor, leftColor, topColor})
		} else {
			c.setColors([4]int{TOP, RIGHT, BOTTOM, LEFT}, [4]int{leftColor, topColor, rightColor, bottomColor})
		}
		return
	}
	if selectedRotation == RAllFront || selectedRotation == RAllBack || selectedRotation == RLeft || selectedRotation == RLrMiddle || selectedRotation == RRight {
		if isForward {
			c.setColors([4]int{FRONT, TOP, BACK, BOTTOM}, [4]int{bottomColor, frontColor, topColor, backColor})
		} else {
			c.setColors([4]int{FRONT, TOP, BACK, BOTTOM}, [4]int{topColor, backColor, bottomColor, frontColor})
		}
		return
	}
	if selectedRotation == RAllLeft || selectedRotation == RAllRight || selectedRotation == RTop || selectedRotation == RTbMiddle || selectedRotation == RBottom {
		if isForward {
			c.setColors([4]int{FRONT, LEFT, BACK, RIGHT}, [4]int{leftColor, backColor, rightColor, frontColor})
		} else {
			c.setColors([4]int{FRONT, LEFT, BACK, RIGHT}, [4]int{rightColor, frontColor, leftColor, backColor})
		}
		return
	}
}

func (c *Cubie) draw(isSelected bool, scaleFactor float32) {
	x := (c.vertices[0].X + c.vertices[6].X) / scaleFactor
	y := (c.vertices[0].Y + c.vertices[6].Y) / scaleFactor
	z := (c.vertices[0].Z + c.vertices[6].Z) / scaleFactor

	rl.PushMatrix()
	rl.Translatef(x, y, z)
	rl.Begin(rl.Quads)
	{
		c.drawFace(FRONT, isSelected, &frontVertIndices)
		c.drawFace(BACK, isSelected, &backVertIndices)
		c.drawFace(TOP, isSelected, &topVertIndices)
		c.drawFace(BOTTOM, isSelected, &bottomVertIndices)
		c.drawFace(LEFT, isSelected, &leftVertIndices)
		c.drawFace(RIGHT, isSelected, &rightVertIndices)
	}
	rl.End()
	rl.PopMatrix()
}

func (c *Cubie) drawFace(face int, isSelected bool, indices *[4]int) {
	textures := If(isSelected, c.application.selectedColorTextures, c.application.colorTextures)
	rl.SetTexture(textures[c.localColors[face]].ID)
	for i := 0; i < len(indices); i++ {
		rl.TexCoord2f(textureCoords[i].X, textureCoords[i].Y)
		rl.Vertex3f(c.vertices[indices[i]].X, c.vertices[indices[i]].Y, c.vertices[indices[i]].Z)
	}
}
