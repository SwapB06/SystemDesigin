package Fs

import "fmt"

type File struct {
	fileName string
}

func NewFile(name string) *File {
	return &File{
		fileName: name,
	}
}

func (f File) LS() {
	fmt.Println("file name ", f.fileName)
}
