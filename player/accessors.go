package player

import (
	"fmt"

	"github.com/jcgraybill/ship-shape/ship"
	"github.com/jcgraybill/ship-shape/structure"
)

func (p *Player) Capitol() (bool, *structure.Structure) {
	if p.capitol != nil {
		return true, p.capitol
	} else {
		return false, nil
	}
}

func (p *Player) MaxPopulation() int {
	return p.maxPopulation
}

func (p *Player) Money() int {
	return p.money
}

func (p *Player) MoneyLabel() string {
	return fmt.Sprintf("bank: $%d", p.money)
}

func (p *Player) Population() int {
	return p.population
}

func (p *Player) PopulationLabel() string {
	return fmt.Sprintf("population: %d/%d (need %d)", p.population, p.maxPopulation, p.workersNeeded)
}

func (p *Player) Ships() map[uint]*ship.Ship {
	return p.ships
}

func (p *Player) Structures() []*structure.Structure {
	return p.structures
}
