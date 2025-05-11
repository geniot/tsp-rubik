package main

const (
	FRONT  = 0
	LEFT   = 1
	BACK   = 2
	RIGHT  = 3
	TOP    = 4
	BOTTOM = 5
)

var (
	xFrom = [4]int{RIGHT, BOTTOM, LEFT, TOP}
	yFrom = [4]int{FRONT, LEFT, BACK, RIGHT}
	zFrom = [4]int{FRONT, BOTTOM, BACK, TOP}
)

type Cubie struct {
	colors [6]int //front,left,back,right,top,bottom
}

func (c *Cubie) rotateX(isForward bool) {
	c.moveColors(xFrom, isForward)
}
func (c *Cubie) rotateY(isForward bool) {
	c.moveColors(yFrom, isForward)
}
func (c *Cubie) rotateZ(isForward bool) {
	c.moveColors(zFrom, isForward)
}

func (c *Cubie) moveColors(fromIndexes [4]int, isForward bool) {
	var (
		fromColor0 = c.colors[fromIndexes[0]]
		fromColor1 = c.colors[fromIndexes[1]]
		fromColor2 = c.colors[fromIndexes[2]]
		fromColor3 = c.colors[fromIndexes[3]]
	)
	if isForward {
		c.colors[fromIndexes[0]] = fromColor1
		c.colors[fromIndexes[1]] = fromColor2
		c.colors[fromIndexes[2]] = fromColor3
		c.colors[fromIndexes[3]] = fromColor0
	} else {
		c.colors[fromIndexes[0]] = fromColor3
		c.colors[fromIndexes[1]] = fromColor0
		c.colors[fromIndexes[2]] = fromColor1
		c.colors[fromIndexes[3]] = fromColor2
	}
}
