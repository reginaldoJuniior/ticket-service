package usecases_test

import (
	"errors"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"ticket-inventory/model"
	"ticket-inventory/usecases"
)

type MockReservationRepository struct {
	bookings []model.Booking
}

func (m *MockReservationRepository) SaveBook(book model.Booking) error {
	for _, b := range m.bookings {
		if b.Seat == book.Seat && b.ServiceID == book.ServiceID {
			return errors.New("seat is already booked")
		}
	}
	m.bookings = append(m.bookings, book)
	return nil
}

func (m *MockReservationRepository) GetAllBookings() []model.Booking {
	return m.bookings
}

func (m *MockReservationRepository) FindSeat(code string, service model.Service) (model.Seat, error) {
	return model.Seat{}, nil
}

func (m *MockReservationRepository) FindBook(bookKey string) (*model.Booking, error) {
	for _, b := range m.bookings {
		if b.ID == bookKey {
			return &b, nil
		}
	}
	return nil, errors.New("booking not found")
}

func (m *MockReservationRepository) FindPassengerByStation(stationName string) ([]model.Passenger, error) {
	var passengers []model.Passenger
	for _, booking := range m.bookings {
		if booking.Origin == stationName {
			passengers = append(passengers, booking.Passenger)
		}
	}
	return passengers, nil
}

func (m *MockReservationRepository) FindPassengerBySeat(serviceID, seatID string) (*model.Passenger, error) {
	for _, booking := range m.bookings {
		if booking.ServiceID == serviceID && booking.Seat == seatID {
			return &booking.Passenger, nil
		}
	}
	return nil, errors.New("passenger not found")
}

var _ = Describe("BookingReservation", func() {
	var (
		bookingReservation *usecases.BookingReservation
		mockRepo           *MockReservationRepository
		booking            model.Booking
	)

	BeforeEach(func() {
		mockRepo = &MockReservationRepository{}
		bookingReservation = usecases.NewBookingReservation(mockRepo)
		booking = model.Booking{
			ID:          "B1",
			Passenger:   model.Passenger{Name: "John Doe"},
			ServiceID:   "5160",
			Seat:        "A11",
			Origin:      "Paris",
			Destination: "London",
		}
	})

	It("CreateBooking successfully creates a booking", func() {
		err := bookingReservation.CreateBooking(booking)
		Expect(err).To(BeNil())
		Expect(mockRepo.GetAllBookings()).To(ContainElement(booking))
	})

	It("CreateBooking returns an error if seat is already booked", func() {
		_ = bookingReservation.CreateBooking(booking)
		err := bookingReservation.CreateBooking(booking)
		Expect(err).To(Equal(errors.New("seat is already booked")))
	})

	It("GetAllBookings returns all bookings", func() {
		_ = bookingReservation.CreateBooking(booking)
		bookings := bookingReservation.GetAllBookings()
		Expect(bookings).To(HaveLen(1))
		Expect(bookings).To(ContainElement(booking))
	})

	It("GetPassengersByStation returns passengers by station ID", func() {
		err := bookingReservation.CreateBooking(booking)
		Expect(err).To(BeNil())
		passengers, err := bookingReservation.GetPassengersByStation("Paris")
		Expect(err).To(BeNil())
		Expect(passengers).To(HaveLen(1))
		Expect(passengers[0].Name).To(Equal("John Doe"))
	})

	It("GetPassengerBySeat returns passenger by service ID and seat ID", func() {
		_ = bookingReservation.CreateBooking(booking)
		passenger, err := bookingReservation.GetPassengerBySeat("5160", "A11")
		Expect(err).To(BeNil())
		Expect(passenger.Name).To(Equal("John Doe"))
	})

	It("GetPassengerBySeat returns an error if passenger is not found", func() {
		_, err := bookingReservation.GetPassengerBySeat("5160", "A12")
		Expect(err).To(Equal(errors.New("passenger not found")))
	})
})
