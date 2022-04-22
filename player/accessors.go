package player

import (
	"github.com/jcgraybill/ship-shape/ship"
	"github.com/jcgraybill/ship-shape/structure"
)

func (p *Player) Structures() []*structure.Structure {
	return p.structures
}

func (p *Player) Ships() map[int]*ship.Ship {
	return p.ships
}
