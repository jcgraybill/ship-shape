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

func showPlanetPanel(panel *panel.Panel, p *planet.Planet, rd [resource.ResourceDataLength]resource.ResourceData) {
	panel.AddLabel(fmt.Sprintf("planet: %s", p.Name()), ui.TtfBold)
	for resource, level := range p.Resources() {
		panel.AddLabel(rd[resource].DisplayName, ui.TtfRegular)
		panel.AddBar(level, rd[resource].Color)
	}
}

func showStructure(panel *panel.Panel, s *structure.Structure, rd [resource.ResourceDataLength]resource.ResourceData) {
	panel.AddLabel(s.Name(), ui.TtfBold)
	if s.WorkersNeeded() > 0 {
		panel.AddLabel(fmt.Sprintf("%d/%d workers ($%d/day)", s.Workers(), s.WorkersNeeded(), s.LaborCost()), ui.TtfRegular)
	}
	if len(s.Storage()) > 0 {
		panel.AddDivider()
		panel.AddLabel("storage:", ui.TtfRegular)
		for _, st := range s.Storage() {
			panel.AddLabel(fmt.Sprintf("%s (%d/%d)", rd[st.Resource].DisplayName, st.Amount, st.Capacity), ui.TtfRegular)
			panel.AddBar(uint8((255*int(st.Amount))/int(st.Capacity)), rd[st.Resource].Color)
		}
	}
}

func showStructurePanel(g *Game, structure *structure.Structure) {
	showStructure(g.panel, structure, g.resourceData)
	g.panel.AddDivider()
	showPlanetPanel(g.panel, structure.Planet(), g.resourceData)
	g.panel.AddDivider()
}
