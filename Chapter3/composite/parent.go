package composite

import "fmt"

type Parent struct {
	SomeField int
}

type Son struct {
	P Parent
}

func GetParentField(p *Parent) {
	fmt.Println(p.SomeField)
}
