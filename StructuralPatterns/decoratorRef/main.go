package main

import "fmt"

// Get the pizza's price [Component interface]
type pizza interface {
	getPrice() int
}

// A kind of pizza [Concrete component]
type veggeMania struct{}

func (p *veggeMania) getPrice() int {
	return 15
}

// The type of Structural Pattern [Concrete Decorator]
type tomatoTopping struct {
	pizza pizza
}

func (c *tomatoTopping) getPrice() int {
	pizzaPrice := c.pizza.getPrice()
	return pizzaPrice + 7
}

// The type of Structural Pattern [Concrete Decorator]
type cheeseTopping struct {
	pizza pizza
}

func (c *cheeseTopping) getPrice() int {
	pizzaPrice := c.pizza.getPrice()
	return pizzaPrice + 10
}

// Main function [Client Code]
func main() {

	pizza := &veggeMania{}

	//Add cheese topping
	pizzaWithCheese := &cheeseTopping{
		pizza: pizza,
	}

	//Add tomato topping
	pizzaWithCheeseAndTomato := &tomatoTopping{
		pizza: pizzaWithCheese,
	}

	fmt.Printf("Price of veggeMani with tomato and cheese topping is %d\n", pizzaWithCheeseAndTomato.getPrice())
}
