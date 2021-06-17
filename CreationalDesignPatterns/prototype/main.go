package main

import "fmt"

func main() {
	file1 := &File{Name: "File1"}
	file2 := &File{Name: "File2"}
	file3 := &File{Name: "File3"}

	folder1 := &Folder{
		children: []inode{file1},
		Name: "Folder1",
	}
	folder2 := &Folder{
		children: []inode{folder1, file2, file3},
		Name: "Folder2",
	}
	fmt.Println("Printing hierarchy for Folder2")
	folder2.print(" ")

	cloneFolder := folder2.clone()
	fmt.Println("Printing hierarchy for clone Folder")
	cloneFolder.print(" ")
}