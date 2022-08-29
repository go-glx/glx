package glx

import (
	"fmt"
	"math"
	"unsafe"
)

// SizeOfVec2 its size for low precision memory data dump (float32)
// dump have size of 8 bytes (x=4 + y=4)
const SizeOfVec2 = SizeOfVec1 * 2

type Vec2 struct {
	X, Y float32
}

func (v *Vec2) String() string {
	return fmt.Sprintf("V2[%.3f, %.3f]", v.X, v.Y)
}

// Data dump for low precision memory representation (GPU, shaders, etc..)
func (v *Vec2) Data() []byte {
	return (*(*[SizeOfVec2]byte)(unsafe.Pointer(v)))[:]
}

// Basic Math
// -----------------------------------------

func (v Vec2) Add(n Vec2) Vec2 {
	return Vec2{
		X: v.X + n.X,
		Y: v.Y + n.Y,
	}
}

func (v Vec2) Sub(n Vec2) Vec2 {
	return Vec2{
		X: v.X - n.X,
		Y: v.Y - n.Y,
	}
}

func (v Vec2) Mul(n Vec2) Vec2 {
	return Vec2{
		X: v.X * n.X,
		Y: v.Y * n.Y,
	}
}

func (v Vec2) Div(n Vec2) Vec2 {
	if n.X == 0 || n.Y == 0 {
		if n.X == 0 && n.Y == 0 {
			return Vec2{InfPositive(), InfPositive()}
		}

		if n.X == 0 {
			return Vec2{InfPositive(), v.Y / n.Y}
		}

		if n.Y == 0 {
			return Vec2{v.X / n.X, InfPositive()}
		}
	}

	return Vec2{
		X: v.X / n.X,
		Y: v.Y / n.Y,
	}
}

func (v Vec2) Plus(n float32) Vec2 {
	return Vec2{
		X: v.X + n,
		Y: v.Y + n,
	}
}

func (v Vec2) Minus(n float32) Vec2 {
	return Vec2{
		X: v.X - n,
		Y: v.Y - n,
	}
}

func (v Vec2) Scale(n float32) Vec2 {
	return Vec2{
		X: v.X * n,
		Y: v.Y * n,
	}
}

func (v Vec2) Decrease(n float32) Vec2 {
	return Vec2{
		X: v.X / n,
		Y: v.Y / n,
	}
}

func (v Vec2) Cross(to Vec2) float32 {
	return v.X*to.Y - v.Y*to.X
}

func (v Vec2) Dot(to Vec2) float32 {
	return v.X*to.X + v.Y*to.Y
}

// Trigonometry
// -----------------------------------------

func (v Vec2) DistanceTo(to Vec2) float32 {
	return Sqrt(
		(v.X-to.X)*(v.X-to.X) +
			(v.Y-to.Y)*(v.Y-to.Y),
	)
}

func (v Vec2) Direction() Angle {
	return Angle(Atan2(-v.Y, v.X))
}

func (v Vec2) AngleBetween(to Vec2) Angle {
	return Angle(Atan2(v.Cross(to), v.Dot(to)))
}

func (v Vec2) AngleTo(to Vec2) Angle {
	return Angle(Atan2(to.Y-v.Y, v.X-to.X) + math.Pi)
}

func (v Vec2) Rotate(angle Angle) Vec2 {
	sin := Sin(angle.Radians())
	cos := Cos(angle.Radians())

	return Vec2{
		X: v.X*cos - v.Y*sin,
		Y: -(v.X*sin + v.Y*cos),
	}
}

func (v Vec2) RotateAround(orig Vec2, angle Angle) Vec2 {
	sin := Sin(angle.Radians())
	cos := Cos(angle.Radians())

	v.X -= orig.X
	v.Y -= orig.Y

	xx := v.X*cos + v.Y*sin
	yy := -(v.X*sin - v.Y*cos)

	v.X = xx + orig.X
	v.Y = yy + orig.Y

	return v
}

func (v Vec2) PolarOffset(distance float32, angle Angle) Vec2 {
	return Vec2{
		X: v.X + distance*Cos(angle.Radians()),
		Y: v.Y - distance*Sin(angle.Radians()),
	}
}
