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
		if g.structureData[s].Class == structure.Tax && g.capitols < ui.MaxCapitols && g.money >= g.structureData[structure.HQ].Cost {
			g.panel.AddButton(fmt.Sprintf("build %s ($%d)", g.structureData[structure.HQ].DisplayName, g.structureData[structure.HQ].Cost), generateConstructionCallback(g, p, structure.HQ))
		}

		if g.money >= g.structureData[s].Cost {
			g.panel.AddButton(fmt.Sprintf("build %s ($%d)", g.structureData[s].DisplayName, g.structureData[s].Cost), generateConstructionCallback(g, p, s))
		}
	}
}

func showPlayerPanel(g *Game) int {
	g.panel.AddInvertedLabel(fmt.Sprintf("population: %d/%d (need %d)", g.pop, g.maxPop, g.workersNeeded), ui.TtfRegular)
	g.panel.AddLabel(fmt.Sprintf("bank: $%d", g.money), ui.TtfRegular)
	g.panel.AddLabel("current day:", ui.TtfRegular)
	g.panel.AddBar(0, color.RGBA{0x00, 0x00, 0x00, 0xff})
	g.panel.AddDivider()
	return 5
}

func updatePlayerPanel(g *Game) {
	g.panel.UpdateLabel(0, fmt.Sprintf("population: %d/%d (need %d)", g.pop, g.maxPop, g.workersNeeded))
	g.panel.UpdateLabel(1, fmt.Sprintf("bank: $%d", g.money))
	var day float32
	day = float32(g.count % ui.DayLength)
	day = day / float32(ui.DayLength)
	day = day * 255
	g.panel.UpdateBar(3, uint8(day))
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

	if g.structureData[s.StructureType()].Prioritize > 0 && g.structureData[s.StructureType()].Prioritize <= g.money && !s.IsPaused() {
		if s.IsPrioritized() {
			g.panel.AddButton("normal deliveries", generateUnPrioritizeCallback(g, s))
		} else {
			g.panel.AddButton(fmt.Sprintf("prioritize deliveries ($%d)", g.structureData[s.StructureType()].Prioritize), generatePrioritizeCallback(g, s))

		}
	}

	if possible, up := s.Upgradeable(); possible && g.structureData[up].Cost <= g.money {
		g.panel.AddButton(fmt.Sprintf("upgrade to %s ($%d)", g.structureData[up].DisplayName, g.structureData[up].Cost), generateUpgradeCallBack(g, s, up))
	}

	g.panel.AddDivider()
	showPlanetPanel(g.panel, s.Planet(), g.resourceData)
	g.panel.AddDivider()
}
