package usecases

import (
	"errors"
	"ticket-inventory/model"
)

// ReservationRepository is an interface for booking reservation repository
type ReservationRepository interface {
	SaveBook(book model.Booking) error
	GetAllBookings() []model.Booking
	FindSeat(code string, service model.Service) (model.Seat, error)
	FindBook(bookKey string) (*model.Booking, error)
	FindPassengerByStation(stationName string) ([]model.Passenger, error)
	FindPassengerBySeat(serviceID, seatID string) (*model.Passenger, error)
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
// Helper methods for booking management
func (b *BookingReservation) CreateBooking(booking model.Booking) error {
	ok, err := booking.Validate()
	if !ok {
		return err
	}
	return b.reservationRepo.SaveBook(booking)
}

func (b *BookingReservation) GetAllBookings() []model.Booking {
	return b.reservationRepo.GetAllBookings()
}

func (b *BookingReservation) GetPassengersByStation(stationName string) ([]model.Passenger, error) {
	if stationName == "" {
		return nil, errors.New("station name is required")
	}
	return b.reservationRepo.FindPassengerByStation(stationName)
}

func (b *BookingReservation) GetPassengerBySeat(serviceID, seatID string) (*model.Passenger, error) {
	if serviceID == "" || seatID == "" {
		return nil, errors.New("service ID and seat ID are required")
	}

	return b.reservationRepo.FindPassengerBySeat(serviceID, seatID)
}
