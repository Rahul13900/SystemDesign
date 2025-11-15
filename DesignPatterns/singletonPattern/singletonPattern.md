# Singleton Pattern in Golang

## ✅ What is the Singleton Pattern?

The **Singleton Pattern** ensures that only **one instance** of a struct exists throughout the application and provides a **global point of access** to this instance.

It is commonly used for:
- Configuration managers  
- Loggers  
- Database connections  
- Caches  

---

## ⚠️ Why Singleton Needs Care in Go

Go is highly concurrent.  
A naive singleton implementation can easily create **multiple instances** when accessed by multiple goroutines at the same time.

Hence, thread-safety is critical.

---

# 1️⃣ Naive Singleton (❌ Not Thread-Safe)

```go
package singleton

type Singleton struct {
    Value string
}

var instance *Singleton

func GetInstance() *Singleton {
    if instance == nil {
        instance = &Singleton{}
    }
    return instance
}
```

### ❌ Problems
- Race conditions
- Multiple instances may get created

---

# 2️⃣ Mutex-Based Singleton (✔️ Thread-Safe, ❌ Slower)

```go
package singleton

import "sync"

type Singleton struct {
    Value string
}

var (
    instance *Singleton
    mu       sync.Mutex
)

func GetInstance() *Singleton {
    mu.Lock()
    defer mu.Unlock()

    if instance == nil {
        instance = &Singleton{}
    }
    return instance
}
```

### ✔️ Benefits
- Thread-safe  

### ❌ Drawbacks
- `mutex.Lock()` happens **every time**  
- Slight unnecessary overhead after initialization

---

# 3️⃣ Best Approach: sync.Once Singleton (🔥 Idiomatic Go)

```go
package singleton

import "sync"

type Singleton struct {
    Value string
}

var (
    instance *Singleton
    once     sync.Once
)

func GetInstance() *Singleton {
    once.Do(func() {
        instance = &Singleton{}
    })
    return instance
}
```

### ✔️ Advantages
- Runs initialization **only once**
- Zero locking after initialization
- Fast and idiomatic
- Perfect for production

---

# 🧠 Example Usage

```go
func main() {
    s1 := GetInstance()
    s2 := GetInstance()

    s1.Value = "Hello"

    fmt.Println(s2.Value) // Output: Hello
    fmt.Println(s1 == s2) // Output: true
}
```

---

# 🎯 Interview-Ready Explanation

> “The Singleton pattern in Go is best implemented using `sync.Once`, which guarantees that initialization happens exactly one time even under heavy concurrency. This avoids race conditions and eliminates unnecessary locking after initialization, making it both thread-safe and efficient.”

---

# 📝 When to Use Singleton

Use Singleton when you need:
- A **single global instance**
- Controlled initialization
- Shared state between components

Examples:
- Logger service  
- DB connection pool  
- Config loader  

---

# ❌ When NOT to Use Singleton

Avoid Singletons when:
- You need multiple independent instances
- You want easy testability  
- You want flexible dependency injection  

Singleton often introduces hidden globals → can make testing harder.

---

# 📌 Summary Table

| Version | Thread-Safe | Performance | Recommended |
|--------|-------------|-------------|-------------|
| Naive | ❌ No | ⭐ Fast | ❌ No |
| Mutex | ✔️ Yes | ⭐⭐ Medium | ❌ Ok but not ideal |
| sync.Once | ✔️ Yes | ⭐⭐⭐ Best | ✅ Yes |

---

# 🎓 Final Tip for Interviews

Always say **sync.Once** because it is the idiomatic and optimal Go approach.

---

