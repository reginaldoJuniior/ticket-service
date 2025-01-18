package repository

import "ticket-inventory/model"

type InventoryRepository interface {
	FindAvailableSeats() (int, error)
	AddStationToRoute(station model.Station, route model.Route) error
}
