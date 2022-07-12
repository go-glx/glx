package glx

import (
	"fmt"
	"unsafe"
)

// SizeOfVec2 its size for low precision memory data dump (float32)
// dump have size of 8 bytes (x=4 + y=4)
const SizeOfVec2 = 8

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
