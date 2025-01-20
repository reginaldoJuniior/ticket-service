package repository

import (
	"errors"
	"fmt"
	"sync"
	"ticket-inventory/model"
)

const BookingNotFoundError = "booking not found"

var instance Reservations
var once sync.Once

// NewReservationRepository function to return a new instance of Reservations struct.
// This function will be used to create a singleton instance of Reservations struct.
func NewReservationRepository() *Reservations {
	once.Do(func() {
		instance = Reservations{
			data: map[string]any{
				"bookings": make([]model.Booking, 0),
				"routes":   loadDefaultRoutes(),
				"services": loadDefaultServices(),
			},
		}

	})
	return &instance
}

type Reservations struct {
	mutex sync.Mutex
	data  map[string]any
}

func loadDefaultRoutes() []model.Route {
	routes := make([]model.Route, 0, 2)
	routes = append(routes,
		model.Route{
			ID:    "route1",
			Stops: []model.Station{{"London"}, {"Ashford"}, {"Calais"}, {"Paris"}},
		},
		model.Route{
			ID:    "route2",
			Stops: []model.Station{{"Paris"}, {"Brussels"}, {"Rotterdam"}, {"Amsterdam"}},
		},
		model.Route{
			ID:    "route3",
			Stops: []model.Station{{"Amsterdam"}, {"Utrecht"}, {"Hanover"}, {"Berlin"}},
		})
	return routes
}

func loadDefaultServices() []model.Service {
	services := make([]model.Service, 0, 3)
	services = append(services,
		model.Service{
			ID:      "service1",
			RouteID: "route1",
		},
		model.Service{
			ID:      "service2",
			RouteID: "route2",
		},
		model.Service{
			ID:      "service3",
			RouteID: "route3",
		})
	return services
}

func (r *Reservations) FindSeat(code string, service model.Service) (model.Seat, error) {
	//TODO implement me
	panic("implement me")
}

func (r *Reservations) SaveBook(booking model.Booking) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	bookings := r.data["bookings"].([]model.Booking)
	for _, b := range bookings {
		if b.ServiceID == booking.ServiceID && b.Seat == booking.Seat {
			return fmt.Errorf("seat %s is already booked", booking.Seat.ID)
		}
	}
	r.data["bookings"] = append(bookings, booking)
	return nil
}

func (r *Reservations) RemoveBook(booking model.Booking) error {
	bookings := r.data["bookings"].([]model.Booking)
	for i, b := range bookings {
		if b.ServiceID == booking.ServiceID && b.Seat == booking.Seat {
			r.data["bookings"] = append(bookings[:i], bookings[i+1:]...)
			return nil
		}
	}
	return errors.New("booking not found")
}

func (r *Reservations) GetBookDetails(passenger model.Passenger) (*model.Booking, error) {
	bookings := r.data["bookings"].([]model.Booking)
	for _, b := range bookings {
		if b.Passenger.Name == passenger.Name {
			return &b, nil
		}
	}
	return nil, errors.New(BookingNotFoundError)
}

func (r *Reservations) GetAllBookings() []model.Booking {
	list := make([]model.Booking, 0, len(r.data["bookings"].([]model.Booking)))
	for _, v := range r.data["bookings"].([]model.Booking) {
		list = append(list, v)
	}
	return list
}

func (r *Reservations) FindBook(bookKey string) (*model.Booking, error) {
	bookings := r.data["bookings"].([]model.Booking)
	for _, b := range bookings {
		if b.ID == bookKey {
			return &b, nil
		}
	}
	return nil, errors.New(BookingNotFoundError)
}

func (r *Reservations) FindPassengerByOrigin(stationName string) ([]model.Passenger, error) {
	var passengers []model.Passenger

	bookings, ok := r.data["bookings"].([]model.Booking)
	if !ok {
		return passengers, errors.New("no bookings available")
	}

	for _, booking := range bookings {
		if booking.Origin == stationName {
			passengers = append(passengers, booking.Passenger)
		}
	}

	if len(passengers) == 0 {
		return passengers, errors.New("passenger not found")
	}

	return passengers, nil
}

func (r *Reservations) FindPassengerBySeat(serviceID, seatID string) (*model.Passenger, error) {
	bookings := r.data["bookings"].([]model.Booking)
	for _, booking := range bookings {
		if booking.ServiceID == serviceID && booking.Seat.ID == seatID {
			return &booking.Passenger, nil
		}
	}
	return nil, errors.New("passenger not found")
}

func (r *Reservations) GetAllServices() []model.Service {
	list := make([]model.Service, 0, len(r.data["services"].([]model.Service)))
	for _, v := range r.data["services"].([]model.Service) {
		list = append(list, v)
	}
	return list
}

func (r *Reservations) GetAllStations() []model.Station {
	list := make([]model.Station, 0, len(r.data["routes"].([]model.Route)))
	for _, route := range r.data["routes"].([]model.Route) {
		for _, station := range route.Stops {
			list = append(list, station)
		}
	}
	return list
}

func (r *Reservations) FindServiceByID(serviceID string) (*model.Service, error) {
	services := r.data["services"].([]model.Service)
	for _, service := range services {
		if service.ID == serviceID {
			return &service, nil
		}
	}
	return nil, errors.New("service not found")
}
