package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	w = 640
	h = 480
)

type Game struct {
	bg *ebiten.Image
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(g.bg, nil)
}

func main() {
	ebiten.SetWindowSize(w, h)
	ebiten.SetWindowTitle("ship shape")

	image := ebiten.NewImage(w, h)
	image.Fill(color.Black)

	g := Game{
		bg: image,
	}

	if err := ebiten.RunGame(&g); err != nil {
		panic(err)
	}
}

func (g *Game) Layout(int, int) (int, int) {
	return w, h
}
