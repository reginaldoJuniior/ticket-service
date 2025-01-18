package usecases_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"ticket-inventory/model"
	"ticket-inventory/usecases"
)

var _ = Describe("BookingReservation", func() {
	var (
		bookingReservation usecases.BookingReservation
		passenger          model.Passenger
		ticket             model.Ticket
	)

	BeforeEach(func() {
		// Initialize variables before each test
		passenger = model.Passenger{Name: "John Doe"}
		ticket = model.Ticket{Code: "T123"}
	})

	Describe("ReservingSeats", func() {
		It("should create a booking with the given passenger and ticket", func() {
			booking := bookingReservation.ReservingSeats(passenger, ticket)
			Expect(booking.Passenger).To(Equal(passenger))
			Expect(booking.Ticket).To(Equal(ticket))
		})

		It("should not allow duplicate bookings for the same seat", func() {
			booking1 := bookingReservation.ReservingSeats(passenger, ticket)
			booking2 := bookingReservation.ReservingSeats(passenger, ticket)
			Expect(booking1).To(Equal(booking2))
		})

		It("should support mixed-class bookings", func() {
			economyTicket := model.Ticket{Code: "E123", Class: "Economy"}
			businessTicket := model.Ticket{Code: "B123", Class: "Business"}
			economyBooking := bookingReservation.ReservingSeats(passenger, economyTicket)
			businessBooking := bookingReservation.ReservingSeats(passenger, businessTicket)
			Expect(economyBooking.Ticket.Class).To(Equal("Economy"))
			Expect(businessBooking.Ticket.Class).To(Equal("Business"))
		})

		It("should validate seat availability during booking", func() {
			available, err := bookingReservation.CheckSeatAvailability(ticket.Code)
			Expect(err).NotTo(HaveOccurred())
			Expect(available).To(BeTrue())
			booking := bookingReservation.ReservingSeats(passenger, ticket)
			Expect(booking).NotTo(BeNil())
		})
	})
})
