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
	repos = [9][4]int{
		{0, 0, 2, 0},
		{2, 0, 2, 2},
		{2, 2, 0, 2},
		{0, 2, 0, 0},
		{1, 0, 2, 1},
		{2, 1, 1, 2},
		{1, 2, 0, 1},
		{0, 1, 1, 0},
		{1, 1, 1, 1},
	}
)

type Cubie struct {
	colors       [6]int //front,left,back,right,top,bottom
	x            int
	y            int
	z            int
	angleX       float32
	angleY       float32
	angleZ       float32
	targetAngleX float32
	targetAngleY float32
	targetAngleZ float32
}

func (r *Cubie) update() bool {
	//X
	if r.targetAngleX > r.angleX {
		r.angleX += rotationSpeed
		if r.angleX >= r.targetAngleX {
			r.angleX = r.targetAngleX
			return false
		}
	}
	if r.targetAngleX < r.angleX {
		r.angleX -= rotationSpeed
		if r.angleX <= r.targetAngleX {
			r.angleX = r.targetAngleX
			return false
		}
	}
	//Y
	if r.targetAngleY > r.angleY {
		r.angleY += rotationSpeed
		if r.angleY >= r.targetAngleY {
			r.angleY = r.targetAngleY
			r.x, r.z = match(r.x, r.z, false)
			return false
		}
	}
	if r.targetAngleY < r.angleY {
		r.angleY -= rotationSpeed
		if r.angleY <= r.targetAngleY {
			r.angleY = r.targetAngleY
			r.x, r.z = match(r.x, r.z, true)
			return false
		}
	}
	//Z
	if r.targetAngleZ > r.angleZ {
		r.angleZ += rotationSpeed
		if r.angleZ >= r.targetAngleZ {
			r.angleZ = r.targetAngleZ
			return false
		}
	}
	if r.targetAngleZ < r.angleZ {
		r.angleZ -= rotationSpeed
		if r.angleZ <= r.targetAngleZ {
			r.angleZ = r.targetAngleZ
			return false
		}
	}
	return true
}

func match(p1 int, p2 int, isForward bool) (int, int) {
	for _, tuple := range repos {
		if isForward && tuple[0] == p1 && tuple[1] == p2 {
			return tuple[2], tuple[3]
		}
		if !isForward && tuple[2] == p1 && tuple[3] == p2 {
			return tuple[0], tuple[1]
		}
	}
	println("match error:", p1, p2)
	return p1, p2
}

func (r *Cubie) rotateX(angleDelta float32) {
	if !r.isRotating() {
		r.targetAngleX = r.angleX + angleDelta
	}
}

func (r *Cubie) rotateY(angleDelta float32) {
	if !r.isRotating() {
		r.targetAngleY = r.angleY + angleDelta
	}
}

func (r *Cubie) rotateZ(angleDelta float32) {
	if !r.isRotating() {
		r.targetAngleZ = r.angleZ + angleDelta
	}
}

func (r *Cubie) isRotating() bool {
	return r.angleX != r.targetAngleX || r.angleY != r.targetAngleY || r.angleZ != r.targetAngleZ
}

func (r *Cubie) startRotation(rotation int, isForward bool) {
	if rotation == R_LEFT || rotation == R_LR_MIDDLE || rotation == R_RIGHT {
		r.rotateX(float32(If(isForward, -90, 90)))
	}
	if rotation == R_TOP || rotation == R_TB_MIDDLE || rotation == R_BOTTOM {
		r.rotateY(float32(If(isForward, -90, 90)))
	}
	if rotation == R_FRONT || rotation == R_FB_MIDDLE || rotation == R_BACK {
		r.rotateZ(float32(If(isForward, 90, -90)))
	}
}
