package main

type Mat4x4 [4]Vec4

func (m *Mat4x4) Fill(d float32) {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			m[i][j] = d
		}
	}
}

func (m *Mat4x4) Identity() {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if i == j {
				m[i][j] = 1
			} else {
				m[i][j] = 0
			}
		}
	}
}

func (m *Mat4x4) Dup(n *Mat4x4) {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			m[i][j] = n[i][j]
		}
	}
}

func (r *Vec4) Mat4x4Row(m *Mat4x4, i int) {
	for k := 0; k < 4; k++ {
		r[k] = m[k][i]
	}
}

func (r *Vec4) Mat4x4Col(m *Mat4x4, i int) {
	for k := 0; k < 4; k++ {
		r[k] = m[i][k]
	}
}

func (m *Mat4x4) Transpose(n *Mat4x4) {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			m[i][j] = n[j][i]
		}
	}
}

func (m *Mat4x4) Add(a, b *Mat4x4) {
	for i := 0; i < 4; i++ {
		m[i].Add(&a[i], &b[i])
	}
}

func (m *Mat4x4) Sub(a, b *Mat4x4) {
	for i := 0; i < 4; i++ {
		m[i].Sub(&a[i], &b[i])
	}
}

func (m *Mat4x4) Scale(a *Mat4x4, k float32) {
	for i := 0; i < 4; i++ {
		m[i].Scale(&a[i], k)
	}
}

func (m *Mat4x4) ScaleAniso(a *Mat4x4, x, y, z float32) {
	m[0].Scale(&a[0], x)
	m[1].Scale(&a[1], y)
	m[2].Scale(&a[2], z)
	for i := 0; i < 4; i++ {
		m[3][i] = a[3][i]
	}
}

func (m *Mat4x4) Mult(a, b *Mat4x4) {
	var temp = new(Mat4x4)

	for c := 0; c < 4; c++ {
		for r := 0; r < 4; r++ {
			temp[c][r] = 0
			for k := 0; k < 4; k++ {
				temp[c][r] += a[k][r] * b[c][k]
			}
		}
	}

	m.Dup(temp)
}

func (r *Vec4) Mat4x4MultVec4(m *Mat4x4, v Vec4) {
	for j := 0; j < 4; j++ {
		r[j] = 0
		for i := 0; i < 4; i++ {
			r[j] += m[i][j] * v[i]
		}
	}
}

func (m *Mat4x4) Translate(x, y, z float32) {
	m.Identity()
	m[3][0] = x
	m[3][1] = y
	m[3][2] = z
}

func (m *Mat4x4) TranslateInPlace(x, y, z float32) {
	var t = &Vec4{x, y, z, 0}
	var r = new(Vec4)
	for i := 0; i < 4; i++ {
		r.Mat4x4Row(m, i)
		m[3][i] += Vec4MultInner(r, t)
	}
}

func (m *Mat4x4) FromVec3MultOuter(a, b *Vec3) {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if i < 3 && j < 3 {
				m[i][j] = a[i] * b[j]
			} else {
				m[i][j] = 0
			}
		}
	}
}

func (r *Mat4x4) Rotate(m *Mat4x4, x, y, z, angle float32) {
	var s = sinf(angle)
	var c = cosf(angle)
	var u = &Vec3{x, y, z}

	if u.Len() > 1e-4 {
		u.Norm(u)
		var T = new(Mat4x4)
		T.FromVec3MultOuter(u, u)

		var S = &Mat4x4{
			{0, u[2], -u[1], 0},
			{-u[2], 0, u[0], 0},
			{u[1], -u[0], 0, 0},
			{0, 0, 0, 0},
		}
		S.Scale(S, s)

		var C = new(Mat4x4)
		C.Identity()
		C.Sub(C, T)

		C.Scale(C, c)

		T.Add(T, C)
		T.Add(T, S)

		T[3][3] = 1
		r.Mult(m, T)
	} else {
		r.Dup(m)
	}
}

