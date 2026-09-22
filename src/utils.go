package main

import (
	gui "github.com/gen2brain/raylib-go/raygui"
)

func orPanic(err interface{}) {
	switch v := err.(type) {
	case error:
		if v != nil {
			panic(err)
		}
	case bool:
		if !v {
			panic("condition failed: != true")
		}
	}
}

func If[T any](cond bool, vTrue, vFalse T) T {
	if cond {
		return vTrue
	}
	return vFalse
}

func IfInt(cond bool, vTrue int, vFalse int) int {
	if cond {
		return vTrue
	}
	return vFalse
}

type TextStyle struct {
	size      gui.PropertyValue
	alignment gui.PropertyValue
	spacing   gui.PropertyValue
	padding   gui.PropertyValue
}

var (
	defaultTextStyle = TextStyle{
		size:      35,
		alignment: gui.TEXT_ALIGN_LEFT,
		spacing:   10,
		padding:   20,
	}
	singleLetterStyle = TextStyle{
		size:      35,
		alignment: gui.TEXT_ALIGN_CENTER,
		spacing:   0,
		padding:   0,
	}
	nextPrevStyle = TextStyle{
		size:      20,
		alignment: gui.TEXT_ALIGN_CENTER,
		spacing:   0,
		padding:   0,
	}
)

func setTextStyle(ts TextStyle) {
	gui.SetStyle(gui.DEFAULT, gui.TEXT_SIZE, ts.size)
	gui.SetStyle(gui.DEFAULT, gui.TEXT_SPACING, ts.spacing)
	gui.SetStyle(gui.DEFAULT, gui.TEXT_ALIGNMENT, ts.alignment)
	gui.SetStyle(gui.DEFAULT, gui.TEXT_PADDING, ts.padding)
	gui.SetStyle(gui.DEFAULT, gui.TEXT_ALIGNMENT_VERTICAL, gui.TEXT_ALIGN_CENTER)
	gui.SetStyle(gui.TEXTBOX, gui.TEXT_ALIGNMENT, gui.TEXT_ALIGN_LEFT)
}

const (
	DELTA = 35
)

func isAround(angle float32, i float32) bool {
	return angle >= i-DELTA && angle <= i+DELTA
}
