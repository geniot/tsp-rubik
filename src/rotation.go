package main

const (
	rotationSpeed = 0.3
)

// rotations
const (
	//forward
	R_NONE = iota
	R_FRONT_F
	R_FB_MIDDLE_F
	R_BACK_F
	R_LEFT_F
	R_LR_MIDDLE_F
	R_RIGHT_F
	R_TOP_F
	R_TB_MIDDLE_F
	R_BOTTOM_F
	//same backwards
	R_FRONT_B
	R_FB_MIDDLE_B
	R_BACK_B
	R_LEFT_B
	R_LR_MIDDLE_B
	R_RIGHT_B
	R_TOP_B
	R_TB_MIDDLE_B
	R_BOTTOM_B
)

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
			//r.angleY = 0
			//r.targetAngleY = 0
			r.angleY = r.targetAngleY
			return false
		}
	}
	if r.targetAngleY < r.angleY {
		r.angleY -= rotationSpeed
		if r.angleY < r.targetAngleY {
			//r.angleY = 0
			//r.targetAngleY = 0
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

func NewRotation() *Rotation {
	return &Rotation{angleX: 0, angleY: 0, angleZ: 0, targetAngleX: 0, targetAngleY: 0, targetAngleZ: 0}
}
