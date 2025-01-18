package model

import (
	"fmt"
)

type NodeRoute struct {
	Station           Station
	DistanceInMinutes int
	Next              *NodeRoute
}

func NewNodeRoute(station Station) *NodeRoute {
	return &NodeRoute{
		Station:           station,
		DistanceInMinutes: 0,
	}
}

func NewNodeRouteWithDistance(station Station, distanceInMinutes int) *NodeRoute {
	return &NodeRoute{
		Station:           station,
		DistanceInMinutes: distanceInMinutes,
	}
}

// Append adds a new node to the end of the list
func (n *NodeRoute) Append(node *NodeRoute, distanceInMinutes int) {
	for n.Next != nil {
		n = n.Next
	}
	n.Next = node
	n.DistanceInMinutes = distanceInMinutes
}

// AppendStation adds a new station to the end of the list
// distanceInMinutes is the time it takes to travel from the previous station to the new station
func (n *NodeRoute) AppendStation(station Station, distanceInMinutes int) {
	n.Append(NewNodeRoute(station), distanceInMinutes)
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
