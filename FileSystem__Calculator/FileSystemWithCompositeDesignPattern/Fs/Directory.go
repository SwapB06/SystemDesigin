package Fs

import "fmt"

type Directory struct {
	directoryName string
	objectList    []FileSystem
}

func NewDirectory(name string) *Directory {
	list := make([]FileSystem, 0)
	return &Directory{
		directoryName: name,
		objectList:    list,
	}
}

func (d *Directory) Add(obj FileSystem) {
	// switch obj.(type) {
	// case File:

	// case Directory:
	// }
	d.objectList = append(d.objectList, obj)
}

// func (d Directory) add(dire Directory) {
// 	d.objectList = append(d.objectList, dire)
// }

func (d Directory) LS() {
	fmt.Println("Directory Name: ", d.directoryName)
	for _, FS := range d.objectList {
		FS.LS()
		// switch fs.(type) {
		// case File:
		// 	fs.ls()
		// case Directory:
		// 	fs.ls()
		// }
	}
}
