package fs

import (
	"io"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
)

type PathError = fs.PathError

type FileReader interface {
	ReadFile(path string) ([]byte, error)
}

type DirReader interface {
	ReadDir(dir string) ([]os.FileInfo, error)
}

type Reader interface {
	FileReader
	DirReader
}

type FileWriter interface {
	WritePath(path string) error
	WriteFile(filename string, content []byte) error
	AppendFile(filename string, content []byte) error
}

type FileReadWriter interface {
	FileReader
	FileWriter
}

type FileCopier interface {
	CopyFile(src, dst string) error
}

type Navigator interface {
	ChDir(path string) error
	CurrentDir() string
	HomeDir() string
}

type FS interface {
	FileReadWriter
	FileCopier
	Navigator
	Exists(path string) string
}

type FileSystem struct{}

func NewFS() *FileSystem {
	return &FileSystem{}
}

// ReadFile reads a file from the FileSystem
func (f *FileSystem) ReadFile(path string) ([]byte, error) {
	absolutePath, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}

	content, err := ioutil.ReadFile(absolutePath)
	if err != nil {
		return nil, err
	}

	return content, nil
}

// WritePath creates a path on the filesystem
// If the path exists, nothing happens
func (f *FileSystem) WritePath(path string) error {
	return os.MkdirAll(path, fs.ModePerm)
}

// WriteFile writes a file to the FileSystem
func (f *FileSystem) WriteFile(filename string, content []byte) error {
	return ioutil.WriteFile(filename, content, 0644)
}

func (f *FileSystem) AppendFile(filename string, content []byte) error {
	var (
		body []byte
		err  error
	)

	if f.Exists(filename) {
		if body, err = f.ReadFile(filename); err != nil {
			return err
		}
	}

	if err = f.WriteFile(filename, append(body, content...)); err != nil {
		return err
	}

	return nil
}

func (f *FileSystem) ReadDir(dir string) ([]os.FileInfo, error) {
	return ioutil.ReadDir(dir)
}

func (f *FileSystem) CopyFile(src, dst string) error {
	if f.Exists(dst) {
		return os.ErrExist
	}

	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	return err

}

func (f *FileSystem) Exists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) || os.IsPermission(err) {
		return false
	}

	return true
}

func (f *FileSystem) HomeDir() string {
	userHome, _ := os.UserHomeDir()
	return userHome
}

func (f *FileSystem) CurrentDir() string {
	currDir, _ := os.Getwd()
	return currDir
}

func (f *FileSystem) ChDir(path string) error {
	return os.Chdir(path)
}
