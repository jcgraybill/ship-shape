package ship

import (
	"image"
	"image/color"
	"math"

	"github.com/fogleman/gg"
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
	x, y          float64
	dx, dy, theta float64
	image         *ebiten.Image
	plume         *ebiten.Image
	plumeVisible  bool
	opts          *ebiten.DrawImageOptions
	Bounds        image.Rectangle

	origin      *structure.Structure
	destination *structure.Structure

	baseCourse *ebiten.Image
	course     *ebiten.Image
	cargo      int
	shipType   int
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

	// TODO - simplify
	var baseX, baseY, w, h int

	if x0 > x1 {
		w = x0 - x1
		baseX = x1
	} else if x1 > x0 {
		w = x1 - x0
		baseX = x0
	} else {
		w = ui.Border
		baseX = x0
	}
	if y0 > y1 {
		h = y0 - y1
		baseY = y1
	} else if y1 > y0 {
		h = y1 - y0
		baseY = y0
	} else {
		h = ui.Border
		baseY = y0
	}

	s.Bounds = image.Rect(baseX, baseY, baseX+w, baseY+h)
	dc := gg.NewContext(w, h)
	dc.SetRGB255(int(ui.NonFocusColor.R), int(ui.NonFocusColor.G), int(ui.NonFocusColor.B))
	dc.DrawLine(float64(x0)-float64(baseX), float64(y0)-float64(baseY), float64(x1)-float64(baseX), float64(y1)-float64(baseY))
	dc.SetLineWidth(ui.Border)
	dc.Stroke()
	s.baseCourse = ebiten.NewImageFromImage(dc.Image())

	dc = gg.NewContext(w, h)
	dc.SetRGB255(int(ui.FocusedColor.R), int(ui.FocusedColor.G), int(ui.FocusedColor.B))
	dc.DrawLine(float64(x0)-float64(baseX), float64(y0)-float64(baseY), float64(x1)-float64(baseX), float64(y1)-float64(baseY))
	dc.SetLineWidth(ui.Border)
	dc.Stroke()
	s.course = ebiten.NewImageFromImage(dc.Image())

	s.theta = math.Atan2(float64(y1-y0), float64(x1-x0))

	s.opts = &ebiten.DrawImageOptions{}
	s.opts.GeoM.Translate(-(plumeW + shipW/2), -shipH/2)
	s.opts.GeoM.Rotate(s.theta)
	s.opts.GeoM.Translate(s.x, s.y)

	s.dx = math.Cos(s.theta)
	s.dy = math.Sin(s.theta)

	return &s
}
