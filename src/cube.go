package main

type Cube struct {
	cubies [3][3][3]*Cubie
}

func (c *Cube) isRotating() bool {
	for xIterator := 0; xIterator < 3; xIterator++ {
		for yIterator := 0; yIterator < 3; yIterator++ {
			for zIterator := 0; zIterator < 3; zIterator++ {
				if c.cubies[xIterator][yIterator][zIterator].isRotating() {
					return true
				}
			}
		}
	}
	return false
}

func (c *Cube) orderRotation(rotation int, isForward bool) {
	if !c.isRotating() {
		for xIterator := 0; xIterator < 3; xIterator++ {
			for yIterator := 0; yIterator < 3; yIterator++ {
				for zIterator := 0; zIterator < 3; zIterator++ {
					if c.shouldSelect(rotation, xIterator, yIterator, zIterator) {
						c.cubies[xIterator][yIterator][zIterator].startRotation(rotation, isForward)
					}
				}
			}
		}
	}
}

func (c *Cube) shouldSelect(rotation int, x int, y int, z int) bool {
	var (
		cubie = c.cubies[x][y][z]
	)
	if rotation == R_BACK && cubie.z == 0 {
		return true
	}
	if rotation == R_FB_MIDDLE && cubie.z == 1 {
		return true
	}
	if rotation == R_FRONT && cubie.z == 2 {
		return true
	}
	if rotation == R_LEFT && cubie.x == 0 {
		return true
	}
	if rotation == R_LR_MIDDLE && cubie.x == 1 {
		return true
	}
	if rotation == R_RIGHT && cubie.x == 2 {
		return true
	}
	if rotation == R_BOTTOM && cubie.y == 0 {
		return true
	}
	if rotation == R_TB_MIDDLE && cubie.y == 1 {
		return true
	}
	if rotation == R_TOP && cubie.y == 2 {
		return true
	}

	return false
}

// front-green, back-blue, left-orange, right-red, top-yellow, bottom-white
func NewCube() *Cube {
	return &Cube{cubies: [3][3][3]*Cubie{
		{
			{
				{colors: [6]int{BL, O, B, BL, BL, W}, x: 0, y: 0, z: 0},
				{colors: [6]int{BL, O, BL, BL, BL, W}, x: 0, y: 0, z: 1},
				{colors: [6]int{G, O, BL, BL, BL, W}, x: 0, y: 0, z: 2},
			},
			{
				{colors: [6]int{BL, O, B, BL, BL, BL}, x: 0, y: 1, z: 0},
				{colors: [6]int{BL, O, BL, BL, BL, BL}, x: 0, y: 1, z: 1},
				{colors: [6]int{G, O, BL, BL, BL, BL}, x: 0, y: 1, z: 2},
			},
			{
				{colors: [6]int{BL, O, B, BL, Y, BL}, x: 0, y: 2, z: 0},
				{colors: [6]int{BL, O, BL, BL, Y, BL}, x: 0, y: 2, z: 1},
				{colors: [6]int{G, O, BL, BL, Y, BL}, x: 0, y: 2, z: 2},
			},
		},
		{
			{
				{colors: [6]int{BL, BL, B, BL, BL, W}, x: 1, y: 0, z: 0},
				{colors: [6]int{BL, BL, BL, BL, BL, W}, x: 1, y: 0, z: 1},
				{colors: [6]int{G, BL, BL, BL, BL, W}, x: 1, y: 0, z: 2},
			},
			{
				{colors: [6]int{BL, BL, B, BL, BL, BL}, x: 1, y: 1, z: 0},
				{colors: [6]int{BL, BL, BL, BL, BL, BL}, x: 1, y: 1, z: 1},
				{colors: [6]int{G, BL, BL, BL, BL, BL}, x: 1, y: 1, z: 2},
			},
			{
				{colors: [6]int{BL, BL, B, BL, Y, BL}, x: 1, y: 2, z: 0},
				{colors: [6]int{BL, BL, BL, BL, Y, BL}, x: 1, y: 2, z: 1},
				{colors: [6]int{G, BL, BL, BL, Y, BL}, x: 1, y: 2, z: 2},
			},
		},
		{
			{
				{colors: [6]int{BL, BL, B, R, BL, W}, x: 2, y: 0, z: 0},
				{colors: [6]int{BL, BL, BL, R, BL, W}, x: 2, y: 0, z: 1},
				{colors: [6]int{G, BL, BL, R, BL, W}, x: 2, y: 0, z: 2},
			},
			{
				{colors: [6]int{BL, BL, B, R, BL, BL}, x: 2, y: 1, z: 0},
				{colors: [6]int{BL, BL, BL, R, BL, BL}, x: 2, y: 1, z: 1},
				{colors: [6]int{G, BL, BL, R, BL, BL}, x: 2, y: 1, z: 2},
			},
			{
				{colors: [6]int{BL, BL, B, R, Y, BL}, x: 2, y: 2, z: 0},
				{colors: [6]int{BL, BL, BL, R, Y, BL}, x: 2, y: 2, z: 1},
				{colors: [6]int{G, BL, BL, R, Y, BL}, x: 2, y: 2, z: 2},
			},
		},
	}}
}
