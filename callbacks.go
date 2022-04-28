package main

import (
	"github.com/jcgraybill/ship-shape/planet"
	"github.com/jcgraybill/ship-shape/structure"
)

func generateConstructionCallback(g *Game, p *planet.Planet, structureType int) func() {
	return func() {
		g.redrawPSLayer = true
		g.panel.Clear()
		s := structure.New(structureType, &g.structureData[structureType], p)
		g.player.RemoveMoney(uint(g.structureData[structureType].Cost))
		g.player.AddStructure(s)
		g.updatePopulation()
		showStructurePanel(g, s)
		s.Highlight()
	}
}

func generateUpgradeCallBack(g *Game, s *structure.Structure, toStructure int) func() {
	return func() {
		g.redrawPSLayer = true
		g.panel.Clear()
		g.player.RemoveMoney(uint(g.structureData[toStructure].Cost))
		s.Upgrade(toStructure, &g.structureData[toStructure])
		g.updatePopulation()
		showStructurePanel(g, s)
		s.Highlight()
	}

}

func generatePauseCallback(g *Game, s *structure.Structure) func() {
	return func() {
		s.Pause()
		g.panel.Clear()
		g.updatePopulation()
		showStructurePanel(g, s)
		s.Highlight()
	}
}

func generateUnPauseCallback(g *Game, s *structure.Structure) func() {
	return func() {
		s.Unpause()
		g.panel.Clear()
		g.updatePopulation()
		showStructurePanel(g, s)
		s.Highlight()
	}
}

func generatePrioritizeCallback(g *Game, s *structure.Structure) func() {
	return func() {
		s.Prioritize()
		g.player.RemoveMoney(uint(g.structureData[s.StructureType()].Prioritize))
		g.panel.Clear()
		g.updatePopulation()
		showStructurePanel(g, s)
		s.Highlight()
	}
}

func generateUnPrioritizeCallback(g *Game, s *structure.Structure) func() {
	return func() {
		s.Deprioritize()
		g.panel.Clear()
		g.updatePopulation()
		showStructurePanel(g, s)
		s.Highlight()
	}
}
