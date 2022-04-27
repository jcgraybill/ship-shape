package main

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jcgraybill/ship-shape/ui"
)

func (g *Game) Draw(screen *ebiten.Image) {

	viewport := image.Rect(-g.offsetX, -g.offsetY, -g.offsetX+g.windowW, -g.offsetY+g.windowH)

	screen.DrawImage(g.starfieldLayer.SubImage(viewport).(*ebiten.Image), nil)

	g.trailsLayer.Clear()
	for _, s := range g.player.Ships() {
		if viewport.Overlaps(s.Bounds) {
			s.DrawCourse(g.trailsLayer)
		}
	}
	screen.DrawImage(g.trailsLayer.SubImage(viewport).(*ebiten.Image), nil)

	if g.redrawPSLayer {
		g.planetsAndStructuresLayer.Clear()
		ui.ShaderOpts.Uniforms["Light"] = []float32{-float32(g.offsetX), -float32(g.offsetY)}

		for _, s := range g.player.Structures() {
			if viewport.Overlaps(s.Bounds) {
				s.Draw(g.planetsAndStructuresLayer)
			}
		}

		for _, p := range g.level.Planets() {
			if viewport.Overlaps(p.Bounds) {
				p.Draw(g.planetsAndStructuresLayer)
			}
		}
		g.redrawPSLayer = false
	}
	screen.DrawImage(g.planetsAndStructuresLayer.SubImage(viewport).(*ebiten.Image), nil)

	g.shipsLayer.Clear()
	for _, s := range g.player.Ships() {
		if viewport.Overlaps(s.Bounds) {
			s.Draw(g.shipsLayer)
		}
	}
	screen.DrawImage(g.shipsLayer.SubImage(viewport).(*ebiten.Image), nil)

	g.panel.Draw(screen)
}
