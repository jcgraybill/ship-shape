package main

import (
	"github.com/jcgraybill/ship-shape/ui"
)

func (g *Game) Update() error {
	g.count++

	handleInputEvents(g)

	produce(g)

	if g.count%ui.BidFrequency == 0 {
		bidForResources(g)
	}

	for key, s := range g.ships {
		if s.Update(g.count) { //ship has arrived
			arrive(s, key, g)
		}
	}

	return nil
}