func (q *Mat4x4) RotateX(m *Mat4x4, angle float32) {
	var s = sinf(angle)
	var c = cosf(angle)
	var R = &Mat4x4{
		{1, 0, 0, 0},
		{0, c, s, 0},
		{0, -s, c, 0},
		{0, 0, 0, 1},
	}
	q.Mult(m, R)
}

func (q *Mat4x4) RotateY(m *Mat4x4, angle float32) {
	var s = sinf(angle)
	var c = cosf(angle)
	var R = &Mat4x4{
		{c, 0, s, 0},
		{0, 1, 0, 0},
		{-s, 0, c, 0},
		{0, 0, 0, 1},
	}
	q.Mult(m, R)
}

func (q *Mat4x4) RotateZ(m *Mat4x4, angle float32) {
	var s = sinf(angle)
	var c = cosf(angle)
	var R = &Mat4x4{
		{c, s, 0, 0},
		{-s, c, 0, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 1},
	}
	q.Mult(m, R)
}

func (t *Mat4x4) Invert(m *Mat4x4) {
	var s = new([6]float32)
	s[0] = m[0][0]*m[1][1] - m[1][0]*m[0][1]
	s[1] = m[0][0]*m[1][2] - m[1][0]*m[0][2]
	s[2] = m[0][0]*m[1][3] - m[1][0]*m[0][3]
	s[3] = m[0][1]*m[1][2] - m[1][1]*m[0][2]
	s[4] = m[0][1]*m[1][3] - m[1][1]*m[0][3]
	s[5] = m[0][2]*m[1][3] - m[1][2]*m[0][3]

	var c = new([6]float32)
	c[0] = m[2][0]*m[3][1] - m[3][0]*m[2][1]
	c[1] = m[2][0]*m[3][2] - m[3][0]*m[2][2]
	c[2] = m[2][0]*m[3][3] - m[3][0]*m[2][3]
	c[3] = m[2][1]*m[3][2] - m[3][1]*m[2][2]
	c[4] = m[2][1]*m[3][3] - m[3][1]*m[2][3]
	c[5] = m[2][2]*m[3][3] - m[3][2]*m[2][3]

	// Assumes it is invertible
	var idet float32 = 1.0 / (s[0]*c[5] - s[1]*c[4] + s[2]*c[3] + s[3]*c[2] - s[4]*c[1] + s[5]*c[0])

	t[0][0] = (m[1][1]*c[5] - m[1][2]*c[4] + m[1][3]*c[3]) * idet
	t[0][1] = (-m[0][1]*c[5] + m[0][2]*c[4] - m[0][3]*c[3]) * idet
	t[0][2] = (m[3][1]*s[5] - m[3][2]*s[4] + m[3][3]*s[3]) * idet
	t[0][3] = (-m[2][1]*s[5] + m[2][2]*s[4] - m[2][3]*s[3]) * idet

	t[1][0] = (-m[1][0]*c[5] + m[1][2]*c[2] - m[1][3]*c[1]) * idet
	t[1][1] = (m[0][0]*c[5] - m[0][2]*c[2] + m[0][3]*c[1]) * idet
	t[1][2] = (-m[3][0]*s[5] + m[3][2]*s[2] - m[3][3]*s[1]) * idet
	t[1][3] = (m[2][0]*s[5] - m[2][2]*s[2] + m[2][3]*s[1]) * idet
	t[2][0] = (m[1][0]*c[4] - m[1][1]*c[2] + m[1][3]*c[0]) * idet
	t[2][1] = (-m[0][0]*c[4] + m[0][1]*c[2] - m[0][3]*c[0]) * idet
	t[2][2] = (m[3][0]*s[4] - m[3][1]*s[2] + m[3][3]*s[0]) * idet
	t[2][3] = (-m[2][0]*s[4] + m[2][1]*s[2] - m[2][3]*s[0]) * idet

	t[3][0] = (-m[1][0]*c[3] + m[1][1]*c[1] - m[1][2]*c[0]) * idet
	t[3][1] = (m[0][0]*c[3] - m[0][1]*c[1] + m[0][2]*c[0]) * idet
	t[3][2] = (-m[3][0]*s[3] + m[3][1]*s[1] - m[3][2]*s[0]) * idet
	t[3][3] = (m[2][0]*s[3] - m[2][1]*s[1] + m[2][2]*s[0]) * idet
}

