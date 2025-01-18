package model

import "time"

type Station struct {
	Name        string
	Connections []Station
}

type Route struct {
	From     Station
	To       Station
	Distance int
}

type Service struct {
	Route    Route
	Duration time.Time
	Train    Train
}

type Train struct {
	Carriage []Carriage
}

type Carriage struct {
	FirstClassSeats []Seat
	EconomicSeats   []Seat
}

type Seat struct {
	Code   string
	Booked bool
}
