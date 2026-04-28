package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

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

func (f *Face) getSide() int {
	vect := rl.NewVector3(0, 0, 0)
	for _, vertex := range f.vertices {
		vect = rl.Vector3Add(vect, vertex)
	}
	isX := round32(f.vertices[0].X) == round32(f.vertices[1].X) &&
		round32(f.vertices[0].X) == round32(f.vertices[2].X) &&
		round32(f.vertices[0].X) == round32(f.vertices[3].X)
	isY := round32(f.vertices[0].Y) == round32(f.vertices[1].Y) &&
		round32(f.vertices[0].Y) == round32(f.vertices[2].Y) &&
		round32(f.vertices[0].Y) == round32(f.vertices[3].Y)
	isZ := round32(f.vertices[0].Z) == round32(f.vertices[1].Z) &&
		round32(f.vertices[0].Z) == round32(f.vertices[2].Z) &&
		round32(f.vertices[0].Z) == round32(f.vertices[3].Z)

	if isZ {
		return IfInt(vect.Z > 0, FRONT, BACK)
	} else if isX {
		return IfInt(vect.X > 0, RIGHT, LEFT)
	} else if isY {
		return IfInt(vect.Y > 0, TOP, BOTTOM)
	}
	panic("unknown side")
}

func NewFace(v [4]rl.Vector3, c int) *Face {
	return &Face{vertices: v, color: c}
}
