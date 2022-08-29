package glx

import (
	"fmt"
	"unsafe"
)

// SizeOfVec1 its size for low precision memory data dump (float32)
// dump have size of 4 bytes (32 bits)
const SizeOfVec1 = 4

type Vec1 struct {
	X float32
}

func (v *Vec1) String() string {
	return fmt.Sprintf("Vec1{%.2f}", v.X)
}

func (v *Vec1) Data() []byte {
	return (*(*[SizeOfVec1]byte)(unsafe.Pointer(v)))[:]
}
