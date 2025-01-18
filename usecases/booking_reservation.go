package usecases

import (
	"ticket-inventory/model"
	"ticket-inventory/repository"
)

// Handle the booking reservation without allowing duplication

// BookingReservation is a use case for booking reservation
type BookingReservation struct {
	reservationRepo repository.ReservationRepository
}

func NewBookingReservation(reservationRepo repository.ReservationRepository) BookingReservation {
	return BookingReservation{
		reservationRepo: reservationRepo,
	}
}

func (b BookingReservation) ReservingSeats(p model.Passenger, t model.Ticket) model.Booking {
	booking := model.Booking{
		Passenger: p,
		Ticket:    t,
	}

	return booking
}

func (b BookingReservation) CheckSeatAvailability(code string) (bool, error) {
	seat, err := b.reservationRepo.FindSeat(code, model.Service{})
	if err != nil {
		return false, err
	}
	return seat.Booked, nil
}
