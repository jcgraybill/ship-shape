package structure

import (
	"github.com/jcgraybill/ship-shape/ui"
)

func (s *Structure) Produce(count int) bool {
	if s.data.Produces.Rate > 0 {
		if s.storage[s.data.Produces.Resource].Resource == s.data.Produces.Resource && s.storage[s.data.Produces.Resource].Amount < s.storage[s.data.Produces.Resource].Capacity {
			if s.Planet().Resources()[s.data.Produces.Requires] > 0 {
				productionRate := 255 - (int(s.data.Produces.Rate) * (int(s.Planet().Resources()[s.data.Produces.Requires]) / 255))
				if count%(productionRate*ui.BaseProductionRate) == 0 {
					s.storage[s.data.Produces.Resource] = Storage{
						Resource: s.data.Produces.Resource,
						Capacity: s.storage[s.data.Produces.Resource].Capacity,
						Amount:   s.storage[s.data.Produces.Resource].Amount + 1,
					}
					return true
				}
			}
		}
	}

	return false
}

func (s *Structure) Bid() (int, uint8) {
	if s.storage[s.data.Consumes].Resource == s.data.Consumes && s.storage[s.data.Consumes].Amount < s.storage[s.data.Consumes].Capacity && s.awaiting != s.data.Consumes {
		urgency := ((float32(s.storage[s.data.Consumes].Capacity) - float32(s.storage[s.data.Consumes].Amount)) / float32(s.storage[s.data.Consumes].Capacity)) * 255
		return s.data.Consumes, uint8(urgency)
	}
	return 0, 0
}

func (s *Structure) LaunchShip(resource int) {

	s.ships -= 1
	s.storage[resource] = Storage{
		Resource: resource,
		Capacity: s.storage[resource].Capacity,
		Amount:   s.storage[resource].Amount - 1,
	}
}

// TODO: lock to a structure, not a resource
func (s *Structure) Await(resource int) {
	s.awaiting = resource
}

func (s *Structure) UnAwait() {
	s.awaiting = -1
}

func (s *Structure) ReceiveCargo(resource int) {
	s.storage[resource] = Storage{
		Resource: resource,
		Capacity: s.storage[resource].Capacity,
		Amount:   s.storage[resource].Amount + 1,
	}
	return
}

func (s *Structure) ReturnShip() {
	if s.ships < s.berths {
		s.ships += 1
	}
}
