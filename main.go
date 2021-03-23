package main

import (
	"fmt"
	"time"
	"github.com/booiykie/bus-service/busservice"
)

type Wallet = busservice.Wallet
type Profile = busservice.Profile
type User = busservice.User
type Entity = busservice.Entity
type Bus = busservice.Bus


func main() {
	// expressLine := busservice.Bus{"Express Line"}
	wallet := Wallet{1, 30.00, time.Now(), true, "12345678"}
	profile := Profile{"Wisani", "ZAID", "###############", "Bus", "Driver"}
	user := User{"1223", profile, wallet}
	driver := Entity{user,  busservice.Operator, 1}
	passengers := make([]Entity, 0)
	
	expressLine := Bus{
		Name: "Famba Kahle",
	    RegNumber: "DX13NMGP",
	    BusNumber: "Famba 1",
	    Passengers: passengers,
	    RouteId: 1,
	    Capacity: 48,
	    Operator: driver,
	}
	// expressLine.Passengers := make([]busservice.Passenger, 0)
	expressLine.Passengers = append(expressLine.Passengers, Entity{User: User{"001", profile, wallet},  UserType: "Passenger"})
	expressLine.Passengers = append(expressLine.Passengers, Entity{User: User{"002", profile, wallet},  UserType: "Passenger"})

	// Get a manifest!
	newPassengers := make([]string, 0)
	for _, p := range expressLine.Passengers {
		newPassengers = append(newPassengers, p.User.ID)
	}
	fmt.Printf("This bus carries %d passengers, here are their userID's: %v\n", len(newPassengers), newPassengers)
}
