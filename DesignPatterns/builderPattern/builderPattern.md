# Builder Pattern in Golang

## ✅ What is the Builder Pattern?

The **Builder Pattern** is a creational design pattern used to construct **complex objects step-by-step**.  
It helps when you have:
- Many optional fields  
- Multiple ways to build the same object  
- Long or telescoping constructors  

---

## 🎯 When to Use the Builder Pattern
Use it when:
- Object creation is complex  
- You want readable, chained initialization  
- You want to avoid constructors with too many parameters  

Examples:
- User creation  
- Configuration objects  
- HTTP requests  
- Query builders  

---

# 🧩 Example: Building a `User` Object

```go
type User struct {
    Name     string
    Age      int
    Email    string
    Address  string
    IsActive bool
}
```

---

# 🔨 Builder Struct

```go
type UserBuilder struct {
    name     string
    age      int
    email    string
    address  string
    isActive bool
}

func NewUserBuilder() *UserBuilder {
    return &UserBuilder{}
}
```

---

# 🔧 Step-by-Step Builder Methods

```go
func (b *UserBuilder) SetName(name string) *UserBuilder {
    b.name = name
    return b
}

func (b *UserBuilder) SetAge(age int) *UserBuilder {
    b.age = age
    return b
}

func (b *UserBuilder) SetEmail(email string) *UserBuilder {
    b.email = email
    return b
}

func (b *UserBuilder) SetAddress(address string) *UserBuilder {
    b.address = address
    return b
}

func (b *UserBuilder) SetActive(isActive bool) *UserBuilder {
    b.isActive = isActive
    return b
}
```

---

# 🧱 Build Method

```go
func (b *UserBuilder) Build() *User {
    return &User{
        Name:     b.name,
        Age:      b.age,
        Email:    b.email,
        Address:  b.address,
        IsActive: b.isActive,
    }
}
```

---

# 🟢 Usage Example

```go
func main() {
    user := NewUserBuilder().
        SetName("Rahul").
        SetAge(25).
        SetEmail("rahul@example.com").
        SetActive(true).
        Build()

    fmt.Println(user)
}
```

---

# 🎓 Interview Explanation (Short Answer)

> “The Builder Pattern constructs complex objects step-by-step. In Go, we use method chaining with a `Build()` method that returns the final struct. It makes object creation clean and helps manage optional fields.”

---

# 📌 Advantages

| Advantage | Description |
|----------|-------------|
| Clean object creation | No long constructors |
| Optional parameters | Set only what you need |
| Fluent chaining | More readable |
| Good for complex objects | Helps maintain clarity |

---

# ❌ Disadvantages

| Drawback | Reason |
|----------|--------|
| Extra code | Builder must be created |
| Overkill for small structs | Simple structs don't need it |

---

# 🔥 Alternative: Functional Options Builder (Go Idiomatic)

```go
type UserOption func(*User)

func WithEmail(email string) UserOption {
    return func(u *User) { u.Email = email }
}

func WithAddress(addr string) UserOption {
    return func(u *User) { u.Address = addr }
}

func NewUser(name string, opts ...UserOption) *User {
    u := &User{Name: name}
    for _, opt := range opts {
        opt(u)
    }
    return u
}
```

### Usage:
```go
user := NewUser("Rahul", WithEmail("rahul@example.com"), WithAddress("Delhi"))
```

---

# 📝 Builder vs Functional Options

| Feature | Builder Pattern | Functional Options |
|--------|------------------|---------------------|
| Method chaining | Yes | No |
| Idiomatic Go | Medium | High |
| Validation before build | Easy | Needs checks |
| Many optional params | Great | Great |

---

# 🎯 Summary

- Builder is great when object creation is **step-by-step and complex**.  
- Functional Options are often more idiomatic in Go but Builder is more explicit.  
- Both help avoid long constructors and improve readability.

---

# 📚 Final Interview Tip

> “If the struct is large and requires step-by-step initialization, I use the Builder Pattern. If I need more Go idiomatic flexibility, I choose the Functional Options pattern.”

