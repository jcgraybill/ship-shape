package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/jcgraybill/ship-shape/structure"
	"github.com/jcgraybill/ship-shape/ui"
)

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
					structure.Highlight()
					showStructurePanel(g, structure)
				} else {
					structure.Unhighlight()
				}
			}

			for _, ship := range g.ships {
				if ship.MouseButton(ebiten.CursorPosition()) {
					g.panel.Clear()
					g.panel.AddLabel("ship", ui.TtfBold)
					cargo, origin, destination := ship.Manifest()
					if cargo > 0 {
						g.panel.AddLabel(fmt.Sprintf("carrying %s\nfrom %s\nat %s\nto %s\nat %s", g.resourceData[cargo].DisplayName, origin.Name(), origin.Planet().Name(), destination.Name(), destination.Planet().Name()), ui.TtfRegular)
					} else {
						g.panel.AddLabel(fmt.Sprintf("returning to %s", destination.Planet().Name()), ui.TtfRegular)
					}
				}
			}

		}
	} else if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		g.panel.LeftMouseButtonRelease(ebiten.CursorPosition())
	}
}
