//go:build desktop

package main

import rl "github.com/gen2brain/raylib-go/raylib"

var (
	keysToRotationsMap = map[int32]int{
		rl.KeyZero:  R_NONE,
		rl.KeyOne:   R_FRONT,
		rl.KeyTwo:   R_FB_MIDDLE,
		rl.KeyThree: R_BACK,
		rl.KeyFour:  R_LEFT,
		rl.KeyFive:  R_LR_MIDDLE,
		rl.KeySix:   R_RIGHT,
		rl.KeySeven: R_TOP,
		rl.KeyEight: R_TB_MIDDLE,
		rl.KeyNine:  R_BOTTOM,
	}
)

func handleUserEvents(c *Cube) {
	for key, rotation := range keysToRotationsMap {
		if rl.IsKeyPressed(key) {
			if c.selectedRotation == rotation {
				c.selectedRotation = R_NONE
				c.isFaceSelectionModeOn = false
			} else {
				c.selectedRotation = rotation
				c.isFaceSelectionModeOn = true
			}
		}
	}
	c.isShuffling = If(rl.IsKeyDown(rl.KeyS), true, false)
	if c.isShuffling {
		c.isFaceSelectionModeOn = false
	}
	if c.isFaceSelectionModeOn {
		if rl.IsKeyDown(rl.KeyUp) {
			rotateAny(c, c.selectedRotation, c.selectedRotation, If(c.selectedRotation <= R_BACK, true, false), false)
		} else if rl.IsKeyDown(rl.KeyDown) {
			rotateAny(c, c.selectedRotation, c.selectedRotation, If(c.selectedRotation <= R_BACK, false, true), false)
		} else if rl.IsKeyDown(rl.KeyLeft) {
			rotateAny(c, c.selectedRotation, c.selectedRotation, If(c.selectedRotation <= R_BACK, true, false), false)
		} else if rl.IsKeyDown(rl.KeyRight) {
			rotateAny(c, c.selectedRotation, c.selectedRotation, If(c.selectedRotation <= R_BACK, false, true), false)
		}
	} else {
		if rl.IsKeyDown(rl.KeyUp) {
			rotateAny(c, R_ALL_BOTTOM, R_ALL_FRONT, false, rl.IsKeyDown(rl.KeyLeftControl))
		} else if rl.IsKeyDown(rl.KeyDown) {
			rotateAny(c, R_ALL_BOTTOM, R_ALL_BACK, true, rl.IsKeyDown(rl.KeyLeftControl))
		} else if rl.IsKeyDown(rl.KeyLeft) {
			rotateAny(c, R_ALL_LEFT, R_ALL_LEFT, false, false)
		} else if rl.IsKeyDown(rl.KeyRight) {
			rotateAny(c, R_ALL_RIGHT, R_ALL_RIGHT, true, false)
		}
	}
}

func rotateAny(c *Cube, rotation1 int, rotation2 int, isForward bool, shouldInverse bool) {
	c.angle = 90
	c.selectedRotation = If(rl.IsKeyDown(rl.KeyLeftControl), rotation1, rotation2)
	c.isForward = If(shouldInverse, !isForward, isForward)
}

func shouldExit() bool {
	return rl.IsKeyPressed(rl.KeyEscape)
}
