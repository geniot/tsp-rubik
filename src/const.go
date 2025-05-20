package main

import rl "github.com/gen2brain/raylib-go/raylib"

//var (
//	//go:embed media/*
//	mediaList embed.FS
//)

// faces
const (
	FRONT = iota
	LEFT
	BACK
	RIGHT
	TOP
	BOTTOM
)

// TSP button codes
const (
	noCode = iota
	upCode
	rightCode
	downCode
	leftCode
	xCode
	aCode
	bCode
	yCode
	l1Code
	l2Code
	r1Code
	r2Code
	selectCode
	menuCode
	startCode
)

const (
	G = iota
	R
	B
	O
	W
	Y
	LB
	BL
)

// colors
// https://www.schemecolor.com/rubik-cube-colors.php
var (
	black      = rl.Color{R: 0, G: 0, B: 0, A: 255}
	lightBlack = rl.Color{R: 77, G: 77, B: 77, A: 255}
	lightGray  = rl.Color{R: 211, G: 211, B: 211, A: 255}
	green      = rl.Color{R: 0, G: 155, B: 72, A: 255}
	red        = rl.Color{R: 185, G: 0, B: 0, A: 255}
	blue       = rl.Color{R: 0, G: 69, B: 173, A: 255}
	orange     = rl.Color{R: 255, G: 89, B: 0, A: 255}
	white      = rl.Color{R: 255, G: 255, B: 255, A: 255}
	yellow     = rl.Color{R: 255, G: 213, B: 0, A: 255}
	allColors  = map[int]rl.Color{
		LB: lightBlack,
		BL: black,
		G:  green,
		R:  red,
		B:  blue,
		O:  orange,
		W:  white,
		Y:  yellow,
	}
)
