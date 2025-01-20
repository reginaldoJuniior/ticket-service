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
		ID:        "B1",
		Passenger: model.Passenger{Name: "John Doe"},
		ServiceID: "5160",
		Seat:      "A11",
	}
	response := client.Post("/bookings", booking)
	fmt.Println(response.GetStatusCode(), response.GetBody())

	// Try booking the same seat again
	response = client.Post("/bookings", booking)
	fmt.Println(response.GetStatusCode(), response.GetBody())

	response = client.Get("/bookings")
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
	GetPassengersByStation(stationName string) []model.Passenger
	GetPassengerBySeat(serviceID, seatID string) *model.Passenger
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
	// extract station name from URL
	var parameter string
	if len(strings.Split(url, "/")) > 2 {
		parameter = strings.Split(url, "/")[1]
		url = strings.Split(url, "/")[2]
	}

	switch url {
	case "/bookings":
		// Retrieve all bookings
		bookings := client.handle.GetAllBookings()
		return &HTTPResponse{StatusCode: 200, Body: bookings}
	case "/station/passengers":
		// Retrieve all passengers
		passengers := client.handle.GetPassengersByStation(parameter)
		return &HTTPResponse{StatusCode: 200, Body: passengers}
	default:
		return &HTTPResponse{StatusCode: 404, Body: "Route not found"}
	}
}
