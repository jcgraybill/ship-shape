package player

import (
	"github.com/jcgraybill/ship-shape/ship"
	"github.com/jcgraybill/ship-shape/structure"
)

type Player struct {
	structures                               []*structure.Structure
	capitol                                  *structure.Structure
	ships                                    map[uint]*ship.Ship
	population, maxPopulation, workersNeeded int
	money                                    int
}

func New() *Player {
	var p Player
	p.structures = make([]*structure.Structure, 0)
	p.ships = make(map[uint]*ship.Ship)
	p.capitol = nil
	p.population, p.maxPopulation, p.workersNeeded = 0, 0, 0
	return &p
}
