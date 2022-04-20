package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/jcgraybill/ship-shape/structure"
	"github.com/jcgraybill/ship-shape/ui"
)

// TODO Click & drag or arrow keys to move screen

func handleInputEvents(g *Game) {

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
					showPlanet(g.panel, planet, g.resourceData)
					g.panel.AddButton("build "+g.structureData[structure.Water].DisplayName, generateConstructionCallback(g, planet, structure.Water))
					g.panel.AddButton("build "+g.structureData[structure.Outpost].DisplayName, generateConstructionCallback(g, planet, structure.Outpost))
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
					g.panel.AddLabel("ship", ui.TtfBold)
					cargo, origin, destination := ship.Manifest()
					if cargo > 0 {
						g.panel.AddLabel(fmt.Sprintf("carrying %s\nfrom %s\nat %s\nto %s\nat %s", g.resourceData[cargo].DisplayName, origin.Name(), origin.Planet().Name(), destination.Name(), destination.Planet().Name()), ui.TtfRegular)
					} else {
						g.panel.AddLabel(fmt.Sprintf("returning to %s", destination.Planet().Name()), ui.TtfRegular)
					}
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

	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) && -1*g.offsetX+g.windowW < ui.W {
		g.opts.GeoM.Translate(-ui.ArrowKeyMoveSpeed, 0)
		g.offsetX -= ui.ArrowKeyMoveSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) && g.offsetX < 0 {
		g.opts.GeoM.Translate(ui.ArrowKeyMoveSpeed, 0)
		g.offsetX += ui.ArrowKeyMoveSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) && g.offsetY < 0 {
		g.opts.GeoM.Translate(0, ui.ArrowKeyMoveSpeed)
		g.offsetY += ui.ArrowKeyMoveSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) && -1*g.offsetY+g.windowH < ui.H {
		g.opts.GeoM.Translate(0, -ui.ArrowKeyMoveSpeed)
		g.offsetY -= ui.ArrowKeyMoveSpeed
	}

}
