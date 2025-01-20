# Ticket Inventory System

This project is a ticket inventory system with a simulated HTTP client. It allows you to manage bookings, services, and stations.

## How to Run the Main Application

1. **Clone the repository:**
   ```sh
   git clone https://github.com/yourusername/ticket-inventory.git
   cd ticket-inventory
   ```

2. **Install dependencies:**
   ```sh
   go mod tidy
   ```

3. **Run the main application:**
   ```sh
   go run main.go
   ```

## How to Run Tests

1. **Install dependencies:**
   ```sh
   go mod tidy
   ```

2. **Run the tests:**
   ```sh
   go test ./...
   ```

## Libraries Used

- **Ginkgo**: BDD Testing Framework for Go.
- **Gomega**: Matcher library for Ginkgo.

## Fake HTTP Implementation

This project includes a simulated HTTP client to handle requests and responses without a real HTTP server. The `SimulatedHTTPClient` class in `client.go` handles various routes and simulates the behavior of an HTTP client. This allows for testing and development without the need for a live server.