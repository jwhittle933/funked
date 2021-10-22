package term

import (
	"errors"
	"fmt"

	"github.com/eiannone/keyboard"
)

var ErrorCanceled = errors.New("wizard was canceled")

type List struct {
	prompt   string
	list     []string
	selected int
}

func NewList(prompt string, opts ...string) *List {
	return &List{prompt, opts, 0}
}

func (l *List) Select() (string, error) {
	fmt.Println(l.prompt)

	l.disableCursor()
	defer l.enableCursor()

	for {
		for i, item := range l.list {
			if i == l.selected {
				fmt.Printf("\033[0;36m❯ %s\033[0m\n", item)
				continue
			}

			fmt.Printf("  %s\n", item)
		}

		key, err := l.ReadKey()
		if err != nil {
			fmt.Println("Read Key error:", err.Error())
		}

		if key == keyboard.KeyArrowDown {
			l.incr()
		}

		if key == keyboard.KeyArrowUp {
			l.decr()
		}

		if key == keyboard.KeyEnter {
			l.clearList()
			break
		}

		if key == keyboard.KeyEsc {
			return l.Current(), ErrorCanceled
		}

		l.clearList()
	}

	l.clearLine()
	return l.Current(), nil
}

func (l *List) ReadKey() (keyboard.Key, error) {
	_, key, err := keyboard.GetSingleKey()

	return key, err
}

func (l *List) Current() string {
	return l.list[l.selected]
}

func (l *List) disableCursor() {
	fmt.Printf("\033[?25l")
}

func (l *List) enableCursor() {
	fmt.Printf("\033[?25h")
}

func (l *List) decr() {
	if l.selected == 0 {
		l.selected = len(l.list) - 1
		return
	}

	l.selected--
}

func (l *List) incr() {
	if l.selected == len(l.list)-1 {
		l.selected = 0
		return
	}

	l.selected++
}

func (l *List) clearList() {
	for i := 0; i < len(l.list); i++ {
		// clear 1 line for every item in list
		// and move up 1 line
		l.clearLine()
	}
}

func (l *List) clearLine() {
	fmt.Printf("\033[1A\033[K")
}
