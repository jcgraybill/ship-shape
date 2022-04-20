package structure

import (
	"github.com/jcgraybill/ship-shape/ui"
)

func (s *Structure) Produce(count int) bool {
	if s.data.Produces.Rate > 0 {
		if s.storage[s.data.Produces.Resource].Resource == s.data.Produces.Resource && s.storage[s.data.Produces.Resource].Amount < s.storage[s.data.Produces.Resource].Capacity {
			if s.Planet().Resources()[s.data.Produces.Requires] > 0 {
				var productionRate float32
				productionRate = float32(s.data.Produces.Rate)
				productionRate *= float32(s.Planet().Resources()[s.data.Produces.Requires]) / 255
				if s.WorkersNeeded() > 0 {
					productionRate *= float32(s.Workers()) / float32(s.WorkersNeeded())
				}
				if productionRate > 0 {
					productionRate = 255 - productionRate
					if count%(int(productionRate)*ui.BaseProductionRate) == 0 {
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
	}

	return false
}

//FIXME race condition allows structures to over-bid
func (s *Structure) Bid() (int, uint8) {
	if s.storage[s.data.Consumes].Resource == s.data.Consumes && s.storage[s.data.Consumes].Amount < s.storage[s.data.Consumes].Capacity {
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

func (s *Structure) ReceiveCargo(resource int) {
	if s.storage[resource].Amount < s.storage[resource].Capacity {
		s.storage[resource] = Storage{
			Resource: resource,
			Capacity: s.storage[resource].Capacity,
			Amount:   s.storage[resource].Amount + 1,
		}
	}
}

func (s *Structure) ReturnShip() {
	if s.ships < s.berths {
		s.ships += 1
	}
}
