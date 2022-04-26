package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/jcgraybill/ship-shape/ui"
)

func (g *Game) Draw(screen *ebiten.Image) {
	g.universe.DrawImage(g.bg, nil)
	ui.ShaderOpts.Uniforms["Light"] = []float32{-float32(g.offsetX), -float32(g.offsetY)}

	for _, s := range g.player.Ships() {
		s.DrawCourse(g.universe)
	}

	for _, s := range g.player.Structures() {
		s.Draw(g.universe)
	}

	for _, p := range g.level.Planets() {
		p.Draw(g.universe)
	}

	for _, s := range g.player.Ships() {
		s.Draw(g.universe)
	}
	screen.DrawImage(g.universe, g.opts)
	g.panel.Draw(screen)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS %f", ebiten.CurrentTPS()))
}
