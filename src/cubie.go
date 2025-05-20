package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	textureCoords     = [4]rl.Vector2{{0.0, 0.0}, {1.0, 0.0}, {1.0, 1.0}, {0.0, 1.0}}
	frontVertIndices  = [4]int{0, 1, 2, 3}
	backVertIndices   = [4]int{4, 5, 6, 7}
	topVertIndices    = [4]int{3, 2, 6, 7}
	bottomVertIndices = [4]int{0, 1, 5, 4}
	leftVertIndices   = [4]int{0, 4, 7, 3}
	rightVertIndices  = [4]int{1, 5, 6, 2}
)

type Cubie struct {
	colors     [6]int        //order: front, left, back, right, top, bottom
	vertices   [8]rl.Vector3 //order: front face, back face, starting from the bottom left corner counterclockwise, see draw()
	isSelected bool
}

func NewCubie(colors [6]int, x, y, z int) *Cubie {
	cubie := &Cubie{}
	return cubie
}

func (c *Cubie) isRotating() bool {
	return false
}

func (c *Cubie) shouldSelect(rotation int) bool {
	return false
}

func (c *Cubie) update() bool {
	return true
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
