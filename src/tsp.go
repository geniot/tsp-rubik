//go:build tsp

package main

import rl "github.com/gen2brain/raylib-go/raylib"

var (
	x1, y1                   float64 = 0, 0
	roundedX1, roundedY1     float64 = 0, 0
	leftRightSequence                = []int{R_LEFT, R_LR_MIDDLE, R_RIGHT, R_FRONT, R_FB_MIDDLE, R_BACK}
	upDownSequence                   = []int{R_TOP, R_TB_MIDDLE, R_BOTTOM}
	leftRightSequencePointer         = -1
	upDownSequencePointer            = -1
	lastFaceSelected                 = R_FRONT
)

func handleUserEvents(c *Cube) {
	if rl.IsGamepadButtonPressed(gamePadId, yCode) {
		c.isFaceSelectionModeOn = !c.isFaceSelectionModeOn
		if c.isFaceSelectionModeOn {
			c.selectedRotation = lastFaceSelected
		} else {
			c.selectedRotation = R_NONE
		}
	}
	c.isShuffling = If(rl.IsGamepadButtonDown(gamePadId, xCode), true, false)
	if c.isShuffling {
		c.isFaceSelectionModeOn = false
	}

	if c.isFaceSelectionModeOn {
		if rl.IsGamepadButtonPressed(gamePadId, aCode) {
			rotateAny(c, c.selectedRotation, c.selectedRotation, true, false)
		}
		if rl.IsGamepadButtonPressed(gamePadId, bCode) {
			rotateAny(c, c.selectedRotation, c.selectedRotation, false, false)
		}
		if rl.IsGamepadButtonPressed(gamePadId, rightCode) {
			leftRightSequencePointer = getNextSelection(leftRightSequence, leftRightSequencePointer, 1)
			c.selectedRotation = leftRightSequence[leftRightSequencePointer]
		}
		if rl.IsGamepadButtonPressed(gamePadId, leftCode) {
			leftRightSequencePointer = getNextSelection(leftRightSequence, leftRightSequencePointer, -1)
			c.selectedRotation = leftRightSequence[leftRightSequencePointer]
		}
		if rl.IsGamepadButtonPressed(gamePadId, upCode) {
			upDownSequencePointer = getNextSelection(upDownSequence, upDownSequencePointer, -1)
			c.selectedRotation = upDownSequence[upDownSequencePointer]
		}
		if rl.IsGamepadButtonPressed(gamePadId, downCode) {
			upDownSequencePointer = getNextSelection(upDownSequence, upDownSequencePointer, 1)
			c.selectedRotation = upDownSequence[upDownSequencePointer]
		}
	} else {
		if rl.IsGamepadButtonPressed(gamePadId, rightCode) || isLeftJoystick(rightCode) {
			rotateAny(c, R_ALL_RIGHT, R_ALL_RIGHT, true, false)
		}
		if rl.IsGamepadButtonPressed(gamePadId, leftCode) || isLeftJoystick(leftCode) {
			rotateAny(c, R_ALL_LEFT, R_ALL_LEFT, false, false)
		}
		if rl.IsGamepadButtonPressed(gamePadId, upCode) || isLeftJoystick(upCode) {
			rotateAny(c, R_ALL_BOTTOM, R_ALL_FRONT, false, rl.IsGamepadButtonDown(gamePadId, startCode))
		}
		if rl.IsGamepadButtonPressed(gamePadId, downCode) || isLeftJoystick(downCode) {
			rotateAny(c, R_ALL_TOP, R_ALL_BACK, true, rl.IsGamepadButtonDown(gamePadId, startCode))
		}
	}
}

func rotateAny(c *Cube, rotation1 int, rotation2 int, isForward bool, shouldInverse bool) {
	c.angle = 90
	c.selectedRotation = If(rl.IsGamepadButtonDown(gamePadId, startCode), rotation1, rotation2)
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

func shouldExit() bool {
	return rl.IsGamepadButtonDown(gamePadId, menuCode) && rl.IsGamepadButtonDown(gamePadId, startCode)
}

func getNextSelection(sequence []int, pointer int, increment int) int {
	pointer += increment
	if pointer >= len(sequence) {
		pointer = 0
	}
	if pointer < 0 {
		pointer = len(sequence) - 1
	}
	return pointer
}
