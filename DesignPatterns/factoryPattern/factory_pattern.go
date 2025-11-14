package main

import "fmt"

type Car interface {
	getCar() string
}

type Sedan struct {
	Name string
}
type Suv struct {
	Name string
}

func NewSedan() *Sedan {
	return &Sedan{}
}

func NewSuv() *Suv {
	return &Suv{}
}

func (sd Sedan) getCar() string {
	sd.Name = "Honda City"
	return sd.Name
}

func (s Suv) getCar() string {
	s.Name = "XUV700"
	return s.Name
}

func CarFactory(n int) { // factory
	var car Car
	switch n {
	case 1:
		car = NewSedan()
	case 2:
		car = NewSuv()
	default:
		fmt.Println("Wrong Input")
	}
	carName := car.getCar()
	fmt.Println(carName)

}

func main() { // client
	CarFactory(1)
	CarFactory(2)
}
