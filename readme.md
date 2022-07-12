# Glx

glx is core utility library for go-glx engine

__NOT READY FOR REAL USE__

core concept is "safety first", so library __immutable__
and will copy vec/mat data for every call.

precision is __float32__ - low memory usage and its native for 
hardware (GPU's shaders, buffers, etc..) anyway use f32,
so not need to convert data every syscall

Its contain:
- f32 math
- vectors/matrix
- angles
- colors

### details

#### unsafe

lib use __unsafe__ package, but only in cast place from float32 to []byte

```go
func (v *Vec2) Data() []byte {
  return (*(*[SizeOfVec2]byte)(unsafe.Pointer(v)))[:]
}
```

This is very useful for fast cast vectors to GPU vertex data 
