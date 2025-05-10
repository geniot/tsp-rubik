package main

const (
	rotationSpeed = 0.3
)

type Rotation struct {
	angleX       float32
	angleY       float32
	angleZ       float32
	targetAngleX float32
	targetAngleY float32
	targetAngleZ float32
}

func (r *Rotation) update() {
	//X
	if r.targetAngleX > r.angleX {
		r.angleX += rotationSpeed
		if r.angleX > r.targetAngleX {
			r.angleX = r.targetAngleX
		}
	}
	if r.targetAngleX < r.angleX {
		r.angleX -= rotationSpeed
		if r.angleX < r.targetAngleX {
			r.angleX = r.targetAngleX
		}
	}
	//Y
	if r.targetAngleY > r.angleY {
		r.angleY += rotationSpeed
		if r.angleY > r.targetAngleY {
			r.angleY = r.targetAngleY
		}
	}
	if r.targetAngleY < r.angleY {
		r.angleY -= rotationSpeed
		if r.angleY < r.targetAngleY {
			r.angleY = r.targetAngleY
		}
	}
	//Z
	if r.targetAngleZ > r.angleZ {
		r.angleZ += rotationSpeed
		if r.angleZ > r.targetAngleZ {
			r.angleZ = r.targetAngleZ
		}
	}
	if r.targetAngleZ < r.angleZ {
		r.angleZ -= rotationSpeed
		if r.angleZ < r.targetAngleZ {
			r.angleZ = r.targetAngleZ
		}
	}
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
