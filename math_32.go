package glx

import "math"

func Floor(x float32) float32 {
	return float32(math.Floor(float64(x)))
}

func Ceil(x float32) float32 {
	return float32(math.Ceil(float64(x)))
}

func Round(x float32) float32 {
	return float32(math.Round(float64(x)))
}

func Sin(x float32) float32 {
	return float32(math.Sin(float64(x)))
}

func Cos(x float32) float32 {
	return float32(math.Cos(float64(x)))
}

func Tan(x float32) float32 {
	return float32(math.Tan(float64(x)))
}

func Atan(x float32) float32 {
	return float32(math.Atan(float64(x)))
}

func Atan2(x, y float32) float32 {
	return float32(math.Atan2(float64(x), float64(y)))
}

func Sqrt(x float32) float32 {
	return float32(math.Sqrt(float64(x)))
}

func Pow(x, y float32) float32 {
	return float32(math.Pow(float64(x), float64(y)))
}

func Remainder(x, y float32) float32 {
	return float32(math.Remainder(float64(x), float64(y)))
}

func Mod(x, y float32) float32 {
	return float32(math.Mod(float64(x), float64(y)))
}
