package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/jcgraybill/ship-shape/planet"
	"github.com/jcgraybill/ship-shape/structure"
	"github.com/jcgraybill/ship-shape/ui"
)

func handleMouseClicks(g *Game) {

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		if !g.panel.LeftMouseButtonPress(ebiten.CursorPosition()) {
			cx, cy := ebiten.CursorPosition()
			cx -= g.offsetX
			cy -= g.offsetY
			g.panel.Clear()
			showPlayerPanel(g.panel, g.money, g.pop, g.maxPop, g.workersNeeded)

			clickedObject := false
			for _, planet := range g.planets {
				if planet.MouseButton(cx, cy) {
					clickedObject = true
					planet.Highlight()
					showPlanetPanel(g.panel, planet, g.resourceData)
					if g.money >= g.structureData[structure.Water].Cost {
						g.panel.AddButton(fmt.Sprintf("build %s ($%d)", g.structureData[structure.Water].DisplayName, g.structureData[structure.Water].Cost), generateConstructionCallback(g, planet, structure.Water))
					}
					if g.money >= g.structureData[structure.Outpost].Cost {
						g.panel.AddButton(fmt.Sprintf("build %s ($%d)", g.structureData[structure.Outpost].DisplayName, g.structureData[structure.Outpost].Cost), generateConstructionCallback(g, planet, structure.Outpost))
					}
				} else {
					planet.Unhighlight()
				}
			}

			for _, structure := range g.structures {
				if structure.MouseButton(cx, cy) {
					clickedObject = true
					structure.Highlight()
					showStructurePanel(g, structure)
				} else {
					structure.Unhighlight()
				}
			}

			for _, ship := range g.ships {
				if ship.MouseButton(cx, cy) {
					clickedObject = true
					g.panel.Clear()
					showPlayerPanel(g.panel, g.money, g.pop, g.maxPop, g.workersNeeded)
					g.panel.AddLabel("ship", ui.TtfBold)
					cargo, origin, destination := ship.Manifest()
					if cargo > 0 {
						g.panel.AddLabel(fmt.Sprintf("carrying %s\nfrom %s\nat %s\nto %s\nat %s", g.resourceData[cargo].DisplayName, origin.Name(), origin.Planet().Name(), destination.Name(), destination.Planet().Name()), ui.TtfRegular)
					} else {
						g.panel.AddLabel(fmt.Sprintf("returning to %s", destination.Planet().Name()), ui.TtfRegular)
					}
				}
			}

			//TODO: click-drag should not cause a selected structure/planet to lose focus
			if !clickedObject {
				g.dragging = true
				g.mouseDragX, g.mouseDragY = ebiten.CursorPosition()
			}
		}
	} else if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		g.dragging = false
		g.panel.LeftMouseButtonRelease(ebiten.CursorPosition())
	} else if g.dragging && ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		newX, newY := ebiten.CursorPosition()

		var xMovement, yMovement int

		if newX < g.mouseDragX && -1*g.offsetX+g.windowW < ui.W {
			xMovement = newX - g.mouseDragX
			g.offsetX += xMovement
		}

		if newX > g.mouseDragX && g.offsetX < 0 {
			xMovement = newX - g.mouseDragX
			g.offsetX += xMovement
		}

		if newY > g.mouseDragY && g.offsetY < 0 {
			yMovement = newY - g.mouseDragY
			g.offsetY += yMovement
		}
		if newY < g.mouseDragY && -1*g.offsetY+g.windowH < ui.H {
			yMovement = newY - g.mouseDragY
			g.offsetY += yMovement
		}

		g.mouseDragX = newX
		g.mouseDragY = newY

		g.opts.GeoM.Translate(float64(xMovement), float64(yMovement))

	}
}

func generateConstructionCallback(g *Game, p *planet.Planet, structureType int) func() {
	return func() {
		g.panel.Clear()
		g.money -= g.structureData[structureType].Cost
		structure := structure.New(g.structureData[structureType], p)
		g.structures = append(g.structures, structure)
		updatePopulation(g)
		showPlayerPanel(g.panel, g.money, g.pop, g.maxPop, g.workersNeeded)
		showStructurePanel(g, structure)
		structure.Highlight()
	}
}
