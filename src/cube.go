package main

type Cube struct {
	size     int
	cubies   [3][3][3]*Cubie
	rotation *Rotation
}

// NewCube front-green, back-blue, left-orange, right-red, top-yellow, bottom-white
func NewCube(size int) *Cube {
	//todo: use size to generate cubie config dynamically, also update possible rotations
	return &Cube{size: size,
		rotation: NewRotation(R_RIGHT, true),
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

func (c *Cube) updateThenDraw() {
	c.rotation.update()
	for xIterator := 0; xIterator < c.size; xIterator++ {
		for yIterator := 0; yIterator < c.size; yIterator++ {
			for zIterator := 0; zIterator < c.size; zIterator++ {
				cubie := c.cubies[xIterator][yIterator][zIterator]
				//if xIterator == 0 && yIterator == 0 && zIterator == 0 {
				//if zIterator == 0 || zIterator == 1 ||
				//	(zIterator == 2 && xIterator == 2 && yIterator == 2) {
				//	(zIterator == 2 && xIterator == 0 && yIterator == 1) ||
				//	(zIterator == 2 && xIterator == 0 && yIterator == 2) ||
				//	(zIterator == 2 && xIterator == 1 && yIterator == 0) ||
				//	(zIterator == 2 && xIterator == 1 && yIterator == 1) ||
				//	(zIterator == 2 && xIterator == 1 && yIterator == 2) ||
				//	(zIterator == 2 && xIterator == 2 && yIterator == 0) ||
				//	(zIterator == 2 && xIterator == 2 && yIterator == 1) ||
				//	(zIterator == 2 && xIterator == 2 && yIterator == 2) {
				cubie.update(c.rotation)
				cubie.draw()
				//}
			}
		}
	}
}
