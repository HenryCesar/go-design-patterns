package composite

import "fmt"

type Athlete struct{}

func (a *Athlete) Train() {
	fmt.Println("Training")
}

func Swim() {
	fmt.Println("Swimming!")
}

type CompositeSwimmerA struct {
	MyAthlete Athlete
	MySwim    func()
}

//----------------------------------------------

type Swimmer interface {
	Swim()
}
type Trainer interface {
	Train()
}

type SwimmerImp1 struct{}

func (s *SwimmerImp1) Swim() {
	fmt.Println("Swimming!")
}

type CompositeSwimmerB struct {
	Trainer
	Swimmer
}

//----------------------------------------------

type Animal struct{}

func (r *Animal) Eat() {
	fmt.Println("Eating")
}

type Shark struct {
	Animal
	Swim func()
}
