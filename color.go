package glx

import (
	"encoding/binary"
	"image/color"
	"math"
)

type Color uint32

//goland:noinspection GoUnusedConst
const (
	ColorWhite Color = 0xFFFFFFFF
	ColorBlack Color = 0x000000FF
	ColorGray  Color = 0x777777FF
	ColorRed   Color = 0xFF0000FF
	ColorGreen Color = 0x00FF00FF
	ColorBlue  Color = 0x0000FFFF
)

// ColorFromRGBA will encode color from RGBA components
// each component is uint8 values (0 .. 255)
func ColorFromRGBA(r, g, b, a uint8) Color {
	return Color(binary.LittleEndian.Uint32([]byte{r, g, b, a}))
}

// ColorFromHex just cast uint32 (hex) to color type
// is alias for cast:
//
//	c := glx.Color(0xff0000ff)
//	c := glx.ColorFromHex(0xff0000ff)
func ColorFromHex(hex uint32) Color {
	return Color(hex)
}

// RGBA will convert color to golang representation
func (c Color) RGBA() color.RGBA {
	vec := c.VecRGBA()
	return color.RGBA{
		R: uint8(vec.X * math.MaxUint8),
		G: uint8(vec.Y * math.MaxUint8),
		B: uint8(vec.Z * math.MaxUint8),
		A: uint8(vec.W * math.MaxUint8),
	}
}

// VecRGBA converts encoded uint32 color into 4 byte values (r,g,b,a)
// and cast in to Vec4 float32 in range of (0 .. 1)
func (c Color) VecRGBA() Vec4 {
	return Vec4{
		X: float32(c>>24&0xff) / math.MaxUint8,
		Y: float32(c>>16&0xff) / math.MaxUint8,
		Z: float32(c>>8&0xff) / math.MaxUint8,
		W: float32(c&0xff) / math.MaxUint8,
	}
}

// Split will decode hex to 4 color components
// each component is uint8 values (0 .. 255)
func (c Color) Split() (r uint8, g uint8, b uint8, a uint8) {
	bytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(bytes, uint32(c))
	return bytes[3], bytes[2], bytes[1], bytes[0]
}

// SplitRGB will decode hex to 3 color components
// each component is uint8 values (0 .. 255)
// alpha component is just ignored
func (c Color) SplitRGB() (r uint8, g uint8, b uint8) {
	bytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(bytes, uint32(c))
	return bytes[3], bytes[2], bytes[1]
}
