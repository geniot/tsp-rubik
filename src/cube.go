package main

import (
	"math"
)

type Cube struct {
	size   int
	cubies [3][3][3]*Cubie
}

// front-green, back-blue, left-orange, right-red, top-yellow, bottom-white
func NewCube(size int) *Cube {
	return &Cube{size: size,
		cubies: [3][3][3]*Cubie{
			{
				{
					NewCubie([6]int{BL, O, B, BL, BL, W}, -1, -1, -1),
					NewCubie([6]int{BL, O, BL, BL, BL, W}, -1, -1, 0),
					NewCubie([6]int{G, O, BL, BL, BL, W}, -1, -1, 1),
				},
				{
					NewCubie([6]int{BL, O, B, BL, BL, BL}, -1, 0, -1),
					NewCubie([6]int{BL, O, BL, BL, BL, BL}, -1, 0, 0),
					NewCubie([6]int{G, O, BL, BL, BL, BL}, -1, 0, 1),
				},
				{
					NewCubie([6]int{BL, O, B, BL, Y, BL}, -1, 1, -1),
					NewCubie([6]int{BL, O, BL, BL, Y, BL}, -1, 1, 0),
					NewCubie([6]int{G, O, BL, BL, Y, BL}, -1, 1, 1),
				},
			},
			{
				{
					NewCubie([6]int{BL, BL, B, BL, BL, W}, 0, -1, -1),
					NewCubie([6]int{BL, BL, BL, BL, BL, W}, 0, -1, 0),
					NewCubie([6]int{G, BL, BL, BL, BL, W}, 0, -1, 1),
				},
				{
					NewCubie([6]int{BL, BL, B, BL, BL, BL}, 0, 0, -1),
					NewCubie([6]int{BL, BL, BL, BL, BL, BL}, 0, 0, 0),
					NewCubie([6]int{G, BL, BL, BL, BL, BL}, 0, 0, 1),
				},
				{
					NewCubie([6]int{BL, BL, B, BL, Y, BL}, 0, 1, -1),
					NewCubie([6]int{BL, BL, BL, BL, Y, BL}, 0, 1, 0),
					NewCubie([6]int{G, BL, BL, BL, Y, BL}, 0, 1, 1),
				},
			},
			{
				{
					NewCubie([6]int{BL, BL, B, R, BL, W}, 1, -1, -1),
					NewCubie([6]int{BL, BL, BL, R, BL, W}, 1, -1, 0),
					NewCubie([6]int{G, BL, BL, R, BL, W}, 1, -1, 1),
				},
				{
					NewCubie([6]int{BL, BL, B, R, BL, BL}, 1, 0, -1),
					NewCubie([6]int{BL, BL, BL, R, BL, BL}, 1, 0, 0),
					NewCubie([6]int{G, BL, BL, R, BL, BL}, 1, 0, 1),
				},
				{
					NewCubie([6]int{BL, BL, B, R, Y, BL}, 1, 1, -1),
					NewCubie([6]int{BL, BL, BL, R, Y, BL}, 1, 1, 0),
					NewCubie([6]int{G, BL, BL, R, Y, BL}, 1, 1, 1),
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
