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

			clickedObject := false
			for _, planet := range g.planets {
				if planet.MouseButton(cx, cy) {
					clickedObject = true
					planet.Highlight()
					showPlanetPanel(g.panel, planet, g.resourceData)
					if g.capitols < ui.MaxCapitols && g.money >= g.structureData[structure.Capitol].Cost {
						g.panel.AddButton(fmt.Sprintf("build %s ($%d)", g.structureData[structure.Capitol].DisplayName, g.structureData[structure.Capitol].Cost), generateConstructionCallback(g, planet, structure.Capitol))
					}
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
		st := structure.New(structureType, g.structureData[structureType], p)
		g.structures = append(g.structures, st)
		updatePopulation(g)
		showStructurePanel(g, st)
		st.Highlight()
		if structureType == structure.Capitol {
			g.capitols += 1
		}
	}
}
