package structure

import (
	"fmt"

	"github.com/jcgraybill/ship-shape/planet"
)

func (s *Structure) Name() string {
	return s.data.DisplayName
}

func (s *Structure) Planet() *planet.Planet {
	return s.planet
}

func (s *Structure) Storage() map[int]*Storage {
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
	return int(s.income)
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

func (s *Structure) UpgradeTo() int {
	return s.data.Upgrade.Structure
}

func (s *Structure) Upgradeable() bool {
	if s.data.Upgrade.Structure > 0 {
		upgradeable := true
		for _, r := range s.data.Upgrade.Required {
			if s.storage[r].Amount < s.storage[r].Capacity {
				upgradeable = false
			}
		}
		if upgradeable {
			return true
		}
	}
	return false
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

func (s *Structure) WorkersLabel() string {
	return fmt.Sprintf("%d/%d workers ($%d/year)", s.ActiveWorkers(), s.WorkerCapacity(), s.LaborCost())
}

func (s *Structure) ResourceLabelCallback(displayName string, resource int) func() string {
	return func() string {
		return fmt.Sprintf("%s (%d/%d)", displayName, s.Storage()[resource].Amount, s.Storage()[resource].Capacity)
	}
}

func (s *Structure) ResourceBarCallback(resource int) func() uint8 {
	return func() uint8 {
		return uint8((255 * int(s.Storage()[resource].Amount)) / int(s.Storage()[resource].Capacity))
	}
}
