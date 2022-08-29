package glx

import (
	"fmt"
	"unsafe"
)

const SizeOfMat4 = SizeOfVec4 * 4

type Mat4 struct {
	A Vec4
	B Vec4
	C Vec4
	D Vec4
}

func Mat4Identity() Mat4 {
	return Mat4{
		A: Vec4{1, 0, 0, 0},
		B: Vec4{0, 1, 0, 0},
		C: Vec4{0, 0, 1, 0},
		D: Vec4{0, 0, 0, 1},
	}
}

func (v *Mat4) String() string {
	return fmt.Sprintf(`Mat4[
     %0.2f, %0.2f, %0.2f, %0.2f,
     %0.2f, %0.2f, %0.2f, %0.2f,
     %0.2f, %0.2f, %0.2f, %0.2f,
     %0.2f, %0.2f, %0.2f, %0.2f,
]`,
		v.A.X, v.A.Y, v.A.Z, v.A.W,
		v.B.X, v.B.Y, v.B.Z, v.B.W,
		v.C.X, v.C.Y, v.C.Z, v.C.W,
		v.D.X, v.D.Y, v.D.Z, v.D.W,
	)
}

// Data dump for low precision memory representation (GPU, shaders, etc..)
func (v *Mat4) Data() []byte {
	return (*(*[SizeOfMat4]byte)(unsafe.Pointer(v)))[:]
}
