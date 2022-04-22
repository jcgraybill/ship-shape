package ship

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jcgraybill/ship-shape/structure"
	"github.com/jcgraybill/ship-shape/ui"
)

const (
	Cargo int = iota
	Income
)

const shipW = 16
const plumeW = 8
const shipH = 10

var shipColor = color.RGBA{0xcc, 0xcc, 0xcc, 0xff}       // silver
var incomeShipColor = color.RGBA{0xd4, 0xaf, 0x47, 0xff} // gold
var plumeOuter = color.RGBA{0xff, 0xa5, 0x00, 0xff}      //orange
var plumeInner = color.RGBA{0xff, 0xff, 0x00, 0xff}      // yellow
const plumeCycleTime = 20
const plumeFrequency = 4

type Ship struct {
	x, y, baseX, baseY float64
	dx, dy, theta      float64
	image              *ebiten.Image
	plume              *ebiten.Image
	plumeVisible       bool
	opts               *ebiten.DrawImageOptions

	origin      *structure.Structure
	destination *structure.Structure
	course      *ebiten.Image
	courseOpts  *ebiten.DrawImageOptions
	cargo       int //TODO - uint
	shipType    int
}

func New(origin, destination *structure.Structure, shipType int) *Ship {
	s := Ship{
		origin:       origin,
		destination:  destination,
		plumeVisible: true,
		cargo:        -1,
		shipType:     Cargo,
	}

	s.shipType = shipType

	x0, y0 := origin.Planet().Center()
	x1, y1 := destination.Planet().Center()
	s.x = float64(x0)
	s.y = float64(y0)

	var v []ebiten.Vertex
	var i []uint16

	if s.shipType == Cargo {
		v, i = ui.Triangle(plumeW, 0, shipW, shipH, shipColor)
	} else {
		v, i = ui.Triangle(plumeW, 0, shipW, shipH, incomeShipColor)
	}
	s.image = ebiten.NewImage(shipW+plumeW, shipH)
	s.image.DrawTriangles(v, i, ui.Src, nil)

	s.plume = ebiten.NewImage(shipW+plumeW, shipH)
	s.plume.DrawImage(s.image, nil)
	v, i = ui.Triangle(plumeW, 1, -plumeW, shipH-2, plumeOuter)
	s.plume.DrawTriangles(v, i, ui.Src, nil)

	v, i = ui.Triangle(plumeW, 4, -plumeW, shipH-8, plumeInner)
	s.plume.DrawTriangles(v, i, ui.Src, nil)

	var w, h int

	if x0 > x1 {
		w = x0 - x1
		s.baseX = float64(x1)
	} else if x1 > x0 {
		w = x1 - x0
		s.baseX = float64(x0)
	} else {
		w = ui.Border
		s.baseX = float64(x0)
	}
	if y0 > y1 {
		h = y0 - y1
		s.baseY = float64(y1)
	} else if y1 > y0 {
		h = y1 - y0
		s.baseY = float64(y0)
	} else {
		h = ui.Border
		s.baseY = float64(y0)
	}
	s.course = ebiten.NewImage(w, h)
	s.updateCourseLine()

	s.theta = math.Atan2(float64(y1-y0), float64(x1-x0))

	s.courseOpts = &ebiten.DrawImageOptions{}
	s.courseOpts.GeoM.Translate(s.baseX, s.baseY)

	s.opts = &ebiten.DrawImageOptions{}
	s.opts.GeoM.Translate(-(plumeW + shipW/2), -shipH/2)
	s.opts.GeoM.Rotate(s.theta)
	s.opts.GeoM.Translate(s.x, s.y)

	s.dx = math.Cos(s.theta)
	s.dy = math.Sin(s.theta)

	return &s
}
