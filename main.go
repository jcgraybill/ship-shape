package main

import (
	"image/color"
	"math/rand"
	"os"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/jcgraybill/ship-shape/planet"
	"github.com/jcgraybill/ship-shape/util"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

const (
	w          = 800
	h          = 480
	starriness = 3000
	ttf        = "fonts/OpenSans_SemiCondensed-Regular.ttf"
)

type Game struct {
	bg      *ebiten.Image
	ttf     font.Face
	planets []*planet.Planet
}

func (g *Game) Update() error {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		for _, planet := range g.planets {
			if foo := planet.In(ebiten.CursorPosition()); foo {
				planet.Highlight()
			} else {
				planet.Unhighlight()
			}
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(g.bg, nil)
	for _, p := range g.planets {
		opts := &ebiten.DrawImageOptions{}
		x, y := p.Location()
		opts.GeoM.Translate(x, y)
		screen.DrawImage(p.Image(), opts)
		cx, _ := p.Center()
		textBounds := text.BoundString(g.ttf, p.Name())

		text.Draw(screen, p.Name(), g.ttf, cx-textBounds.Dx()/2, int(y), color.White)
	}
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

	g := Game{
		bg:      util.StarField(w, h, starriness),
		ttf:     loadFont(ttf),
		planets: planets,
	}

	if err := ebiten.RunGame(&g); err != nil {
		panic(err)
	}
}

func (g *Game) Layout(int, int) (int, int) {
	return w, h
}

func loadFont(path string) font.Face {
	ttbytes, err := os.ReadFile(path)
	if err == nil {
		tt, err := opentype.Parse(ttbytes)
		if err == nil {
			fontface, err := opentype.NewFace(tt, &opentype.FaceOptions{
				Size:    12,
				DPI:     72,
				Hinting: font.HintingFull,
			})
			if err == nil {
				return fontface
			}
			panic(err)
		}
		panic(err)
	}
	panic(err)
}
