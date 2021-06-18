package composite

import "testing"

func TestAthlete_Train(t *testing.T) {
	athlete := Athlete{}
	athlete.Train()
}

func TestSwimmer_Swim(t *testing.T) {
	localSwim := Swim
	swimmer := CompositeSwimmerA{
		MySwim: localSwim,
	}
	swimmer.MyAthlete.Train()
	swimmer.MySwim()
}

func TestAnimal_Swim(t *testing.T) {
	fish := Shark{
		Swim: Swim,
	}

	fish.Eat()
	fish.Swim()
}

func TestSwimmer_SwimB(t *testing.T) {
	swimmerb := CompositeSwimmerB{
		&Athlete{},
		&SwimmerImp1{},
	}

	swimmerb.Train()
	swimmerb.Swim()
}
