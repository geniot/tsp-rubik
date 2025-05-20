package main

import (
	"bufio"
	"bytes"
	"github.com/fogleman/gg"
	rl "github.com/gen2brain/raylib-go/raylib"
	"math"
	"unsafe"
)

func prepareTextures() {
	var (
		width  = 100
		height = 100
		//padding = 5
	)
	//for colorKey, color := range allColors {
	//	pngBytes := makePng(width, height, padding, black, color, false)
	//	colorTextures[colorKey] = rl.LoadTextureFromImage(rl.LoadImageFromMemory(".png", pngBytes, int32(len(pngBytes))))
	//}
	//for colorKey, color := range allColors {
	//	pngBytes := makePng(width, height, padding, black, color, true)
	//	selectedColorTextures[colorKey] = rl.LoadTextureFromImage(rl.LoadImageFromMemory(".png", pngBytes, int32(len(pngBytes))))
	//}

	//combined texture
	bytesBuffer := new(bytes.Buffer)
	dc := gg.NewContext(width*6, height)
	colors := [6]rl.Color{red, green, yellow, blue, orange, lightBlack}
	for i, color := range colors {
		dc.DrawRectangle(float64(i*width), 0, float64(width), float64(height))
		dc.SetRGBA255(int(color.R), int(color.G), int(color.B), int(color.A))
		dc.Fill()
	}
	w := bufio.NewWriter(bytesBuffer)
	//dc.SavePNG("out.png")
	orPanic(dc.EncodePNG(w))
	orPanic(w.Flush())
	//pngBytes := bytesBuffer.Bytes()
	//combinedTexture = rl.LoadTextureFromImage(rl.LoadImageFromMemory(".png", pngBytes, int32(len(pngBytes))))
}

func genTextureFromColors(colors [6]int, paddingColor int, isSelected bool) *rl.Texture2D {
	var (
		width               = 100
		height              = 100
		padding             = 5
		selectionColor      = white
		subSelectionColor   = black
		subSelectionPadding = 2
	)
	bytesBuffer := new(bytes.Buffer)
	dc := gg.NewContext(width*6, height)
	for i, color := range colors {
		dc.DrawRectangle(float64(i*width), 0, float64(width), float64(height))
		dc.SetRGBA255(int(allColors[paddingColor].R), int(allColors[paddingColor].G), int(allColors[paddingColor].B), int(allColors[paddingColor].A))
		dc.Fill()
		if isSelected {
			dc.DrawRectangle(float64(i*width)+float64(padding), float64(padding), float64(width-padding*2), float64(height-padding*2))
			dc.SetRGBA255(int(selectionColor.R), int(selectionColor.G), int(selectionColor.B), int(selectionColor.A))
			dc.Fill()
			dc.DrawRectangle(float64(i*width)+float64(padding*2-subSelectionPadding),
				float64(padding*2-subSelectionPadding),
				float64(width-padding*4+subSelectionPadding*2),
				float64(height-padding*4+subSelectionPadding*2))
			dc.SetRGBA255(int(subSelectionColor.R), int(subSelectionColor.G), int(subSelectionColor.B), int(subSelectionColor.A))
			dc.Fill()
			dc.DrawRectangle(float64(i*width+padding*2), float64(padding*2), float64(width-padding*4), float64(height-padding*4))
			dc.SetRGBA255(int(allColors[color].R), int(allColors[color].G), int(allColors[color].B), int(allColors[color].A))
			dc.Fill()
		} else {
			dc.DrawRectangle(float64(i*width+padding), float64(padding), float64(width-padding*2), float64(height-padding*2))
			dc.SetRGBA255(int(allColors[color].R), int(allColors[color].G), int(allColors[color].B), int(allColors[color].A))
			dc.Fill()
		}

	}
	w := bufio.NewWriter(bytesBuffer)
	//dc.SavePNG("out.png")
	orPanic(dc.EncodePNG(w))
	orPanic(w.Flush())
	pngBytes := bytesBuffer.Bytes()
	texture := rl.LoadTextureFromImage(rl.LoadImageFromMemory(".png", pngBytes, int32(len(pngBytes))))
	return &texture
}

