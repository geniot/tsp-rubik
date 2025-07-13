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
	buttonsToRotationsMap = map[int32]int{
		xCode:  R_TOP,
		bCode:  R_BOTTOM,
		yCode:  R_LEFT,
		aCode:  R_RIGHT,
		r1Code: R_FRONT,
		r2Code: R_BACK,
	}
	x1, y1               float64 = 0, 0
	roundedX1, roundedY1 float64 = 0, 0
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
	for key, rotation := range buttonsToRotationsMap {
		if rl.IsGamepadButtonPressed(gamePadId, key) {
			if c.selectedRotation == rotation || c.selectedRotation == R_NONE || c.selectedRotation > R_BOTTOM {
				c.isFaceSelectionModeOn = !c.isFaceSelectionModeOn
			}
			if c.isFaceSelectionModeOn {
				c.selectedRotation = rotation
			}
		}
	}
	//if rl.IsKeyPressed(rl.KeyZ) { //can be used for testing single shuffling
	//	c.Shuffle(1)
	//}
	c.isShuffling = If(rl.IsKeyDown(rl.KeyS) || rl.IsGamepadButtonDown(gamePadId, l1Code), true, false)
	if c.isShuffling {
		c.isFaceSelectionModeOn = false
	}
	if c.isFaceSelectionModeOn {
		if rl.IsKeyDown(rl.KeyUp) || rl.IsGamepadButtonPressed(gamePadId, upCode) {
			rotateAny(c, c.selectedRotation, If(c.selectedRotation <= R_BACK, true, false), false)
		} else if rl.IsKeyDown(rl.KeyDown) || rl.IsGamepadButtonPressed(gamePadId, downCode) {
			rotateAny(c, c.selectedRotation, If(c.selectedRotation <= R_BACK, false, true), false)
		} else if rl.IsKeyDown(rl.KeyLeft) || rl.IsGamepadButtonPressed(gamePadId, leftCode) {
			rotateAny(c, c.selectedRotation, If(c.selectedRotation <= R_BACK, true, false), c.selectedRotation == R_LEFT || c.selectedRotation == R_RIGHT)
		} else if rl.IsKeyDown(rl.KeyRight) || rl.IsGamepadButtonPressed(gamePadId, rightCode) {
			rotateAny(c, c.selectedRotation, If(c.selectedRotation <= R_BACK, false, true), c.selectedRotation == R_LEFT || c.selectedRotation == R_RIGHT)
		}
	} else {
		if rl.IsKeyDown(rl.KeyUp) || (rl.IsGamepadButtonPressed(gamePadId, upCode) || isLeftJoystick(upCode)) {
			rotateAny(c, If(rl.IsKeyDown(rl.KeyLeftControl) || rl.IsGamepadButtonDown(gamePadId, startCode), R_ALL_BOTTOM, R_ALL_FRONT), false, rl.IsKeyDown(rl.KeyLeftControl) || rl.IsGamepadButtonDown(gamePadId, startCode))
		} else if rl.IsKeyDown(rl.KeyDown) || (rl.IsGamepadButtonPressed(gamePadId, downCode) || isLeftJoystick(downCode)) {
			rotateAny(c, If(rl.IsKeyDown(rl.KeyLeftControl) || rl.IsGamepadButtonDown(gamePadId, startCode), R_ALL_BOTTOM, R_ALL_BACK), true, rl.IsKeyDown(rl.KeyLeftControl) || rl.IsGamepadButtonDown(gamePadId, startCode))
		} else if rl.IsKeyDown(rl.KeyLeft) || (rl.IsGamepadButtonPressed(gamePadId, leftCode) || isLeftJoystick(leftCode)) {
			rotateAny(c, If(rl.IsKeyDown(rl.KeyLeftControl) || rl.IsGamepadButtonDown(gamePadId, startCode), R_ALL_LEFT, R_ALL_LEFT), false, false)
		} else if rl.IsKeyDown(rl.KeyRight) || (rl.IsGamepadButtonPressed(gamePadId, rightCode) || isLeftJoystick(rightCode)) {
			rotateAny(c, If(rl.IsKeyDown(rl.KeyLeftControl) || rl.IsGamepadButtonDown(gamePadId, startCode), R_ALL_RIGHT, R_ALL_RIGHT), true, false)
		}
	}
}

func rotateAny(c *Cube, rotation int, isForward bool, shouldInverse bool) {
	c.angle = 90
	c.selectedRotation = rotation
	c.isForward = If(shouldInverse, !isForward, isForward)
}

func isLeftJoystick(code int) bool {
	//joysticks
	x1 = float64(rl.GetGamepadAxisMovement(gamePadId, rl.GamepadAxisLeftX))
	y1 = float64(rl.GetGamepadAxisMovement(gamePadId, rl.GamepadAxisLeftY))

	roundedX1 = toFixed(x1, 3)
	roundedY1 = toFixed(y1, 3)

	if code == upCode && roundedY1 < -0.5 {
		return true
	}
	if code == downCode && roundedY1 > 0.5 {
		return true
	}
	if code == rightCode && roundedX1 > 0.5 {
		return true
	}
	if code == leftCode && roundedX1 < -0.5 {
		return true
	}
	return false
}
