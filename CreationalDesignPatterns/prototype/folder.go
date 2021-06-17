package main

import "fmt"

type Folder struct {
	children []inode
	Name     string
}

func (f *Folder) print(indentation string) {
	fmt.Println(indentation + f.Name)
	for _, i := range f.children {
		i.print(indentation + indentation)
	}
}

func (f *Folder) clone() inode {
	cloneFolder := &Folder{Name: f.Name + "_clone"}
	var tempChildren []inode
	for _, i := range f.children {
		copy := i.clone()
		tempChildren = append(tempChildren, copy)
	}
	cloneFolder.children = tempChildren
	return cloneFolder

}
