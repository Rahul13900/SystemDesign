# Strategy Pattern in Go – Interview Quick Reference

This document demonstrates the **Strategy Design Pattern** using Go, with a real-world use case: vehicle driving behavior. It's tailored to help you **explain and implement the pattern during interviews**, with clean code and practical explanations.

---

## 🎯 Why Strategy Pattern?

The **Strategy Pattern** is used to define a family of interchangeable algorithms (strategies), encapsulate each one, and make them interchangeable at runtime. This promotes:

- ✅ **Separation of concerns** – Behavior logic is separate from core object structure
- ✅ **Runtime flexibility** – Behavior can change without modifying structs
- ✅ **Open/Closed Principle** – Open to new strategies, closed to modifying base types
- ✅ **Testability** – Strategies can be independently tested

---

## 🛠 Use Case: Vehicle Drive Behavior

We model three types of vehicles:

- **PassengerVehicle** → drives normally  
- **OffroadVehicle** → drives off-road  
- **SportsVehicle** → drives fast/agile  

Each vehicle has a drive capability, which varies based on type.

---

## 🚫 Without Strategy Pattern

All logic resides in child structs overriding parent behavior (tight coupling):

```go
package main

import "fmt"

type Vehicle interface {
	Drive()
}

type PassengerVehicle struct{}

func (p *PassengerVehicle) Drive() {
	fmt.Println("Driving on regular roads")
}

type OffroadVehicle struct{}

func (o *OffroadVehicle) Drive() {
	fmt.Println("Driving on rough off-road terrain")
}

type SportsVehicle struct{}

func (s *SportsVehicle) Drive() {
	fmt.Println("Driving at high speed with agility")
}

func main() {
	var v Vehicle

	v = &PassengerVehicle{}
	v.Drive()

	v = &OffroadVehicle{}
	v.Drive()

	v = &SportsVehicle{}
	v.Drive()
}
