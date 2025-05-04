package main

type Vec4 [4]float32

func (r *Vec4) Add(a, b *Vec4) {
	for i := 0; i < 4; i++ {
		r[i] = a[i] + b[i]
	}
}

func (r *Vec4) Sub(a, b *Vec4) {
	for i := 0; i < 4; i++ {
		r[i] = a[i] - b[i]
	}
}

func (r *Vec4) SubVec3(a *Vec4, b *Vec3) {
	for i := 0; i < 3; i++ {
		r[i] = a[i] - b[i]
	}
}

func (r *Vec4) Scale(v *Vec4, s float32) {
	for i := 0; i < 4; i++ {
		r[i] = v[i] * s
	}
}

func (v *Vec4) Len() float32 {
	return sqrtf(Vec4MultInner(v, v))
}

func (r *Vec4) Norm(v *Vec4) {
	var k float32 = 1.0 / v.Len()
	r.Scale(v, k)
}

func (r *Vec4) Min(a, b *Vec4) {
	for i := 0; i < 4; i++ {
		if a[i] < b[i] {
			r[i] = a[i]
		} else {
			r[i] = b[i]
		}
	}
}

func (r *Vec4) Max(a, b *Vec4) {
	for i := 0; i < 4; i++ {
		if a[i] > b[i] {
			r[i] = a[i]
		} else {
			r[i] = b[i]
		}
	}
}

func Vec4MultInner(a, b *Vec4) (p float32) {
	for i := 0; i < 4; i++ {
		p += b[i] * a[i]
	}
	return p
}

func Vec4MultInner3(a, b *Vec4) (p float32) {
	for i := 0; i < 3; i++ {
		p += b[i] * a[i]
	}
	return p
}

func (r *Vec4) MultCross(a, b *Vec4) {
	r[0] = a[1]*b[2] - a[2]*b[1]
	r[1] = a[2]*b[0] - a[0]*b[2]
	r[2] = a[0]*b[1] - a[1]*b[0]
	r[3] = 1
}

func (r *Vec4) Reflect(v, n *Vec4) {
	var p float32 = 2 * Vec4MultInner(v, n)
	for i := 0; i < 4; i++ {
		r[i] = v[i] - p*n[i]
	}
}
