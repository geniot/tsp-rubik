package main

import "math"

type Cube struct {
	size   int
	cubies [3][3][3]*Cubie
}

type AngleDelta struct {
	x      float32
	y      float32
	z      float32
	deltaX float32
	deltaY float32
	deltaZ float32
}

var (
	angleDeltas = []AngleDelta{
		{x: 0, y: 0, z: 1, deltaX: 0, deltaY: 0, deltaZ: 0},
		{x: 1, y: 0, z: 1, deltaX: 0, deltaY: 0, deltaZ: 0},
		{x: 1, y: 1, z: 1, deltaX: 0, deltaY: 0, deltaZ: 0.5 * 90},
		{x: 0, y: 1, z: 1, deltaX: 0, deltaY: 0, deltaZ: 1 * 90},
		{x: -1, y: 1, z: 1, deltaX: 0, deltaY: 0, deltaZ: 1.5 * 90},
		{x: -1, y: 0, z: 1, deltaX: 0, deltaY: 0, deltaZ: 2 * 90},
		{x: -1, y: -1, z: 1, deltaX: 0, deltaY: 0, deltaZ: 2.5 * 90},
		{x: 0, y: -1, z: 1, deltaX: 0, deltaY: 0, deltaZ: 3 * 90},
		{x: 1, y: -1, z: 1, deltaX: 0, deltaY: 0, deltaZ: 3.5 * 90},
	}
)

// front-green, back-blue, left-orange, right-red, top-yellow, bottom-white
func NewCube(size int) *Cube {
	return &Cube{size: size, cubies: [3][3][3]*Cubie{
		{
			{
				{colors: [6]int{BL, O, B, BL, BL, W}, x: -1, y: -1, z: -1},
				{colors: [6]int{BL, O, BL, BL, BL, W}, x: -1, y: -1, z: 0},
				{colors: [6]int{G, O, BL, BL, BL, W}, x: -1, y: -1, z: 1, r: math.Sqrt(2)},
			},
			{
				{colors: [6]int{BL, O, B, BL, BL, BL}, x: -1, y: 0, z: -1},
				{colors: [6]int{BL, O, BL, BL, BL, BL}, x: -1, y: 0, z: 0},
				{colors: [6]int{G, O, BL, BL, BL, BL}, x: -1, y: 0, z: 1, r: 1},
			},
			{
				{colors: [6]int{BL, O, B, BL, Y, BL}, x: -1, y: 1, z: -1},
				{colors: [6]int{BL, O, BL, BL, Y, BL}, x: -1, y: 1, z: 0},
				{colors: [6]int{G, O, BL, BL, Y, BL}, x: -1, y: 1, z: 1, r: math.Sqrt(2)},
			},
		},
		{
			{
				{colors: [6]int{BL, BL, B, BL, BL, W}, x: 0, y: -1, z: -1},
				{colors: [6]int{BL, BL, BL, BL, BL, W}, x: 0, y: -1, z: 0},
				{colors: [6]int{G, BL, BL, BL, BL, W}, x: 0, y: -1, z: 1, r: 1},
			},
			{
				{colors: [6]int{BL, BL, B, BL, BL, BL}, x: 0, y: 0, z: -1},
				{colors: [6]int{BL, BL, BL, BL, BL, BL}, x: 0, y: 0, z: 0},
				{colors: [6]int{G, BL, BL, BL, BL, BL}, x: 0, y: 0, z: 1, r: 0},
			},
			{
				{colors: [6]int{BL, BL, B, BL, Y, BL}, x: 0, y: 1, z: -1},
				{colors: [6]int{BL, BL, BL, BL, Y, BL}, x: 0, y: 1, z: 0},
				{colors: [6]int{G, BL, BL, BL, Y, BL}, x: 0, y: 1, z: 1, r: 1},
			},
		},
		{
			{
				{colors: [6]int{BL, BL, B, R, BL, W}, x: 1, y: -1, z: -1, r: math.Sqrt(2)},
				{colors: [6]int{BL, BL, BL, R, BL, W}, x: 1, y: -1, z: 0, r: 1},
				{colors: [6]int{G, BL, BL, R, BL, W}, x: 1, y: -1, z: 1, r: math.Sqrt(2)},
			},
			{
				{colors: [6]int{BL, BL, B, R, BL, BL}, x: 1, y: 0, z: -1, r: 1},
				{colors: [6]int{BL, BL, BL, R, BL, BL}, x: 1, y: 0, z: 0, r: 0},
				{colors: [6]int{G, BL, BL, R, BL, BL}, x: 1, y: 0, z: 1, r: 1},
			},
			{
				{colors: [6]int{BL, BL, B, R, Y, BL}, x: 1, y: 1, z: -1, r: math.Sqrt(2)},
				{colors: [6]int{BL, BL, BL, R, Y, BL}, x: 1, y: 1, z: 0, r: 1},
				{colors: [6]int{G, BL, BL, R, Y, BL}, x: 1, y: 1, z: 1, r: math.Sqrt(2)},
			},
		},
	}}
}

