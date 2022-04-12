package main

import (
	"image/color"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/jcgraybill/ship-shape/panel"
	"github.com/jcgraybill/ship-shape/planet"
	"github.com/jcgraybill/ship-shape/util"
	"golang.org/x/image/font"
)

const (
	w = 800
	h = 480
)

type Game struct {
	bg      *ebiten.Image
	ttf     font.Face
	planets []*planet.Planet
	panel   *panel.Panel
}

func (g *Game) Update() error {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		if !g.panel.MouseButton(ebiten.CursorPosition()) {
			planetClicked := false
			for _, planet := range g.planets {
				if planet.In(ebiten.CursorPosition()) {
					planet.Highlight()
					g.panel.ShowPlanet(planet)
					planetClicked = true
				} else {
					planet.Unhighlight()
				}
				if !planetClicked {
					g.panel.Clear()
				}
			}
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(g.bg, nil)
	for _, p := range g.planets {
		screen.DrawImage(p.Image(), p.Location())
		cx, cy := p.Center()
		textBounds := text.BoundString(g.ttf, p.Name())
		text.Draw(screen, p.Name(), g.ttf, cx-textBounds.Dx()/2, cy-16, color.White)
	}
	screen.DrawImage(g.panel.Image(), g.panel.Location())
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	ebiten.SetWindowSize(w, h)
	ebiten.SetWindowTitle("ship shape")

	planets := make([]*planet.Planet, 2)
	planets[0] = planet.New(100, 100, 255, 1)
	planets[1] = planet.New(300, 300, 1, 128)

	panel := panel.New(w, h)

	g := Game{
		bg:      util.StarField(w, h),
		ttf:     util.Font(),
		planets: planets,
		panel:   panel,
	}

	if err := ebiten.RunGame(&g); err != nil {
		panic(err)
	}
}

func (g *Game) Layout(int, int) (int, int) {
	return w, h
}
