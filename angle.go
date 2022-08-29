package glx

import "math"

type Angle float32

func Radians(rad float32) Angle {
	return Angle(rad)
}

func Degrees(deg float32) Angle {
	return Angle(deg2rad(clampDeg(deg)))
}

func (a Angle) Normalize() Angle {
	if a >= 0 {
		return a
	}

	return a + circleFullRad
}

func (a Angle) Flip() Angle {
	return Angle(Remainder(a.Radians()+math.Pi, circleFullRad)).Normalize()
}

func (a Angle) Degrees() float32 {
	return RoundTo(clampDeg(rad2deg(float32(a))))
}

func (a Angle) Radians() float32 {
	return float32(a)
}

func (a Angle) Unit() Vec2 {
	return Vec2{
		X: Cos(a.Radians()),
		Y: -Sin(a.Radians()),
	}
}

func (a Angle) Add(t Angle) Angle {
	return a + t
}
