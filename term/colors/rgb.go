package colors

import (
	"encoding/binary"
	"fmt"
)

const (
	fmtRGBLayerForeground = "\033[38;2;%sm%s\033[0m"
	fmtRGBLayerBackground = "\033[48;2;%sm%s\033[0m"
)

type RGB struct {
	Red,
	Green,
	Blue uint8
	format string
}

// NewRGB create an RGB struct from individual rgb values.
func NewRGB(r, g, b uint8) RGB {
	return RGB{r, g, b, fmtRGBLayerForeground}
}

// RGBFrom32Bit read from a 32-bit integer into RGB.
// Bits are read from right to left.
func RGBFrom32Bit(rgb uint32) RGB {
	return NewRGB(Red(rgb), Green(rgb), Blue(rgb))
}

// Uint converts the RGB value to uint.
func (r RGB) Uint() uint {
	return uint(binary.BigEndian.Uint32([]byte{r.Red, r.Green, r.Blue}))
}

// RGB returns the underlying rgb values.
func (r RGB) RGB() (uint8, uint8, uint8) {
	return r.Red, r.Green, r.Blue
}

// ANSI converts the RGB to an ANSI 8-bit integer.
func (r RGB) ANSI() ANSI {
	return ANSI{RGBToANSI(r.Red, r.Green, r.Blue), r.format}
}

// Bg sets the background color instead of foreground color
func (r RGB) Bg() RGB {
	r.format = fmtRGBLayerBackground
	return r
}

// Fg sets the foreground color instead of background color.
// Foreground is the default behavior.
func (r RGB) Fg() RGB {
	r.format = fmtRGBLayerForeground
	return r
}

func (r RGB) String() string {
	return fmt.Sprintf("%d;%d;%d", r.Red, r.Green, r.Blue)
}

func (r RGB) Printf(format string, args ...interface{}) {
	fmt.Printf(r.Sprintf(format, args...))
}

func (r RGB) Println(format string, args ...interface{}) {
	fmt.Println(r.Sprintf(format, args...))
}

func (r RGB) Sprintf(format string, args ...interface{}) string {
	return fmt.Sprintf(
		r.format,
		r.String(),
		fmt.Sprintf(format, args...),
	)
}
