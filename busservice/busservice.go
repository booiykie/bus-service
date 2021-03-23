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
    Passengers []Entity
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

