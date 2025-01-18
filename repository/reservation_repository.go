package repository

import (
	"sync"
	"ticket-inventory/model"
)

type ReservationRepository interface {
	SaveBook(book model.Booking)
	GetAllBookings() []model.Booking
	FindSeat(code string, service model.Service) (model.Seat, error)
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
	Bookings []model.Booking
	mutex    sync.Mutex
}

func (r *Reservations) FindSeat(code string, service model.Service) (model.Seat, error) {
	//TODO implement me
	panic("implement me")
}

func (r *Reservations) SaveBook(book model.Booking) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	r.Bookings = append(r.Bookings, book)
}

func (r *Reservations) GetAllBookings() []model.Booking {
	return r.Bookings
}
