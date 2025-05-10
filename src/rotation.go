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
}

func (r *Rotation) rotateRight() {
	if !r.isRotating() {
		r.targetAngleX = r.angleX + 90
	}
}

func (r *Rotation) rotateLeft() {
	if !r.isRotating() {
		r.targetAngleX = r.angleX - 90
	}
}

func (r *Rotation) isRotating() bool {
	return r.angleX != r.targetAngleX || r.angleY != r.targetAngleY || r.angleZ != r.targetAngleZ
}

func NewRotation() *Rotation {
	return &Rotation{angleX: 0, angleY: 0, angleZ: 0, targetAngleX: 0, targetAngleY: 0, targetAngleZ: 0}
}
