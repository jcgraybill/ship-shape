package main

import (
	"bytes"
	"fmt"
	"log"
	"math/rand"
	"os"
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
	count          int
	year           uint
	level          *level.Level
	player         *player.Player
	bg             *ebiten.Image
	universe       *ebiten.Image
	redrawUniverse bool

	panel                              *panel.Panel
	endOfLevelPlayerPanel              bool
	structureData                      *[structure.StructureDataLength]structure.StructureData
	resourceData                       *[resource.ResourceDataLength]resource.ResourceData
	opts                               *ebiten.DrawImageOptions
	offsetX, offsetY, windowW, windowH int
	mouseDragX, mouseDragY             int
	dragging                           bool
}

var (
	InfoLogger *log.Logger
	ambient    *audio.Player
)

func init() {
	rand.Seed(time.Now().UnixNano())
	audio.NewContext(44100)

	file, err := os.OpenFile("info.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
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
	playAmbientAudio()

	if err := ebiten.RunGame(&g); err != nil {
		panic(err)
	}
}

func (g *Game) load(lvl *level.Level) {
	g.count = 0
	g.player = player.New()
	g.level = level.New(lvl)
	g.bg = ui.StarField(lvl.W, lvl.H)
	g.universe = ebiten.NewImage(lvl.W, lvl.H)
	g.redrawUniverse = true
	g.year = g.level.StartingYear()
	g.player.AddMoney(lvl.StartingMoney())
	g.endOfLevelPlayerPanel = false

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
	}
	return w, h
}
