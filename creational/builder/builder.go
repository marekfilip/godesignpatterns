package builder

// BuildProcess defines steps needed to create a vehicle
type BuildProcess interface {
	SetWheels() BuildProcess
	SetSeats() BuildProcess
	SetStructure() BuildProcess
	GetVehicle() VehicleProduct
}

// Director accepts builders to create a vehicles
type Director struct {
	builder BuildProcess
}

// Construct method uses the stored builder
// to create a vehicle
func (f *Director) Construct() {
	f.builder.
		SetWheels().
		SetSeats().
		SetStructure()
}

// SetBuilder allows to provide a builder
// for later creation
func (f *Director) SetBuilder(b BuildProcess) {
	f.builder = b
}

// VehicleProduct is the final product of builders
type VehicleProduct struct {
	Wheels    int
	Seats     int
	Structure string
}

// CarBuilder is a implementation of BuildProcess
// which final product is a car
type CarBuilder struct {
	v VehicleProduct
}

// SetWheels sets a required number of wheels for car
func (c *CarBuilder) SetWheels() BuildProcess {
	c.v.Wheels = 4
	return c
}

// SetSeats sets a required number of seats for car
func (c *CarBuilder) SetSeats() BuildProcess {
	c.v.Seats = 5
	return c
}

// SetStructure sets a structore of a car
func (c *CarBuilder) SetStructure() BuildProcess {
	c.v.Structure = "Car"
	return c
}

// GetVehicle creates a car
func (c *CarBuilder) GetVehicle() VehicleProduct {
	return c.v
}

// BikeBuilder is a implementation of BuildProcess
// which final product is a bike
type BikeBuilder struct {
	v VehicleProduct
}

// SetWheels sets a required number of wheels for bike
func (b *BikeBuilder) SetWheels() BuildProcess {
	b.v.Wheels = 2
	return b
}

// SetSeats sets a required number of seats for bike
func (b *BikeBuilder) SetSeats() BuildProcess {
	b.v.Seats = 1
	return b
}

// SetStructure sets a structore of a bike
func (b *BikeBuilder) SetStructure() BuildProcess {
	b.v.Structure = "Bike"
	return b
}

// GetVehicle creates a bike
func (b *BikeBuilder) GetVehicle() VehicleProduct {
	return b.v
}
