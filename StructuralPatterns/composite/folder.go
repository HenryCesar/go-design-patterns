package composite

import "fmt"

type Folder struct {
	Components []Component
	Name       string
}

func (f *Folder) search(keyword string) {
	fmt.Printf("Serching recursively for keyword %s in folder %s\n", keyword, f.Name)
	for _, composite := range f.Components {
		composite.search(keyword)
	}
}

func (f *Folder) add(c component) {
	f.Components = append(f.Components, c)
}
