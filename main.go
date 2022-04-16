package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/jcgraybill/ship-shape/panel"
	"github.com/jcgraybill/ship-shape/planet"
	"github.com/jcgraybill/ship-shape/structure"
	"github.com/jcgraybill/ship-shape/util"
	"golang.org/x/image/font"
)

const (
	w = 800
	h = 480
)

type Game struct {
	bg            *ebiten.Image
	ttf           font.Face
	planets       []*planet.Planet
	structures    []*structure.Structure
	panel         *panel.Panel
	structureData map[string]structure.StructureData
}

func (g *Game) Update() error {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		if !g.panel.MouseButton(ebiten.CursorPosition()) {
			g.panel.Clear()

			for _, planet := range g.planets {
				if planet.MouseButton(ebiten.CursorPosition()) {
					planet.Highlight()
					g.panel.AddLabel(planet.Describe())
					// FIXME - by the time this fires, planet is always a pointer to the last element in the slice.
					// What I want is for each function to be a pointer to its specific planet
					bh := func() { fmt.Println("build habitat on ", planet.Name()) }
					bdp := func() { fmt.Println("build desalinization plant on ", planet.Name()) }
					g.panel.AddButton("build habitat", bh)
					g.panel.AddButton("build desalinization plant", bdp)
				} else {
					planet.Unhighlight()
				}
			}

			for _, structure := range g.structures {
				if structure.MouseButton(ebiten.CursorPosition()) {
					g.panel.AddLabel(structure.Describe())
				}
			}

		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(g.bg, nil)
	for _, p := range g.planets {
		p.Draw(screen)
	}

	for _, s := range g.structures {
		s.Draw(screen)
	}
	screen.DrawImage(g.panel.Image(), g.panel.Location())
}

func init() {
	rand.Seed(time.Now().UnixNano())
	audio.NewContext(48000)
}

func main() {
	ebiten.SetWindowSize(w, h)
	ebiten.SetWindowTitle("ship shape")

	planets := make([]*planet.Planet, 3)
	planets[0] = planet.New(100, 100, 255, 1)
	planets[1] = planet.New(300, 300, 1, 128)
	planets[2] = planet.New(500, 200, 128, 128)
	panel := panel.New(w, h)

	sd, err := structure.GetStructureData()
	if err != nil {
		panic(err)
	}

	structures := make([]*structure.Structure, 1)
	structures[0] = structure.New(sd["admin"], planets[2])

	g := Game{
		bg:            util.StarField(w, h),
		ttf:           util.Font(),
		planets:       planets,
		panel:         panel,
		structureData: sd,
		structures:    structures,
	}

	if err := ebiten.RunGame(&g); err != nil {
		panic(err)
	}
}

func (g *Game) Layout(int, int) (int, int) {
	return w, h
}
