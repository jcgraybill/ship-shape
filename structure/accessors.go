package structure

import (
	"github.com/jcgraybill/ship-shape/planet"
)

func (s *Structure) Name() string {
	return s.data.DisplayName
}

func (s *Structure) Planet() *planet.Planet {
	return s.planet
}

func (s *Structure) Storage() map[int]Storage {
	return s.storage
}

func (s *Structure) Produces() int {
	return s.data.Produces.Resource
}

func (s *Structure) HasShips() bool {
	if s.IsPaused() {
		return false
	}

	if s.Class() == Tax {
		maxShips := s.workers
		if s.berths < maxShips {
			maxShips = s.berths
		}
		s.ships = maxShips - s.inFlight
	}

	if !s.IsPaused() && s.ships == 0 && s.inFlight == 0 {
		s.ships = s.data.MinShips
	}

	if s.ships == 0 {
		return false
	}

	return true
}

func (s *Structure) LaborCost() int {
	return s.workers * s.data.WorkerCost
}

func (s *Structure) WorkerCost() int {
	return s.data.WorkerCost
}

func (s *Structure) CanProduce() bool {
	if s.Class() == Tax {
		return true
	}
	if s.Storage()[s.data.Produces.Resource].Amount < s.Storage()[s.data.Produces.Resource].Capacity {
		return true
	}
	return false
}

func (s *Structure) Class() int {
	return s.data.Class
}

func (s *Structure) Income() int {
	return s.income
}

func (s *Structure) StructureType() int {
	return s.structureType
}

func (s *Structure) IsPaused() bool {
	return s.paused
}

func (s *Structure) IsPrioritized() bool {
	return s.prioritized
}

func (s *Structure) Upgradeable() (bool, int) {
	if s.data.Upgrade.Structure > 0 {
		upgradeable := true
		for _, r := range s.data.Upgrade.Required {
			if s.storage[r].Amount < s.storage[r].Capacity {
				upgradeable = false
			}
		}
		if upgradeable {
			return true, s.data.Upgrade.Structure
		}
	}
	return false, s.data.Upgrade.Structure
}

func (s *Structure) ActiveWorkers() int {
	return s.workers
}
func (s *Structure) WorkerCapacity() int {
	if s.paused {
		return 0
	}
	return s.data.Workers
}
