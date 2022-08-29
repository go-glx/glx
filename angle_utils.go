package glx

import "math"

const circleFullDeg = 360.0
const circleHalfDeg = 180.0
const circleFullRad = math.Pi * 2
const convDeg2Rad = math.Pi / circleHalfDeg
const convRad2Deg = circleHalfDeg / math.Pi
const floatRoundPow = 10000

func clampDeg(deg float32) float32 {
	return Mod(circleFullDeg+Mod(deg, circleFullDeg), circleFullDeg)
}

func deg2rad(deg float32) float32 {
	return deg * convDeg2Rad
}

func rad2deg(rad float32) float32 {
	return rad * convRad2Deg
}
