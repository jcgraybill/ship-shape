package main

import (
	"fmt"

	"github.com/jcgraybill/ship-shape/panel"
	"github.com/jcgraybill/ship-shape/planet"
	"github.com/jcgraybill/ship-shape/resource"
	"github.com/jcgraybill/ship-shape/structure"
	"github.com/jcgraybill/ship-shape/ui"
)

func showPopulationPanel(panel *panel.Panel, pop, maxPop, workersNeeded int) {
	panel.AddInvertedLabel(fmt.Sprintf("Population: %d/%d (need %d)", pop, maxPop, workersNeeded), ui.TtfRegular)
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
		panel.AddLabel(fmt.Sprintf("%d/%d workers", s.Workers(), s.WorkersNeeded()), ui.TtfRegular)
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
