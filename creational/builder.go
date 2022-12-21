package creational

type ManufacturingDirector struct{}

type BuildProcess interface {
	SetWheels() BuildProcess
	SetSeats() BuildProcess
	SetStructure() BuildProcess
	// GetVehicle() BuildProcess
}

func (f *ManufacturingDirector) Construct() {
	// Implementation goes here
}

func (f *ManufacturingDirector) SetBuilder(b BuildProcess) {
	// Implementation goes here
}

type VehicleProduct struct {
	Wheels    int
	Seats     int
	Structure string
}

type CarBuilder struct{}

func (c *CarBuilder) SetWheels() BuildProcess {
	return nil
}
func (c *CarBuilder) SetSeats() BuildProcess {
	return nil
}
func (c *CarBuilder) SetStructure() BuildProcess {
	return nil
}
func (c *CarBuilder) Build() VehicleProduct {
	return VehicleProduct{}
}

type BikeBuilder struct{}

func (b *BikeBuilder) SetWheels() BuildProcess {
	return nil
}
func (b *BikeBuilder) SetSeats() BuildProcess {
	return nil
}
func (b *BikeBuilder) SetStructure() BuildProcess {
	return nil
}
func (b *BikeBuilder) Build() VehicleProduct {
	return VehicleProduct{}
}
