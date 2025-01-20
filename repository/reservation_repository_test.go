package repository_test

import (
	"errors"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"ticket-inventory/model"
	"ticket-inventory/repository"
)

var _ = Describe("Reservations", func() {
	var (
		reservations *repository.Reservations
		booking      model.Booking
	)

	BeforeEach(func() {
		reservations = repository.NewReservationRepository()
		booking = model.Booking{
			ID:        "booking1",
			Passenger: model.Passenger{Name: "John Doe"},
			ServiceID: "service1",
			Seat:      "seat1",
			Origin:    "station1",
		}
	})

	It("FindPassengerByStation returns passengers by station name", func() {
		_ = reservations.SaveBook(booking)
		passengers, err := reservations.FindPassengerByStation("station1")
		Expect(err).To(BeNil())
		Expect(passengers).To(HaveLen(1))
		Expect(passengers[0].Name).To(Equal("John Doe"))
	})

	It("FindPassengerByStation returns an error if no passengers found", func() {
		_, err := reservations.FindPassengerByStation("station2")
		Expect(err).To(Equal(errors.New("passenger not found")))
	})

	It("FindPassengerBySeat returns passenger by service ID and seat ID", func() {
		_ = reservations.SaveBook(booking)
		passenger, err := reservations.FindPassengerBySeat("service1", "seat1")
		Expect(err).To(BeNil())
		Expect(passenger.Name).To(Equal("John Doe"))
	})

	It("FindPassengerBySeat returns an error if passenger is not found", func() {
		_, err := reservations.FindPassengerBySeat("service1", "seat2")
		Expect(err).To(Equal(errors.New("passenger not found")))
	})
})
