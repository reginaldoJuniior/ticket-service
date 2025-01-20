package main

import (
	"fmt"
	"strings"
	"ticket-inventory/model"
	"ticket-inventory/repository"
	"ticket-inventory/usecases"
)

func main() {
	client := NewSimulatedHTTPClient() // Initialize empty bookings list

	// Create a booking
	booking := model.Booking{
		ID:          "B1",
		Passenger:   model.Passenger{Name: "John Doe"},
		ServiceID:   "5160",
		Seat:        "A11",
		Origin:      "Paris",
		Destination: "London",
	}
	response := client.Post("/bookings", booking)
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

	booking = model.Booking{ID: "B1", Passenger: model.Passenger{Name: "John"}, ServiceID: "5160", Seat: "A11", Origin: "London", Destination: "Amsterdam"}
	response = client.Post("/bookings", booking)
	fmt.Println(response.GetStatusCode(), response.GetBody())

	booking = model.Booking{ID: "B2", Passenger: model.Passenger{Name: "Mary"}, ServiceID: "5160", Seat: "A12", Origin: "Paris", Destination: "Berlin"}
	response = client.Post("/bookings", booking)
	fmt.Println(response.GetStatusCode(), response.GetBody())

	// Query passengers boarding at London
	response = client.Get("/London/passengers")
	fmt.Println(response.GetStatusCode(), response.GetBody())
}

// Response interface simulates an HTTP response
type Response interface {
	GetStatusCode() int
	GetBody() any
}

// HTTPClient interface simulates REST-like client methods
type HTTPClient interface {
	Post(url string, body any) Response
	Get(url string) Response
}

// HTTPResponse implements the Response interface
type HTTPResponse struct {
	StatusCode int
	Body       any
}

func (r *HTTPResponse) GetStatusCode() int {
	return r.StatusCode
}

func (r *HTTPResponse) GetBody() any {
	return r.Body
}

type BookHandle interface {
	GetAllBookings() []model.Booking
	CreateBooking(booking model.Booking) error
	GetPassengersByStation(stationName string) ([]model.Passenger, error)
	GetPassengerBySeat(serviceID, seatID string) (*model.Passenger, error)
}

// SimulatedHTTPClient implements the HTTPClient interface
type SimulatedHTTPClient struct {
	handle BookHandle
}

func NewSimulatedHTTPClient() *SimulatedHTTPClient {
	return &SimulatedHTTPClient{
		handle: usecases.NewBookingReservation(repository.NewReservationRepository()),
	}
}

func (client *SimulatedHTTPClient) Post(url string, body any) *HTTPResponse {
	switch url {
	case "/bookings":
		// Process booking creation
		booking := body.(model.Booking)
		if err := client.handle.CreateBooking(booking); err != nil {
			return &HTTPResponse{StatusCode: 400, Body: err.Error()}
		}
		return &HTTPResponse{StatusCode: 201, Body: "Booking created successfully"}
	default:
		return &HTTPResponse{StatusCode: 404, Body: "Route not found"}
	}
}

func (client *SimulatedHTTPClient) Get(url string) *HTTPResponse {
	parts := strings.Split(url, "/")

	switch {
	case len(parts) == 2 && parts[1] == "bookings":
		// URL: /bookings
		// Retrieve all bookings
		bookings := client.handle.GetAllBookings()
		return &HTTPResponse{StatusCode: 200, Body: bookings}
	case len(parts) == 3 && parts[2] == "passengers":
		// URL: /{stationID}/passengers
		// Retrieve all passengers by station ID
		stationID := parts[1]
		passengers, err := client.handle.GetPassengersByStation(stationID)
		if err != nil {
			return &HTTPResponse{
				StatusCode: 400,
				Body:       err.Error(),
			}
		}
		return &HTTPResponse{StatusCode: 200, Body: passengers}
	case len(parts) == 4 && parts[3] == "passengers":
		// URL: /{serviceID}/{seatID}/passengers
		// Retrieve passenger by service ID and seat ID
		serviceID := parts[1]
		seatID := parts[2]
		passenger, err := client.handle.GetPassengerBySeat(serviceID, seatID)
		if err != nil {
			return &HTTPResponse{
				StatusCode: 400,
				Body:       err.Error(),
			}
		}
		if passenger != nil {
			return &HTTPResponse{StatusCode: 200, Body: passenger}
		}
		return &HTTPResponse{StatusCode: 404, Body: "Passenger not found"}
	default:
		// URL not found
		return &HTTPResponse{StatusCode: 404, Body: "Route not found"}
	}
}
