package main

// Get the pizza's price [Component interface]
type pizza interface {
	getPrice() int
}

// A kind of pizza [Concrete component]
type veggeMania struct{}

func (p *veggeMania) getPrice() int {
	return 15
}
