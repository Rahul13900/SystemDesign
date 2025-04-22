
# Abstract Factory Design Pattern in Go

## Overview

The **Abstract Factory Design Pattern** is a creational design pattern that provides an interface for creating **families of related or dependent objects** without specifying their concrete classes.

It is essentially a **factory of factories**, where each factory is responsible for creating a set of related objects.

---

## Real-World Analogy

Imagine you're furnishing a house and want everything to match a certain style — like **Modern** or **Victorian**. Instead of picking each item individually, you go to a store section and say:

> “Give me the Modern furniture set.”

The store (Abstract Factory) gives you a **Modern Chair**, **Modern Sofa**, and **Modern Coffee Table** — all matching and designed to go together.

You don’t worry about how they’re made. You just know that you’re getting a set that works together.

---

## When to Use

- When your system needs to create multiple related objects (families).
- When you want to ensure consistency among related objects.
- When the system should be independent of how products are created, composed, and represented.

---

## Structure

### Interfaces

- Define interfaces for each type of product (e.g., `Chair`, `Sofa`).
- Define an interface for the factory that creates those products.

### Concrete Products

- Implement the product interfaces for each variant (e.g., `ModernChair`, `VictorianChair`).

### Concrete Factories

- Implement the abstract factory interface, creating specific variants of products (e.g., `ModernFurnitureFactory`, `VictorianFurnitureFactory`).

### Client

- Uses only the abstract interfaces and factories to get the required product families.

---

## Benefits

1. **Encapsulation of Object Creation**: Client code doesn’t depend on specific classes.
2. **Product Consistency**: Guarantees that products from the same family match.
3. **Scalability**: You can easily introduce new families of products without changing existing code.
4. **Adherence to SOLID Principles**: Especially the Open/Closed and Dependency Inversion principles.

---

## Example in Go

You define:

- `Chair` and `Sofa` interfaces
- `ModernChair`, `ModernSofa`, `VictorianChair`, and `VictorianSofa` structs
- `FurnitureFactory` interface with `CreateChair()` and `CreateSofa()`
- `ModernFurnitureFactory` and `VictorianFurnitureFactory` structs implementing that factory interface

The **client** calls only the factory interface and doesn't know what actual structs are being used.

```go
factory := ModernFurnitureFactory{}
chair := factory.CreateChair()
sofa := factory.CreateSofa()
```

---

## Summary

The **Abstract Factory Pattern** provides a way to encapsulate a group of individual factories with a common theme. It promotes consistency and scalability when creating families of related objects, especially in applications that must remain independent of how their objects are created and represented.

---

## Further Reading

- [Go Patterns Reference](https://github.com/tmrts/go-patterns)
- [Refactoring Guru: Abstract Factory](https://refactoring.guru/design-patterns/abstract-factory)