func (r *Mat4x4) OrthoNormalize(m *Mat4x4) {
	r.Dup(m)
	var s float32
	var h = new(Vec3)
	r[2].Norm(&r[2])

	s = Vec4MultInner3(&r[1], &r[2])
	h.ScaleVec4(&r[2], s)
	r[1].SubVec3(&r[1], h)
	r[2].Norm(&r[2])

	s = Vec4MultInner3(&r[1], &r[2])
	h.ScaleVec4(&r[2], s)
	r[1].SubVec3(&r[1], h)
	r[1].Norm(&r[1])

	s = Vec4MultInner3(&r[0], &r[1])
	h.ScaleVec4(&r[1], s)
	r[0].SubVec3(&r[0], h)
	r[0].Norm(&r[0])
}

func (m *Mat4x4) Frustum(l, r, b, t, n, f float32) {
	m[0][0] = 2 * n / (r - l)
	m[0][1] = 0
	m[0][2] = 0
	m[0][3] = 0

	m[1][1] = 2 * n / (t - b)
	m[1][0] = 0
	m[1][2] = 0
	m[1][3] = 0

	m[2][0] = (r + l) / (r - l)
	m[2][1] = (t + b) / (t - b)
	m[2][2] = -(f + n) / (f - n)
	m[2][3] = -1

	m[3][2] = -2 * (f * n) / (f - n)
	m[3][0] = 0
	m[3][1] = 0
	m[3][3] = 0
}

func (m *Mat4x4) Ortho(l, r, b, t, n, f float32) {
	m[0][0] = 2 / (r - l)
	m[0][1] = 0
	m[0][2] = 0
	m[0][3] = 0

	m[1][1] = 2 / (t - b)
	m[1][0] = 0
	m[1][2] = 0
	m[1][3] = 0

	m[2][2] = -2 / (f - n)
	m[2][0] = 0
	m[2][1] = 0
	m[2][3] = 0

	m[3][0] = -(r + l) / (r - l)
	m[3][1] = -(t + b) / (t - b)
	m[3][2] = -(f + n) / (f - n)
	m[3][3] = 1
}

func (m *Mat4x4) Perspective(y_fov, aspect, n, f float32) {
	// NOTE: Degrees are an unhandy unit to work with.
	// linmath.go uses radians for everything!
	var a float32 = 1 / tanf(y_fov/2)

	m[0][0] = a / aspect
	m[0][1] = 0
	m[0][2] = 0
	m[0][3] = 0

	m[1][0] = 0
	m[1][1] = a
	m[1][2] = 0
	m[1][3] = 0

	m[2][0] = 0
	m[2][1] = 0
	m[2][2] = -((f + n) / (f - n))
	m[2][3] = -1

	m[3][0] = 0
	m[3][1] = 0
	m[3][2] = -((2 * f * n) / (f - n))
	m[3][3] = 0
}

func (m *Mat4x4) LookAt(eye, center, up *Vec3) {
	// Adapted from Android's OpenGL Matrix.java.
	// See the OpenGL GLUT documentation for gluLookAt for a description
	// of the algorithm. We implement it in a straightforward way:
	//
	// TODO: The negation of of can be spared by swapping the order of
	//       operands in the following cross products in the right way.
	var f = new(Vec3)
	f.Sub(center, eye)
	f.Norm(f)

	var s = new(Vec3)
	s.MultCross(f, up)
	s.Norm(s)

	var t = new(Vec3)
	t.MultCross(s, f)

	m[0][0] = s[0]
	m[0][1] = t[0]
	m[0][2] = -f[0]
	m[0][3] = 0

	m[1][0] = s[1]
	m[1][1] = t[1]
	m[1][2] = -f[1]
	m[1][3] = 0

	m[2][0] = s[2]
	m[2][1] = t[2]
	m[2][2] = -f[2]
	m[2][3] = 0

	m[3][0] = 0
	m[3][1] = 0
	m[3][2] = 0
	m[3][3] = 1

	m.TranslateInPlace(-eye[0], -eye[1], -eye[2])
}

