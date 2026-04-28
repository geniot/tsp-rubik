package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Face struct {
	vertices [4]rl.Vector3
	color    int
}

func (f *Face) draw(c *Cubie, isSelected bool, textureCoords [4]rl.Vector2) {
	textures := If(isSelected, c.application.selectedColorTextures, c.application.colorTextures)
	rl.SetTexture(textures[f.color].ID)
	for i, vertex := range f.vertices {
		rl.TexCoord2f(textureCoords[i].X, textureCoords[i].Y)
		rl.Vertex3f(vertex.X, vertex.Y, vertex.Z)
	}
}

func NewFace(v [4]rl.Vector3, c int) *Face {
	return &Face{vertices: v, color: c}
}
