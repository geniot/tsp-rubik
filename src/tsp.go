//go:build tsp

package main

import rl "github.com/gen2brain/raylib-go/raylib"

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

var (
	x1, y1               float64 = 0, 0
	roundedX1, roundedY1 float64 = 0, 0
)

var (
	leftRightSequence        = []int{R_LEFT, R_LR_MIDDLE, R_RIGHT, R_FRONT, R_FB_MIDDLE, R_BACK}
	upDownSequence           = []int{R_TOP, R_TB_MIDDLE, R_BOTTOM}
	leftRightSequencePointer = -1
	upDownSequencePointer    = -1
)

func drawHelp() {
	rl.DrawRectangle(helpPadding, winHeight-helpHeight-helpPadding, helpWidth, helpHeight, rl.Fade(rl.SkyBlue, 0.5))
	rl.DrawRectangleLines(helpPadding, winHeight-helpHeight-helpPadding, helpWidth, helpHeight, rl.Blue)

	rl.DrawText("TrimUI Smart Pro controls:", helpPadding*2, winHeight-helpHeight-helpPadding, helpFontSize, rl.Black)
	rl.DrawText("use arrow joystick to select", helpPadding*2, winHeight-helpHeight-helpPadding+helpLineHeight*1, helpFontSize, rl.DarkGray)
	rl.DrawText("use left analogue joystick", helpPadding*2, winHeight-helpHeight-helpPadding+helpLineHeight*2, helpFontSize, rl.DarkGray)
	rl.DrawText("A/B to rotate, Y to deselect", helpPadding*2, winHeight-helpHeight-helpPadding+helpLineHeight*3, helpFontSize, rl.DarkGray)
	rl.DrawText("hold X to shuffle, start+up/down", helpPadding*2, winHeight-helpHeight-helpPadding+helpLineHeight*4, helpFontSize, rl.DarkGray)
	rl.DrawText("menu+start -> exit", helpPadding*6, winHeight-helpHeight-helpPadding+helpLineHeight*5, helpFontSize, rl.DarkGray)
}

func isShuffling() bool {
	return rl.IsGamepadButtonDown(gamePadId, xCode)
}

func handleUserEvents(c *Cube) {
	if rl.IsGamepadButtonPressed(gamePadId, yCode) {
		c.selectedRotation = R_NONE
	}
	if rl.IsGamepadButtonPressed(gamePadId, leftCode) || rl.IsGamepadButtonPressed(gamePadId, rightCode) {
		leftRightSequencePointer += If(rl.IsGamepadButtonDown(gamePadId, leftCode), -1, 1)
		if leftRightSequencePointer >= len(leftRightSequence) {
			leftRightSequencePointer = 0
		}
		if leftRightSequencePointer < 0 {
			leftRightSequencePointer = len(leftRightSequence) - 1
		}
		c.selectedRotation = leftRightSequence[leftRightSequencePointer]
	}
	if rl.IsGamepadButtonPressed(gamePadId, upCode) || rl.IsGamepadButtonPressed(gamePadId, downCode) {
		upDownSequencePointer += If(rl.IsGamepadButtonDown(gamePadId, downCode), 1, -1)
		if upDownSequencePointer >= len(upDownSequence) {
			upDownSequencePointer = 0
		}
		if upDownSequencePointer < 0 {
			upDownSequencePointer = len(upDownSequence) - 1
		}
		c.selectedRotation = upDownSequence[upDownSequencePointer]
	}
	if isLeftJoystick(upCode) || rl.IsGamepadButtonPressed(gamePadId, aCode) {
		c.angle = 90
		c.selectedRotation = If(c.selectedRotation == R_NONE, If(rl.IsGamepadButtonDown(gamePadId, startCode), R_ALL_TOP, R_ALL_FRONT), c.selectedRotation)
		c.isForward = If(c.selectedRotation <= R_BACK, true, false)
		c.isForward = If(rl.IsGamepadButtonDown(gamePadId, startCode), !c.isForward, c.isForward)
	}
	if isLeftJoystick(downCode) || rl.IsGamepadButtonPressed(gamePadId, bCode) {
		c.angle = 90
		c.selectedRotation = If(c.selectedRotation == R_NONE, If(rl.IsGamepadButtonDown(gamePadId, startCode), R_ALL_BOTTOM, R_ALL_BACK), c.selectedRotation)
		c.isForward = If(c.selectedRotation <= R_BACK, false, true)
		c.isForward = If(rl.IsGamepadButtonDown(gamePadId, startCode), !c.isForward, c.isForward)
	}
	if isLeftJoystick(leftCode) || rl.IsGamepadButtonPressed(gamePadId, aCode) {
		c.angle = 90
		c.selectedRotation = If(c.selectedRotation == R_NONE, R_ALL_LEFT, c.selectedRotation)
		c.isForward = If(c.selectedRotation <= R_BACK, true, false)
	}
	if isLeftJoystick(rightCode) || rl.IsGamepadButtonPressed(gamePadId, bCode) {
		c.angle = 90
		c.selectedRotation = If(c.selectedRotation == R_NONE, R_ALL_RIGHT, c.selectedRotation)
		c.isForward = If(c.selectedRotation <= R_BACK, false, true)
	}
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
