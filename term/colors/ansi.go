package colors

import "fmt"

const (
	fmtANSILayerForeground = "\033[38;5;%sm%s\033[0m"
	fmtANSILayerBackground = "\033[48;5;%sm%s\033[0m"
)

type ANSI struct {
	a      uint8
	format string
}

// NewANSI returns an ANSI struct.
func NewANSI(a uint8) ANSI {
	return ANSI{a, fmtANSILayerForeground}
}

// Uint converts the ANSI uint8 to uint.
func (a ANSI) Uint() uint {
	return uint(a.a)
}

// RGB converts an ANSI to RGB. Some variation will be observed
// in the underlying rgb values when converting from RGB to ANSI
// and back. The variation are within 1 point.
func (a ANSI) RGB() RGB {
	return RGBFrom32Bit(ANSIToRGB(a.a))
}

// Bg sets the background color instead of foreground color
func (a ANSI) Bg() Color {
	a.format = fmtANSILayerBackground
	return a
}

// Fg sets the foreground color instead of background color.
// Foreground is the default behavior.
func (a ANSI) Fg() Color {
	a.format = fmtANSILayerForeground
	return a
}

func (a ANSI) String() string {
	return fmt.Sprintf("%d", a.a)
}

func (a ANSI) Printf(format string, args ...interface{}) {
	fmt.Printf(a.Sprintf(format, args...))
}

func (a ANSI) Println(format string, args ...interface{}) {
	fmt.Println(a.Sprintf(format, args...))
}

func (a ANSI) Sprintf(format string, args ...interface{}) string {
	return fmt.Sprintf(
		a.format,
		a.String(),
		fmt.Sprintf(format, args...),
	)
}
