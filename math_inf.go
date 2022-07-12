package glx

import "math"

const (
	uPositiveInf = 0x7f800000
	uNegativeInf = 0xff800000
)

// InfPositive returns positive infinity
func InfPositive() float32 {
	return math.Float32frombits(uPositiveInf)
}

// InfNegative returns negative infinity
func InfNegative() float32 {
	return math.Float32frombits(uNegativeInf)
}

// IsPositiveInf reports whether f is a positive infinity
func IsPositiveInf(f float32) bool {
	return f > math.MaxFloat32
}

// IsNegativeInf reports whether f is a negative infinity
func IsNegativeInf(f float32) bool {
	return f < -math.MaxFloat32
}
