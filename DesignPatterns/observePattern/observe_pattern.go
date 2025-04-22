package main

import "fmt"

////////////////////////////
// Observer Pattern Components
////////////////////////////

// Observer interface defines the method that all observers must implement.
type Observer interface {
	Update(orderID string, status string)
}

// Subject interface defines methods for managing observers and notifying them.
type Subject interface {
	Register(observer Observer)
	Deregister(observer Observer)
	NotifyAll()
}

////////////////////////////
// Concrete Subject
////////////////////////////

// OrderService is the concrete subject that holds order data and notifies observers.
type OrderService struct {
	observers []Observer
	orderID   string
	status    string
}

// Register adds a new observer to the list.
func (o *OrderService) Register(observer Observer) {
	o.observers = append(o.observers, observer)
}

// Deregister removes an observer from the list.
func (o *OrderService) Deregister(observer Observer) {
	for i, obs := range o.observers {
		if obs == observer {
			o.observers = append(o.observers[:i], o.observers[i+1:]...)
			break
		}
	}
}

// NotifyAll informs all registered observers of the current status.
func (o *OrderService) NotifyAll() {
	for _, observer := range o.observers {
		observer.Update(o.orderID, o.status)
	}
}

// UpdateStatus sets a new status and notifies observers.
// This is a domain-specific method (not part of Subject interface).
func (o *OrderService) UpdateStatus(status string) {
	o.status = status
	fmt.Println("\n Order status updated to:", status)
	o.NotifyAll()
}

////////////////////////////
// Concrete Observers
////////////////////////////

// EmailNotifier is a concrete observer that simulates sending an email.
type EmailNotifier struct {
	email string
}

// Update method for EmailNotifier.
func (e *EmailNotifier) Update(orderID string, status string) {
	fmt.Printf("📧 Email to %s: Order %s is now '%s'\n", e.email, orderID, status)
}

// SMSNotifier is another concrete observer that simulates sending an SMS.
type SMSNotifier struct {
	phone string
}

// Update method for SMSNotifier.
func (s *SMSNotifier) Update(orderID string, status string) {
	fmt.Printf("📱 SMS to %s: Order %s is now '%s'\n", s.phone, orderID, status)
}

////////////////////////////
// Usage Example
////////////////////////////

func main() {
	// Create an order with ID "ORD123"
	orderService := &OrderService{orderID: "ORD123"}

	// Create two observers
	emailNotif := &EmailNotifier{email: "user@example.com"}
	smsNotif := &SMSNotifier{phone: "+1234567890"}

	// Register observers to the order service
	orderService.Register(emailNotif)
	orderService.Register(smsNotif)

	// Simulate order status changes
	orderService.UpdateStatus("Placed")
	orderService.UpdateStatus("Shipped")

	// Deregister one observer (e.g., user opted out of SMS)
	orderService.Deregister(smsNotif)

	// Another status update, only email should receive this
	orderService.UpdateStatus("Delivered")
}
