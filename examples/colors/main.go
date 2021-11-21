package main

import (
	"fmt"

	"github.com/jwhittle933/funked/term/colors"
)

func main() {
	for i := uint8(16); i < 232; i++ {
		ansi := colors.NewANSI(i)
		rgb := ansi.RGB()

		fmt.Printf(
			"%s -- %s\t%s -- %s\n",
			ansi.Sprintf("ANSI"),
			rgb.Sprintf("RGB"),
			ansi.Bg().Sprintf("ANSI"),
			rgb.Bg().Sprintf("RGB"),
		)
	}
}
