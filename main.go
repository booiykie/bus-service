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
	wallet := Wallet{1, 30.00, time.Now(), true, "12345678"}
	profile := Profile{"Wisani", "ZAID", "###############", "Bus", "Driver"}
	user := User{"1223", profile, wallet}
	driver := Entity{user,  busservice.Operator, 1}
	
	expressLine := Bus{
		Name: "Famba Kahle",
	    RegNumber: "DX13NMGP",
	    BusNumber: "Famba 1",
	    RouteId: 1,
	    Capacity: 48,
	    Operator: driver,
	}

	// Assign seat numbers to passengers.
	var seatNumber uint8
	
	entity1 := Entity{User: User{"001", profile, wallet},  UserType: "Passenger"}
	expressLine.Add(entity1)
	expressLine.UpdatePassengers(func(p *busservice.Entity) {
		seatNumber++
		p.SeatNumber = seatNumber
	})

	entity2 := Entity{User: User{"002", profile, wallet},  UserType: "Passenger"}
	expressLine.Add(entity2)
	expressLine.UpdatePassengers(func(p *busservice.Entity) {
		seatNumber++
		p.SeatNumber = seatNumber
	})

	userIDs := make([]string, 0)

	expressLine.VisitPassengers(func(p busservice.Entity) { userIDs = append(userIDs, p.User.ID) })


	fmt.Printf("This bus carries %d passengers, here are their userID's: %v\n", len(userIDs), userIDs)
}
