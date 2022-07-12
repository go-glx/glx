package glx

import (
	"fmt"
	"unsafe"
)

// SizeOfVec4 its size for low precision memory data dump (float32)
// dump have size of 16 bytes (x=4 + y=4 + z=4 + a=4)
const SizeOfVec4 = 16

type Vec4 struct {
	X, Y, Z, R float32
}

func NewVec4(x, y, z, r float32) Vec4 {
	return Vec4{
		X: x,
		Y: y,
		Z: z,
		R: r,
	}
}

func (v *Vec4) String() string {
	return fmt.Sprintf("V4[%.3f, %.3f, %.3f, %.3f]", v.X, v.Y, v.Z, v.R)
}

// Data dump for low precision memory representation (GPU, shaders, etc..)
func (v *Vec4) Data() []byte {
	return (*(*[SizeOfVec4]byte)(unsafe.Pointer(v)))[:]
}

// Basic Math
// -----------------------------------------

func (v Vec4) Add(n Vec4) Vec4 {
	return Vec4{
		X: v.X + n.X,
		Y: v.Y + n.Y,
		Z: v.Z + n.Z,
		R: v.R + n.R,
	}
}

func (v Vec4) Sub(n Vec4) Vec4 {
	return Vec4{
		X: v.X - n.X,
		Y: v.Y - n.Y,
		Z: v.Z - n.Z,
		R: v.R - n.R,
	}
}

func (v Vec4) Mul(n Vec4) Vec4 {
	return Vec4{
		X: v.X * n.X,
		Y: v.Y * n.Y,
		Z: v.Z * n.Z,
		R: v.R * n.R,
	}
}

func (v Vec4) Div(n Vec4) Vec4 {
	if n.X == 0 || n.Y == 0 || n.Z == 0 || n.R == 0 {
		var resX, resY, resZ, resR float32

		if n.X == 0 {
			resX = InfPositive()
		} else {
			resX = v.X / n.X
		}

		if n.Y == 0 {
			resY = InfPositive()
		} else {
			resY = v.Y / n.Y
		}

		if n.Z == 0 {
			resZ = InfPositive()
		} else {
			resZ = v.Z / n.Z
		}

		if n.R == 0 {
			resR = InfPositive()
		} else {
			resR = v.R / n.R
		}

		return Vec4{
			X: resX,
			Y: resY,
			Z: resZ,
			R: resR,
		}
	}

	return Vec4{
		X: v.X / n.X,
		Y: v.Y / n.Y,
		Z: v.Z / n.Z,
		R: v.R / n.R,
	}
}