func (r *Vec3) QuatMultVec3(q *Quat, v *Vec3) {
	// Method by Fabian 'ryg' Giessen (of Farbrausch)
	//   t = 2 * cross(q.xyz, v)
	//   v' = v + q.w * t + cross(q.xyz, t)
	var t = new(Vec3)
	var q_xyz = &Vec3{q[0], q[1], q[2]}
	var u = &Vec3{q[0], q[1], q[2]}

	t.MultCross(q_xyz, v)
	t.Scale(t, 2)

	u.MultCross(q_xyz, t)
	t.Scale(t, q[3])

	r.Add(v, t)
	r.Add(r, u)
}

func (r *Vec4) QuatMultVec4(q *Quat, v *Vec4) {
	// Method by Fabian 'ryg' Giessen (of Farbrausch)
	//   t = 2 * cross(q.xyz, v)
	//   v' = v + q.w * t + cross(q.xyz, t)
	var t = new(Vec4)
	var q_xyz = &Vec4{q[0], q[1], q[2]}
	var u = &Vec4{q[0], q[1], q[2]}

	t.MultCross(q_xyz, v)
	t.Scale(t, 2)

	u.MultCross(q_xyz, t)
	t.Scale(t, q[3])

	r.Add(v, t)
	r.Add(r, u)
}

func (m *Mat4x4) FromQuat(q *Quat) {
	var a float32 = q[3]
	var b float32 = q[0]
	var c float32 = q[1]
	var d float32 = q[2]
	var a2 float32 = a * a
	var b2 float32 = b * b
	var c2 float32 = c * c
	var d2 float32 = d * d

	m[0][0] = a2 + b2 - c2 - d2
	m[0][1] = 2 * (b*c + a*d)
	m[0][2] = 2 * (b*d - a*c)
	m[0][3] = 0

	m[1][0] = 2 * (b*c - a*d)
	m[1][1] = a2 - b2 + c2 - d2
	m[1][2] = 2 * (c*d + a*b)
	m[1][3] = 0

	m[2][0] = 2 * (b*d + a*c)
	m[2][1] = 2 * (c*d - a*b)
	m[2][2] = a2 - b2 - c2 + d2
	m[2][3] = 0

	m[3][0] = 0
	m[3][1] = 0
	m[3][2] = 0
	m[3][3] = 1
}

func (r *Mat4x4) MultQuat(m *Mat4x4, q *Quat) {
	// XXX: The way this is written only works for othogonal matrices.
	// TODO: Take care of non-orthogonal case.
	r[0].QuatMultVec4(q, &m[0])
	r[1].QuatMultVec4(q, &m[1])
	r[2].QuatMultVec4(q, &m[2])

	r[3][0] = 0
	r[3][1] = 0
	r[3][2] = 0
	r[3][3] = 1
}

func (q *Quat) FromMat4x4(m *Mat4x4) {
	var r float32
	var p = []int{0, 1, 2, 0, 1}
	var idx int

	for i := 0; i < 3; i++ {
		var m float32 = m[i][i]
		if m < r {
			continue
		}
		m = r
		idx = i
	}

	p = p[idx:] // reslice p starting from idx
	r = sqrtf(1 + m[p[0]][p[0]] - m[p[1]][p[1]] - m[p[2]][p[2]])

	if r < 1e-6 {
		q[0] = 1
		q[1] = 0
		q[2] = 0
		q[3] = 0
		return
	}

	q[0] = r / 2
	q[1] = (m[p[0]][p[1]] - m[p[1]][p[0]]) / (2 * r)
	q[2] = (m[p[2]][p[0]] - m[p[0]][p[2]]) / (2 * r)
	q[3] = (m[p[2]][p[1]] - m[p[1]][p[2]]) / (2 * r)
}
