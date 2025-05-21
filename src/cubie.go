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

var (
	textureCoords = [4]rl.Vector2{{0.0, 0.0}, {1.0, 0.0}, {1.0, 1.0}, {0.0, 1.0}}
	//
	frontVertIndices  = [4]int{0, 1, 2, 3}
	backVertIndices   = [4]int{4, 5, 6, 7}
	topVertIndices    = [4]int{3, 2, 6, 7}
	bottomVertIndices = [4]int{0, 1, 5, 4}
	leftVertIndices   = [4]int{0, 4, 7, 3}
	rightVertIndices  = [4]int{1, 5, 6, 2}
	//used for quick comparison, ideally should be initialized from the values above! front and back are already sorted
	topVertIndicesSorted    = [4]int{2, 3, 6, 7}
	bottomVertIndicesSorted = [4]int{0, 1, 4, 5}
	leftVertIndicesSorted   = [4]int{0, 3, 4, 7}
	rightVertIndicesSorted  = [4]int{1, 2, 5, 6}

	cWidth, cHeight, cLength = float32(cubeSideLength), float32(cubeSideLength), float32(cubeSideLength)
	vecX                     = rl.NewVector3(1, 0, 0)
	vecY                     = rl.NewVector3(0, 1, 0)
	vecZ                     = rl.NewVector3(0, 0, 1)
)

var (
	faceToIndices = map[int][4]int{
		FRONT:  frontVertIndices,
		BACK:   backVertIndices,
		TOP:    topVertIndicesSorted,
		BOTTOM: bottomVertIndicesSorted,
		LEFT:   leftVertIndicesSorted,
		RIGHT:  rightVertIndicesSorted,
	}
)

var (
	rotationsToVectors = map[int]*rl.Vector3{
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

func (c *Cubie) update(selectedRotation int, isRotating bool, isForward bool) {
	c.isSelected = If(c.shouldSelect(selectedRotation), true, false)
	if c.isSelected && isRotating {
		angleDelta := If(isForward, rotationSpeed, -rotationSpeed)
		vec := rotationsToVectors[selectedRotation]
		for _, vertex := range c.vertices {
			res := rl.Vector3RotateByAxisAngle(*vertex, *vec, rl.Deg2rad*angleDelta)
			vertex.X = res.X
			vertex.Y = res.Y
			vertex.Z = res.Z
		}
	}
}

func (c *Cubie) draw(scaleFactor float32) {
	x := (c.vertices[0].X + c.vertices[6].X) / scaleFactor
	y := (c.vertices[0].Y + c.vertices[6].Y) / scaleFactor
	z := (c.vertices[0].Z + c.vertices[6].Z) / scaleFactor

	rl.PushMatrix()
	rl.Translatef(x, y, z)
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
	rl.PopMatrix()
}

func (c *Cubie) drawFace(face int, indices *[4]int) {
	textures := If(c.isSelected, selectedColorTextures, colorTextures)
	rl.SetTexture(textures[c.colors[face]].ID)
	for i := 0; i < len(indices); i++ {
		rl.TexCoord2f(textureCoords[i].X, textureCoords[i].Y)
		rl.Vertex3f(c.vertices[indices[i]].X, c.vertices[indices[i]].Y, c.vertices[indices[i]].Z)
	}
}

var (
	faceToMinMax = map[int]float32{
		FRONT:  float32(-math.MaxFloat32),
		BACK:   float32(math.MaxFloat32),
		RIGHT:  float32(-math.MaxFloat32),
		LEFT:   float32(math.MaxFloat32),
		TOP:    float32(-math.MaxFloat32),
		BOTTOM: float32(math.MaxFloat32),
	}
)

func (c *Cubie) getFacePoint(face int) int {
	//use face to define filter min/max x/y/z
	point := faceToMinMax[face]
	for _, vertex := range c.vertices {
		if face == LEFT {
			point = float32(math.Min(float64(point), float64(vertex.X)))
		}
		if face == RIGHT {
			point = float32(math.Max(float64(point), float64(vertex.X)))
		}
		if face == TOP {
			point = float32(math.Max(float64(point), float64(vertex.Y)))
		}
		if face == BOTTOM {
			point = float32(math.Min(float64(point), float64(vertex.Y)))
		}
		if face == FRONT {
			point = float32(math.Max(float64(point), float64(vertex.Z)))
		}
		if face == BACK {
			point = float32(math.Min(float64(point), float64(vertex.Z)))
		}
	}
	return int(math.Round(float64(point)))
}

func (c *Cubie) getFaceIndices(face int, point int) [4]int {
	faceIndices := [4]int{0, 0, 0, 0}
	counter := 0
	for index, vertex := range c.vertices {
		if ((face == LEFT || face == RIGHT) && int(math.Round(float64(vertex.X))) == point) ||
			((face == TOP || face == BOTTOM) && int(math.Round(float64(vertex.Y))) == point) ||
			((face == FRONT || face == BACK) && int(math.Round(float64(vertex.Z))) == point) {
			faceIndices[counter] = index
			counter++
		}
	}
	return faceIndices
}

func (c *Cubie) getFaceColor(face int) int {
	point := c.getFacePoint(face)
	faceIndices := c.getFaceIndices(face, point)
	actualFace := FRONT
	for key, value := range faceToIndices {
		if arraysEqual(faceIndices, value) {
			actualFace = key
			break
		}
	}
	return c.colors[actualFace]
}

func arraysEqual(a, b [4]int) bool {
	for i := 0; i < 4; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
