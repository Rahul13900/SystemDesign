# SOLID Principles Simplified with Examples

This guide breaks down the SOLID principles in a simplified manner with easy-to-understand Golang examples and helpful comments.

---

## 🧱 What is SOLID?

**SOLID** is an acronym for **five design principles** that help you write **better, more maintainable, and scalable code** — especially in object-oriented programming, but they’re useful in all paradigms.

---

## 1. S – Single Responsibility Principle (SRP)

> A class/module should have **only one reason to change**.
> It should have only one reponsiblity
>If one component does too much, a change in one responsibility might accidentally affect the other. It’s hard to test, harder to maintain.

### ✅ Do one thing well. Don't mix responsibilities.

### 🔴 Bad Example:
```go
// Violates SRP: User has multiple responsibilities

type User struct {
    Name string
}

func (u *User) SaveToDB() {
    // DB logic here
}

func (u *User) SendEmail() {
    // Email logic here
}
```

### 🟢 Good Example:
```go
// Each class/module has one job only

type User struct {
    Name string
}

type UserRepository struct{}

func (r *UserRepository) Save(u User) {
    // DB logic here
}

type EmailService struct{}

func (e *EmailService) SendWelcomeEmail(u User) {
    // Email logic here
}
```

---

## 2. O – Open/Closed Principle (OCP)

> Software entities should be **open for extension, but closed for modification**.

### ✅ You should be able to add new features without breaking existing code. This avoids bugs in already-tested code.

### 🔴 Bad Example:
```go
func CalculateDiscount(userType string) float64 {
    if userType == "regular" {
        return 5
    } else if userType == "premium" {
        return 10
    }
    return 0
}
```
>>Every time you want to add a new user type, you have to modify this function, which risks breaking it.

### 🟢 Good Example:
```go
// We can now extend this easily by adding new types

type DiscountStrategy interface {
    GetDiscount() float64
}

type RegularUser struct{}
func (r RegularUser) GetDiscount() float64 { return 5 }

type PremiumUser struct{}
func (p PremiumUser) GetDiscount() float64 { return 10 }
```
>You can add a VIPUser or EmployeeUser later — just implement DiscountStrategy without touching any existing logic.
---

## 3. L – Liskov Substitution Principle (LSP)

> Subtypes must be substitutable for their base types **without breaking behavior**.

### ✅ If you use a base class/interface, it should work with **any subclass** of that base.

### 🔴 Bad Example:
```go
type Bird interface {
    Fly()
}

type Ostrich struct{}

func (o Ostrich) Fly() {
    panic("Ostrich can't fly") // Violates LSP
}
```
>Ostrich technically is a bird, but it can’t fly — calling Fly() causes a runtime panic. Bad behavior!

### 🟢 Good Example:
```go
type Bird interface {
    MakeSound()
}

type FlyingBird interface {
    Bird
    Fly()
}

type Sparrow struct{}
func (s Sparrow) Fly() {}
func (s Sparrow) MakeSound() {}

type Ostrich struct{}
func (o Ostrich) MakeSound() {}
```

---

## 4. I – Interface Segregation Principle (ISP)

> Don’t force clients to depend on methods they don’t use.
> Large interfaces force unnecessary code and reduce flexibility.

### ✅ Create **small, specific interfaces**.

### 🔴 Bad Example:
```go
type Worker interface {
    Work()
    Eat()
}

type Robot struct{}

func (r Robot) Work() {}
func (r Robot) Eat() { panic("I don't eat") } // ❌
```

### 🟢 Good Example:
```go
type Workable interface {
    Work()
}

type Eatable interface {
    Eat()
}

type Robot struct{}
func (r Robot) Work() {}
```

---

## 5. D – Dependency Inversion Principle (DIP)

> High-level modules should not depend on low-level modules. Both should depend on abstractions.
>Code becomes more testable, modular, and swappable — especially for things like databases, payment providers, etc.

### ✅ Depend on interfaces, not concrete implementations.

### 🔴 Bad Example:
```go
type MySQL struct{}
func (db MySQL) Save(data string) {}

type App struct {
    db MySQL // tightly coupled with MySQL
}
```
>If you want to switch to PostgreSQL, you have to change the App code.

### 🟢 Good Example:
```go
type Database interface {
    Save(data string)
}

type MySQL struct{}
func (db MySQL) Save(data string) {}

type App struct {
    db Database
}
```
>Now App works with any Database implementation — MySQL, PostgreSQL, even mock DBs for testing!

---

## 🎯 Easy Way to Remember SOLID:

| Letter | Principle                    | Quick Hint                |
|--------|------------------------------|---------------------------|
| **S**  | Single Responsibility         | One job per class         |
| **O**  | Open/Closed                  | Add, don’t change         |
| **L**  | Liskov Substitution          | Subtypes must behave      |
| **I**  | Interface Segregation        | Keep interfaces lean      |
| **D**  | Dependency Inversion         | Depend on abstractions    |

---
>LSP = "Don’t lie about what your type can do."
>ISP = "Don’t force types to do more than they should."

