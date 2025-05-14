package main

type Rotation struct {
	angleX       float32
	angleY       float32
	angleZ       float32
	targetAngleX float32
	targetAngleY float32
	targetAngleZ float32
}

func (r *Rotation) update() bool {
	//X
	if r.targetAngleX > r.angleX {
		r.angleX += rotationSpeed
		if r.angleX > r.targetAngleX {
			r.angleX = r.targetAngleX
			return false
		}
	}
	if r.targetAngleX < r.angleX {
		r.angleX -= rotationSpeed
		if r.angleX < r.targetAngleX {
			r.angleX = r.targetAngleX
			return false
		}
	}
	//Y
	if r.targetAngleY > r.angleY {
		r.angleY += rotationSpeed
		if r.angleY > r.targetAngleY {
			r.angleY = r.targetAngleY
			return false
		}
	}
	if r.targetAngleY < r.angleY {
		r.angleY -= rotationSpeed
		if r.angleY < r.targetAngleY {
			r.angleY = r.targetAngleY
			return false
		}
	}
	//Z
	if r.targetAngleZ > r.angleZ {
		r.angleZ += rotationSpeed
		if r.angleZ > r.targetAngleZ {
			r.angleZ = r.targetAngleZ
			return false
		}
	}
	if r.targetAngleZ < r.angleZ {
		r.angleZ -= rotationSpeed
		if r.angleZ < r.targetAngleZ {
			r.angleZ = r.targetAngleZ
			return false
		}
	}
	return true
}

func (r *Rotation) rotateX(angleDelta float32) {
	if !r.isRotating() {
		r.targetAngleX = r.angleX + angleDelta
	}
}

func (r *Rotation) rotateY(angleDelta float32) {
	if !r.isRotating() {
		r.targetAngleY = r.angleY + angleDelta
	}
}

func (r *Rotation) rotateZ(angleDelta float32) {
	if !r.isRotating() {
		r.targetAngleZ = r.angleZ + angleDelta
	}
}

func (r *Rotation) isRotating() bool {
	return r.angleX != r.targetAngleX || r.angleY != r.targetAngleY || r.angleZ != r.targetAngleZ
}

func (r *Rotation) rotate(rotation int, isForward bool) {
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

func NewRotation() *Rotation {
	return &Rotation{angleX: 0, angleY: 0, angleZ: 0, targetAngleX: 0, targetAngleY: 0, targetAngleZ: 0}
}
