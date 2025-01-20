package model_test

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
	reservations = repository.NewReservationRepository()

	BeforeEach(func() {
		booking = model.Booking{
			ID:        "booking1",
			Passenger: model.Passenger{Name: "John Doe"},
			ServiceID: "service1",
			Seat:      model.Seat{ID: "seat1", ComfortZone: "second-class"},
			Origin:    "station1",
			Date:      "2025-04-01",
		}
	})

	It("SaveBook saves a booking successfully", func() {
		err := reservations.SaveBook(booking)
		Expect(err).To(BeNil())
		Expect(reservations.GetAllBookings()).To(ContainElement(booking))
	})

	It("SaveBook returns an error if seat is already booked", func() {
		_ = reservations.SaveBook(booking)
		err := reservations.SaveBook(booking)
		Expect(err).To(Equal(errors.New("seat seat1 is already booked")))
	})

	It("RemoveBook removes a booking successfully", func() {
		_ = reservations.SaveBook(booking)
		err := reservations.RemoveBook(booking)
		Expect(err).To(BeNil())
		Expect(reservations.GetAllBookings()).NotTo(ContainElement(booking))
	})

	It("RemoveBook returns an error if booking is not found", func() {
		err := reservations.RemoveBook(booking)
		Expect(err).To(Equal(errors.New("booking not found")))
	})

	It("GetBookDetails returns booking details for a passenger", func() {
		err := reservations.SaveBook(booking)
		Expect(err).To(BeNil())
		foundBooking, err := reservations.GetBookDetails(model.Passenger{Name: "John Doe"})
		Expect(err).To(BeNil())
		Expect(foundBooking).To(Equal(&booking))
	})

	It("GetBookDetails returns an error if booking is not found", func() {
		_, err := reservations.GetBookDetails(model.Passenger{Name: "Jane Doe"})
		Expect(err).To(Equal(errors.New("booking not found")))
	})

	It("GetAllBookings returns all bookings", func() {
		_ = reservations.SaveBook(booking)
		bookings := reservations.GetAllBookings()
		Expect(bookings).To(HaveLen(1))
		Expect(bookings).To(ContainElement(booking))
	})

	It("FindBook returns a booking if found", func() {
		//err := reservations.SaveBook(booking)
		//Expect(err).To(BeNil())
		foundBooking, err := reservations.FindBook("booking1")
		Expect(err).To(BeNil())
		Expect(foundBooking).To(Equal(&booking))
	})

	It("FindBook returns an error if booking is not found", func() {
		_, err := reservations.FindBook("booking2")
		Expect(err).To(Equal(errors.New("booking not found")))
	})
})
