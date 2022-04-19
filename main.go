package main

import (
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/jcgraybill/ship-shape/panel"
	"github.com/jcgraybill/ship-shape/planet"
	"github.com/jcgraybill/ship-shape/resource"
	"github.com/jcgraybill/ship-shape/ship"
	"github.com/jcgraybill/ship-shape/structure"
	"github.com/jcgraybill/ship-shape/ui"
	"golang.org/x/image/font"
)

type Game struct {
	count         int
	bg            *ebiten.Image
	ttf           font.Face
	planets       []*planet.Planet
	structures    []*structure.Structure
	panel         *panel.Panel
	structureData [structure.StructureDataLength]structure.StructureData
	resourceData  [resource.ResourceDataLength]resource.ResourceData

	ship *ship.Ship
}

func init() {
	rand.Seed(time.Now().UnixNano())
	audio.NewContext(48000)
}

func main() {
	ebiten.SetWindowSize(ui.W, ui.H)
	ebiten.SetWindowTitle("ship shape")

	rd := resource.GetResourceData()

	planets := make([]*planet.Planet, 4)
	planets[0] = planet.New(100, 100, map[int]uint8{resource.Ice: 196}, rd)
	planets[1] = planet.New(300, 300, map[int]uint8{resource.Habitability: 200}, rd)
	planets[2] = planet.New(500, 140, map[int]uint8{resource.Ice: 128, resource.Habitability: 128, resource.Iron: 32}, rd)
	planets[3] = planet.New(550, 250, map[int]uint8{resource.Habitability: 60, resource.Ice: 196}, rd)
	panel := panel.New(ui.W, ui.H)

	sd := structure.GetStructureData()

	structures := make([]*structure.Structure, 1)
	structures[0] = structure.New(sd[structure.Home], planets[2])

	g := Game{
		count:         0,
		bg:            ui.StarField(ui.W, ui.H),
		ttf:           ui.Font(),
		planets:       planets,
		panel:         panel,
		structureData: sd,
		resourceData:  rd,
		structures:    structures,
		ship:          ship.New(320, 240),
	}

	if err := ebiten.RunGame(&g); err != nil {
		panic(err)
	}
}

func (g *Game) Layout(int, int) (int, int) {
	return ui.W, ui.H
}
