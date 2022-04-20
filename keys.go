package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jcgraybill/ship-shape/ui"
)

func handleKeyPresses(g *Game) {

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
