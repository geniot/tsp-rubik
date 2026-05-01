package main

import rl "github.com/gen2brain/raylib-go/raylib"

func isLeftJoystick(code int) bool {
	return isJoystick(code, rl.GamepadAxisLeftX, rl.GamepadAxisLeftY)
}

func isRightJoystick(code int) bool {
	return isJoystick(code, rl.GamepadAxisRightX, rl.GamepadAxisRightY)
}

func isJoystick(code int, joystickX int32, joystickY int32) bool {
	//joysticks
	x1 = float64(rl.GetGamepadAxisMovement(gamePadId, joystickX))
	y1 = float64(rl.GetGamepadAxisMovement(gamePadId, joystickY))

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
