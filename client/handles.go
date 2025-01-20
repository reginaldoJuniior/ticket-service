package client

import "ticket-inventory/model"

type HTTPResponse struct {
	StatusCode int
	Body       any
}

func (r *HTTPResponse) GetStatusCode() int {
	return r.StatusCode
}

func (r *HTTPResponse) GetBody() any {
	return r.Body
}

type BookHandle interface {
	GetAllBookings() []model.Booking
	GetAllServices() []model.Service
	GetAllStations() []model.Station
	CreateBooking(booking model.Booking) error
	GetPassengersByOrigin(stationName string) ([]model.Passenger, error)
	GetPassengerBySeat(serviceID, seatID string) (*model.Passenger, error)
	GetPassengersByDestination(stationName string) ([]model.Passenger, error)
}
