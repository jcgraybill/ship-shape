package main

import (
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/jcgraybill/ship-shape/panel"
	"github.com/jcgraybill/ship-shape/planet"
	"github.com/jcgraybill/ship-shape/structure"
	"github.com/jcgraybill/ship-shape/ui"
	"golang.org/x/image/font"
)

type Game struct {
	count         int32
	bg            *ebiten.Image
	ttf           font.Face
	planets       []*planet.Planet
	structures    []*structure.Structure
	panel         *panel.Panel
	structureData map[string]structure.StructureData
}

func init() {
	rand.Seed(time.Now().UnixNano())
	audio.NewContext(48000)
}

func main() {
	ebiten.SetWindowSize(ui.W, ui.H)
	ebiten.SetWindowTitle("ship shape")

	planets := make([]*planet.Planet, 3)
	planets[0] = planet.New(100, 100, 255, 1)
	planets[1] = planet.New(300, 300, 1, 128)
	planets[2] = planet.New(500, 200, 128, 128)
	panel := panel.New(ui.W, ui.H)

	sd, err := structure.GetStructureData()
	if err != nil {
		panic(err)
	}

	structures := make([]*structure.Structure, 1)
	structures[0] = structure.New(sd["home"], planets[2])

	g := Game{
		count:         0,
		bg:            ui.StarField(ui.W, ui.H),
		ttf:           ui.Font(),
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
	return ui.W, ui.H
}
