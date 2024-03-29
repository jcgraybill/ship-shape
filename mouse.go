package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func (g *Game) handleMouseClicks() {

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		if !g.panel.LeftMouseButtonPress(ebiten.CursorPosition()) {
			g.redrawPSLayer = true
			cx, cy := ebiten.CursorPosition()
			cx -= g.offsetX
			cy -= g.offsetY
			g.panel.Clear()

			clickedObject := false
			for _, p := range g.level.Planets() {
				if p.MouseButton(cx, cy) {
					clickedObject = true
					p.Highlight()
					showPlanetPanel(g.panel, p, g.resourceData, g.level.AllowedResources())
					showBuildOptionsPanel(p, g)
				} else {
					p.Unhighlight()
				}
			}

			for _, s := range g.player.Structures() {
				if s.MouseButton(cx, cy) {
					clickedObject = true
					s.Highlight()
					showStructurePanel(g, s)
				} else {
					s.Unhighlight()
				}
			}

			if !clickedObject {
				g.dragging = true
				g.mouseDragX, g.mouseDragY = ebiten.CursorPosition()
			}
		}
	} else if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		g.redrawPSLayer = true
		g.dragging = false
		g.panel.LeftMouseButtonRelease(ebiten.CursorPosition())
	} else if g.dragging && ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		g.redrawPSLayer = true
		// FIXME continuing to drag mouse after you reach the screen edge moves
		// the cursor position but not the viewport, so planet selection targets
		// are wrong
		newX, newY := ebiten.CursorPosition()

		var xMovement, yMovement int

		if newX < g.mouseDragX && -1*g.offsetX+g.windowW < g.level.W {
			xMovement = newX - g.mouseDragX
			g.offsetX += xMovement
			g.mouseDragX = newX
		}

		if newX > g.mouseDragX && g.offsetX < 0 {
			xMovement = newX - g.mouseDragX
			g.offsetX += xMovement
			g.mouseDragX = newX
		}

		if newY > g.mouseDragY && g.offsetY < 0 {
			yMovement = newY - g.mouseDragY
			g.offsetY += yMovement
			g.mouseDragY = newY
		}
		if newY < g.mouseDragY && -1*g.offsetY+g.windowH < g.level.H {
			yMovement = newY - g.mouseDragY
			g.offsetY += yMovement
			g.mouseDragY = newY
		}

	}
}
