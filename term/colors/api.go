package colors

import "fmt"

var (
	ANSI256FromGrey = []uint8{
		16, 16, 16, 16, 16, 232, 232, 232, 232, 232, 232, 232, 232, 232, 233, 233, 233, 233, 233, 233,
		233, 233, 233, 233, 234, 234, 234, 234, 234, 234, 234, 234, 234, 234, 235, 235, 235, 235, 235,
		235, 235, 235, 235, 235, 236, 236, 236, 236, 236, 236, 236, 236, 236, 236, 237, 237, 237, 237,
		237, 237, 237, 237, 237, 237, 238, 238, 238, 238, 238, 238, 238, 238, 238, 238, 239, 239, 239,
		239, 239, 239, 239, 239, 239, 239, 240, 240, 240, 240, 240, 240, 240, 240, 59, 59, 59, 59, 59,
		241, 241, 241, 241, 241, 241, 241, 242, 242, 242, 242, 242, 242, 242, 242, 242, 242, 243, 243,
		243, 243, 243, 243, 243, 243, 243, 244, 244, 244, 244, 244, 244, 244, 244, 244, 102, 102, 102,
		102, 102, 245, 245, 245, 245, 245, 245, 246, 246, 246, 246, 246, 246, 246, 246, 246, 246, 247,
		247, 247, 247, 247, 247, 247, 247, 247, 247, 248, 248, 248, 248, 248, 248, 248, 248, 248, 145,
		145, 145, 145, 145, 249, 249, 249, 249, 249, 249, 250, 250, 250, 250, 250, 250, 250, 250, 250,
		250, 251, 251, 251, 251, 251, 251, 251, 251, 251, 251, 252, 252, 252, 252, 252, 252, 252, 252,
		252, 188, 188, 188, 188, 188, 253, 253, 253, 253, 253, 253, 254, 254, 254, 254, 254, 254, 254,
		254, 254, 254, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 231, 231,
		231, 231, 231, 231, 231, 231, 231,
	}

	SystemColors = [...]uint32{
		0x000000, 0xcd0000, 0x00cd00, 0xcdcd00, 0x0000ee, 0xcd00cd, 0x00cdcd, 0xe5e5e5,
		0x7f7f7f, 0xff0000, 0x00ff00, 0xffff00, 0x5c5cff, 0xff00ff, 0x00ffff, 0xffffff,
	}
)

type Printer interface {
	Printf(format string, args ...interface{})
	Println(format string, args ...interface{})
}

type Sprinter interface {
	Sprintf(format string, args ...interface{}) string
}

type Layer interface {
	Fg() Color
	Bg() Color
}

// Color interface represents a colorized print utility
type Color interface {
	Uint() uint
	Printer
	Sprinter
	Layer
	fmt.Stringer
}

// Printf formatted printer for Color.
func Printf(c Color, format string, args ...interface{}) {
	c.Printf(format, args...)
}

// Println formatted printer for Color.
func Println(c Color, format string, args ...interface{}) {
	c.Println(format, args...)
}

// Sprintf formats `args` with color codes from `c`.
func Sprintf(c Color, format string, args ...interface{}) string {
	return c.Sprintf(format, args...)
}

// Red read the R-value from the rgb value.
func Red(rgb uint32) uint8 {
	return uint8((rgb >> 16) & 0xff)
}

// Green reads the G-value from the rgb value.
func Green(rgb uint32) uint8 {
	return uint8((rgb >> 8) & 0xff)
}

// Blue reads the B-value from the rgb value.
func Blue(rgb uint32) uint8 {
	return uint8(rgb & 0xff)
}

func Luminance(rgb uint32) uint8 {
	x := uint32(3567664)*(rgb>>16&0xff) +
		uint32(11998547)*(rgb>>8&0xff) +
		1211005*(rgb&0xff)
	return uint8((x + (1 << 23)) >> 24)
}

// RGBToANSI converts rgb values into an ANSI 8-bit integer.
func RGBToANSI(r, g, b uint8) uint8 {
	rgb := (uint32(r) << 16) + (uint32(g) << 8) + uint32(b)
	if r == g && g == b {
		return ANSI256FromGrey[uint(rgb)&0xff]
	}

	greyIndex := ANSI256FromGrey[Luminance(rgb)]
	greyDistance := distance(rgb, ANSIToRGB(greyIndex))

	cube := cubeIndexRed(r) + cubeIndexGreen(g) + cubeIndexBlue(b)
	if distance(rgb, cube) < greyDistance {
		return uint8(cube >> 24)
	}

	return greyIndex
}

// ANSIToRGB converts the ANSI code to RGB values,
// encoded as uint32.
func ANSIToRGB(ansi uint8) uint32 {
	if ansi < 16 {
		return SystemColors[ansi]
	}

	if ansi < 232 {
		ansi = ansi - 16
		return (cubeValue(ansi/36) << 16) | (cubeValue(ansi/6%6) << 8) | (cubeValue(ansi % 6))
	}

	ansi = (ansi-232)*10 + 8
	return uint32(ansi) * 0x010101
}

func distance(x, y uint32) uint32 {
	rSum := int32(Red(x) + Red(y))
	red := int32(Red(x) - Red(y))
	green := int32(Green(x) - Green(y))
	blue := int32(Blue(x) - Blue(y))

	return uint32((1024+rSum)*red*red + 2048*green*green + (1534-rSum)*blue*blue)
}

func cubeValue(idx uint8) uint32 {
	return [...]uint32{0, 95, 135, 175, 215, 255}[idx]
}

func cubeIndexRed(r uint8) uint32 {
	return cubeThresholds(r, 38, 115, 155, 196, 235, func(i uint32, v uint32) uint32 {
		return ((i*36 + 16) << 24) | (v << 16)
	})
}

func cubeIndexGreen(g uint8) uint32 {
	return cubeThresholds(g, 36, 116, 154, 195, 235, func(i uint32, v uint32) uint32 {
		return ((i * 6) << 24) | (v << 8)
	})
}

func cubeIndexBlue(b uint8) uint32 {
	return cubeThresholds(b, 36, 116, 154, 195, 235, func(i uint32, v uint32) uint32 {
		return (i << 24) | v
	})
}

func cubeThresholds(v, a, b, c, d, e uint8, idx func(uint32, uint32) uint32) uint32 {
	if v < a {
		return idx(0, 0)
	}

	if v < b {
		return idx(1, 95)
	}

	if v < c {
		return idx(2, 135)
	}

	if v < d {
		return idx(3, 175)
	}

	if v < e {
		return idx(4, 215)
	}

	return idx(5, 255)
}
