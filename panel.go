package main

import (
	"fmt"

	"github.com/jcgraybill/ship-shape/panel"
	"github.com/jcgraybill/ship-shape/planet"
	"github.com/jcgraybill/ship-shape/resource"
	"github.com/jcgraybill/ship-shape/structure"
	"github.com/jcgraybill/ship-shape/ui"
)

func showBuildOptionsPanel(p *planet.Planet, g *Game) {
	for _, s := range g.level.AllowedStructures() {
		if g.structureData[s].Buildable {
			if g.structureData[s].Class == structure.Tax {
				g.panel.AddButton(fmt.Sprintf("build %s ($%d)", g.structureData[structure.HQ].DisplayName, g.structureData[structure.HQ].Cost),
					generateConstructionCallback(g, p, structure.HQ),
					func() bool {
						if hasCapitol, _ := g.player.Capitol(); !hasCapitol && g.player.Money() >= g.structureData[structure.HQ].Cost {
							return true
						} else {
							return false
						}
					},
				)
			} else {
				cost := g.structureData[s].Cost
				g.panel.AddButton(fmt.Sprintf("build %s ($%d)", g.structureData[s].DisplayName, cost),
					generateConstructionCallback(g, p, s),
					func() bool {
						if g.player.Money() >= cost {
							return true
						} else {
							return false
						}
					},
				)
			}
		}
	}
}

func showPlanetPanel(pl *panel.Panel, p *planet.Planet, rd *[resource.ResourceDataLength]resource.ResourceData, allowed []int) {
	pl.AddLabel(func() string { return fmt.Sprintf("planet: %s", p.Name()) }, ui.TtfBold)
	for _, resource := range allowed {
		if level, exists := p.Resources()[resource]; exists {
			lvl := level
			name := rd[resource].DisplayName
			pl.AddLabel(func() string { return name }, ui.TtfRegular)
			pl.AddBar(func() uint8 { return lvl }, rd[resource].Color)
		}
	}
}

func showPlayerPanel(g *Game) int {
	g.panel.AddInvertedLabel(g.level.Title, ui.TtfBold)
	g.panel.AddLabel(g.player.PopulationLabel, ui.TtfRegular)
	g.panel.AddLabel(g.player.MoneyLabel, ui.TtfRegular)
	g.panel.AddDivider()
	g.panel.AddLabel(g.level.Message, ui.TtfRegular)
	g.panel.AddDivider()
	g.panel.AddLabel(g.level.Label, ui.TtfRegular)
	g.panel.AddBar(g.level.ProgressBarValue, ui.LevelProgressColor)
	g.panel.AddDivider()
	return 9
}

func showStructure(pl *panel.Panel, s *structure.Structure, rd *[resource.ResourceDataLength]resource.ResourceData, allowed []int) {
	pl.AddLabel(func() string { return s.Name() }, ui.TtfBold)
	if s.WorkerCapacity() > 0 {
		pl.AddLabel(s.WorkersLabel, ui.TtfRegular)
	}
	if len(s.Storage()) > 0 {
		pl.AddDivider()

		for _, resource := range allowed {
			if st, exists := s.Storage()[resource]; exists {
				pl.AddLabel(s.ResourceLabelCallback(rd[resource].DisplayName, resource), ui.TtfRegular)
				pl.AddBar(s.ResourceBarCallback(resource), rd[st.Resource].Color)
			}
		}
	}
}

func showStructurePanel(g *Game, s *structure.Structure) {
	showStructure(g.panel, s, g.resourceData, g.level.AllowedResources())

	if g.structureData[s.StructureType()].Workers > 0 {
		if s.IsPaused() {
			g.panel.AddButton("resume production", generateUnPauseCallback(g, s), func() bool { return true })
		} else {
			g.panel.AddButton("pause production", generatePauseCallback(g, s), func() bool { return true })
		}
	}

	if g.structureData[s.StructureType()].Prioritize > 0 && !s.IsPaused() {
		if s.IsPrioritized() {
			g.panel.AddButton("normal deliveries", generateUnPrioritizeCallback(g, s), func() bool { return true })
		} else {
			g.panel.AddButton(fmt.Sprintf("prioritize deliveries ($%d)", g.structureData[s.StructureType()].Prioritize),
				generatePrioritizeCallback(g, s),
				func() bool {
					if g.player.Money() >= g.structureData[s.StructureType()].Prioritize {
						return true
					} else {
						return false
					}
				},
			)

		}
	}

	for _, st := range g.level.AllowedStructures() {
		if st > 0 && st == s.UpgradeTo() {
			cost := g.structureData[s.UpgradeTo()].Cost
			s1 := s
			g.panel.AddButton(fmt.Sprintf("upgrade to %s ($%d)", g.structureData[s.UpgradeTo()].DisplayName, cost),
				generateUpgradeCallBack(g, s, s.UpgradeTo()),
				func() bool {
					if g.player.Money() >= cost && s1.Upgradeable() {
						return true
					} else {
						return false
					}
				},
			)
		}
	}
	g.panel.AddDivider()
	showPlanetPanel(g.panel, s.Planet(), g.resourceData, g.level.AllowedResources())
	g.panel.AddDivider()
}
