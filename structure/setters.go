package structure

import (
	"math"

	"github.com/jcgraybill/ship-shape/resource"
	"github.com/jcgraybill/ship-shape/ui"
)

func (s *Structure) Pause() {
	s.paused = true
	s.workers = 0
}
func (s *Structure) Unpause() {
	s.paused = false
}

func (s *Structure) Prioritize() {
	s.prioritized = true
}
func (s *Structure) Deprioritize() {
	s.prioritized = false
}

func (s *Structure) adjustPopulationCapacity() {
	if s.data.Class == Residential {
		cap := float64(s.storage[resource.Population].Capacity) * (float64(s.planet.Resources()[resource.Habitability]) / 255)
		s.storage[resource.Population] = Storage{
			Resource: resource.Population,
			Amount:   s.storage[resource.Population].Amount,
			Capacity: uint8(math.Ceil(cap)),
		}
	}
}

func (s *Structure) AssignWorkers(workers int) {
	s.workers = workers
}

func (s *Structure) CollectIncome() int {
	income := s.income
	s.income = 0
	return income
}

func (s *Structure) Upgrade(st int, sd StructureData) {
	s.structureType = st
	s.data = sd
	s.image, _, _, s.w, s.h = s.generateImage(s.x, s.x, ui.NonFocusColor)
	s.highlightedImage, _, _, _, _ = s.generateImage(s.x, s.x, ui.FocusedColor)
	s.berths, s.ships = sd.Berths, sd.Berths

	carryover := make(map[int]uint8)

	for _, st := range s.storage {
		carryover[st.Resource] = st.Amount
	}

	s.storage = make(map[int]Storage)
	s.resourcesWanted = make([]int, 0)

	for _, st := range s.data.Storage {
		amount := carryover[st.Resource]
		if amount > st.Capacity {
			amount = st.Capacity
		}

		s.storage[st.Resource] = Storage{
			Resource: st.Resource,
			Capacity: st.Capacity,
			Amount:   amount,
		}

		if st.Resource != s.data.Produces.Resource {
			s.resourcesWanted = append(s.resourcesWanted, st.Resource)
		}
	}
	s.adjustPopulationCapacity()
}