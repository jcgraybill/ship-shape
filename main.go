package main

import (
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/jcgraybill/ship-shape/level"
	"github.com/jcgraybill/ship-shape/panel"
	"github.com/jcgraybill/ship-shape/player"
	"github.com/jcgraybill/ship-shape/resource"
	"github.com/jcgraybill/ship-shape/structure"
	"github.com/jcgraybill/ship-shape/ui"
)

type Game struct {
	count    int
	year     uint
	level    *level.Level
	player   *player.Player
	bg       *ebiten.Image
	universe *ebiten.Image

	panel                              *panel.Panel
	structureData                      [structure.StructureDataLength]structure.StructureData
	resourceData                       [resource.ResourceDataLength]resource.ResourceData
	opts                               *ebiten.DrawImageOptions
	offsetX, offsetY, windowW, windowH int
	mouseDragX, mouseDragY             int
	dragging                           bool
}

func init() {
	rand.Seed(time.Now().UnixNano())
	audio.NewContext(48000)
}

func main() {

	lvl := level.New(1)

	g := Game{
		level:         lvl,
		player:        player.New(),
		count:         0,
		bg:            ui.StarField(lvl.W, lvl.H),
		universe:      ebiten.NewImage(lvl.W, lvl.H),
		opts:          &ebiten.DrawImageOptions{},
		panel:         panel.New(ui.WindowW, ui.WindowH),
		structureData: structure.GetStructureData(),
		resourceData:  resource.GetResourceData(),

		offsetX: 0,
		offsetY: 0,
		windowW: ui.WindowW,
		windowH: ui.WindowH,
	}

	ebiten.SetWindowSize(ui.WindowW, ui.WindowH)
	ebiten.SetWindowTitle(ui.NameofGame)
	ebiten.SetWindowResizable(true)
	g.player.AddMoney(g.level.StartingMoney())
	g.year = g.level.StartingYear()
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
