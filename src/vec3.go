package main

type Vec3 [3]float32

func (r *Vec3) Add(a, b *Vec3) {
	for i := 0; i < 3; i++ {
		r[i] = a[i] + b[i]
	}
}

func (r *Vec3) Sub(a, b *Vec3) {
	for i := 0; i < 3; i++ {
		r[i] = a[i] - b[i]
	}
}

func (r *Vec3) Scale(v *Vec3, s float32) {
	for i := 0; i < 3; i++ {
		r[i] = v[i] * s
	}
}

func (r *Vec3) ScaleVec4(v *Vec4, s float32) {
	for i := 0; i < 3; i++ {
		r[i] = v[i] * s
	}
}

func (r *Vec3) ScaleQuat(q *Quat, s float32) {
	for i := 0; i < 3; i++ {
		r[i] = q[i] * s
	}
}

func (v *Vec3) Len() float32 {
	return sqrtf(Vec3MultInner(v, v))
}

func (r *Vec3) Norm(v *Vec3) {
	var k float32 = 1.0 / v.Len()
	r.Scale(v, k)
}

func (r *Vec3) Min(a, b *Vec3) {
	for i := 0; i < 3; i++ {
		if a[i] < b[i] {
			r[i] = a[i]
		} else {
			r[i] = b[i]
		}
	}
}

func (r *Vec3) Max(a, b *Vec3) {
	for i := 0; i < 3; i++ {
		if a[i] > b[i] {
			r[i] = a[i]
		} else {
			r[i] = b[i]
		}
	}
}

func Vec3MultInner(a, b *Vec3) (p float32) {
	for i := 0; i < 3; i++ {
		p += b[i] * a[i]
	}
	return p
}

func (r *Vec3) MultCross(a, b *Vec3) {
	r[0] = a[1]*b[2] - a[2]*b[1]
	r[1] = a[2]*b[0] - a[0]*b[2]
	r[2] = a[0]*b[1] - a[1]*b[0]
}

func (r *Vec3) Reflect(v, n *Vec3) {
	var p float32 = 2 * Vec3MultInner(v, n)
	for i := 0; i < 3; i++ {
		r[i] = v[i] - p*n[i]
	}
}
