package main

import "fmt"

type File struct {
	Name string
}

func (f *File) print(indentation string) {
	fmt.Println(indentation + f.Name)
}

func (f *File) clone() inode {
	return &File{Name: f.Name + "_clone"}
}
