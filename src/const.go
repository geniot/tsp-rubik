package main

import rl "github.com/gen2brain/raylib-go/raylib"

//var (
//	//go:embed media/*
//	mediaList embed.FS
//)

// TSP button codes
const (
	upCode     = 1
	rightCode  = 2
	downCode   = 3
	leftCode   = 4
	xCode      = 5
	aCode      = 6
	bCode      = 7
	yCode      = 8
	l1Code     = 9
	l2Code     = 10
	r1Code     = 11
	r2Code     = 12
	selectCode = 13
	menuCode   = 14
	startCode  = 15
)

const (
	G = iota
	R
	B
	O
	W
	Y
	BL
)

// https://www.schemecolor.com/rubik-cube-colors.php
var (
	black      = rl.Color{R: 0, G: 0, B: 0, A: 255}
	lightBlack = rl.Color{R: 77, G: 77, B: 77, A: 255}
	green      = rl.Color{R: 0, G: 155, B: 72, A: 255}
	red        = rl.Color{R: 185, G: 0, B: 0, A: 255}
	blue       = rl.Color{R: 0, G: 69, B: 173, A: 255}
	orange     = rl.Color{R: 255, G: 89, B: 0, A: 255}
	white      = rl.Color{R: 255, G: 255, B: 255, A: 255}
	yellow     = rl.Color{R: 255, G: 213, B: 0, A: 255}
	allColors  = map[int]rl.Color{
		BL: lightBlack,
		G:  green,
		R:  red,
		B:  blue,
		O:  orange,
		W:  white,
		Y:  yellow,
	}
)

var (
	colorTextures         = make(map[int]rl.Texture2D)
	selectedColorTextures = make(map[int]rl.Texture2D)
)
