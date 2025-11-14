# Adapter Pattern in Go (Interview Revision)

## 1. What is the Adapter Pattern?
The **Adapter Pattern** is a structural design pattern that allows incompatible interfaces to work together.  
It converts the interface of one class into another interface clients expect.

**Use When:**
- You want to use an existing class but its interface doesn’t match what your code expects.
- You want to integrate third‑party libraries without modifying them.

---

## 2. Real‑world Analogy
A laptop charger plug does not fit every country’s socket.  
An **adapter** converts the plug shape so your laptop can be used anywhere.

---

## 3. Components of Adapter Pattern
| Component | Description |
|----------|-------------|
| **Target Interface** | The interface your code expects |
| **Adaptee** | Existing/legacy library with incompatible interface |
| **Adapter** | Converts the Adaptee to the Target Interface |

---

## 4. Go Example (Clean & Interview Friendly)

### 🧩 Step 1: Define the Target Interface
```go
// Target interface expected by the client
type Payment interface {
    Pay(amount float64)
}
```

### 🧩 Step 2: Existing System (Adaptee)
```go
// Adaptee with incompatible method
type Razorpay struct {}

func (r *Razorpay) MakePaymentInPaise(amountInPaise int) {
    fmt.Println("Paid using Razorpay:", amountInPaise)
}
```

### 🧩 Step 3: Create the Adapter
```go
// Adapter converts Pay(float64) to MakePaymentInPaise(int)
type RazorpayAdapter struct {
    razor *Razorpay
}

func (ra *RazorpayAdapter) Pay(amount float64) {
    paise := int(amount * 100)
    ra.razor.MakePaymentInPaise(paise)
}
```

### 🧩 Step 4: Client Code
```go
func main() {
    r := &Razorpay{}
    adapter := &RazorpayAdapter{razor: r}

    adapter.Pay(499.99)
}
```

---

## 5. What Interviewers Want to Hear
- Adapter pattern helps **integrate incompatible systems** without modifying existing code.
- Encourages **loose coupling** and **clean integration**.
- Common use cases:
  - Payment gateway integrations
  - Logging frameworks
  - Wrapping external APIs
  - Legacy system migration

---

## 6. Adapter Pattern UML (Simple)
```
Client → Target Interface → Adapter → Adaptee
```

---

## 7. When Not to Use Adapter Pattern
- When you have control over both interfaces (better to refactor).
- When too many adapters cause complexity → consider a facade.

---

## 8. One‑Line Summary for Interview
**Adapter Pattern lets incompatible interfaces work together by providing a wrapper (adapter) that translates one interface into another.**
