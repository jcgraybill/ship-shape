package main

import (
	"fmt"
	"image/color"

	"github.com/jcgraybill/ship-shape/panel"
	"github.com/jcgraybill/ship-shape/planet"
	"github.com/jcgraybill/ship-shape/resource"
	"github.com/jcgraybill/ship-shape/structure"
	"github.com/jcgraybill/ship-shape/ui"
)

func showBuildOptionsPanel(p *planet.Planet, g *Game) {

	for _, s := range g.level.AllowedStructures() {

		if g.structureData[s].Class == structure.Tax {
			if hasCapitol, _ := g.player.Capitol(); !hasCapitol && g.player.Money() >= g.structureData[structure.HQ].Cost {
				g.panel.AddButton(fmt.Sprintf("build %s ($%d)", g.structureData[structure.HQ].DisplayName, g.structureData[structure.HQ].Cost), generateConstructionCallback(g, p, structure.HQ))
			}
		} else if g.player.Money() >= g.structureData[s].Cost {
			g.panel.AddButton(fmt.Sprintf("build %s ($%d)", g.structureData[s].DisplayName, g.structureData[s].Cost), generateConstructionCallback(g, p, s))
		}
	}
}

func showPlayerPanel(g *Game) int {
	g.panel.AddBar(0, color.RGBA{0x00, 0x00, 0x00, 0xff})
	g.panel.AddInvertedLabel(fmt.Sprintf("year: %d", g.year), ui.TtfBold)
	pop, maxPop, workersNeeded := g.player.Population()
	g.panel.AddLabel(fmt.Sprintf("population: %d/%d (need %d)", pop, maxPop, workersNeeded), ui.TtfRegular)
	g.panel.AddLabel(fmt.Sprintf("bank: $%d", g.player.Money()), ui.TtfRegular)
	g.panel.AddDivider()
	g.panel.AddLabel(g.level.Message(), ui.TtfRegular)
	g.panel.AddDivider()
	return 7
}

func updatePlayerPanel(g *Game) {
	var year float32
	year = float32(g.count % ui.YearLength)
	year = year / float32(ui.YearLength)
	year = year * 255
	g.panel.UpdateBar(0, uint8(year))
	g.panel.UpdateLabel(1, fmt.Sprintf("year: %d", g.year))
	pop, maxPop, workersNeeded := g.player.Population()
	g.panel.UpdateLabel(2, fmt.Sprintf("population: %d/%d (need %d)", pop, maxPop, workersNeeded))
	g.panel.UpdateLabel(3, fmt.Sprintf("bank: $%d", g.player.Money()))
	g.panel.UpdateLabel(5, g.level.Message())
}

func showPlanetPanel(pl *panel.Panel, p *planet.Planet, rd [resource.ResourceDataLength]resource.ResourceData) {
	pl.AddLabel(fmt.Sprintf("planet: %s", p.Name()), ui.TtfBold)
	for resource, level := range p.Resources() {
		pl.AddLabel(rd[resource].DisplayName, ui.TtfRegular)
		pl.AddBar(level, rd[resource].Color)
	}
}

func showStructure(pl *panel.Panel, s *structure.Structure, rd [resource.ResourceDataLength]resource.ResourceData) {
	pl.AddLabel(s.Name(), ui.TtfBold)
	if s.WorkerCapacity() > 0 {
		pl.AddLabel(fmt.Sprintf("%d/%d workers ($%d/day)", s.ActiveWorkers(), s.WorkerCapacity(), s.LaborCost()), ui.TtfRegular)
	}
	if len(s.Storage()) > 0 {
		pl.AddDivider()
		for _, st := range s.Storage() {
			pl.AddLabel(fmt.Sprintf("%s (%d/%d)", rd[st.Resource].DisplayName, st.Amount, st.Capacity), ui.TtfRegular)
			pl.AddBar(uint8((255*int(st.Amount))/int(st.Capacity)), rd[st.Resource].Color)
		}
	}
}

func showStructurePanel(g *Game, s *structure.Structure) {
	showStructure(g.panel, s, g.resourceData)

	if g.structureData[s.StructureType()].Workers > 0 {
		if s.IsPaused() {
			g.panel.AddButton("resume production", generateUnPauseCallback(g, s))
		} else {
			g.panel.AddButton("pause production", generatePauseCallback(g, s))
		}
	}

	if g.structureData[s.StructureType()].Prioritize > 0 && g.structureData[s.StructureType()].Prioritize <= g.player.Money() && !s.IsPaused() {
		if s.IsPrioritized() {
			g.panel.AddButton("normal deliveries", generateUnPrioritizeCallback(g, s))
		} else {
			g.panel.AddButton(fmt.Sprintf("prioritize deliveries ($%d)", g.structureData[s.StructureType()].Prioritize), generatePrioritizeCallback(g, s))

		}
	}

	if possible, up := s.Upgradeable(); possible && g.structureData[up].Cost <= g.player.Money() {
		g.panel.AddButton(fmt.Sprintf("upgrade to %s ($%d)", g.structureData[up].DisplayName, g.structureData[up].Cost), generateUpgradeCallBack(g, s, up))
	}

	g.panel.AddDivider()
	showPlanetPanel(g.panel, s.Planet(), g.resourceData)
	g.panel.AddDivider()
}
