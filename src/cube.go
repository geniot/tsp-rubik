package main

type Cube struct {
	size   int
	cubies [3][3][3]*Cubie
}

// front-green, back-blue, left-orange, right-red, top-yellow, bottom-white
func NewCube(size int) *Cube {
	return &Cube{size: size, cubies: [3][3][3]*Cubie{
		{
			{
				{colors: [6]int{BL, O, B, BL, BL, W}, x: -1, y: -1, z: -1},
				{colors: [6]int{BL, O, BL, BL, BL, W}, x: -1, y: -1, z: 0},
				{colors: [6]int{G, O, BL, BL, BL, W}, x: -1, y: -1, z: 1},
			},
			{
				{colors: [6]int{BL, O, B, BL, BL, BL}, x: -1, y: 0, z: -1},
				{colors: [6]int{BL, O, BL, BL, BL, BL}, x: -1, y: 0, z: 0},
				{colors: [6]int{G, O, BL, BL, BL, BL}, x: -1, y: 0, z: 1, angleDelta: 180},
			},
			{
				{colors: [6]int{BL, O, B, BL, Y, BL}, x: -1, y: 1, z: -1},
				{colors: [6]int{BL, O, BL, BL, Y, BL}, x: -1, y: 1, z: 0},
				{colors: [6]int{G, O, BL, BL, Y, BL}, x: -1, y: 1, z: 1},
			},
		},
		{
			{
				{colors: [6]int{BL, BL, B, BL, BL, W}, x: 0, y: -1, z: -1},
				{colors: [6]int{BL, BL, BL, BL, BL, W}, x: 0, y: -1, z: 0},
				{colors: [6]int{G, BL, BL, BL, BL, W}, x: 0, y: -1, z: 1, angleDelta: 270},
			},
			{
				{colors: [6]int{BL, BL, B, BL, BL, BL}, x: 0, y: 0, z: -1},
				{colors: [6]int{BL, BL, BL, BL, BL, BL}, x: 0, y: 0, z: 0},
				{colors: [6]int{G, BL, BL, BL, BL, BL}, x: 0, y: 0, z: 1},
			},
			{
				{colors: [6]int{BL, BL, B, BL, Y, BL}, x: 0, y: 1, z: -1},
				{colors: [6]int{BL, BL, BL, BL, Y, BL}, x: 0, y: 1, z: 0},
				{colors: [6]int{G, BL, BL, BL, Y, BL}, x: 0, y: 1, z: 1, angleDelta: 90},
			},
		},
		{
			{
				{colors: [6]int{BL, BL, B, R, BL, W}, x: 1, y: -1, z: -1},
				{colors: [6]int{BL, BL, BL, R, BL, W}, x: 1, y: -1, z: 0},
				{colors: [6]int{G, BL, BL, R, BL, W}, x: 1, y: -1, z: 1},
			},
			{
				{colors: [6]int{BL, BL, B, R, BL, BL}, x: 1, y: 0, z: -1},
				{colors: [6]int{BL, BL, BL, R, BL, BL}, x: 1, y: 0, z: 0},
				{colors: [6]int{G, BL, BL, R, BL, BL}, x: 1, y: 0, z: 1},
			},
			{
				{colors: [6]int{BL, BL, B, R, Y, BL}, x: 1, y: 1, z: -1},
				{colors: [6]int{BL, BL, BL, R, Y, BL}, x: 1, y: 1, z: 0},
				{colors: [6]int{G, BL, BL, R, Y, BL}, x: 1, y: 1, z: 1},
			},
		},
	}}
}

func (c *Cube) shouldSelect(rotation int, x int, y int, z int) bool {
	if rotation == R_BACK && z == 0 {
		return true
	}
	if rotation == R_FB_MIDDLE && z == 1 {
		return true
	}
	if rotation == R_FRONT && z == 2 {
		return true
	}
	if rotation == R_LEFT && x == 0 {
		return true
	}
	if rotation == R_LR_MIDDLE && x == 1 {
		return true
	}
	if rotation == R_RIGHT && x == 2 {
		return true
	}
	if rotation == R_BOTTOM && y == 0 {
		return true
	}
	if rotation == R_TB_MIDDLE && y == 1 {
		return true
	}
	if rotation == R_TOP && y == 2 {
		return true
	}
	return false
}

func (c *Cube) startRotation(rotation int, isForward bool) {
	for xIterator := 0; xIterator < c.size; xIterator++ {
		for yIterator := 0; yIterator < c.size; yIterator++ {
			for zIterator := 0; zIterator < c.size; zIterator++ {
				cubie := c.cubies[xIterator][yIterator][zIterator]
				if c.shouldSelect(rotation, xIterator, yIterator, zIterator) && !cubie.isRotating() {
					cubie.targetAngleZ += float32(If(isForward, 90, -90))
				}
			}
		}
	}
}
