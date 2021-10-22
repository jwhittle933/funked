// Package hiccup is a port of Clojure's hiccup library
package hiccup

import (
	"bufio"
	"bytes"

	"github.com/jwhittle933/funked/tree"
)

type Renderer interface {
	Render() string
}

type Document struct {
	tree    *tree.Tree
	changes chan struct{}
}

type Wither func()

func HTML(rr ...Renderer) Document {
	return Document{
		tree.New(tree.NewBranch("html")),
		make(chan struct{}),
	}
}

func RenderHTML(rr ...Renderer) string {
	return (Document{
		tree.New(tree.NewBranch("html")),
		make(chan struct{}),
	}).Render()
}

func (d Document) Render() string {
	buf := bytes.NewBuffer([]byte{})
	w := bufio.NewWriter(buf)

	write(w, "<html>")
	write(w, "</html>")

	return buf.String()
}

func write(w *bufio.Writer, s string) {
	_, _ = w.WriteString(s)
}
