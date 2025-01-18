package model

import "fmt"

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
	Service   Service
}

func (b *Booking) Key() string {
	return fmt.Sprintf("%s-%s-%s-%s",
		b.Passenger.Name,
		b.Ticket.Code,
		b.Ticket.Seat.Code,
		b.Service.Key())
}
