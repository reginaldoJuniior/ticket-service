package model_test

import (
	"ticket-inventory/model"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Booking", func() {
	Context("when generating a service key", func() {
		It("returns the correct key for a valid booking", func() {
			passenger := model.Passenger{Name: "John Doe"}
			seat := model.Seat{Code: "A1"}
			ticket := model.Ticket{Code: "T123", Seat: seat, Buyer: passenger, Class: "First"}

			layout := "15h04"
			hour, err := time.Parse(layout, "10h30")
			Expect(err).To(BeNil())
			service := model.Service{
				Code: "S456",
				Hour: hour,
				Train: model.Train{
					Code: "ATR",
				},
				Route: model.Route{
					To: model.Station{Name: "Station A"}, From: model.Station{Name: "Station B"},
				}}

			booking := model.Booking{Passenger: passenger, Ticket: ticket, Service: service}

			result := booking.Key()

			expected := "John Doe-T123-A1-S456-ATR-Station A-Station B-10:30"
			Expect(result).To(Equal(expected))
		})

		It("returns a key with empty fields if booking details are empty", func() {
			booking := model.Booking{}

			result := booking.Key()

			expected := "-------0:0"
			Expect(result).To(Equal(expected))
		})
	})

	It("returns the correct key for a valid booking with hour and minute", func() {
		passenger := model.Passenger{Name: "John Doe"}
		seat := model.Seat{Code: "A1"}
		ticket := model.Ticket{Code: "T123", Seat: seat, Buyer: passenger, Class: "First"}

		layout := "15h04"
		hour, err := time.Parse(layout, "10h30")
		Expect(err).To(BeNil())
		service := model.Service{
			Code: "S456",
			Hour: hour,
			Train: model.Train{
				Code: "ATR",
			},
			Route: model.Route{
				To: model.Station{Name: "Station A"}, From: model.Station{Name: "Station B"},
			}}

		booking := model.Booking{Passenger: passenger, Ticket: ticket, Service: service}

		result := booking.Key()

		expected := "John Doe-T123-A1-S456-ATR-Station A-Station B-10:30"
		Expect(result).To(Equal(expected))
	})

	It("returns a key with empty fields if booking details are empty", func() {
		booking := model.Booking{}

		result := booking.Key()

		expected := "-------0:0"
		Expect(result).To(Equal(expected))
	})

	It("returns a key with missing hour and minute if service hour is not set", func() {
		passenger := model.Passenger{Name: "John Doe"}
		seat := model.Seat{Code: "A1"}
		ticket := model.Ticket{Code: "T123", Seat: seat, Buyer: passenger, Class: "First"}
		service := model.Service{
			Code: "S456",
			Train: model.Train{
				Code: "ATR",
			},
			Route: model.Route{
				To: model.Station{Name: "Station A"}, From: model.Station{Name: "Station B"},
			}}

		booking := model.Booking{Passenger: passenger, Ticket: ticket, Service: service}

		result := booking.Key()

		expected := "John Doe-T123-A1-S456-ATR-Station A-Station B-0:0"
		Expect(result).To(Equal(expected))
	})
})
