package model_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"ticket-inventory/model"
)

var _ = Describe("ListRoutes", func() {
	Context("when listing a single node route", func() {
		It("returns the route correctly", func() {
			station := model.Station{Name: "Station A"}
			node := &model.NodeRoute{Station: station, DistanceInMinutes: 10}

			result := model.ListRoutes(node)

			expected := "Station A"
			Expect(result).To(Equal(expected))
		})
	})

	Context("when listing multiple node routes", func() {
		It("returns the route correctly", func() {
			stationA := model.Station{Name: "Station A"}
			stationB := model.Station{Name: "Station B"}
			stationC := model.Station{Name: "Station C"}

			nodeA := model.NewNodeRoute(stationA)
			nodeA.AppendStation(stationB, 10)
			nodeA.AppendStation(stationC, 20)

			result := model.ListRoutes(nodeA)

			expected := "Station A (10 min) -> Station B (20 min) -> Station C"
			Expect(result).To(Equal(expected))
		})
	})

	Context("when listing an empty route list", func() {
		It("returns an empty string", func() {
			result := model.ListRoutes(nil)

			expected := ""
			Expect(result).To(Equal(expected))
		})
	})
})
