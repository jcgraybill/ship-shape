package player

import (
	"github.com/jcgraybill/ship-shape/ship"
	"github.com/jcgraybill/ship-shape/structure"
)

func (p *Player) Structures() []*structure.Structure {
	return p.structures
}

func (p *Player) Ships() map[uint]*ship.Ship {
	return p.ships
}

func (p *Player) Money() int {
	return p.money
}

func (p *Player) Population() (int, int, int) {
	return p.population, p.maxPopulation, p.workersNeeded
}

func (p *Player) Capitol() (bool, *structure.Structure) {
	if p.capitol != nil {
		return true, p.capitol
	} else {
		return false, nil
	}
}
