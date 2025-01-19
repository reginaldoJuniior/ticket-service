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
			},
		}
	})
	return &instance
}

type Reservations struct {
	Bookings      map[string]model.Booking
	ReservedSeats map[string]struct{}
	mutex         sync.Mutex
	data          map[string]any
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
			return fmt.Errorf("seat %s is already booked", booking.Seat)
		}
	}
	r.data["bookings"] = append(bookings, booking)
	return nil
}

func (r *Reservations) GetBookDetails(passenger model.Passenger) (*model.Booking, error) {
	for _, b := range r.Bookings {
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
	if _, ok := r.Bookings[bookKey]; !ok {
		return nil, errors.New(BookingNotFoundError)
	}
	book, _ := r.Bookings[bookKey]
	return &book, nil
}
