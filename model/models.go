package model

import "errors"

type Station struct {
	Name string
}

type Route struct {
	ID    string
	Stops []Station
}

type Service struct {
	ID      string
	RouteID string
	Date    string
}

type Carriage struct {
	ID    string
	Seats []Seat
}

type Seat struct {
	ID          string
	ComfortZone string // First-class, Second-class
}

type Passenger struct {
	Name string
}

func (p Passenger) Validate() (bool, error) {
	if p.Name == "" {
		return false, errors.New("passenger name is required")
	}
	return true, nil
}

type Booking struct {
	ID          string
	Passenger   Passenger
	ServiceID   string
	Seat        string
	Origin      string
	Destination string
}

func (b *Booking) Validate() (bool, error) {
	if b.ID == "" {
		return false, errors.New("booking ID is required")
	}
	if b.ServiceID == "" {
		return false, errors.New("service ID is required")
	}
	if b.Seat == "" {
		return false, errors.New("seat is required")
	}
	if b.Origin == "" {
		return false, errors.New("origin is required")
	}
	if b.Destination == "" {
		return false, errors.New("destination is required")
	}

	validate, err := b.Passenger.Validate()
	if err != nil {
		return false, err
	}

	return validate, nil
}
