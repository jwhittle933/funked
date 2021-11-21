package colors

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func inRangeOf(val, exact, delta uint8) bool {
	return (val > exact-delta) || val < exact+delta
}

func TestColor_ANSI(t *testing.T) {
	testCases := []struct {
		ansi    uint8
		r, g, b uint8
	}{
		{ansi: 28, r: 2, g: 135, b: 0},
		{ansi: 29, r: 3, g: 135, b: 95},
		{ansi: 30, r: 2, g: 135, b: 135},
		{ansi: 31, r: 0, g: 135, b: 175},
		{ansi: 32, r: 1, g: 135, b: 215},
		{ansi: 33, r: 4, g: 135, b: 255},
		{ansi: 172, r: 216, g: 135, b: 0},
		{ansi: 173, r: 216, g: 136, b: 94},
		{ansi: 174, r: 216, g: 135, b: 135},
		{ansi: 175, r: 215, g: 136, b: 174},
		{ansi: 220, r: 255, g: 215, b: 0},
		{ansi: 221, r: 255, g: 214, b: 95},
	}

	for _, testCase := range testCases {
		name := fmt.Sprintf("%d-%d-%d -- %d", testCase.r, testCase.g, testCase.b, testCase.ansi)
		t.Run(name, func(t *testing.T) {
			rgb := NewRGB(testCase.r, testCase.g, testCase.b)
			ansi := rgb.ANSI()

			assert.Equal(t, testCase.ansi, uint8(ansi), "ANSI values do not match")
			rgb.Printf("RGB(%s)", rgb)
			fmt.Print(" -- ")
			ansi.Printf("ANSI(%s)", ansi)
			fmt.Printf("\n")

			rgbFromANSI := ansi.RGB()
			assert.True(t, inRangeOf(rgbFromANSI.Red, testCase.r, 1), "Converted R-value out of range")
			assert.True(t, inRangeOf(rgbFromANSI.Green, testCase.g, 1), "Converted G-value out of range")
			assert.True(t, inRangeOf(rgbFromANSI.Blue, testCase.b, 1), "Converted B-value out of range")
		})
	}
}
