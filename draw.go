package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jcgraybill/ship-shape/ui"
)

func (g *Game) Draw(screen *ebiten.Image) {

	screen.DrawImage(g.starfieldLayer, g.opts)

	g.trailsLayer.Clear()
	for _, s := range g.player.Ships() {
		s.DrawCourse(g.trailsLayer)
	}
	screen.DrawImage(g.trailsLayer, g.opts)

	if g.redrawPSLayer {
		g.planetsAndStructuresLayer.Clear()
		ui.ShaderOpts.Uniforms["Light"] = []float32{-float32(g.offsetX), -float32(g.offsetY)}

		for _, s := range g.player.Structures() {
			s.Draw(g.planetsAndStructuresLayer)
		}

		for _, p := range g.level.Planets() {
			p.Draw(g.planetsAndStructuresLayer)
		}
		g.redrawPSLayer = false
	}
	screen.DrawImage(g.planetsAndStructuresLayer, g.opts)

	g.shipsLayer.Clear()
	for _, s := range g.player.Ships() {
		s.Draw(g.shipsLayer)
	}
	screen.DrawImage(g.shipsLayer, g.opts)

	g.panel.Draw(screen)
}
