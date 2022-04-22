package player

import (
	"github.com/jcgraybill/ship-shape/ship"
	"github.com/jcgraybill/ship-shape/structure"
)

type Player struct {
	structures                 []*structure.Structure
	capitols                   []*structure.Structure
	ships                      map[int]*ship.Ship
	pop, maxPop, workersNeeded int
	money                      int
}

func New() *Player {
	var p Player
	p.structures = make([]*structure.Structure, 0)
	p.capitols = make([]*structure.Structure, 0)
	p.ships = make(map[int]*ship.Ship)

	p.pop, p.maxPop, p.workersNeeded = 0, 0, 0
	return &p
}
