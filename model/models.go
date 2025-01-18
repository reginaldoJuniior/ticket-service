package model

// Define structs in Golang for Station, Route, Service, Carriage, Seat,
// Passenger, Ticket, and Booking.
type Passenger struct {
	Name    string
	Tickets []Ticket
}

type Ticket struct {
	Code  string
	Seat  Seat
	Buyer Passenger
	Class string
}

type Booking struct {
	Passenger Passenger
	Ticket    Ticket
}
