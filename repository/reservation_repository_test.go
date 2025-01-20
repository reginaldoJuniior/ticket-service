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
			Seat:      model.Seat{ID: "seat1", ComfortZone: "second-class"},
			Origin:    "station1",
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

	Context("Assessment requirements", func() {
		It("should allow making a reservation for 2 passengers with 2 first-class seats from Paris to Amsterdam on service 5160 on April 1st 2021 with seats A11 & A12", func() {
			booking1 := model.Booking{
				ID:          "booking1",
				Passenger:   model.Passenger{Name: "Alice"},
				ServiceID:   "5160",
				Seat:        model.Seat{ID: "A11", ComfortZone: "first-class"},
				Origin:      "Paris",
				Destination: "Amsterdam",
				Date:        "2021-04-01",
			}
			booking2 := model.Booking{
				ID:          "booking2",
				Passenger:   model.Passenger{Name: "Bob"},
				ServiceID:   "5160",
				Seat:        model.Seat{ID: "A12", ComfortZone: "first-class"},
				Origin:      "Paris",
				Destination: "Amsterdam",
				Date:        "2021-04-01",
			}
			Expect(reservations.SaveBook(booking1)).To(BeNil())
			Expect(reservations.SaveBook(booking2)).To(BeNil())

			// Attempt to make the same booking again, should fail
			Expect(reservations.SaveBook(booking1)).ToNot(BeNil())
			Expect(reservations.SaveBook(booking2)).ToNot(BeNil())
		})

		It("should allow making a reservation for 2 passengers where 1 passenger in second-class, one in first-class from London to Amsterdam. London to Paris, seat H1 and N5, Paris to Amsterdam, seat A1 & T7", func() {
			booking1 := model.Booking{
				ID:          "booking3",
				Passenger:   model.Passenger{Name: "Charlie"},
				ServiceID:   "5161",
				Seat:        model.Seat{ID: "H1", ComfortZone: "second-class"},
				Origin:      "London",
				Destination: "Paris",
				Date:        "2021-04-01",
			}
			booking2 := model.Booking{
				ID:          "booking4",
				Passenger:   model.Passenger{Name: "Dave"},
				ServiceID:   "5161",
				Seat:        model.Seat{ID: "N5", ComfortZone: "first-class"},
				Origin:      "London",
				Destination: "Paris",
				Date:        "2021-04-01",
			}
			booking3 := model.Booking{
				ID:          "booking5",
				Passenger:   model.Passenger{Name: "Charlie"},
				ServiceID:   "5162",
				Seat:        model.Seat{ID: "A1", ComfortZone: "first-class"},
				Origin:      "Paris",
				Destination: "Amsterdam",
				Date:        "2021-04-01",
			}
			booking4 := model.Booking{
				ID:          "booking6",
				Passenger:   model.Passenger{Name: "Dave"},
				ServiceID:   "5162",
				Seat:        model.Seat{ID: "T7", ComfortZone: "second-class"},
				Origin:      "Paris",
				Destination: "Amsterdam",
				Date:        "2021-04-01",
			}
			Expect(reservations.SaveBook(booking1)).To(BeNil())
			Expect(reservations.SaveBook(booking2)).To(BeNil())
			Expect(reservations.SaveBook(booking3)).To(BeNil())
			Expect(reservations.SaveBook(booking4)).To(BeNil())

			// Attempt to make the same booking again, should fail
			Expect(reservations.SaveBook(booking1)).ToNot(BeNil())
			Expect(reservations.SaveBook(booking2)).ToNot(BeNil())
			Expect(reservations.SaveBook(booking3)).ToNot(BeNil())
			Expect(reservations.SaveBook(booking4)).ToNot(BeNil())
		})
	})
})
