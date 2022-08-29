package glx

func Clamp(v, min, max float32) float32 {
	if v < min {
		v = min
	}

	if v > max {
		v = max
	}

	return v
}

func RoundTo(n float32) float32 {
	if n == 0 {
		return 0
	}

	return Round(floatRoundPow*n) / floatRoundPow
}

func Lerp(a, b, t float32) float32 {
	if t <= 0 {
		return a
	}

	if t >= 1 {
		return b
	}

	return a + t*(b-a)
}

func LerpInverse(a, b, t float32) float32 {
	if a == t {
		return 0
	}

	if b == t {
		return 1
	}

	return (t - a) / (b - a)
}

// Lerpf will remap values v=0:originA->targetA, v=1:originA->targetB
func Lerpf(oA, oB, tA, tB, o float32) float32 {
	return Lerp(tA, tB, LerpInverse(oA, oB, o))
}
