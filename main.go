package main

import (
	"fmt"
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
	endOfLevelPlayerPanel              bool
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
	g := Game{
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

	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	g.load(level.StartingLevel())
	ui.InitializeShader()
	if err := ebiten.RunGame(&g); err != nil {
		panic(err)
	}
}

func (g *Game) load(lvl *level.Level) {
	g.count = 0
	g.level = lvl
	g.bg = ui.StarField(lvl.W, lvl.H)
	g.universe = ebiten.NewImage(lvl.W, lvl.H)
	g.year = g.level.StartingYear()
	g.player = player.New()
	g.player.AddMoney(lvl.StartingMoney())
	g.endOfLevelPlayerPanel = false

	ebiten.SetWindowTitle(fmt.Sprintf("%s: %s", ui.NameofGame, lvl.Title()))
	g.panel.Lock(0)
	g.panel.Clear()
	g.panel.Lock(showPlayerPanel(g))
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
