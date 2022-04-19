package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/jcgraybill/ship-shape/panel"
	"github.com/jcgraybill/ship-shape/planet"
	"github.com/jcgraybill/ship-shape/resource"
	"github.com/jcgraybill/ship-shape/structure"
)

func (g *Game) Update() error {
	g.count++

	handleInputEvents(g)

	for _, structure := range g.structures {
		if structure.Update(g.count) {
			// TODO only do this if it's the currently visible structure
			g.panel.Clear()
			showStructurePanel(g, structure)
		}
	}

	return nil
}

func handleInputEvents(g *Game) {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		if !g.panel.LeftMouseButtonPress(ebiten.CursorPosition()) {
			g.panel.Clear()

			for _, planet := range g.planets {
				if planet.MouseButton(ebiten.CursorPosition()) {
					planet.Highlight()
					showPlanet(g.panel, planet, g.resourceData)
					g.panel.AddButton("build "+g.structureData[structure.Water].DisplayName, generateConstructionCallback(g, planet, structure.Water))
					g.panel.AddButton("build "+g.structureData[structure.Outpost].DisplayName, generateConstructionCallback(g, planet, structure.Outpost))
				} else {
					planet.Unhighlight()
				}
			}

			for _, structure := range g.structures {
				if structure.MouseButton(ebiten.CursorPosition()) {
					showStructurePanel(g, structure)
				}
			}

		}
	} else if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		g.panel.LeftMouseButtonRelease(ebiten.CursorPosition())
	}
}

func generateConstructionCallback(g *Game, p *planet.Planet, structureType int) func() {
	return func() {
		g.panel.Clear()
		structure := structure.New(g.structureData[structureType], p)
		showStructurePanel(g, structure)
		g.structures = append(g.structures, structure)
	}
}

func showPlanet(panel *panel.Panel, p *planet.Planet, rd [resource.ResourceDataLength]resource.ResourceData) {
	panel.AddLabel(fmt.Sprintf("planet: %s", p.Name()))
	for resource, level := range p.Resources() {
		panel.AddLabel(rd[resource].DisplayName)
		panel.AddBar(level, rd[resource].Color)
	}
}

func showStructure(panel *panel.Panel, s *structure.Structure, rd [resource.ResourceDataLength]resource.ResourceData) {
	panel.AddLabel(s.Name())

	if s.Storage().Capacity > 0 {
		panel.AddDivider()
		panel.AddLabel("storage:")
		panel.AddLabel(fmt.Sprintf("%s (%d/%d)", rd[s.Storage().Resource].DisplayName, s.Storage().Amount, s.Storage().Capacity))
		panel.AddBar(uint8((255*int(s.Storage().Amount))/int(s.Storage().Capacity)), rd[s.Storage().Resource].Color)
	}
}

func showStructurePanel(g *Game, structure *structure.Structure) {
	showStructure(g.panel, structure, g.resourceData)
	g.panel.AddDivider()
	showPlanet(g.panel, structure.Planet(), g.resourceData)
	g.panel.AddDivider()
}
