package main

import (
	"bytes"
	"fmt"
	"image"
	"math/rand"
	"runtime"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
	"github.com/jcgraybill/ship-shape/level"
	"github.com/jcgraybill/ship-shape/panel"
	"github.com/jcgraybill/ship-shape/player"
	"github.com/jcgraybill/ship-shape/resource"
	"github.com/jcgraybill/ship-shape/structure"
	"github.com/jcgraybill/ship-shape/ui"
)

type Game struct {
	count                     uint
	level                     *level.Level
	player                    *player.Player
	trailsLayer               *ebiten.Image
	starfieldLayer            *ebiten.Image
	planetsAndStructuresLayer *ebiten.Image
	shipsLayer                *ebiten.Image
	panelLayer                *ebiten.Image
	viewport                  image.Rectangle
	redrawPSLayer             bool

	panel                              *panel.Panel
	endOfLevelPlayerPanel              bool
	structureData                      *[structure.StructureDataLength]structure.StructureData
	resourceData                       *[resource.ResourceDataLength]resource.ResourceData
	offsetX, offsetY, windowW, windowH int
	mouseDragX, mouseDragY             int
	dragging                           bool
}

var (
	ambient *audio.Player
)

func init() {
	rand.Seed(time.Now().UnixNano())
	audio.NewContext(44100)
}

func main() {
	g := Game{
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
	playAmbientAudio()

	if err := ebiten.RunGame(&g); err != nil {
		panic(err)
	}
}

func (g *Game) load(lvl *level.Level) {
	g.count = 0
	g.player = player.New()
	g.level = level.New(lvl)
	g.starfieldLayer = ui.StarField(lvl.W, lvl.H)
	g.trailsLayer = ebiten.NewImage(lvl.W, lvl.H)
	g.planetsAndStructuresLayer = ebiten.NewImage(lvl.W, lvl.H)
	g.shipsLayer = ebiten.NewImage(lvl.W, lvl.H)
	g.panelLayer = ebiten.NewImage(ui.WindowW, ui.WindowH)
	g.player.AddMoney(lvl.StartingMoney())
	g.endOfLevelPlayerPanel = false
	g.redrawPSLayer = true

	ebiten.SetWindowTitle(fmt.Sprintf("%s: %s", ui.NameofGame, lvl.Title()))
	g.panel.Lock(0)
	g.panel.Clear()
	g.panel.Lock(showPlayerPanel(g))
	runtime.GC()
}

func playAmbientAudio() {
	audioContext := audio.CurrentContext()
	audioBytes, err := ui.GameData("audio/ambient.ogg")
	if err == nil {
		d, err := vorbis.Decode(audioContext, bytes.NewReader(audioBytes))
		if err == nil {
			s := audio.NewInfiniteLoop(d, d.Length())
			ambient, err = audioContext.NewPlayer(s)
			if err == nil {
				ambient.Play()
			} else {
				panic(err)
			}
		} else {
			panic(err)
		}
	} else {
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
	if g.windowW != w && g.windowH != h {
		g.panel.Resize(w, h)
		g.windowW = w
		g.windowH = h
		g.panelLayer = ebiten.NewImage(w, h)
		g.redrawPSLayer = true
	}
	return w, h
}
