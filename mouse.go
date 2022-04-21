package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
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
			for _, p := range g.planets {
				if p.MouseButton(cx, cy) {
					clickedObject = true
					p.Highlight()
					showPlanetPanel(g.panel, p, g.resourceData)
					showBuildOptionsPanel(p, g)
				} else {
					p.Unhighlight()
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
