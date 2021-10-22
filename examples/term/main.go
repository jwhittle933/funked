package main

import (
	"fmt"

	"github.com/jwhittle933/funked/term"
)

func main() {
	ins := "(Use arrow keys)"
	prompt := "\033[0;32m?\033[0m Please select a platform: "
	l := term.NewList(fmt.Sprintf("%s%s", prompt, ins), "vue", "express", "fastify", "dotnet")
	selected, err := l.Select()
	if err != nil {
		return
	}

	fmt.Printf("%s\033[36m%s\033[0m\n", prompt, selected)
	// list := []string{"vue", "express", "fastify", "dotnet"}
	// selected := 0

	// fmt.Printf("\033[?25l")       // turn cursor off
	// defer fmt.Printf("\033[?25h") // turn cursor on
	// for {
	// 	for i, l := range list {
	// 		if i == selected {
	// 			fmt.Printf("\033[0;36m%s\033[0m <-\n", l)
	// 			continue
	// 		}

	// 		fmt.Println(l)
	// 	}

	// 	// GetSingleKey opens and closes the keyboard
	// 	_, key, err := keyboard.GetSingleKey()
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	if key == keyboard.KeyArrowDown {
	// 		selected = up(selected, len(list)-1)
	// 	}

	// 	if key == keyboard.KeyArrowUp {
	// 		selected = down(selected, len(list)-1)
	// 	}

	// 	if key == keyboard.KeyEnter {
	// 		clearN(len(list))
	// 		fmt.Printf("Selected \033[36m%s\033[0m\n", list[selected])
	// 		break
	// 	}

	// 	if key == keyboard.KeyEsc {
	// 		break
	// 	}

	// 	clearN(len(list))
	// }
}

func down(selected int, max int) int {
	if selected == 0 {
		return max
	}

	return selected - 1
}

func up(selected int, max int) int {
	if selected == max {
		return 0
	}

	return selected + 1
}

func clearN(delta int) {
	for i := 0; i < delta; i++ {
		// clear 1 line for every item in list
		// and move up 1 line
		fmt.Printf("\033[1A\033[K")
	}
}
