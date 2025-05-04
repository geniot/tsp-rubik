package main

import "math"

func sqrtf(v float32) float32 {
	// implemented in assembly on some systems
	return float32(math.Sqrt(float64(v)))
}

func sinf(v float32) float32 {
	// implemented in assembly on some systems
	return float32(math.Sin(float64(v)))
}

func cosf(v float32) float32 {
	// implemented in assembly on some systems
	return float32(math.Cos(float64(v)))
}

func tanf(v float32) float32 {
	// implemented in assembly on some systems
	return float32(math.Tan(float64(v)))
}

// DegreesToRadians converts degrees to radians.
func DegreesToRadians(angleDegrees float32) float32 {
	return angleDegrees * float32(math.Pi) / 180.0
}

// RadiansToDegrees converts radians to degrees.
func RadiansToDegrees(angleRadians float32) float32 {
	return angleRadians * 180.0 / float32(math.Pi)
}
