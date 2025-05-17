package main

import "math"

type Cubie struct {
	colors       [6]int
	x, y, z      float32
	r            float64
	angleX       float32
	angleY       float32
	angleZ       float32
	angleDelta   float32
	targetAngleX float32
	targetAngleY float32
	targetAngleZ float32
}

func (c *Cubie) isRotating() bool {
	return c.angleX != c.targetAngleX || c.angleY != c.targetAngleY || c.angleZ != c.targetAngleZ
}

func (c *Cubie) update() bool {
	//X
	if c.targetAngleX > c.angleX {
		c.angleX += rotationSpeed
		if c.angleX >= c.targetAngleX {
			c.angleX = c.targetAngleX
			return false
		}
	}
	if c.targetAngleX < c.angleX {
		c.angleX -= rotationSpeed
		if c.angleX <= c.targetAngleX {
			c.angleX = c.targetAngleX
			return false
		}
	}
	//Y
	if c.targetAngleY > c.angleY {
		c.angleY += rotationSpeed
		if c.angleY >= c.targetAngleY {
			c.angleY = c.targetAngleY
			return false
		}
	}
	if c.targetAngleY < c.angleY {
		c.angleY -= rotationSpeed
		if c.angleY <= c.targetAngleY {
			c.angleY = c.targetAngleY
			return false
		}
	}
	//Z
	if c.targetAngleZ > c.angleZ {
		c.angleZ += rotationSpeed
		sinDelta := float32(c.r * math.Sin(float64((c.angleZ+c.angleDelta)*math.Pi/180)))
		cosDelta := float32(c.r * math.Cos(float64((c.angleZ+c.angleDelta)*math.Pi/180)))
		c.x = cosDelta
		c.y = sinDelta
		if c.angleZ >= c.targetAngleZ {
			c.angleZ = c.targetAngleZ
			return false
		}
	}
	if c.targetAngleZ < c.angleZ {
		c.angleZ -= rotationSpeed
		sinDelta := float32(c.r * math.Sin(float64((c.angleZ+c.angleDelta)*math.Pi/180)))
		cosDelta := float32(c.r * math.Cos(float64((c.angleZ+c.angleDelta)*math.Pi/180)))
		c.x = cosDelta
		c.y = sinDelta
		if c.angleZ <= c.targetAngleZ {
			c.angleZ = c.targetAngleZ
			return false
		}
	}
	return true
}
