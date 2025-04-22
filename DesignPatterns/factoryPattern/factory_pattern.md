
# Factory Design Pattern in Go

## Overview

The **Factory Design Pattern** is a creational design pattern that provides a way to **create objects without exposing the instantiation logic** to the client.

Instead of using the `new` keyword directly, the client calls a factory function which returns the appropriate object.

---

## Real-World Analogy

Imagine you go to a bakery and ask for a specific pastry — say a croissant or a muffin. You don’t make it yourself; the **baker** makes it and gives it to you.

Here, the **baker is the factory**, and you just get what you asked for.

---

## When to Use

- When you don’t want the client to know the specific class it needs to create.
- When you want to centralize object creation logic.
- When your code needs flexibility and extensibility with different types.

---

## Structure

### Step 1: Define an Interface

Define a common interface or base class for the products.

```go
type Vehicle interface {
	Drive() string
}
```

### Step 2: Implement the Interface

Create concrete types that implement this interface.

```go
type Car struct{}
func (c Car) Drive() string { return "Driving a car" }

type Bike struct{}
func (b Bike) Drive() string { return "Riding a bike" }
```

### Step 3: Create a Factory Function

The factory function takes input and returns the appropriate type.

```go
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
```

### Step 4: Use the Factory in Client Code

```go
func main() {
	vehicle := VehicleFactory("car")
	if vehicle != nil {
		fmt.Println(vehicle.Drive())  // Output: Driving a car
	} else {
		fmt.Println("Unknown vehicle type")
	}
}
```

---

## Benefits

1. **Encapsulation of Object Creation**: The creation logic is separated from usage.
2. **Easy to Introduce New Types**: You can add new types without changing the client code.
3. **Simplifies Testing**: You can mock the factory for testing.

---

## Summary

The **Factory Pattern** is a simple yet powerful way to manage object creation. It promotes loose coupling and enhances code maintainability by abstracting the object creation process from the business logic.

---

## Further Reading

- [Go Patterns Reference](https://github.com/tmrts/go-patterns)
- [Refactoring Guru: Factory Method](https://refactoring.guru/design-patterns/factory-method)
