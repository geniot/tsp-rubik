package main

import (
	"sort"

	rl "github.com/gen2brain/raylib-go/raylib"
)

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
		l1Code: RFront,
		r1Code: RBack,
		l2Code: RFront,
		r2Code: RBack,
	}
	x1, y1               float64 = 0, 0
	roundedX1, roundedY1 float64 = 0, 0
)

func (c *Cube) handleMouseEvents() {
	//mouse
	if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
		c.fromMouseSelectedPosition = rl.GetMousePosition()
		ray := rl.GetScreenToWorldRay(c.fromMouseSelectedPosition, *c.application.camera)
		hitFaces := c.getAllHitFaces(ray)
		sort.Slice(hitFaces, func(i, j int) bool {
			centerI := hitFaces[i].getCenter()
			centerJ := hitFaces[j].getCenter()
			return centerI.X+centerI.Y+centerI.Z > centerJ.X+centerJ.Y+centerJ.Z
		})
		if len(hitFaces) > 0 {
			c.mouseSelectedFace = hitFaces[0]
		} else {
			c.mouseSelectedFace = nil
		}
	}
	if rl.IsMouseButtonReleased(rl.MouseLeftButton) {
		c.toMouseSelectedPosition = rl.GetMousePosition()
		if c.mouseSelectedFace != nil && rl.Vector2Distance(c.fromMouseSelectedPosition, c.toMouseSelectedPosition) > 0 {
			angle := rl.Vector2LineAngle(c.fromMouseSelectedPosition, c.toMouseSelectedPosition) * rl.Rad2deg
			center := c.mouseSelectedFace.getCenter()
			centerX, centerY, centerZ := round32(center.X), round32(center.Y), round32(center.Z)
			c.isFaceSelectionModeOn = false

			//FRONT
			if centerZ == 3 && (isAround(angle, -45) || isAround(angle, 135)) {
				c.selectedRotation = IfInt(centerY == 2, RTop, IfInt(centerY == 0, RTbMiddle, RBottom))
				c.RotateAny(c.selectedRotation, If(isAround(angle, -45), true, false), false)
			}
			if centerZ == 3 && (isAround(angle, -90) || isAround(angle, 90)) {
				c.selectedRotation = IfInt(centerX == 2, RRight, IfInt(centerX == 0, RLrMiddle, RLeft))
				c.RotateAny(c.selectedRotation, If(isAround(angle, -90), true, false), false)
			}
			//RIGHT
			if centerX == 3 && (isAround(angle, 45) || isAround(angle, -135)) {
				c.selectedRotation = IfInt(centerY == 2, RTop, IfInt(centerY == 0, RTbMiddle, RBottom))
				c.RotateAny(c.selectedRotation, If(isAround(angle, 45), true, false), false)
			}
			if centerX == 3 && (isAround(angle, -90) || isAround(angle, 90)) {
				c.selectedRotation = IfInt(centerZ == 2, RFront, IfInt(centerZ == 0, RFbMiddle, RBack))
				c.RotateAny(c.selectedRotation, If(isAround(angle, 90), true, false), false)
			}
			//TOP
			if centerY == 3 && (isAround(angle, 135) || isAround(angle, -45)) {
				c.selectedRotation = IfInt(centerZ == 2, RFront, IfInt(centerZ == 0, RFbMiddle, RBack))
				c.RotateAny(c.selectedRotation, If(isAround(angle, 135), true, false), false)
			}
			if centerY == 3 && (isAround(angle, 45) || isAround(angle, -135)) {
				c.selectedRotation = IfInt(centerX == 2, RRight, IfInt(centerX == 0, RLrMiddle, RLeft))
				c.RotateAny(c.selectedRotation, If(isAround(angle, -135), true, false), false)
			}
		}
	}
}

