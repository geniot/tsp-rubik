package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"math"
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
		R_LEFT:       &vecX,
		R_LR_MIDDLE:  &vecX,
		R_RIGHT:      &vecX,
		R_TOP:        &vecY,
		R_TB_MIDDLE:  &vecY,
		R_BOTTOM:     &vecY,
		R_FRONT:      &vecZ,
		R_BACK:       &vecZ,
		R_FB_MIDDLE:  &vecZ,
		R_ALL_LEFT:   &vecY,
		R_ALL_RIGHT:  &vecY,
		R_ALL_FRONT:  &vecX,
		R_ALL_BACK:   &vecX,
		R_ALL_TOP:    &vecZ,
		R_ALL_BOTTOM: &vecZ,
	}
)

type Cubie struct {
	localColors  [6]int         //order: front, left, back, right, top, bottom
	globalColors [6]int         //colors relative to the viewer
	vertices     [8]*rl.Vector3 //order: front face, back face, starting from the bottom left corner counterclockwise, see draw()
}

func NewCubie(localColors [6]int, x, y, z int) *Cubie {
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
	cubie := &Cubie{vertices: vertices, localColors: localColors, globalColors: globalColors}
	return cubie
}

func (c *Cubie) shouldSelect(rotation int) bool {
	x := float64(c.vertices[0].X+c.vertices[6].X) / 2
	y := float64(c.vertices[0].Y+c.vertices[6].Y) / 2
	z := float64(c.vertices[0].Z+c.vertices[6].Z) / 2

	if rotation == R_ALL_LEFT || rotation == R_ALL_RIGHT ||
		rotation == R_ALL_FRONT || rotation == R_ALL_BACK ||
		rotation == R_ALL_TOP || rotation == R_ALL_BOTTOM {
		return true
	}

	return (rotation == R_LEFT && math.Round(x) == -float64(cWidth)) ||
		(rotation == R_BOTTOM && math.Round(y) == -float64(cHeight)) ||
		(rotation == R_BACK && math.Round(z) == -float64(cLength)) ||
		(rotation == R_LR_MIDDLE && math.Round(x) == 0) ||
		(rotation == R_TB_MIDDLE && math.Round(y) == 0) ||
		(rotation == R_FB_MIDDLE && math.Round(z) == 0) ||
		(rotation == R_RIGHT && math.Round(x) == float64(cWidth)) ||
		rotation == R_TOP && math.Round(y) == float64(cHeight) ||
		(rotation == R_FRONT && math.Round(z) == float64(cLength))
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

func (c *Cubie) updateGlobalColors(selectedRotation int, isForward bool) {

	frontColor := c.globalColors[FRONT]
	leftColor := c.globalColors[LEFT]
	backColor := c.globalColors[BACK]
	rightColor := c.globalColors[RIGHT]
	topColor := c.globalColors[TOP]
	bottomColor := c.globalColors[BOTTOM]

	if selectedRotation == R_ALL_TOP || selectedRotation == R_ALL_BOTTOM || selectedRotation == R_FRONT || selectedRotation == R_FB_MIDDLE || selectedRotation == R_BACK {
		if isForward {
			c.globalColors[TOP] = rightColor
			c.globalColors[RIGHT] = bottomColor
			c.globalColors[BOTTOM] = leftColor
			c.globalColors[LEFT] = topColor
		} else {
			c.globalColors[TOP] = leftColor
			c.globalColors[RIGHT] = topColor
			c.globalColors[BOTTOM] = rightColor
			c.globalColors[LEFT] = bottomColor
		}
		return
	}
	if selectedRotation == R_ALL_FRONT || selectedRotation == R_ALL_BACK || selectedRotation == R_LEFT || selectedRotation == R_LR_MIDDLE || selectedRotation == R_RIGHT {
		if isForward {
			c.globalColors[FRONT] = bottomColor
			c.globalColors[TOP] = frontColor
			c.globalColors[BACK] = topColor
			c.globalColors[BOTTOM] = backColor
		} else {
			c.globalColors[FRONT] = topColor
			c.globalColors[TOP] = backColor
			c.globalColors[BACK] = bottomColor
			c.globalColors[BOTTOM] = frontColor
		}
		return
	}
	if selectedRotation == R_ALL_LEFT || selectedRotation == R_ALL_RIGHT || selectedRotation == R_TOP || selectedRotation == R_TB_MIDDLE || selectedRotation == R_BOTTOM {
		if isForward {
			c.globalColors[FRONT] = leftColor
			c.globalColors[LEFT] = backColor
			c.globalColors[BACK] = rightColor
			c.globalColors[RIGHT] = frontColor
		} else {
			c.globalColors[FRONT] = rightColor
			c.globalColors[LEFT] = frontColor
			c.globalColors[BACK] = leftColor
			c.globalColors[RIGHT] = backColor
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
	textures := If(isSelected, selectedColorTextures, colorTextures)
	rl.SetTexture(textures[c.localColors[face]].ID)
	for i := 0; i < len(indices); i++ {
		rl.TexCoord2f(textureCoords[i].X, textureCoords[i].Y)
		rl.Vertex3f(c.vertices[indices[i]].X, c.vertices[indices[i]].Y, c.vertices[indices[i]].Z)
	}
}
