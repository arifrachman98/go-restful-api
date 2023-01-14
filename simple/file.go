package simple

import "fmt"

type File struct {
	Name string
}

func (f File) Close() {
	fmt.Println("Close File", f.Name)
}

func NewFile(name string) (*File, func()) {
	files := &File{
		Name: name,
	}
	return files, func() {
		files.Close()
	}
}
