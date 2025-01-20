package main

import (
	"fmt"
	http "ticket-inventory/client"
	"ticket-inventory/model"
)

func main() {
	client := http.NewSimulatedHTTPClient() // Initialize empty bookings list

	// My tests
	myTestRequests(client)

	// Assessment requirements
	assessmentRequests(client)
}

func myTestRequests(client *http.SimulatedHTTPClient) {
	response := client.Get("/services")
	fmt.Println(response.GetStatusCode(), response.GetBody())

	response = client.Get("/stations")
	fmt.Println(response.GetStatusCode(), response.GetBody())

	// Create a booking
	booking := model.Booking{
		ID:        "B1",
		Passenger: model.Passenger{Name: "John Doe"},
		ServiceID: "5160",
		Seat: model.Seat{
			ID:          "A11",
			ComfortZone: "first-class",
		},
		Origin:      "Paris",
		Destination: "London",
		Date:        "2025-10-01",
	}
	response = client.Post("/bookings", booking)
	fmt.Println(response.GetStatusCode(), response.GetBody())

	// Try booking the same seat again
	response = client.Post("/bookings", booking)
	fmt.Println(response.GetStatusCode(), response.GetBody())

	response = client.Get("/bookings")
	fmt.Println(response.GetStatusCode(), response.GetBody())

	response = client.Get("/5160/A11/passengers")
	fmt.Println(response.GetStatusCode(), response.GetBody())

	response = client.Get("/5160/passengers")
	fmt.Println(response.GetStatusCode(), response.GetBody())

	// Same seat as before but different service ID
	// Should be allowed
	booking = model.Booking{
		ID: "B1",
		Passenger: model.Passenger{
			Name: "Michael",
		},
		ServiceID:   "5161",
		Seat:        model.Seat{ID: "A12", ComfortZone: "second-class"},
		Origin:      "London",
		Destination: "Amsterdam",
		Date:        "2025-10-01",
	}
	response = client.Post("/bookings", booking)
	fmt.Println(response.GetStatusCode(), response.GetBody())

	booking = model.Booking{
		ID: "B2",
		Passenger: model.Passenger{
			Name: "Mary",
		},
		ServiceID:   "5161",
		Seat:        model.Seat{ID: "A12", ComfortZone: "second-class"},
		Origin:      "Paris",
		Destination: "Berlin",
		Date:        "2025-10-01",
	}
	response = client.Post("/bookings", booking)
	fmt.Println(response.GetStatusCode(), response.GetBody())

	// Query passengers boarding at London
	response = client.Get("/London/passengers")
	fmt.Println(response.GetStatusCode(), response.GetBody())
}

