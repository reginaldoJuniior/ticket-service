package client

import (
	"strings"
	"ticket-inventory/model"
	"ticket-inventory/repository"
	"ticket-inventory/usecases"
)

type SimulatedHTTPClient struct {
	handle BookHandle
}

func NewSimulatedHTTPClient() *SimulatedHTTPClient {
	return &SimulatedHTTPClient{
		handle: usecases.NewBookingReservation(
			repository.NewReservationRepository(),
		),
	}
}

func (client *SimulatedHTTPClient) Post(url string, body any) *HTTPResponse {
	switch url {
	case "/bookings":
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
		bookings := client.handle.GetAllBookings()
		return &HTTPResponse{StatusCode: 200, Body: bookings}
	case len(parts) == 2 && parts[1] == "services":
		services := client.handle.GetAllServices()
		return &HTTPResponse{StatusCode: 200, Body: services}
	case len(parts) == 2 && parts[1] == "stations":
		stations := client.handle.GetAllStations()
		return &HTTPResponse{StatusCode: 200, Body: stations}
	case len(parts) == 3 && parts[2] == "boarding":
		stationID := parts[1]
		passengers, err := client.handle.GetPassengersByOrigin(stationID)
		if err != nil {
			return &HTTPResponse{StatusCode: 400, Body: err.Error()}
		}
		return &HTTPResponse{StatusCode: 200, Body: passengers}
	case len(parts) == 3 && parts[2] == "leaving":
		stationID := parts[1]
		passengers, err := client.handle.GetPassengersByDestination(stationID)
		if err != nil {
			return &HTTPResponse{StatusCode: 400, Body: err.Error()}
		}
		return &HTTPResponse{StatusCode: 200, Body: passengers}
	case len(parts) == 4 && parts[3] == "passengers":
		serviceID := parts[1]
		seatID := parts[2]
		passenger, err := client.handle.GetPassengerBySeat(serviceID, seatID)
		if err != nil {
			return &HTTPResponse{StatusCode: 400, Body: err.Error()}
		}
		if passenger != nil {
			return &HTTPResponse{StatusCode: 200, Body: passenger}
		}
		return &HTTPResponse{StatusCode: 404, Body: "Passenger not found"}
	case len(parts) == 6 && parts[5] == "passenger":
		serviceID := parts[1]
		seatID := parts[2]
		date := parts[3]
		passenger, err := client.handle.GetPassengerByServiceSeatDate(serviceID, seatID, date)
		if err != nil {
			return &HTTPResponse{StatusCode: 400, Body: err.Error()}
		}
		return &HTTPResponse{StatusCode: 200, Body: passenger}
	case len(parts) == 3 && parts[2] == "passengers":
		origin := strings.Split(parts[1], "-")[0]
		destination := strings.Split(parts[1], "-")[1]
		passengers, err := client.handle.GetPassengersByOriginDestination(origin, destination)
		if err != nil {
			return &HTTPResponse{StatusCode: 400, Body: err.Error()}
		}
		return &HTTPResponse{StatusCode: 200, Body: map[string]any{
			"numberOfPassengers": len(passengers),
			"passengers":         passengers,
		}}
	default:
		return &HTTPResponse{StatusCode: 404, Body: "Route not found"}
	}
}
