package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/jwhittle933/funked/async/process"
	"github.com/jwhittle933/funked/term/colors"
)

// FileProcessor example struct for reading and decoding a file.
// Not safe for concurrent use.
type FileProcessor struct {
	filename string
	f        *os.File
	decode   func([]byte)
}

// NewFileProcessor factory for FileProcessor.
func NewFileProcessor(file string, decoder func([]byte)) *FileProcessor {
	return &FileProcessor{file, nil, decoder}
}

// Process processes the file contents.
func (p *FileProcessor) Process() {
	home, _ := os.UserHomeDir()
	location := filepath.Join(home, "Development", "lexica", p.filename)

	f, err := os.Open(location)
	fmt.Printf("\t%s: %s\n", colors.NewANSI(50).Sprintf("Open"), location)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	content, err := ioutil.ReadAll(f)
	fmt.Printf("\t%s: %s\n", colors.NewANSI(200).Sprintf("Read"), location)
	if err != nil {
		panic(err)
	}

	p.decode(content)
}

func main() {
	greek := NewFileProcessor("grc/lsj/dictionary.json", func(content []byte) {
		printSize("grc/lsj/dictionary.json", len(content))
	})
	hebrew := NewFileProcessor("heb1/BDB/DictBDB.json", func(content []byte) {
		printSize("heb/BDB/DictBDB.json", len(content))
	})
	latin := NewFileProcessor("lat/ls/ls_A.json", func(content []byte) {
		printSize("lat/ls/ls_A.json", len(content))
	})

	gkPID := process.Start(process.New(greek.Process))
	hebPID := process.Start(process.New(hebrew.Process))
	latinPID := process.Start(process.New(latin.Process))

	process.Await(gkPID)
	process.Await(hebPID)
	process.Await(latinPID)
}

func printSize(filename string, size int) {
	fmt.Printf(
		"\t%s: %s: %d%s\n",
		colors.NewANSI(120).Sprintf("Size"),
		filename,
		size/1024/1024,
		"MB",
	)
}
