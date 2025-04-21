package main

import "fmt"

type DriveFunc func()

type Vehicle struct {
	DriveBehavior DriveFunc
}

func (v *Vehicle) Drive() {
	v.DriveBehavior()
}

func NewPassengerVehicle() *Vehicle {
	return &Vehicle{
		DriveBehavior: func() {
			fmt.Println("Driving on regular roads")
		},
	}
}

func NewOffroadVehicle() *Vehicle {
	return &Vehicle{
		DriveBehavior: func() {
			fmt.Println("Driving on rough off-road terrain")
		},
	}
}

func NewSportsVehicle() *Vehicle {
	return &Vehicle{
		DriveBehavior: func() {
			fmt.Println("Driving at high speed with agility")
		},
	}
}

func main() {
	pv := NewPassengerVehicle()
	or := NewOffroadVehicle()
	sv := NewSportsVehicle()

	pv.Drive()
	or.Drive()
	sv.Drive()

	// Change behavior dynamically
	pv.DriveBehavior = func() {
		fmt.Println("Modified: Now driving off-road")
	}
	pv.Drive()
}
