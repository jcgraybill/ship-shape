package main

import (
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/jcgraybill/ship-shape/level"
	"github.com/jcgraybill/ship-shape/panel"
	"github.com/jcgraybill/ship-shape/resource"
	"github.com/jcgraybill/ship-shape/ship"
	"github.com/jcgraybill/ship-shape/structure"
	"github.com/jcgraybill/ship-shape/ui"
)

type Game struct {
	count                              int
	level                              *level.Level
	bg                                 *ebiten.Image
	universe                           *ebiten.Image
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

	lvl := level.New(1)

	g := Game{
		level:         lvl,
		count:         0,
		bg:            ui.StarField(lvl.W, lvl.H),
		universe:      ebiten.NewImage(lvl.W, lvl.H),
		opts:          &ebiten.DrawImageOptions{},
		panel:         panel.New(ui.WindowW, ui.WindowH),
		structureData: structure.GetStructureData(),
		resourceData:  resource.GetResourceData(),
		structures:    make([]*structure.Structure, 0),
		ships:         make(map[int]*ship.Ship),
		offsetX:       0,
		offsetY:       0,
		windowW:       ui.WindowW,
		windowH:       ui.WindowH,
		money:         lvl.StartingMoney,
		capitols:      0,
	}

	ebiten.SetWindowSize(ui.WindowW, ui.WindowH)
	ebiten.SetWindowTitle(ui.NameofGame)
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
