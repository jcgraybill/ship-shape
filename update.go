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

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		handleLeftMouseButtonPress(g)
	} else if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		handleLeftMouseButtonRelease(g)
	}

	//TODO let every structure update its state
	for _, structure := range g.structures {
		if err := structure.Update(); err != nil {
			return err
		}
	}

	return nil
}

func handleLeftMouseButtonRelease(g *Game) {
	g.panel.LeftMouseButtonRelease(ebiten.CursorPosition())
}

func handleLeftMouseButtonPress(g *Game) {
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
				g.panel.AddLabel(structure.Name())
				showPlanet(g.panel, structure.Planet(), g.resourceData)
			}
		}

	}
}

func generateConstructionCallback(g *Game, p *planet.Planet, structureType int) func() {
	return func() {
		g.panel.Clear()
		structure := structure.New(g.structureData[structureType], p)
		g.panel.AddLabel(structure.Name())
		showPlanet(g.panel, structure.Planet(), g.resourceData)
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
