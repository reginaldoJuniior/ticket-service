package repository

import (
	"errors"
	"sync"
	"ticket-inventory/model"
)

type ReservationRepository interface {
	SaveBook(book model.Booking)
	GetAllBookings() []model.Booking
	FindSeat(code string, service model.Service) (model.Seat, error)
	CheckSeatAvailability(code string)
}

var instance Reservations
var once sync.Once

// NewReservationRepository function to return a new instance of Reservations struct.
// This function will be used to create a singleton instance of Reservations struct.
func NewReservationRepository() ReservationRepository {
	once.Do(func() {
		instance = Reservations{}
	})
	return &instance
}

type Reservations struct {
	Bookings      []model.Booking
	ReservedSeats map[string]struct{}
	mutex         sync.Mutex
}

func (r *Reservations) FindSeat(code string, service model.Service) (model.Seat, error) {
	//TODO implement me
	panic("implement me")
}

func (r *Reservations) SaveBook(book model.Booking) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	r.Bookings = append(r.Bookings, book)
	r.ReservedSeats[book.Key()] = struct{}{}
}

func (r *Reservations) GetBookDetails(passenger model.Passenger) (*model.Booking, error) {
	for _, b := range r.Bookings {
		if b.Passenger.Name == passenger.Name {
			return &b, nil
		}
	}
	return nil, errors.New("booking not found")
}

func (r *Reservations) GetAllBookings() []model.Booking {
	return r.Bookings
}

func (r *Reservations) CheckSeatAvailability(code string) {
	//TODO implement me
	panic("implement me")
}
