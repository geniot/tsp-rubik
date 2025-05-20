package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"math"
)

var (
	textureCoords            = [4]rl.Vector2{{0.0, 0.0}, {1.0, 0.0}, {1.0, 1.0}, {0.0, 1.0}}
	frontVertIndices         = [4]int{0, 1, 2, 3}
	backVertIndices          = [4]int{4, 5, 6, 7}
	topVertIndices           = [4]int{3, 2, 6, 7}
	bottomVertIndices        = [4]int{0, 1, 5, 4}
	leftVertIndices          = [4]int{0, 4, 7, 3}
	rightVertIndices         = [4]int{1, 5, 6, 2}
	cWidth, cHeight, cLength = float32(2), float32(2), float32(2)
)

type Cubie struct {
	colors     [6]int        //order: front, left, back, right, top, bottom
	vertices   [8]rl.Vector3 //order: front face, back face, starting from the bottom left corner counterclockwise, see draw()
	isSelected bool
}

func NewCubie(colors [6]int, x, y, z int) *Cubie {
	wX := float32(x) * cWidth
	hY := float32(y) * cHeight
	lZ := float32(z) * cLength
	vertices := [8]rl.Vector3{
		rl.NewVector3(wX-cWidth/2, hY-cHeight/2, lZ+cLength/2),
		rl.NewVector3(wX+cWidth/2, hY-cHeight/2, lZ+cLength/2),
		rl.NewVector3(wX+cWidth/2, hY+cHeight/2, lZ+cLength/2),
		rl.NewVector3(wX-cWidth/2, hY+cHeight/2, lZ+cLength/2),
		rl.NewVector3(wX-cWidth/2, hY-cHeight/2, lZ-cLength/2),
		rl.NewVector3(wX+cWidth/2, hY-cHeight/2, lZ-cLength/2),
		rl.NewVector3(wX+cWidth/2, hY+cHeight/2, lZ-cLength/2),
		rl.NewVector3(wX-cWidth/2, hY+cHeight/2, lZ-cLength/2),
	}
	cubie := &Cubie{vertices: vertices, colors: colors, isSelected: false}
	return cubie
}

func (c *Cubie) isRotating() bool {
	return false
}

func (c *Cubie) shouldSelect(rotation int) bool {
	x := c.vertices[0].X + cWidth/2
	y := c.vertices[0].Y + cHeight/2
	z := c.vertices[0].Z - cLength/2
	return (rotation == R_LEFT && math.Round(float64(x)) == -float64(cWidth)) ||
		(rotation == R_BOTTOM && math.Round(float64(y)) == -float64(cHeight)) ||
		(rotation == R_BACK && math.Round(float64(z)) == -float64(cLength)) ||
		(rotation == R_LR_MIDDLE && math.Round(float64(x)) == 0) ||
		(rotation == R_TB_MIDDLE && math.Round(float64(y)) == 0) ||
		(rotation == R_FB_MIDDLE && math.Round(float64(z)) == 0) ||
		(rotation == R_RIGHT && math.Round(float64(x)) == float64(cWidth)) ||
		rotation == R_TOP && math.Round(float64(y)) == float64(cHeight) ||
		(rotation == R_FRONT && math.Round(float64(z)) == float64(cLength))
}

func (c *Cubie) update(selectedRotation int) {
	c.isSelected = If(c.shouldSelect(selectedRotation), true, false)
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
