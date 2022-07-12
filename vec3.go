package glx

import (
	"fmt"
	"unsafe"
)

// SizeOfVec3 its size for low precision memory data dump (float32)
// dump have size of 8 bytes (x=4 + y=4 + z=4)
const SizeOfVec3 = 12

type Vec3 struct {
	X, Y, Z float32
}

func NewVec3(x, y, z float32) Vec3 {
	return Vec3{
		X: x,
		Y: y,
		Z: z,
	}
}

func (v *Vec3) String() string {
	return fmt.Sprintf("V3[%.3f, %.3f, %.3f]", v.X, v.Y, v.Z)
}

// Data dump for low precision memory representation (GPU, shaders, etc..)
func (v *Vec3) Data() []byte {
	return (*(*[SizeOfVec3]byte)(unsafe.Pointer(v)))[:]
}

// Basic Math
// -----------------------------------------

func (v Vec3) Add(n Vec3) Vec3 {
	return Vec3{
		X: v.X + n.X,
		Y: v.Y + n.Y,
		Z: v.Z + n.Z,
	}
}

func (v Vec3) Sub(n Vec3) Vec3 {
	return Vec3{
		X: v.X - n.X,
		Y: v.Y - n.Y,
		Z: v.Z - n.Z,
	}
}

func (v Vec3) Mul(n Vec3) Vec3 {
	return Vec3{
		X: v.X * n.X,
		Y: v.Y * n.Y,
		Z: v.Z * n.Z,
	}
}

func (v Vec3) Div(n Vec3) Vec3 {
	if n.X == 0 || n.Y == 0 || n.Z == 0 {
		var resX, resY, resZ float32

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

		return Vec3{
			X: resX,
			Y: resY,
			Z: resZ,
		}
	}

	return Vec3{
		X: v.X / n.X,
		Y: v.Y / n.Y,
		Z: v.Z / n.Z,
	}
}
