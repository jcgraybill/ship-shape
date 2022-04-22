package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func (g *Game) Draw(screen *ebiten.Image) {
	g.universe.DrawImage(g.bg, nil)

	for _, s := range g.ships {
		s.DrawCourse(g.universe)
	}

	for _, p := range g.level.Planets() {
		p.Draw(g.universe)
	}

	for _, s := range g.structures {
		s.Draw(g.universe)
	}

	for _, s := range g.ships {
		s.Draw(g.universe)
	}
	screen.DrawImage(g.universe, g.opts)
	g.panel.Draw(screen)
}
