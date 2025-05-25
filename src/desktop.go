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

func drawHelp() {
	rl.DrawRectangle(helpPadding, winHeight-helpHeight-helpPadding, helpWidth, helpHeight, rl.Fade(rl.SkyBlue, 0.5))
	rl.DrawRectangleLines(helpPadding, winHeight-helpHeight-helpPadding, helpWidth, helpHeight, rl.Blue)

	rl.DrawText("Desktop controls:", helpPadding*2, winHeight-helpHeight-helpPadding, helpFontSize, rl.Black)
	rl.DrawText("use arrow keys to rotate", helpPadding*2, winHeight-helpHeight-helpPadding+helpLineHeight*1, helpFontSize, rl.DarkGray)
	rl.DrawText("1-9 to (de)select faces", helpPadding*2, winHeight-helpHeight-helpPadding+helpLineHeight*2, helpFontSize, rl.DarkGray)
	rl.DrawText("hold 'S' to shuffle", helpPadding*2, winHeight-helpHeight-helpPadding+helpLineHeight*3, helpFontSize, rl.DarkGray)
	rl.DrawText("'Left Control' (hold) + Up/Down", helpPadding*2, winHeight-helpHeight-helpPadding+helpLineHeight*4, helpFontSize, rl.DarkGray)
	rl.DrawText(" - rotate around the Z-axis", helpPadding*6, winHeight-helpHeight-helpPadding+helpLineHeight*5, helpFontSize, rl.DarkGray)
}

func isShuffling() bool {
	return rl.IsKeyDown(rl.KeyS)
}

func handleUserEvents(c *Cube) {
	for key, rotation := range keysToRotationsMap {
		if rl.IsKeyPressed(key) {
			if c.selectedRotation == rotation {
				c.selectedRotation = R_NONE
			} else {
				c.selectedRotation = rotation
			}
		}
	}
	if rl.IsKeyDown(rl.KeyUp) {
		c.angle = 90
		c.selectedRotation = If(c.selectedRotation == R_NONE, If(rl.IsKeyDown(rl.KeyLeftControl), R_ALL_TOP, R_ALL_FRONT), c.selectedRotation)
		c.isForward = If(c.selectedRotation <= R_BACK, true, false)
		c.isForward = If(rl.IsKeyDown(rl.KeyLeftControl), !c.isForward, c.isForward)
	}
	if rl.IsKeyDown(rl.KeyDown) {
		c.angle = 90
		c.selectedRotation = If(c.selectedRotation == R_NONE, If(rl.IsKeyDown(rl.KeyLeftControl), R_ALL_BOTTOM, R_ALL_BACK), c.selectedRotation)
		c.isForward = If(c.selectedRotation <= R_BACK, false, true)
		c.isForward = If(rl.IsKeyDown(rl.KeyLeftControl), !c.isForward, c.isForward)
	}
	if rl.IsKeyDown(rl.KeyLeft) {
		c.angle = 90
		c.selectedRotation = If(c.selectedRotation == R_NONE, R_ALL_LEFT, c.selectedRotation)
		c.isForward = If(c.selectedRotation <= R_BACK, true, false)
	}
	if rl.IsKeyDown(rl.KeyRight) {
		c.angle = 90
		c.selectedRotation = If(c.selectedRotation == R_NONE, R_ALL_RIGHT, c.selectedRotation)
		c.isForward = If(c.selectedRotation <= R_BACK, false, true)
	}
}

func shouldExit() bool {
	return rl.IsKeyPressed(rl.KeyEscape)
}
