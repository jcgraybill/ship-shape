package structure

import (
	"github.com/jcgraybill/ship-shape/resource"
	"github.com/jcgraybill/ship-shape/ui"
)

func (s *Structure) Produce(count int) bool {
	if s.IsPaused() {
		return false
	}
	if s.data.Produces.Rate > 0 {
		if s.storage[s.data.Produces.Resource].Resource == s.data.Produces.Resource {
			if s.storage[s.data.Produces.Resource].Amount < s.storage[s.data.Produces.Resource].Capacity {
				productionRate := float32(s.data.Produces.Rate)

				for _, ingredient := range s.data.Produces.Requires {
					if s.Planet().Resources()[ingredient.Resource] > 0 {
						productionRate *= float32(s.Planet().Resources()[ingredient.Resource]) / 255
					} else if s.storage[ingredient.Resource].Amount < ingredient.Quantity {
						productionRate = 0
					}
				}

				if s.WorkerCapacity() > 0 {
					productionRate *= float32(s.ActiveWorkers()) / float32(s.WorkerCapacity())
				}

				if productionRate > 0 {
					productionRate = ui.BaseProductionRate / productionRate
					if count%int(productionRate) == 0 {
						s.storage[s.data.Produces.Resource] = &Storage{
							Resource: s.data.Produces.Resource,
							Capacity: s.storage[s.data.Produces.Resource].Capacity,
							Amount:   s.storage[s.data.Produces.Resource].Amount + 1,
						}

						for _, ingredient := range s.data.Produces.Requires {
							if s.Planet().Resources()[ingredient.Resource] == 0 {
								if s.storage[ingredient.Resource].Resource == ingredient.Resource {
									s.storage[ingredient.Resource] = &Storage{
										Resource: s.storage[ingredient.Resource].Resource,
										Capacity: s.storage[ingredient.Resource].Capacity,
										Amount:   s.storage[ingredient.Resource].Amount - ingredient.Quantity,
									}
								}
							}
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
const PriorityBidValue = 255

func (s *Structure) Bid() map[int]uint8 {
	bids := make(map[int]uint8)
	if s.IsPaused() {
		return bids
	}
	for _, r := range s.resourcesWanted {
		if s.storage[r].Amount < s.storage[r].Capacity {
			if s.prioritized {
				bids[r] = PriorityBidValue
			} else {
				bids[r] = uint8(((float32(s.storage[r].Capacity) - float32(s.storage[r].Amount)) / float32(s.storage[r].Capacity)) * 255)
			}
		}
	}
	return bids
}

func (s *Structure) LaunchShip(resource int) {

	s.ships -= 1
	s.inFlight += 1
	if s.Class() != Tax {
		s.storage[resource] = &Storage{
			Resource: resource,
			Capacity: s.storage[resource].Capacity,
			Amount:   s.storage[resource].Amount - 1,
		}
	}
}

func (s *Structure) ReceiveCargo(resource int) {
	if s.storage[resource].Amount < s.storage[resource].Capacity {
		s.storage[resource] = &Storage{
			Resource: resource,
			Capacity: s.storage[resource].Capacity,
			Amount:   s.storage[resource].Amount + 1,
		}
	}
}

func (s *Structure) ReturnShip() {
	s.inFlight -= 1
	if s.ships < s.berths {
		s.ships += 1
	}
}

func (s *Structure) GenerateIncome() {
	s.income += (float64(s.Storage()[resource.Population].Amount) * ui.IncomeRate) / ui.YearLength
}

func (s *Structure) Consume(count int) (bool, int) {
	var consumed bool
	var downgrade int
	for _, c := range s.data.Consumes {
		productionRate := float32(c.Rate)
		productionRate = ui.BaseProductionRate / productionRate
		if count%int(productionRate) == 0 {
			if s.storage[c.Resource].Amount > 0 {
				s.storage[c.Resource] = &Storage{
					Resource: c.Resource,
					Capacity: s.storage[c.Resource].Capacity,
					Amount:   s.storage[c.Resource].Amount - 1,
				}
				consumed = true
			} else {
				// TODO test this
				if s.data.Downgrade.Structure > 0 {
					for _, r := range s.data.Downgrade.Required {
						if r == c.Resource {
							downgrade = s.data.Downgrade.Structure
						}
					}
				}
			}
		}
	}
	return consumed, downgrade
}
