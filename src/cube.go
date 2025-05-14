package main

type Cube struct {
	size int
	//cubies []*Cubie
	colors [6][]int
}

// NewCube front-green, back-blue, left-orange, right-red, top-yellow, bottom-white
func NewCube(size int) *Cube {
	model := [6][]int{
		newInts(size*size, G), //front-green
		newInts(size*size, O), //left-orange
		newInts(size*size, B), //back-blue
		newInts(size*size, R), //right-red
		newInts(size*size, Y), //top-yellow
		newInts(size*size, W), //bottom-white
	}
	return &Cube{size: size, colors: model}
}

func (c *Cube) getColor(x int, y int, z int, face int) int {
	if face == FRONT && z == c.size-1 {
		return c.colors[0][x+y*(c.size-1)]
	}
	if face == LEFT && x == 0 {
		return c.colors[1][x+y*(c.size-1)]
	}
	if face == BACK && z == 0 {
		return c.colors[2][x+y*(c.size-1)]
	}
	if face == RIGHT && x == c.size-1 {
		return c.colors[3][x+y*(c.size-1)]
	}
	if face == TOP && y == c.size-1 {
		return c.colors[4][x+y*(c.size-1)]
	}
	if face == BOTTOM && y == 0 {
		return c.colors[5][x+y*(c.size-1)]
	}
	return BL
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
