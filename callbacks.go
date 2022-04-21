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
		if structureType == structure.HQ {
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

func generatePauseCallback(g *Game, s *structure.Structure) func() {
	return func() {
		s.Pause()
		g.panel.Clear()
		updatePopulation(g)
		showStructurePanel(g, s)
		s.Highlight()
	}
}

func generateUnPauseCallback(g *Game, s *structure.Structure) func() {
	return func() {
		s.Unpause()
		g.panel.Clear()
		updatePopulation(g)
		showStructurePanel(g, s)
		s.Highlight()
	}
}

func generatePrioritizeCallback(g *Game, s *structure.Structure) func() {
	return func() {
		s.Prioritize()
		g.money -= g.structureData[s.StructureType()].Prioritize
		g.panel.Clear()
		updatePopulation(g)
		showStructurePanel(g, s)
		s.Highlight()
	}
}

func generateUnPrioritizeCallback(g *Game, s *structure.Structure) func() {
	return func() {
		s.Deprioritize()
		g.panel.Clear()
		updatePopulation(g)
		showStructurePanel(g, s)
		s.Highlight()
	}
}
