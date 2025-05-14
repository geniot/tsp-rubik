package main

import "slices"

type Cube struct {
	size   int
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
		return c.colors[FRONT][x+y*c.size]
	}
	if face == LEFT && x == 0 {
		return c.colors[LEFT][y+z*c.size]
	}
	if face == BACK && z == 0 {
		return c.colors[BACK][x+y*c.size]
	}
	if face == RIGHT && x == c.size-1 {
		return c.colors[RIGHT][y+z*c.size]
	}
	if face == TOP && y == c.size-1 {
		return c.colors[TOP][x+z*c.size]
	}
	if face == BOTTOM && y == 0 {
		return c.colors[BOTTOM][x+z*c.size]
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

func (c *Cube) rotate(rotation int, isForward bool) {
	if rotation == R_FRONT {
		var tmpTop = slices.Clone(c.colors[TOP])
		if isForward {
			copyInts(c.colors[RIGHT], c.colors[TOP], c.size*2, c.size*2, c.size)
			copyInts(c.colors[BOTTOM], c.colors[RIGHT], c.size*2, c.size*2, c.size)
			copyInts(c.colors[LEFT], c.colors[BOTTOM], c.size*2, c.size*2, c.size)
			copyInts(tmpTop, c.colors[LEFT], c.size*2, c.size*2, c.size)
		} else {
			copyInts(c.colors[LEFT], c.colors[TOP], c.size*2, c.size*2, c.size)
			copyInts(c.colors[BOTTOM], c.colors[LEFT], c.size*2, c.size*2, c.size)
			copyInts(c.colors[RIGHT], c.colors[BOTTOM], c.size*2, c.size*2, c.size)
			copyInts(tmpTop, c.colors[RIGHT], c.size*2, c.size*2, c.size)
		}
	}
}
