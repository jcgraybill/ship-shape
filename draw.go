package main

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jcgraybill/ship-shape/ui"
)

var dt [5]uint64
var dm [5]uint64

func (g *Game) Draw(screen *ebiten.Image) {

	g.viewport = image.Rect(-g.offsetX, -g.offsetY, -g.offsetX+g.windowW, -g.offsetY+g.windowH)

	screen.DrawImage(g.starfieldLayer.SubImage(g.viewport).(*ebiten.Image), nil)

	dt[0], dm[0] = g.measure(g.drawTrails)
	screen.DrawImage(g.trailsLayer.SubImage(g.viewport).(*ebiten.Image), nil)

	if g.redrawPSLayer {
		g.planetsAndStructuresLayer.Clear()
		ui.ShaderOpts.Uniforms["Light"] = []float32{-float32(g.offsetX), -float32(g.offsetY)}
		dt[1], dm[1] = g.measure(g.drawStructures)
		dt[2], dm[2] = g.measure(g.drawPlanets)
		g.redrawPSLayer = false
	} else {
		ut[1], um[1], ut[2], um[2] = 0, 0, 0, 0
	}
	screen.DrawImage(g.planetsAndStructuresLayer.SubImage(g.viewport).(*ebiten.Image), nil)

	dt[3], dm[3] = g.measure(g.drawShips)
	screen.DrawImage(g.shipsLayer.SubImage(g.viewport).(*ebiten.Image), nil)

	g.panelLayer.Clear()
	dt[4], dm[4] = g.measure(g.drawPanel)
	screen.DrawImage(g.panelLayer, nil)
	UpdateLogger.Printf("draw tick %d time %v", g.count, dt)
	UpdateLogger.Printf("draw tick %d mem %v", g.count, dm)

}

func (g *Game) drawTrails() {
	g.trailsLayer.Clear()
	for _, s := range g.player.Ships() {
		if g.viewport.Overlaps(s.Bounds) {
			s.DrawCourse(g.trailsLayer)
		}
	}
}

func (g *Game) drawStructures() {
	for _, s := range g.player.Structures() {
		if g.viewport.Overlaps(s.Bounds) {
			s.Draw(g.planetsAndStructuresLayer)
		}
	}
}

func (g *Game) drawPlanets() {
	for _, p := range g.level.Planets() {
		if g.viewport.Overlaps(p.Bounds) {
			p.Draw(g.planetsAndStructuresLayer)
		}
	}
}
func (g *Game) drawShips() {
	g.shipsLayer.Clear()
	for _, s := range g.player.Ships() {
		if g.viewport.Overlaps(s.Bounds) {
			s.Draw(g.shipsLayer)
		}
	}
}
func (g *Game) drawPanel() {
	g.panel.Draw(g.panelLayer)
}