func assessmentRequests(client *http.SimulatedHTTPClient) {
	response := client.Post("/bookings", model.Booking{
		ID:          "B3",
		Passenger:   model.Passenger{Name: "Alice"},
		ServiceID:   "5160",
		Seat:        model.Seat{ID: "A11", ComfortZone: "first-class"},
		Origin:      "Paris",
		Destination: "Amsterdam",
		Date:        "2021-04-01",
	})
	fmt.Println(response.GetStatusCode(), response.GetBody())

	response = client.Post("/bookings", model.Booking{
		ID:          "B4",
		Passenger:   model.Passenger{Name: "Bob"},
		ServiceID:   "5160",
		Seat:        model.Seat{ID: "A12", ComfortZone: "first-class"},
		Origin:      "Paris",
		Destination: "Amsterdam",
		Date:        "2021-04-01",
	})
	fmt.Println(response.GetStatusCode(), response.GetBody())

	// Attempt to book the same seats again
	response = client.Post("/bookings", model.Booking{
		ID:          "B3",
		Passenger:   model.Passenger{Name: "Alice"},
		ServiceID:   "5160",
		Seat:        model.Seat{ID: "A11", ComfortZone: "first-class"},
		Origin:      "Paris",
		Destination: "Amsterdam",
		Date:        "2021-04-01",
	})
	fmt.Println(response.GetStatusCode(), response.GetBody())

	response = client.Post("/bookings", model.Booking{
		ID:          "B4",
		Passenger:   model.Passenger{Name: "Bob"},
		ServiceID:   "5160",
		Seat:        model.Seat{ID: "A12", ComfortZone: "first-class"},
		Origin:      "Paris",
		Destination: "Amsterdam",
		Date:        "2021-04-01",
	})
	fmt.Println(response.GetStatusCode(), response.GetBody())

	response = client.Post("/bookings", model.Booking{
		ID:          "B5",
		Passenger:   model.Passenger{Name: "Charlie"},
		ServiceID:   "5161",
		Seat:        model.Seat{ID: "H1", ComfortZone: "second-class"},
		Origin:      "London",
		Destination: "Paris",
		Date:        "2021-04-01",
	})
	fmt.Println(response.GetStatusCode(), response.GetBody())

	response = client.Post("/bookings", model.Booking{
		ID:          "B6",
		Passenger:   model.Passenger{Name: "Dave"},
		ServiceID:   "5161",
		Seat:        model.Seat{ID: "N5", ComfortZone: "second-class"},
		Origin:      "London",
		Destination: "Paris",
		Date:        "2021-04-01",
	})
	fmt.Println(response.GetStatusCode(), response.GetBody())

	response = client.Post("/bookings", model.Booking{
		ID:          "B7",
		Passenger:   model.Passenger{Name: "Charlie"},
		ServiceID:   "5161",
		Seat:        model.Seat{ID: "A1", ComfortZone: "first-class"},
		Origin:      "Paris",
		Destination: "Amsterdam",
		Date:        "2021-04-01",
	})
	fmt.Println(response.GetStatusCode(), response.GetBody())

	response = client.Post("/bookings", model.Booking{
		ID:          "B8",
		Passenger:   model.Passenger{Name: "Dave"},
		ServiceID:   "5161",
		Seat:        model.Seat{ID: "T7", ComfortZone: "first-class"},
		Origin:      "Paris",
		Destination: "Amsterdam",
		Date:        "2021-04-01",
	})
	fmt.Println(response.GetStatusCode(), response.GetBody())

	// Attempt to book the same seats again
	response = client.Post("/bookings", model.Booking{
		ID:          "B5",
		Passenger:   model.Passenger{Name: "Charlie"},
		ServiceID:   "5161",
		Seat:        model.Seat{ID: "H1", ComfortZone: "second-class"},
		Origin:      "London",
		Destination: "Paris",
		Date:        "2021-04-01",
	})
	fmt.Println(response.GetStatusCode(), response.GetBody())

	response = client.Post("/bookings", model.Booking{
		ID:          "B6",
		Passenger:   model.Passenger{Name: "Dave"},
		ServiceID:   "5161",
		Seat:        model.Seat{ID: "N5", ComfortZone: "second-class"},
		Origin:      "London",
		Destination: "Paris",
		Date:        "2021-04-01",
	})
	fmt.Println(response.GetStatusCode(), response.GetBody())

	response = client.Post("/bookings", model.Booking{
		ID:          "B7",
		Passenger:   model.Passenger{Name: "Charlie"},
		ServiceID:   "5161",
		Seat:        model.Seat{ID: "A1", ComfortZone: "first-class"},
		Origin:      "Paris",
		Destination: "Amsterdam",
		Date:        "2021-04-01",
	})
	fmt.Println(response.GetStatusCode(), response.GetBody())

	response = client.Post("/bookings", model.Booking{
		ID:          "B8",
		Passenger:   model.Passenger{Name: "Dave"},
		ServiceID:   "5161",
		Seat:        model.Seat{ID: "T7", ComfortZone: "first-class"},
		Origin:      "Paris",
		Destination: "Amsterdam",
		Date:        "2021-04-01",
	})
	fmt.Println(response.GetStatusCode(), response.GetBody())
}
