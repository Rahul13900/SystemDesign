package main

import (
	"fmt"
	"time"
)

////////////////////////////
// Observer via Go Channels
////////////////////////////

// Notification represents the event/message that gets sent.
type Notification struct {
	OrderID string
	Status  string
}

// ObserverChannel is an observer with its own channel to receive notifications.
type ObserverChannel struct {
	Name    string
	Channel chan Notification
	Done    chan bool // optional: for graceful shutdown
}

// Listen continuously waits for updates and handles them.
func (o *ObserverChannel) Listen() {
	for {
		select {
		case notif := <-o.Channel:
			fmt.Printf("%s received update: Order %s is now '%s'\n", o.Name, notif.OrderID, notif.Status)
		case <-o.Done:
			fmt.Printf("%s shutting down...\n", o.Name)
			return
		}
	}
}

////////////////////////////
// Subject (OrderService)
////////////////////////////

type OrderService struct {
	OrderID    string
	Status     string
	Subscribers []*ObserverChannel
}

// Register adds an observer to the list.
func (o *OrderService) Register(observer *ObserverChannel) {
	o.Subscribers = append(o.Subscribers, observer)
}

// NotifyAll pushes the new status to all observers via their channels.
func (o *OrderService) NotifyAll() {
	for _, obs := range o.Subscribers {
		obs.Channel <- Notification{OrderID: o.OrderID, Status: o.Status}
	}
}

// UpdateStatus sets the status and notifies all observers.
func (o *OrderService) UpdateStatus(status string) {
	o.Status = status
	fmt.Println("\nOrder status updated to:", status)
	o.NotifyAll()
}

////////////////////////////
// Main
////////////////////////////

func main() {
	// Initialize the order service
	order := &OrderService{OrderID: "ORD987"}

	// Create observer instances with their own channels
	emailObserver := &ObserverChannel{
		Name:    "📧 EmailNotifier",
		Channel: make(chan Notification),
		Done:    make(chan bool),
	}
	smsObserver := &ObserverChannel{
		Name:    "📱 SMSNotifier",
		Channel: make(chan Notification),
		Done:    make(chan bool),
	}

	// Register the observers
	order.Register(emailObserver)
	order.Register(smsObserver)

	// Start observers in separate goroutines
	go emailObserver.Listen()
	go smsObserver.Listen()

	// Simulate status updates
	order.UpdateStatus("Order Placed")
	time.Sleep(time.Millisecond * 500)

	order.UpdateStatus("Dispatched")
	time.Sleep(time.Millisecond * 500)

	order.UpdateStatus("Delivered")
	time.Sleep(time.Millisecond * 500)

	// Gracefully stop the observers
	emailObserver.Done <- true
	smsObserver.Done <- true

	// Wait to let goroutines print shutdown message
	time.Sleep(time.Millisecond * 200)
}
