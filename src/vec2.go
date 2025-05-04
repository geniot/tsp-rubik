package main

type Vec2 [2]float32

func (r *Vec2) Add(a, b *Vec2) {
	for i := 0; i < 2; i++ {
		r[i] = a[i] + b[i]
	}
}

func (r *Vec2) Sub(a, b *Vec2) {
	for i := 0; i < 2; i++ {
		r[i] = a[i] - b[i]
	}
}

func (r *Vec2) Scale(v *Vec2, s float32) {
	for i := 0; i < 2; i++ {
		r[i] = v[i] * s
	}
}

func (v *Vec2) Len() float32 {
	return sqrtf(Vec2MultInner(v, v))
}

func (r *Vec2) Norm(v *Vec2) {
	var k float32 = 1.0 / v.Len()
	r.Scale(v, k)
}

func (r *Vec2) Min(a, b *Vec2) {
	for i := 0; i < 2; i++ {
		if a[i] < b[i] {
			r[i] = a[i]
		} else {
			r[i] = b[i]
		}
	}
}

func (r *Vec2) Max(a, b *Vec2) {
	for i := 0; i < 2; i++ {
		if a[i] > b[i] {
			r[i] = a[i]
		} else {
			r[i] = b[i]
		}
	}
}

func Vec2MultInner(a, b *Vec2) (p float32) {
	for i := 0; i < 2; i++ {
		p += b[i] * a[i]
	}
	return p
}
