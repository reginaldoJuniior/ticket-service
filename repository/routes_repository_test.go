package repository_test

import (
	"ticket-inventory/model"
	"ticket-inventory/repository"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("ListRoutes", func() {
	Context("when listing a single node route", func() {
		It("returns the route correctly", func() {
			station := model.Station{Name: "Station A"}
			node := &repository.NodeRoute{Station: station, DistanceInMinutes: 10}

			result := repository.ListRoutes(node)

			expected := "Station A"
			Expect(result).To(Equal(expected))
		})
	})

	Context("when listing multiple node routes", func() {
		It("returns the route correctly", func() {
			stationA := model.Station{Name: "Station A"}
			stationB := model.Station{Name: "Station B"}
			stationC := model.Station{Name: "Station C"}
			nodeC := &repository.NodeRoute{Station: stationC, DistanceInMinutes: 30}
			nodeB := &repository.NodeRoute{Station: stationB, DistanceInMinutes: 20, Next: nodeC}
			nodeA := &repository.NodeRoute{Station: stationA, DistanceInMinutes: 10, Next: nodeB}

			result := repository.ListRoutes(nodeA)

			expected := "Station A (10 min) -> Station B (20 min) -> Station C"
			Expect(result).To(Equal(expected))
		})
	})

	Context("when listing an empty route list", func() {
		It("returns an empty string", func() {
			result := repository.ListRoutes(nil)

			expected := ""
			Expect(result).To(Equal(expected))
		})
	})
})
