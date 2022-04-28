package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jcgraybill/ship-shape/ui"
)

func (g *Game) handleKeyPresses() {

	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) && -1*g.offsetX+g.windowW < g.level.W {
		g.redrawPSLayer = true
		g.offsetX -= ui.ArrowKeyMoveSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) && g.offsetX < 0 {
		g.redrawPSLayer = true
		g.offsetX += ui.ArrowKeyMoveSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) && g.offsetY < 0 {
		g.redrawPSLayer = true
		g.offsetY += ui.ArrowKeyMoveSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) && -1*g.offsetY+g.windowH < g.level.H {
		g.redrawPSLayer = true
		g.offsetY -= ui.ArrowKeyMoveSpeed
	}
}
