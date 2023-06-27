package Fs

import "fmt"

type Directory struct {
	directoryName string
	objectList    []interface{}
}

func NewDirectory(name string) *Directory {
	list := make([]interface{}, 0)
	return &Directory{
		directoryName: name,
		objectList:    list,
	}
}

func (d *Directory) Add(obj interface{}) {
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
		//FS.LS()
		switch FS.(type) {
		case File:
			FS.LS()
		case Directory:
			FS.LS()
		}
	}
}
