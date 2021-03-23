package busservice

import (
    // "fmt"
    "time"
)

type EntityType string

const(
    Passenger EntityType = "Passenger"
    Operator = "Operator"
    Marshall = "Marshall"
    Admin = "Administrator"
    PayMaster = "Pay Master"
)

type Entity struct {
    User User
    UserType string
    SeatNumber uint8
}

type Bus struct {
    Name  string
    RegNumber string
    BusNumber string
    passengers map[string]Entity
    RouteId uint8
    Capacity uint8
    Operator Entity
}

type Profile struct {
    Username string
    IDType string
    IdentificactionNumber string
    FirstName string
    Surname string
}

type Wallet struct {
    ID int32
    Balance float32
    LastUpdate time.Time
    Active bool
    AccountId string
}

type User struct {
    ID string
    Profile Profile
    Wallet Wallet
}

func (bus *Bus) Add(passenger Entity) {
    if bus.passengers == nil {
        bus.passengers = make(map[string]Entity)
    }
    bus.passengers[passenger.User.ID] = passenger
}

func (bus *Bus) VisitPassengers(visitor func(Entity)) {
    for _, p := range bus.passengers {
        visitor(p)
    }
}

// FindPassenger returns the Passenger that matches the given SSN, if found. Otherwise, an empty Passenger is returned.
func (bus *Bus) FindPassenger(userId string) Entity {
    if p, ok := bus.passengers[userId]; ok {
        return p
    }
    return Entity{} // A nobody.
}

// UpdatePassengers calls function visitor for each Passenger on the bus. Passengers are passed by reference and may be modified.
func (b *Bus) UpdatePassengers(visitor func(*Entity)) {
    ps := make(map[string]Entity, len(b.passengers))
    for userId, p := range b.passengers {
        visitor(&p)
        ps[userId] = p
    }
    b.passengers = ps
}

