package main

type Cube struct {
	cubies [3][3][3]*Cubie
}

func (c *Cube) rotate(rotation int) {
	var (
		cubie0 = c.cubies[0][2][0]
		cubie1 = c.cubies[0][2][1]
		cubie2 = c.cubies[0][2][2]
		cubie3 = c.cubies[1][2][0]
		cubie4 = c.cubies[1][2][1]
		cubie5 = c.cubies[1][2][2]
		cubie6 = c.cubies[2][2][0]
		cubie7 = c.cubies[2][2][1]
		cubie8 = c.cubies[2][2][2]
	)
	cubiesToRotate := [9]*Cubie{
		cubie0, cubie1, cubie2, cubie3, cubie4, cubie5, cubie6, cubie7, cubie8,
	}
	for _, cubie := range cubiesToRotate {
		if rotation == R_TOP_B {
			cubie.rotateY(true)
			c.cubies[0][2][0] = cubie6
			c.cubies[0][2][1] = cubie3
			c.cubies[0][2][2] = cubie0
			c.cubies[1][2][0] = cubie7
			c.cubies[1][2][1] = cubie4
			c.cubies[1][2][2] = cubie1
			c.cubies[2][2][0] = cubie8
			c.cubies[2][2][1] = cubie5
			c.cubies[2][2][2] = cubie2

		}
		if rotation == R_TOP_F {
			cubie.rotateY(false)
			c.cubies[0][2][0] = cubie2
			c.cubies[0][2][1] = cubie5
			c.cubies[0][2][2] = cubie8
			c.cubies[1][2][0] = cubie1
			c.cubies[1][2][1] = cubie4
			c.cubies[1][2][2] = cubie7
			c.cubies[2][2][0] = cubie0
			c.cubies[2][2][1] = cubie3
			c.cubies[2][2][2] = cubie6

		}
	}

}

// front-green, back-blue, left-orange, right-red, top-yellow, bottom-white
func NewCube() *Cube {
	return &Cube{cubies: [3][3][3]*Cubie{
		{
			{
				{colors: [6]int{BL, O, B, BL, BL, W}},
				{colors: [6]int{BL, O, BL, BL, BL, W}},
				{colors: [6]int{G, O, BL, BL, BL, W}},
			},
			{
				{colors: [6]int{BL, O, B, BL, BL, BL}},
				{colors: [6]int{BL, O, BL, BL, BL, BL}},
				{colors: [6]int{G, O, BL, BL, BL, BL}},
			},
			{
				{colors: [6]int{BL, O, B, BL, Y, BL}},
				{colors: [6]int{BL, O, BL, BL, Y, BL}},
				{colors: [6]int{G, O, BL, BL, Y, BL}},
			},
		},
		{
			{
				{colors: [6]int{BL, BL, B, BL, BL, W}},
				{colors: [6]int{BL, BL, BL, BL, BL, W}},
				{colors: [6]int{G, BL, BL, BL, BL, W}},
			},
			{
				{colors: [6]int{BL, BL, B, BL, BL, BL}},
				{colors: [6]int{BL, BL, BL, BL, BL, BL}},
				{colors: [6]int{G, BL, BL, BL, BL, BL}},
			},
			{
				{colors: [6]int{BL, BL, B, BL, Y, BL}},
				{colors: [6]int{BL, BL, BL, BL, Y, BL}},
				{colors: [6]int{G, BL, BL, BL, Y, BL}},
			},
		},
		{
			{
				{colors: [6]int{BL, BL, B, R, BL, W}},
				{colors: [6]int{BL, BL, BL, R, BL, W}},
				{colors: [6]int{G, BL, BL, R, BL, W}},
			},
			{
				{colors: [6]int{BL, BL, B, R, BL, BL}},
				{colors: [6]int{BL, BL, BL, R, BL, BL}},
				{colors: [6]int{G, BL, BL, R, BL, BL}},
			},
			{
				{colors: [6]int{BL, BL, B, R, Y, BL}},
				{colors: [6]int{BL, BL, BL, R, Y, BL}},
				{colors: [6]int{G, BL, BL, R, Y, BL}},
			},
		},
	}}
}