func makePng(width int, height int, padding int, paddingColor rl.Color, color rl.Color, isSelected bool) []byte {
	selectionColor := white
	subSelectionColor := black
	subSelectionPadding := 2
	bytesBuffer := new(bytes.Buffer)
	dc := gg.NewContext(width, height)
	dc.DrawRectangle(0, 0, float64(width), float64(height))
	dc.SetRGBA255(int(paddingColor.R), int(paddingColor.G), int(paddingColor.B), int(paddingColor.A))
	dc.Fill()
	if isSelected {
		dc.DrawRectangle(float64(padding), float64(padding), float64(width-padding*2), float64(height-padding*2))
		dc.SetRGBA255(int(selectionColor.R), int(selectionColor.G), int(selectionColor.B), int(selectionColor.A))
		dc.Fill()
		dc.DrawRectangle(float64(padding*2-subSelectionPadding),
			float64(padding*2-subSelectionPadding),
			float64(width-padding*4+subSelectionPadding*2),
			float64(height-padding*4+subSelectionPadding*2))
		dc.SetRGBA255(int(subSelectionColor.R), int(subSelectionColor.G), int(subSelectionColor.B), int(subSelectionColor.A))
		dc.Fill()
		dc.DrawRectangle(float64(padding*2), float64(padding*2), float64(width-padding*4), float64(height-padding*4))
		dc.SetRGBA255(int(color.R), int(color.G), int(color.B), int(color.A))
		dc.Fill()
	} else {
		dc.DrawRectangle(float64(padding), float64(padding), float64(width-padding*2), float64(height-padding*2))
		dc.SetRGBA255(int(color.R), int(color.G), int(color.B), int(color.A))
		dc.Fill()
	}
	w := bufio.NewWriter(bytesBuffer)
	//dc.SavePNG("out.png")
	orPanic(dc.EncodePNG(w))
	orPanic(w.Flush())
	return bytesBuffer.Bytes()
}

func orPanic(err interface{}) {
	switch v := err.(type) {
	case error:
		if v != nil {
			panic(err)
		}
	case bool:
		if !v {
			panic("condition failed: != true")
		}
	}
}

func orPanicRes[T any](res T, err interface{}) T {
	orPanic(err)
	return res
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}

func If[T any](cond bool, vTrue, vFalse T) T {
	if cond {
		return vTrue
	}
	return vFalse
}

func newInts(n, v int) []int {
	s := make([]int, n)
	for i := range s {
		s[i] = v
	}
	return s
}

func copyInts(fromArray []int, toArray []int, fromOffset int, toOffset int, length int) {
	for i := 0; i < length; i++ {
		toArray[toOffset+i] = fromArray[fromOffset+i]
	}
}

