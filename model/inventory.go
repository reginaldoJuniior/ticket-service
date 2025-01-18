package model

import (
	"fmt"
	"time"
)

type Station struct {
	Name        string
	Connections []Station
}

type Route struct {
	From      Station
	To        Station
	NodeRoute NodeRoute
}

type Service struct {
	Code     string
	Route    Route
	Duration time.Time
	Train    Train
	Hour     time.Time
}

func (s *Service) Key() string {
	return fmt.Sprintf("%s-%s-%s-%s-%d:%d",
		s.Code,
		s.Train.Code,
		s.Route.To.Name,
		s.Route.From.Name,
		s.Hour.Hour(),
		s.Hour.Minute())
}

type Train struct {
	Code     string
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
