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
	count                              int
	bg                                 *ebiten.Image
	universe                           *ebiten.Image
	ttf                                font.Face
	planets                            []*planet.Planet
	structures                         []*structure.Structure
	ships                              map[int]*ship.Ship
	panel                              *panel.Panel
	structureData                      [structure.StructureDataLength]structure.StructureData
	resourceData                       [resource.ResourceDataLength]resource.ResourceData
	opts                               *ebiten.DrawImageOptions
	offsetX, offsetY, windowW, windowH int
	mouseDragX, mouseDragY             int
	dragging                           bool
	pop, maxPop, workersNeeded         int
	money                              int
	capitols                           int
}

func init() {
	rand.Seed(time.Now().UnixNano())
	audio.NewContext(48000)
}

func main() {

	g := Game{
		count:         0,
		bg:            ui.StarField(ui.W, ui.H),
		universe:      ebiten.NewImage(ui.W, ui.H),
		opts:          &ebiten.DrawImageOptions{},
		planets:       generatePlanets(ui.W, ui.H),
		panel:         panel.New(ui.WindowW, ui.WindowH),
		structureData: structure.GetStructureData(),
		resourceData:  resource.GetResourceData(),
		structures:    make([]*structure.Structure, 0),
		ships:         make(map[int]*ship.Ship),
		offsetX:       0,
		offsetY:       0,
		windowW:       ui.WindowW,
		windowH:       ui.WindowH,
		money:         ui.StartingMoney,
		capitols:      0,
	}

	ebiten.SetWindowSize(ui.WindowW, ui.WindowH)
	ebiten.SetWindowTitle("ship shape")
	ebiten.SetWindowResizable(true)
	g.panel.Lock(showPlayerPanel(&g))

	if err := ebiten.RunGame(&g); err != nil {
		panic(err)
	}
}

func (g *Game) Layout(w, h int) (int, int) {
	if w < ui.WindowW {
		w = ui.WindowW
	}
	if h < ui.WindowH {
		h = ui.WindowH
	}
	g.panel.Resize(w, h)
	g.windowW = w
	g.windowH = h
	return w, h
}

// TODO avoid placing planets underneath the panel

func generatePlanets(w, h int) []*planet.Planet {
	cellsize := ui.PlanetSize * ui.PlanetDistance
	planets := make([]*planet.Planet, 0)
	rd := resource.GetResourceData()

	for i := 0; i < h/cellsize; i++ {
		for j := 0; j < w/cellsize; j++ {
			x := j*cellsize + rand.Intn(cellsize-ui.PlanetSize*2) + ui.PlanetSize
			y := i*cellsize + rand.Intn(cellsize-ui.PlanetSize*2) + ui.PlanetSize
			ice := uint8(rand.Intn(255))
			habitability := uint8(rand.Intn(255))
			planets = append(planets, planet.New(x, y, map[int]uint8{resource.Ice: ice, resource.Habitability: habitability}, rd))
		}
	}
	return planets
}
