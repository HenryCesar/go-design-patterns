package abstract_factory

import "testing"

func TestCarFactory(t *testing.T) {
	carF, err := GetVehicleFactory(CarFactoryType)
	if err != nil {
		t.Fatal(err)
	}

	carVehicle, err := carF.GetVehicle(LuxuryCarType)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Car vehicle has %d seats\n", carVehicle.NumWheels())

	luxuryCar, ok := carVehicle.(Car)
	if !ok {
		t.Fatal("Struct assertion has failed")
	}

	t.Logf("Luxury car has %d doors.\n", luxuryCar.NumDoors())
}
