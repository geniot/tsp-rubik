package main

import rl "github.com/gen2brain/raylib-go/raylib"

var (
	keysToRotationsMap = map[int32]int{
		rl.KeyZero:  RNone,
		rl.KeyOne:   RFront,
		rl.KeyTwo:   RFbMiddle,
		rl.KeyThree: RBack,
		rl.KeyFour:  RLeft,
		rl.KeyFive:  RLrMiddle,
		rl.KeySix:   RRight,
		rl.KeySeven: RTop,
		rl.KeyEight: RTbMiddle,
		rl.KeyNine:  RBottom,
	}
	buttonsToRotationsMap = map[int32]int{
		xCode:  RTop,
		bCode:  RBottom,
		yCode:  RLeft,
		aCode:  RRight,
		r1Code: RFront,
		r2Code: RBack,
	}
	x1, y1               float64 = 0, 0
	roundedX1, roundedY1 float64 = 0, 0
)

func (c *Cube) handleUserEvents() {
	for key, rotation := range keysToRotationsMap {
		if rl.IsKeyPressed(key) {
			if c.selectedRotation == rotation {
				c.selectedRotation = RNone
				c.isFaceSelectionModeOn = false
			} else {
				c.selectedRotation = rotation
				c.isFaceSelectionModeOn = true
			}
		}
	}
	for key, rotation := range buttonsToRotationsMap {
		if rl.IsGamepadButtonPressed(gamePadId, key) {
			if c.selectedRotation == rotation || c.selectedRotation == RNone || c.selectedRotation > RBottom {
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
	if rl.IsKeyPressed(rl.KeyQ) || rl.IsGamepadButtonPressed(gamePadId, menuCode) {
		c.application.currentSceneIndex = menuSceneKey
	}

	c.isShuffling = If(rl.IsKeyDown(rl.KeyS) || rl.IsGamepadButtonDown(gamePadId, l1Code), true, false)
	if c.isShuffling {
		c.isFaceSelectionModeOn = false
	}
	if c.isFaceSelectionModeOn {
		if rl.IsKeyDown(rl.KeyUp) || rl.IsGamepadButtonPressed(gamePadId, upCode) {
			rotateAny(c, c.selectedRotation, If(c.selectedRotation <= RBack, true, false), false)
		} else if rl.IsKeyDown(rl.KeyDown) || rl.IsGamepadButtonPressed(gamePadId, downCode) {
			rotateAny(c, c.selectedRotation, If(c.selectedRotation <= RBack, false, true), false)
		} else if rl.IsKeyDown(rl.KeyLeft) || rl.IsGamepadButtonPressed(gamePadId, leftCode) {
			rotateAny(c, c.selectedRotation, If(c.selectedRotation <= RBack, true, false), c.selectedRotation == RLeft || c.selectedRotation == RRight)
		} else if rl.IsKeyDown(rl.KeyRight) || rl.IsGamepadButtonPressed(gamePadId, rightCode) {
			rotateAny(c, c.selectedRotation, If(c.selectedRotation <= RBack, false, true), c.selectedRotation == RLeft || c.selectedRotation == RRight)
		}
	} else {
		if rl.IsKeyDown(rl.KeyUp) || (rl.IsGamepadButtonPressed(gamePadId, upCode) || isLeftJoystick(upCode)) {
			rotateAny(c, IfInt(rl.IsKeyDown(rl.KeyLeftControl) || rl.IsGamepadButtonDown(gamePadId, startCode), RAllBottom, RAllFront), false, rl.IsKeyDown(rl.KeyLeftControl) || rl.IsGamepadButtonDown(gamePadId, startCode))
		} else if rl.IsKeyDown(rl.KeyDown) || (rl.IsGamepadButtonPressed(gamePadId, downCode) || isLeftJoystick(downCode)) {
			rotateAny(c, IfInt(rl.IsKeyDown(rl.KeyLeftControl) || rl.IsGamepadButtonDown(gamePadId, startCode), RAllBottom, RAllBack), true, rl.IsKeyDown(rl.KeyLeftControl) || rl.IsGamepadButtonDown(gamePadId, startCode))
		} else if rl.IsKeyDown(rl.KeyLeft) || (rl.IsGamepadButtonPressed(gamePadId, leftCode) || isLeftJoystick(leftCode)) {
			rotateAny(c, IfInt(rl.IsKeyDown(rl.KeyLeftControl) || rl.IsGamepadButtonDown(gamePadId, startCode), RAllLeft, RAllLeft), false, false)
		} else if rl.IsKeyDown(rl.KeyRight) || (rl.IsGamepadButtonPressed(gamePadId, rightCode) || isLeftJoystick(rightCode)) {
			rotateAny(c, IfInt(rl.IsKeyDown(rl.KeyLeftControl) || rl.IsGamepadButtonDown(gamePadId, startCode), RAllRight, RAllRight), true, false)
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