// GenMeshCustom generates a simple cube mesh from code
func GenMeshCustom() *rl.Mesh {
	mesh := rl.Mesh{}
	var vertices, normals, texCoords []float32

	w := float32(2)
	h := float32(2)
	l := float32(2)
	x := -(w / 2)
	y := -(h / 2)
	z := -(l / 2)

	//to make a square, we always draw 2 triangles, starting from 0,0 (bottom left)
	//the first triangle is drawn up,right,down diagonally
	//the second triangle is drawn right, up, down diagonally
	//this important because texture mapping should be done with the same triangles and vectors

	//front,left,back,right,top,bottom - is the order colors in the configuration

	//front triangles z+l
	vertices = addCoordinate(vertices, x, y, z+l)
	vertices = addCoordinate(vertices, x, y+h, z+l)
	vertices = addCoordinate(vertices, x+w, y+h, z+l)
	vertices = addCoordinate(vertices, x, y, z+l)
	vertices = addCoordinate(vertices, x+w, y, z+l)
	vertices = addCoordinate(vertices, x+w, y+h, z+l)

	//left triangles x
	vertices = addCoordinate(vertices, x, y, z)
	vertices = addCoordinate(vertices, x, y, z+l)
	vertices = addCoordinate(vertices, x, y+h, z+l)
	vertices = addCoordinate(vertices, x, y, z)
	vertices = addCoordinate(vertices, x, y+h, z)
	vertices = addCoordinate(vertices, x, y+h, z+l)

	//back triangles z
	vertices = addCoordinate(vertices, x, y, z)
	vertices = addCoordinate(vertices, x, y+h, z)
	vertices = addCoordinate(vertices, x+w, y+h, z)
	vertices = addCoordinate(vertices, x, y, z)
	vertices = addCoordinate(vertices, x+w, y, z)
	vertices = addCoordinate(vertices, x+w, y+h, z)

	//right triangles x+w
	vertices = addCoordinate(vertices, x+w, y, z)
	vertices = addCoordinate(vertices, x+w, y, z+l)
	vertices = addCoordinate(vertices, x+w, y+h, z+l)
	vertices = addCoordinate(vertices, x+w, y, z)
	vertices = addCoordinate(vertices, x+w, y+h, z)
	vertices = addCoordinate(vertices, x+w, y+h, z+l)

	//top triangles y+h
	vertices = addCoordinate(vertices, x, y+h, z)
	vertices = addCoordinate(vertices, x, y+h, z+l)
	vertices = addCoordinate(vertices, x+w, y+h, z+l)
	vertices = addCoordinate(vertices, x, y+h, z)
	vertices = addCoordinate(vertices, x+w, y+h, z)
	vertices = addCoordinate(vertices, x+w, y+h, z+l)

	//bottom triangles y
	vertices = addCoordinate(vertices, x, y, z)
	vertices = addCoordinate(vertices, x, y, z+l)
	vertices = addCoordinate(vertices, x+w, y, z+l)
	vertices = addCoordinate(vertices, x, y, z)
	vertices = addCoordinate(vertices, x+w, y, z)
	vertices = addCoordinate(vertices, x+w, y, z+l)

	mesh.Vertices = unsafe.SliceData(vertices)
	mesh.VertexCount = int32(len(vertices) / 3)
	mesh.TriangleCount = int32(len(vertices) / 3 / 2)

	//https://en.wikipedia.org/wiki/Vertex_normal
	//directional vectors used for shading
	//front,left,back,right,top,bottom - is the order colors in the configuration
	normals = addCoordinate(normals, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0, 0, 1)       //front
	normals = addCoordinate(normals, -1, 0, 0, -1, 0, 0, -1, 0, 0, -1, 0, 0, -1, 0, 0, -1, 0, 0) //left
	normals = addCoordinate(normals, 0, 0, -1, 0, 0, -1, 0, 0, -1, 0, 0, -1, 0, 0, -1, 0, 0, -1) //back
	normals = addCoordinate(normals, 1, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0, 0)       //right
	normals = addCoordinate(normals, 0, 1, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0)       //top
	normals = addCoordinate(normals, 0, -1, 0, 0, -1, 0, 0, -1, 0, 0, -1, 0, 0, -1, 0, 0, -1, 0) //bottom

	mesh.Normals = unsafe.SliceData(normals)

	// 4 texCoords

	os := float32(1) / 6 // one sixth of the texture
	// front,left,back,right,top,bottom
	tX := os * 0
	texCoords = addCoordinate(texCoords, tX, 0, tX, 1, tX+os, 1, tX, 0, tX+os, 0, tX+os, 1)
	tX = os * 1
	texCoords = addCoordinate(texCoords, tX, 0, tX, 1, tX+os, 1, tX, 0, tX+os, 0, tX+os, 1)
	tX = os * 2
	texCoords = addCoordinate(texCoords, tX, 0, tX, 1, tX+os, 1, tX, 0, tX+os, 0, tX+os, 1)
	tX = os * 3
	texCoords = addCoordinate(texCoords, tX, 0, tX, 1, tX+os, 1, tX, 0, tX+os, 0, tX+os, 1)
	tX = os * 4
	texCoords = addCoordinate(texCoords, tX, 0, tX, 1, tX+os, 1, tX, 0, tX+os, 0, tX+os, 1)
	tX = os * 5
	texCoords = addCoordinate(texCoords, tX, 0, tX, 1, tX+os, 1, tX, 0, tX+os, 0, tX+os, 1)

	mesh.Texcoords = unsafe.SliceData(texCoords)

	// Upload mesh data from CPU (RAM) to GPU (VRAM) memory
	rl.UploadMesh(&mesh, false)

	return &mesh
}

func addCoordinate(slice []float32, values ...float32) []float32 {
	for _, value := range values {
		slice = append(slice, value)
	}
	return slice
}