func (c *Cube) handleUserEvents() {
	requestedRotation := RNone
	//keyboard
	for key, rotation := range keysToRotationsMap {
		if rl.IsKeyPressed(key) {
			requestedRotation = rotation
		}
	}
	for key, rotation := range buttonsToRotationsMap {
		if rl.IsGamepadButtonPressed(gamePadId, key) {
			requestedRotation = rotation
		}
	}
	//gamepad
	if requestedRotation != RNone {
		if requestedRotation > RBottom {
			c.isFaceSelectionModeOn = false
		} else {
			if c.selectedRotation == requestedRotation {
				c.isFaceSelectionModeOn = !c.isFaceSelectionModeOn
			} else {
				c.isFaceSelectionModeOn = true
			}
		}
		c.selectedRotation = requestedRotation
	}

	//if rl.IsKeyPressed(rl.KeyZ) { //can be used for testing single shuffling
	//	c.Shuffle(1)
	//}
	//if rl.IsKeyPressed(rl.KeyQ) || rl.IsGamepadButtonPressed(gamePadId, menuCode) {
	//	c.application.currentSceneIndex = menuSceneKey
	//}

	c.isShuffling = rl.IsKeyDown(rl.KeyS)
	if c.isShuffling {
		c.isFaceSelectionModeOn = false
	}
	if c.isFaceSelectionModeOn {
		if rl.IsKeyDown(rl.KeyUp) || rl.IsGamepadButtonDown(gamePadId, upCode) || isLeftJoystick(upCode) {
			c.RotateAny(c.selectedRotation, If(c.selectedRotation <= RBack, true, false), false)
		} else if rl.IsKeyDown(rl.KeyDown) || rl.IsGamepadButtonDown(gamePadId, downCode) || isLeftJoystick(downCode) {
			c.RotateAny(c.selectedRotation, If(c.selectedRotation <= RBack, false, true), false)
		} else if rl.IsKeyDown(rl.KeyLeft) || rl.IsGamepadButtonDown(gamePadId, leftCode) || isLeftJoystick(leftCode) {
			c.RotateAny(c.selectedRotation, If(c.selectedRotation <= RBack, true, false), c.selectedRotation == RLeft || c.selectedRotation == RRight)
		} else if rl.IsKeyDown(rl.KeyRight) || rl.IsGamepadButtonDown(gamePadId, rightCode) || isLeftJoystick(rightCode) {
			c.RotateAny(c.selectedRotation, If(c.selectedRotation <= RBack, false, true), c.selectedRotation == RLeft || c.selectedRotation == RRight)
		}
	} else {
		if rl.IsKeyDown(rl.KeyUp) || rl.IsGamepadButtonDown(gamePadId, upCode) || isLeftJoystick(upCode) {
			c.RotateAny(IfInt(rl.IsKeyDown(rl.KeyLeftControl) || rl.IsGamepadButtonDown(gamePadId, startCode), RAllBottom, RAllFront), false, rl.IsKeyDown(rl.KeyLeftControl) || rl.IsGamepadButtonDown(gamePadId, startCode))
		} else if rl.IsKeyDown(rl.KeyDown) || rl.IsGamepadButtonDown(gamePadId, downCode) || isLeftJoystick(downCode) {
			c.RotateAny(IfInt(rl.IsKeyDown(rl.KeyLeftControl) || rl.IsGamepadButtonDown(gamePadId, startCode), RAllBottom, RAllBack), true, rl.IsKeyDown(rl.KeyLeftControl) || rl.IsGamepadButtonDown(gamePadId, startCode))
		} else if rl.IsKeyDown(rl.KeyLeft) || rl.IsGamepadButtonDown(gamePadId, leftCode) || isLeftJoystick(leftCode) {
			c.RotateAny(IfInt(rl.IsKeyDown(rl.KeyLeftControl) || rl.IsGamepadButtonDown(gamePadId, startCode), RAllLeft, RAllLeft), false, false)
		} else if rl.IsKeyDown(rl.KeyRight) || rl.IsGamepadButtonDown(gamePadId, rightCode) || isLeftJoystick(rightCode) {
			c.RotateAny(IfInt(rl.IsKeyDown(rl.KeyLeftControl) || rl.IsGamepadButtonDown(gamePadId, startCode), RAllRight, RAllRight), true, false)
		}
	}
}

func (c *Cube) RotateAny(rotation int, isForward bool, shouldInverse bool) {
	c.angle = 90
	c.selectedRotation = rotation
	c.isForward = If(shouldInverse, !isForward, isForward)
}
