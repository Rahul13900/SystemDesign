package main

import "fmt"

type Vehicle interface {
	Drive() string
}

type Car struct{}

func (c Car) Drive() string {
	return "Driving a car"
}

type Bike struct{}

func (b Bike) Drive() string {
	return "Riding a bike"
}

func VehicleFactory(vehicleType string) Vehicle {
	switch vehicleType {
	case "car":
		return Car{}
	case "bike":
		return Bike{}
	default:
		return nil
	}
}

func main() {
	vehicle := VehicleFactory("car")
	if vehicle != nil {
		fmt.Println(vehicle.Drive()) // Output: Driving a car
	} else {
		fmt.Println("Unknown vehicle type")
	}
}
