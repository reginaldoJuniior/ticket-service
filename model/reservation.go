package model

type Reservation interface {
	Book() Booking
	GetReservations() []Booking
	GetBookDetails(passenger Passenger) Booking
}

type BookingValidator interface {
	CheckBookingAvailability() (bool, error)
	ValidateTicketAndRoute() (bool, error)
	CheckSeatAvailability(seatCode string, service Service) (bool, error)
}
