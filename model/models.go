package model

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

type Booking struct {
	ID        string
	Passenger Passenger
	ServiceID string
	Seat      string
}
