package main

import (
	"github.com/jcgraybill/ship-shape/ship"
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
			cargo, origin, destination := s.Manifest()
			if cargo > 0 {
				destination.ReceiveCargo(cargo)
				if destination.IsHighlighted() {
					g.panel.Clear()
					showStructurePanel(g, destination)
				}
				returnShip := ship.New(destination, origin)
				g.ships[key] = returnShip
			} else {
				origin.UnAwait()
				destination.ReturnShip()
				delete(g.ships, key)
			}
		}
	}

	return nil
}

func produce(g *Game) {
	for _, structure := range g.structures {
		if structure.Produce(g.count) && structure.IsHighlighted() {
			g.panel.Clear()
			showStructurePanel(g, structure)
		}
	}
}
