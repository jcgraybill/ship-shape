package main

import "github.com/hajimehoshi/ebiten/v2"

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(g.bg, nil)
	for _, p := range g.planets {
		p.Draw(screen)
	}

	for _, s := range g.structures {
		s.Draw(screen)
	}

	g.panel.Draw(screen)
}
