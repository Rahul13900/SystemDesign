package main

import "fmt"

// Strategy interface
type DriveStrategy interface {
	Drive()
}

// Concrete strategies
type NormalDrive struct{}
func (n *NormalDrive) Drive() { fmt.Println("Driving on regular roads") }

type OffroadDrive struct{}
func (o *OffroadDrive) Drive() { fmt.Println("Driving on rough off-road terrain") }

type SportsDrive struct{}
func (s *SportsDrive) Drive() { fmt.Println("Driving at high speed with agility") }

// Vehicle uses strategy
type Vehicle struct {
	driveStrategy DriveStrategy
}

func (v *Vehicle) SetDriveStrategy(ds DriveStrategy) {
	v.driveStrategy = ds
}

func (v *Vehicle) Drive() {
	v.driveStrategy.Drive()
}

func main() {
	v := &Vehicle{}

	v.SetDriveStrategy(&NormalDrive{})
	v.Drive()

	v.SetDriveStrategy(&OffroadDrive{})
	v.Drive()

	v.SetDriveStrategy(&SportsDrive{})
	v.Drive()
}
