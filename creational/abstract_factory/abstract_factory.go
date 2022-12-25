package abstract_factory

import "fmt"

type AbstractFactory interface {
	Build(v int) (Vehicle, error)
}

const (
	CarFactoryType       = 1
	MotorbikeFactoryType = 2
)

func BuildFactory(f int) (VehicleFactory, error) {
	switch f {
	default:
		return nil, fmt.Errorf("factory with id %d not recognized", f)
	}
}
