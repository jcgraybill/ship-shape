package ship

import (
	"image/color"
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jcgraybill/ship-shape/structure"
	"github.com/jcgraybill/ship-shape/ui"
)

const shipW = 16
const plumeW = 8
const shipH = 10

var shipColor = color.RGBA{0xcc, 0xcc, 0xcc, 0xff}  // silver
var plumeOuter = color.RGBA{0xff, 0xa5, 0x00, 0xff} //orange
var plumeInner = color.RGBA{0xff, 0xff, 0x00, 0xff} // yellow
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
	cargo       int
}

func New(origin, destination *structure.Structure) *Ship {
	s := Ship{
		origin:       origin,
		destination:  destination,
		plumeVisible: true,
		cargo:        -1,
	}

	x0, y0 := origin.Planet().Center()
	x1, y1 := destination.Planet().Center()
	s.x = float64(x0)
	s.y = float64(y0)

	s.image = ebiten.NewImage(shipW+plumeW, shipH)
	v, i := ui.Triangle(plumeW, 0, shipW, shipH, shipColor)
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

func (s *Ship) Draw(image *ebiten.Image) {
	if s.plumeVisible {
		image.DrawImage(s.plume, s.opts)
	} else {
		image.DrawImage(s.image, s.opts)
	}
}

func (s *Ship) Update(count int) bool {
	if count%plumeCycleTime == 0 {
		if rand.Intn(plumeFrequency) == 0 {
			s.plumeVisible = false
		} else {
			s.plumeVisible = true
		}
	}

	if s.destination.Planet().In(int(s.x), int(s.y)) {
		return true
	} else {
		if count%ui.ShipSpeed == 0 {
			s.x += s.dx
			s.y += s.dy
			s.opts.GeoM.Translate(s.dx, s.dy)
			s.updateCourseLine()
		}
	}
	return false
}

func (s *Ship) updateCourseLine() {
	s.course.Clear()
	x0, y0 := s.origin.Planet().Center()
	x1, y1 := s.destination.Planet().Center()
	v, i := ui.Line(float32(s.x)-float32(s.baseX), float32(s.y)-float32(s.baseY), float32(x1)-float32(s.baseX), float32(y1)-float32(s.baseY), ui.Border, ui.FocusedColor)
	s.course.DrawTriangles(v, i, ui.Src, nil)
	v, i = ui.Line(float32(x0)-float32(s.baseX), float32(y0)-float32(s.baseY), float32(s.x)-float32(s.baseX), float32(s.y)-float32(s.baseY), ui.Border, ui.NonFocusColor)
	s.course.DrawTriangles(v, i, ui.Src, nil)

}

func (s *Ship) DrawCourse(image *ebiten.Image) {
	image.DrawImage(s.course, s.courseOpts)
}

//FIXME rotation causes this to not be the actual ship coordinates
func (s *Ship) MouseButton(x, y int) bool {
	if int(s.x)+plumeW < x && int(s.x)+shipW+plumeW > x {
		if int(s.y) < y && int(s.y)+shipH > y {
			return true
		}
	}
	return false
}

func (s *Ship) LoadCargo(resource int, cargoColor color.RGBA) {
	v, i := ui.Triangle(plumeW+2, 2, shipW-6, shipH-4, cargoColor)
	s.image.DrawTriangles(v, i, ui.Src, nil)
	s.plume.DrawTriangles(v, i, ui.Src, nil)
	s.cargo = resource
}

func (s *Ship) Manifest() (int, *structure.Structure, *structure.Structure) {
	return s.cargo, s.origin, s.destination
}
