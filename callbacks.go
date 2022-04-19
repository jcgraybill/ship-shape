package main

import (
	"github.com/jcgraybill/ship-shape/planet"
	"github.com/jcgraybill/ship-shape/structure"
)

func generateConstructionCallback(g *Game, p *planet.Planet, structureType int) func() {
	return func() {
		g.panel.Clear()
		structure := structure.New(g.structureData[structureType], p)
		showStructurePanel(g, structure)
		structure.Highlight()
		g.structures = append(g.structures, structure)
	}
}
