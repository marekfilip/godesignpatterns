package builder

import "testing"

func TestBuilderBuildCar(t *testing.T) {
	manufacturingComplex := Director{}
	carBuilder := &CarBuilder{}
	manufacturingComplex.SetBuilder(carBuilder)
	manufacturingComplex.Construct()
	car := carBuilder.GetVehicle()

	if car.Wheels != 4 {
		t.Errorf("Car should have 4 wheels, had %d\n", car.Wheels)
	}
	if car.Structure != "Car" {
		t.Errorf("Car structure must be 'Car' and was %s\n", car.Structure)
	}
	if car.Seats != 4 {
		t.Errorf("Car should have 4 seats, had %d\n", car.Seats)
	}
}

func TestBuilderBuildBike(t *testing.T) {
	manufacturingComplex := Director{}
	bikeBuilder := &BikeBuilder{}
	manufacturingComplex.SetBuilder(bikeBuilder)
	manufacturingComplex.Construct()
	bike := bikeBuilder.GetVehicle()

	if bike.Wheels != 2 {
		t.Errorf("Bike should have 2 wheels, had %d\n", bike.Wheels)
	}
	if bike.Structure != "Bike" {
		t.Errorf("Bike structure must be 'Bike' and was %s\n", bike.Structure)
	}
	if bike.Seats != 1 {
		t.Errorf("Car should have 1 seat, had %d\n", bike.Seats)
	}
}
