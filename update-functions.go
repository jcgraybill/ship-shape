package main

import (
	"math"

	"github.com/jcgraybill/ship-shape/planet"
	"github.com/jcgraybill/ship-shape/ship"
	"github.com/jcgraybill/ship-shape/structure"
)

func produce(g *Game) {
	for _, structure := range g.structures {
		if structure.Produce(g.count) && structure.IsHighlighted() {
			g.panel.Clear()
			showStructurePanel(g, structure)
		}
	}
}

func arrive(s *ship.Ship, key int, g *Game) {
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
		destination.ReturnShip()
		delete(g.ships, key)
	}
}

func generateConstructionCallback(g *Game, p *planet.Planet, structureType int) func() {
	return func() {
		g.panel.Clear()
		structure := structure.New(g.structureData[structureType], p)
		showStructurePanel(g, structure)
		structure.Highlight()
		g.structures = append(g.structures, structure)
	}
}

func distance(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(math.Pow(math.Abs(x1-x2), 2) + math.Pow(math.Abs(y1-y2), 2))
}
