package repository

import (
	"fmt"
	"ticket-inventory/model"
)

type NodeRoute struct {
	Station           model.Station
	DistanceInMinutes int
	Next              *NodeRoute
}

func ListRoutes(n *NodeRoute) string {
	if n == nil {
		return ""
	}

	result := ""
	for n.Next != nil {
		result = result + fmt.Sprintf("%s (%d min) -> ", n.Station.Name, n.DistanceInMinutes)
		n = n.Next
	}
	// Last station
	result = result + fmt.Sprintf("%s", n.Station.Name)

	return result
}