func (c *Cube) startRotation(rotation int, isForward bool) {
	for xIterator := 0; xIterator < c.size; xIterator++ {
		for yIterator := 0; yIterator < c.size; yIterator++ {
			for zIterator := 0; zIterator < c.size; zIterator++ {
				cubie := c.cubies[xIterator][yIterator][zIterator]
				if cubie.shouldSelect(rotation) && !cubie.isRotating() {
					if rotation == R_FRONT || rotation == R_BACK || rotation == R_FB_MIDDLE {
						if math.Round(float64(cubie.x)) == 1 && math.Round(float64(cubie.y)) == 0 {
							cubie.actualAngleZ = 0 * 90
						}
						if math.Round(float64(cubie.x)) == 1 && math.Round(float64(cubie.y)) == 1 {
							cubie.actualAngleZ = 0.5 * 90
						}
						if math.Round(float64(cubie.x)) == 0 && math.Round(float64(cubie.y)) == 1 {
							cubie.actualAngleZ = 1 * 90
						}
						if math.Round(float64(cubie.x)) == -1 && math.Round(float64(cubie.y)) == 1 {
							cubie.actualAngleZ = 1.5 * 90
						}
						if math.Round(float64(cubie.x)) == -1 && math.Round(float64(cubie.y)) == 0 {
							cubie.actualAngleZ = 2 * 90
						}
						if math.Round(float64(cubie.x)) == -1 && math.Round(float64(cubie.y)) == -1 {
							cubie.actualAngleZ = 2.5 * 90
						}
						if math.Round(float64(cubie.x)) == 0 && math.Round(float64(cubie.y)) == -1 {
							cubie.actualAngleZ = 3 * 90
						}
						if math.Round(float64(cubie.x)) == 1 && math.Round(float64(cubie.y)) == -1 {
							cubie.actualAngleZ = 3.5 * 90
						}
						cubie.targetAngleZ += float32(If(isForward, 90, -90))
					}
					if rotation == R_LEFT || rotation == R_RIGHT || rotation == R_LR_MIDDLE {
						if math.Round(float64(cubie.y)) == 1 && math.Round(float64(cubie.z)) == 0 {
							cubie.actualAngleX = 0 * 90
						}
						if math.Round(float64(cubie.y)) == 1 && math.Round(float64(cubie.z)) == 1 {
							cubie.actualAngleX = 0.5 * 90
						}
						if math.Round(float64(cubie.y)) == 0 && math.Round(float64(cubie.z)) == 1 {
							cubie.actualAngleX = 1 * 90
						}
						if math.Round(float64(cubie.y)) == -1 && math.Round(float64(cubie.z)) == 1 {
							cubie.actualAngleX = 1.5 * 90
						}
						if math.Round(float64(cubie.y)) == -1 && math.Round(float64(cubie.z)) == 0 {
							cubie.actualAngleX = 2 * 90
						}
						if math.Round(float64(cubie.y)) == -1 && math.Round(float64(cubie.z)) == -1 {
							cubie.actualAngleX = 2.5 * 90
						}
						if math.Round(float64(cubie.y)) == 0 && math.Round(float64(cubie.z)) == -1 {
							cubie.actualAngleX = 3 * 90
						}
						if math.Round(float64(cubie.y)) == 1 && math.Round(float64(cubie.z)) == -1 {
							cubie.actualAngleX = 3.5 * 90
						}
						cubie.targetAngleX += float32(If(isForward, 90, -90))
					}
					if rotation == R_TOP || rotation == R_BOTTOM || rotation == R_TB_MIDDLE {
						cubie.targetAngleY += float32(If(isForward, 90, -90))
					}
				}
			}
		}
	}
}

func getDeltaZ(cubie *Cubie) float32 {
	for _, angleDelta := range angleDeltas {
		if math.Round(float64(cubie.x)) == float64(angleDelta.x) &&
			math.Round(float64(cubie.y)) == float64(angleDelta.y) &&
			math.Round(float64(cubie.z)) == float64(angleDelta.z) {
			return angleDelta.deltaZ
		}
	}
	return 0
}
