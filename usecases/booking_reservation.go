package usecases

import (
	"errors"
	"ticket-inventory/model"
)

// ReservationRepository is an interface for booking reservation repository
type ReservationRepository interface {
	SaveBook(book model.Booking) error
	GetAllBookings() []model.Booking
	GetAllServices() []model.Service
	FindServiceByID(serviceID string) (*model.Service, error)
	GetAllStations() []model.Station
	FindPassengerByOrigin(stationName string) ([]model.Passenger, error)
	FindPassengerByDestination(stationName string) ([]model.Passenger, error)
	FindPassengerBySeat(serviceID, seatID string) (*model.Passenger, error)
	FindPassengerByServiceSeatDate(serviceID, seatID, date string) (model.Passenger, error)
	FindPassengerByOriginDestination(origin string, destination string) ([]model.Passenger, error)
}

// BookingReservation is a use case for booking reservation
type BookingReservation struct {
	reservationRepo ReservationRepository
}

func NewBookingReservation(reservationRepo ReservationRepository) *BookingReservation {
	return &BookingReservation{
		reservationRepo: reservationRepo,
	}
}

// CreateBooking simulates a booking creation
func (b *BookingReservation) CreateBooking(booking model.Booking) error {
	ok, err := booking.Validate()
	if !ok {
		return err
	}

	// Check if service exists
	_, err = b.reservationRepo.FindServiceByID(booking.ServiceID)
	if err != nil {
		return err
	}

	return b.reservationRepo.SaveBook(booking)
}

func (b *BookingReservation) GetAllBookings() []model.Booking {
	return b.reservationRepo.GetAllBookings()
}

func (b *BookingReservation) GetPassengersByOrigin(stationName string) ([]model.Passenger, error) {
	if stationName == "" {
		return nil, errors.New("station name is required")
	}
	return b.reservationRepo.FindPassengerByOrigin(stationName)
}

func (b *BookingReservation) GetPassengersByDestination(stationName string) ([]model.Passenger, error) {
	if stationName == "" {
		return nil, errors.New("station name is required")
	}
	return b.reservationRepo.FindPassengerByDestination(stationName)
}

func (b *BookingReservation) GetPassengerBySeat(serviceID, seatID string) (*model.Passenger, error) {
	if serviceID == "" || seatID == "" {
		return nil, errors.New("service ID and seat ID are required")
	}

	return b.reservationRepo.FindPassengerBySeat(serviceID, seatID)
}

func (b *BookingReservation) GetAllServices() []model.Service {
	return b.reservationRepo.GetAllServices()
}

func (b *BookingReservation) GetAllStations() []model.Station {
	return b.reservationRepo.GetAllStations()
}

func (b *BookingReservation) GetPassengerByServiceSeatDate(serviceID string, seatID string, date string) (model.Passenger, error) {
	if serviceID == "" || seatID == "" || date == "" {
		return model.Passenger{}, errors.New("service ID, seat ID, and date are required")
	}
	return b.reservationRepo.FindPassengerByServiceSeatDate(serviceID, seatID, date)
}

func (b *BookingReservation) GetPassengersByOriginDestination(origin string, destination string) ([]model.Passenger, error) {
	if origin == "" || destination == "" {
		return nil, errors.New("origin and destination are required")
	}

	return b.reservationRepo.FindPassengerByOriginDestination(origin, destination)
}
