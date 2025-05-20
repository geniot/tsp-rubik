package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"math"
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
	rotsToVectors = map[int]*rl.Vector3{
		R_LEFT:      &vecX,
		R_LR_MIDDLE: &vecX,
		R_RIGHT:     &vecX,
		R_TOP:       &vecY,
		R_TB_MIDDLE: &vecY,
		R_BOTTOM:    &vecY,
		R_FRONT:     &vecZ,
		R_BACK:      &vecZ,
		R_FB_MIDDLE: &vecZ,
	}
)

type Cubie struct {
	colors     [6]int         //order: front, left, back, right, top, bottom
	vertices   [8]*rl.Vector3 //order: front face, back face, starting from the bottom left corner counterclockwise, see draw()
	isSelected bool
}

func NewCubie(colors [6]int, x, y, z int) *Cubie {
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
	cubie := &Cubie{vertices: vertices, colors: colors, isSelected: false}
	return cubie
}

func (c *Cubie) shouldSelect(rotation int) bool {
	x := float64(c.vertices[0].X+c.vertices[6].X) / 2
	y := float64(c.vertices[0].Y+c.vertices[6].Y) / 2
	z := float64(c.vertices[0].Z+c.vertices[6].Z) / 2
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

func (c *Cubie) update(rotation *Rotation) {
	c.isSelected = If(c.shouldSelect(rotation.selectedRotation), true, false)
	if c.isSelected && rotation.isRotating() {
		angleDelta := If(rotation.isForward, rotationSpeed, -rotationSpeed)
		vec := rotsToVectors[rotation.selectedRotation]
		for _, vertex := range c.vertices {
			res := rl.Vector3RotateByAxisAngle(*vertex, *vec, rl.Deg2rad*angleDelta)
			vertex.X = res.X
			vertex.Y = res.Y
			vertex.Z = res.Z
		}
	}
}

func (c *Cubie) draw() {
	rl.Begin(rl.Quads)
	{
		c.drawFace(FRONT, &frontVertIndices)
		c.drawFace(BACK, &backVertIndices)
		c.drawFace(TOP, &topVertIndices)
		c.drawFace(BOTTOM, &bottomVertIndices)
		c.drawFace(LEFT, &leftVertIndices)
		c.drawFace(RIGHT, &rightVertIndices)
	}
	rl.End()
}

func (c *Cubie) drawFace(face int, indices *[4]int) {
	textures := If(c.isSelected, selectedColorTextures, colorTextures)
	rl.SetTexture(textures[c.colors[face]].ID)
	for i := 0; i < 4; i++ {
		rl.TexCoord2f(textureCoords[i].X, textureCoords[i].Y)
		rl.Vertex3f(c.vertices[indices[i]].X, c.vertices[indices[i]].Y, c.vertices[indices[i]].Z)
	}
}
