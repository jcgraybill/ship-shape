package main

import (
	"github.com/jcgraybill/ship-shape/planet"
	"github.com/jcgraybill/ship-shape/structure"
)

func generateConstructionCallback(g *Game, p *planet.Planet, structureType int) func() {
	return func() {
		g.panel.Clear()
		g.money -= g.structureData[structureType].Cost
		st := structure.New(structureType, g.structureData[structureType], p)
		g.structures = append(g.structures, st)
		updatePopulation(g)
		showStructurePanel(g, st)
		st.Highlight()
		if structureType == structure.Capitol {
			g.capitols += 1
		}
	}
}

func generateUpgradeCallBack(g *Game, s *structure.Structure, structureType int) func() {
	return func() {
		g.panel.Clear()
		g.money -= g.structureData[structureType].Cost
		s.Upgrade(structureType, g.structureData[structureType])
		updatePopulation(g)
		showStructurePanel(g, s)
		s.Highlight()
	}

}
