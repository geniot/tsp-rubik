package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"math"
)

type Cubie struct {
	colors       [6]int
	x, y, z      float32
	r            float64
	angleX       float32
	angleY       float32
	angleZ       float32
	targetAngleX float32
	targetAngleY float32
	targetAngleZ float32
	actualAngleX float32
	actualAngleY float32
	actualAngleZ float32
	vecX         rl.Vector3
	vecY         rl.Vector3
	vecZ         rl.Vector3
}

func NewCubie(colors [6]int, x, y, z int) *Cubie {
	cubie := &Cubie{colors: colors, x: float32(x), y: float32(y), z: float32(z)}
	sum := math.Round(math.Abs(float64(cubie.x))) + math.Round(math.Abs(float64(cubie.y))) + math.Round(math.Abs(float64(cubie.z)))
	cubie.r = If(sum == 3, math.Sqrt(2), If(sum == 2, float64(1), 0))
	cubie.vecX = rl.Vector3{X: 1}
	cubie.vecY = rl.Vector3{Y: 1}
	cubie.vecZ = rl.Vector3{Z: 1}
	return cubie
}

func (c *Cubie) isRotating() bool {
	return c.angleX != c.targetAngleX || c.angleY != c.targetAngleY || c.angleZ != c.targetAngleZ
}

func (c *Cubie) shouldSelect(rotation int) bool {
	if rotation == R_BACK && math.Round(float64(c.z)) == -1 {
		return true
	}
	if rotation == R_FB_MIDDLE && math.Round(float64(c.z)) == 0 {
		return true
	}
	if rotation == R_FRONT && math.Round(float64(c.z)) == 1 {
		return true
	}
	if rotation == R_LEFT && math.Round(float64(c.x)) == -1 {
		return true
	}
	if rotation == R_LR_MIDDLE && math.Round(float64(c.x)) == 0 {
		return true
	}
	if rotation == R_RIGHT && math.Round(float64(c.x)) == 1 {
		return true
	}
	if rotation == R_BOTTOM && math.Round(float64(c.y)) == -1 {
		return true
	}
	if rotation == R_TB_MIDDLE && math.Round(float64(c.y)) == 0 {
		return true
	}
	if rotation == R_TOP && math.Round(float64(c.y)) == 1 {
		return true
	}
	return false
}

func (c *Cubie) update() bool {
	//X
	if c.targetAngleX > c.angleX {
		c.angleX += rotationSpeed
		angleDelta := 90 - (c.targetAngleX - c.angleX)
		sinDelta := float32(c.r * math.Sin(float64((c.actualAngleX+angleDelta)*math.Pi/180)))
		cosDelta := float32(c.r * math.Cos(float64((c.actualAngleX+angleDelta)*math.Pi/180)))
		c.y = cosDelta
		c.z = sinDelta
		if c.angleX >= c.targetAngleX {
			c.angleX = c.targetAngleX
			return false
		}
	}
	if c.targetAngleX < c.angleX {
		c.angleX -= rotationSpeed
		angleDelta := float32(90 - math.Abs(float64(c.targetAngleX-c.angleX)))
		sinDelta := float32(c.r * math.Sin(float64((c.actualAngleX-angleDelta)*math.Pi/180)))
		cosDelta := float32(c.r * math.Cos(float64((c.actualAngleX-angleDelta)*math.Pi/180)))
		c.y = cosDelta
		c.z = sinDelta
		if c.angleX <= c.targetAngleX {
			c.angleX = c.targetAngleX
			return false
		}
	}
	//Y
	if c.targetAngleY > c.angleY {
		c.angleY += rotationSpeed
		//sinDelta := float32(c.r * math.Sin(float64((c.angleY+c.angleDeltaY)*math.Pi/180)))
		//cosDelta := float32(c.r * math.Cos(float64((c.angleY+c.angleDeltaY)*math.Pi/180)))
		//c.x = cosDelta
		//c.y = sinDelta
		if c.angleY >= c.targetAngleY {
			c.angleY = c.targetAngleY
			return false
		}
	}
	if c.targetAngleY < c.angleY {
		c.angleY -= rotationSpeed
		//sinDelta := float32(c.r * math.Sin(float64((c.angleY+c.angleDeltaY)*math.Pi/180)))
		//cosDelta := float32(c.r * math.Cos(float64((c.angleY+c.angleDeltaY)*math.Pi/180)))
		//c.x = cosDelta
		//c.y = sinDelta
		if c.angleY <= c.targetAngleY {
			c.angleY = c.targetAngleY
			return false
		}
	}
	//Z
	if c.targetAngleZ > c.angleZ {
		c.angleZ += rotationSpeed
		angleDelta := 90 - (c.targetAngleZ - c.angleZ)
		sinDelta := float32(c.r * math.Sin(float64((c.actualAngleZ+angleDelta)*math.Pi/180)))
		cosDelta := float32(c.r * math.Cos(float64((c.actualAngleZ+angleDelta)*math.Pi/180)))
		c.x = cosDelta
		c.y = sinDelta
		if c.angleZ >= c.targetAngleZ {
			c.angleZ = c.targetAngleZ
			return false
		}
	}
	if c.targetAngleZ < c.angleZ {
		c.angleZ -= rotationSpeed
		angleDelta := float32(90 - math.Abs(float64(c.targetAngleZ-c.angleZ)))
		sinDelta := float32(c.r * math.Sin(float64((c.actualAngleZ-angleDelta)*math.Pi/180)))
		cosDelta := float32(c.r * math.Cos(float64((c.actualAngleZ-angleDelta)*math.Pi/180)))
		c.x = cosDelta
		c.y = sinDelta
		if c.angleZ <= c.targetAngleZ {
			c.angleZ = c.targetAngleZ
			return false
		}
	}
	return true
}

func (c *Cubie) updateVecs() {
	//restX := math.Round(math.Mod(float64(c.targetAngleX), 360))
	//restY := math.Round(math.Mod(float64(c.targetAngleY), 360))
	//restZ := math.Round(math.Mod(float64(c.targetAngleZ), 360))
	//c.vecX = rl.Vector3{X: 1}
	//c.vecY = rl.Vector3{Y: 1}
	//c.vecZ = rl.Vector3{Z: 1}
	//println(restX, restY, restZ)
	//if restX == 0 {
	//	c.vecY = rl.Vector3{Y: 1}
	//	c.vecZ = rl.Vector3{Z: 1}
	//}
	//if restX == 90 {
	//	c.vecY = rl.Vector3{Z: -1}
	//	c.vecZ = rl.Vector3{Y: 1}
	//}
	//if restX == 180 {
	//	c.vecY = rl.Vector3{Y: -1}
	//	c.vecZ = rl.Vector3{Z: -1}
	//}
	//if restX == 270 {
	//	c.vecY = rl.Vector3{Y: -1}
	//	c.vecZ = rl.Vector3{Y: -1}
	//}
	//if restX == -90 {
	//	c.vecZ = rl.Vector3{Y: -1}
	//}
	//if restX == -180 {
	//	c.vecZ = rl.Vector3{Z: -1}
	//}
	//if restX == -270 {
	//	c.vecZ = rl.Vector3{Y: 1}
	//}
}
