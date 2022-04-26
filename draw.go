package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jcgraybill/ship-shape/ui"
)

func (g *Game) Draw(screen *ebiten.Image) {

	screen.DrawImage(g.bg, g.opts)

	for _, s := range g.player.Ships() {
		s.DrawCourse(screen)
	}

	if g.redrawUniverse {
		g.universe.Clear()
		ui.ShaderOpts.Uniforms["Light"] = []float32{-float32(g.offsetX), -float32(g.offsetY)}

		for _, s := range g.player.Structures() {
			s.Draw(g.universe)
		}

		for _, p := range g.level.Planets() {
			p.Draw(g.universe)
		}
		g.redrawUniverse = false
	}

	screen.DrawImage(g.universe, g.opts)

	for _, s := range g.player.Ships() {
		s.Draw(screen)
	}

	g.panel.Draw(screen)
}
