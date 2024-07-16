//--Summary:
//  Create a program that directs vehicles at a mechanic shop
//  to the correct vehicle lift, based on vehicle size.
//
//--Requirements:
//* The shop has lifts for multiple vehicle sizes/types:
//  - Motorcycles: small lifts
//  - Cars: standard lifts
//  - Trucks: large lifts
//* Write a single function to handle all of the vehicles
//  that the shop works on.
//* Vehicles have a model name in addition to the vehicle type:
//  - Example: "Truck" is the vehicle type, "Road Devourer" is a model name
//* Direct at least 1 of each vehicle type to the correct
//  lift, and print out the vehicle information.
//
//--Notes:
//* Use any names for vehicle models

package main

import "fmt"

//* The shop has lifts for multiple vehicle sizes/types:
//  - Motorcycles: small lifts
//  - Cars: standard lifts
//  - Trucks: large lifts

const (
	SmollLift = iota
	StandardLift
	LargeLift
)

type Lift int

type LiftDirecter interface {
	LiftSize() Lift
}

type Model interface {
	ModelName() string
}
type Motorcycle string
type Car        string
type Truck      string

//* Vehicles have a model name in addition to the vehicle type:
func (m Motorcycle) String() string {
	return fmt.Sprintf("Motorcycle: %v", string(m))
}

//* Vehicles have a model name in addition to the vehicle type:
func (c Car) String() string {
	return fmt.Sprintf("Car: %v", string(c))
}

//* Vehicles have a model name in addition to the vehicle type:
func (t Truck) String() string {
	return fmt.Sprintf("Truck: %v", string(t))
}

func (m Motorcycle) LiftSize() Lift {
	return SmollLift
}

func (c Car) LiftSize() Lift {
	return StandardLift
}

func (t Truck) LiftSize() Lift {
	return LargeLift
}

//* Write a single function to handle all of the vehicles
//  that the shop works on.
func sendToLift(p LiftDirecter) {
	switch p.LiftSize() {
	case SmollLift:
		fmt.Printf("send %v to small lift\n", p)
	case StandardLift:
		fmt.Printf("send %v to standard lift\n", p)
	case LargeLift:
		fmt.Printf("send %v to large lift\n", p)
	}
}

func main() {
	car := Car("Sporty")
	truck := Truck("MountainCrusher")
	motorcycle := Motorcycle("Croozer")
	//* Direct at least 1 of each vehicle type to the correct
	//  lift, and print out the vehicle information.
	sendToLift(car)
	sendToLift(truck)
	sendToLift(motorcycle)
}
