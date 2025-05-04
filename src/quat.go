package main

type Quat [4]float32

func (q *Quat) Identity() {
	q[0] = 0
	q[1] = 0
	q[2] = 0
	q[3] = 1
}

func (r *Quat) Add(a, b *Quat) {
	for i := 0; i < 4; i++ {
		r[i] = a[i] + b[i]
	}
}

func (r *Quat) AddVec3(a *Quat, v *Vec3) {
	for i := 0; i < 3; i++ {
		r[i] = a[i] + v[i]
	}
}

func (r *Quat) Sub(a, b *Quat) {
	for i := 0; i < 4; i++ {
		r[i] = a[i] - b[i]
	}
}

// MultCross the same as for Vec3
func (r *Quat) MultCross3(a, b *Quat) {
	r[0] = a[1]*b[2] - a[2]*b[1]
	r[1] = a[2]*b[0] - a[0]*b[2]
	r[2] = a[0]*b[1] - a[1]*b[0]
}

// QuatMultInner the same as for Vec3
func QuatMultInner3(a, b *Quat) (p float32) {
	for i := 0; i < 3; i++ {
		p += b[i] * a[i]
	}
	return p
}

func (r *Quat) Mult(p, q *Quat) {
	var w = new(Vec3)
	r.MultCross3(p, q)
	w.ScaleQuat(p, q[3])
	r.AddVec3(r, w)
	w.ScaleQuat(q, p[3])
	r.AddVec3(r, w)

	r[3] = p[3]*q[3] - QuatMultInner3(p, q)
}

func (r *Quat) Scale(q *Quat, s float32) {
	for i := 0; i < 4; i++ {
		r[i] = q[i] * s
	}
}

func QuatInnerProduct(a, b *Quat) (p float32) {
	for i := 0; i < 4; i++ {
		p += b[i] * a[i]
	}
	return p
}

func (r *Quat) Conj(q *Quat) {
	for i := 0; i < 3; i++ {
		r[i] = -q[i]
	}
	r[3] = q[3]
}

func (q *Quat) Len() float32 {
	return sqrtf(QuatInnerProduct(q, q))
}

func (r *Quat) Norm(q *Quat) {
	var k float32 = 1.0 / q.Len()
	r.Scale(q, k)
}
