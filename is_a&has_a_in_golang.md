# Modeling "Is-A" and "Has-A" Relationships in Golang

Golang doesn't have traditional OOP inheritance, but you can still model relationships like **"is-a"** and **"has-a"** effectively using **interfaces** and **composition**.

---

## 🔹 "Is-A" Relationship (Inheritance-like)

### 🧠 Concept:
> A `Dog` **is-a** `Animal`

### ✅ In Go: Use **interfaces**

Interfaces let you define behavior that types must implement, modeling subtype relationships without inheritance.

```go
// Define an interface
type Animal interface {
	Speak() string
}

// Dog implements Animal
type Dog struct{}

func (d Dog) Speak() string {
	return "Woof"
}

// Cat implements Animal
type Cat struct{}

func (c Cat) Speak() string {
	return "Meow"
}

// Function that accepts Animal interface
func MakeItSpeak(a Animal) {
	fmt.Println(a.Speak())
}
```

```go
MakeItSpeak(Dog{}) // Woof
MakeItSpeak(Cat{}) // Meow
```

### ✅ Summary:
- **"Is-A"** is achieved using interfaces.
- If a type implements an interface, it "is-a" that interface type.

---

## 🔸 "Has-A" Relationship (Composition)

### 🧠 Concept:
> A `Car` **has-a** `Engine`

### ✅ In Go: Use **struct fields or embedding**

#### 📌 Example 1: Composition via fields
```go
type Engine struct {
	HorsePower int
}

type Car struct {
	Engine Engine // Car has-a Engine
}
```

#### 📌 Example 2: Composition via embedding
```go
type Logger struct{}

func (l Logger) Log(msg string) {
	fmt.Println("[LOG]:", msg)
}

type Service struct {
	Logger // embedded
}

func main() {
	s := Service{}
	s.Log("Service started!") // can call Logger methods directly
}
```

### ✅ Summary:
- **"Has-A"** is modeled using:
  - Struct fields (for ownership/containment)
  - Struct embedding (for behavior reuse)

---

## 🔄 Quick Comparison Table

| Relationship | Traditional OOP | Go Equivalent | Notes |
|--------------|------------------|----------------|-------|
| **Is-A**     | Inheritance (`extends`) | Interfaces | Substitution of types via behavior |
| **Has-A**    | Composition (`has-a` via fields) | Struct fields or embedding | Reuse and contain functionality |

---

## 🔥 TL;DR Cheat Sheet

| Concept       | Traditional OOP       | Go Equivalent             |
|---------------|------------------------|----------------------------|
| Inheritance   | `extends`             | Interfaces + composition   |
| Interface     | `implements`          | Implicit interface match   |
| Composition   | `has-a` via fields    | Struct fields or embedding |

---

