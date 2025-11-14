package main

import "fmt"

type mobile interface {
	chargeAppleMobile()
}

type apple struct{}

func (a *apple) chargeAppleMobile() {
	fmt.Println("Charging APPLE mobile")
}

// adaptee, which actually implements the specific functionality
type andriod struct{}

func (a *andriod) chargeAndriodMobile() {
	fmt.Println("Charging ANDRIOD mobile")
}

// adapter
type andriodAdapter struct {
	andriod *andriod
}

// now this adapter has to implement interface so that we can reach it through same chargeMobile rather then writing a new
func (ad *andriodAdapter) chargeAppleMobile() {
	ad.andriod.chargeAndriodMobile()
}

type client struct{}

func (c *client) chargeMobile(mob mobile) { // we are declaring it as mobile type and sending apple type this is possible becuase of interface implementation
	mob.chargeAppleMobile()
}

func main() {
	// 1st requirement
	apple := &apple{}
	client := &client{}
	client.chargeMobile(apple)

	// implemented Adapter pattern to satisfy 2nd requirement
	andriod := &andriodAdapter{}
	client.chargeMobile(andriod)
}
