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
			ID:          "booking1",
			Passenger:   model.Passenger{Name: "John Doe"},
			ServiceID:   "service1",
			Seat:        model.Seat{ID: "seat1", ComfortZone: "second-class"},
			Origin:      "station1",
			Destination: "station2",
			Date:        "2025-04-01",
		}
	})

	It("FindPassengerByOrigin returns passengers by station name", func() {
		_ = reservations.SaveBook(booking)
		passengers, err := reservations.FindPassengerByOrigin("station1")
		Expect(err).To(BeNil())
		Expect(passengers).To(HaveLen(1))
		Expect(passengers[0].Name).To(Equal("John Doe"))
	})

	It("FindPassengerByOrigin returns an error if no passengers found", func() {
		_, err := reservations.FindPassengerByOrigin("station2")
		Expect(err).To(Equal(errors.New("passenger not found")))
	})

	It("Remove book removes a booking successfully", func() {
		Expect(reservations.GetAllBookings()).To(ContainElement(booking))
		err := reservations.RemoveBook(booking)
		Expect(err).To(BeNil())
		Expect(reservations.GetAllBookings()).NotTo(ContainElement(booking))

	})

	It("FindPassengerByDestination returns passengers by destination station name", func() {
		_ = reservations.SaveBook(booking)
		passengers, err := reservations.FindPassengerByDestination("station2")
		Expect(err).To(BeNil())
		Expect(passengers).To(HaveLen(1))
		Expect(passengers[0].Name).To(Equal("John Doe"))
	})

	It("FindPassengerByDestination returns an error if no passengers found", func() {
		_, err := reservations.FindPassengerByDestination("station1")
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

	It("FindPassengerByServiceSeatDate returns passenger by service ID, seat ID, and date", func() {
		_ = reservations.SaveBook(booking)
		passenger, err := reservations.FindPassengerByServiceSeatDate("service1", "seat1", "2025-04-01")
		Expect(err).To(BeNil())
		Expect(passenger.Name).To(Equal("John Doe"))
	})

	It("FindPassengerByServiceSeatDate returns an error if passenger is not found", func() {
		_, err := reservations.FindPassengerByServiceSeatDate("service1", "seat2", "2025-04-01")
		Expect(err).To(Equal(errors.New("passenger not found")))
	})

	It("FindPassengerByServiceSeatDate returns an error if date does not match", func() {
		_ = reservations.SaveBook(booking)
		_, err := reservations.FindPassengerByServiceSeatDate("service1", "seat1", "2025-04-02")
		Expect(err).To(Equal(errors.New("passenger not found")))
	})

	It("FindPassengerByOriginDestination returns passengers by origin and destination", func() {
		_ = reservations.SaveBook(booking)
		passengers, err := reservations.FindPassengerByOriginDestination("station1", "station2")
		Expect(err).To(BeNil())
		Expect(passengers).To(HaveLen(1))
		Expect(passengers[0].Name).To(Equal("John Doe"))
	})

	It("FindPassengerByOriginDestination returns an error if no passengers found", func() {
		_, err := reservations.FindPassengerByOriginDestination("station1", "station3")
		Expect(err).To(Equal(errors.New("passenger not found")))
	})
})
