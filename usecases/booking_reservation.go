package usecases

import (
	"ticket-inventory/model"
)

// ReservationRepository is an interface for booking reservation repository
type ReservationRepository interface {
	SaveBook(book model.Booking) error
	GetAllBookings() []model.Booking
	FindSeat(code string, service model.Service) (model.Seat, error)
	FindBook(bookKey string) (*model.Booking, error)
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
	return b.reservationRepo.SaveBook(booking)
}

func (b *BookingReservation) GetAllBookings() []model.Booking {
	return b.reservationRepo.GetAllBookings()
}
