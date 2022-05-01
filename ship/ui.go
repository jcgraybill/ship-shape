package ship

import (
	"image"

	"github.com/fogleman/gg"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jcgraybill/ship-shape/ui"
)

func (s *Ship) Draw(targetImage *ebiten.Image) {
	if s.image == nil || s.plume == nil {
		s.createShipImages()
	}
	if s.plumeVisible {
		targetImage.DrawImage(s.plume, s.opts)
	} else {
		targetImage.DrawImage(s.image, s.opts)
	}
}

func (s *Ship) DrawCourse(targetImage *ebiten.Image) {
	if s.baseCourse == nil || s.course == nil {
		s.createCourseImages()
	}

	x0, y0 := s.origin.Planet().Center()
	x1, y1 := s.destination.Planet().Center()

	trailOpts := &ebiten.DrawImageOptions{}
	trailRect := image.Rect(x0-s.Bounds.Min.X, y0-s.Bounds.Min.Y, int(s.x)-s.Bounds.Min.X, int(s.y)-s.Bounds.Min.Y)
	trailOpts.GeoM.Translate(float64(s.Bounds.Min.X+trailRect.Min.X), float64(s.Bounds.Min.Y+trailRect.Min.Y))
	targetImage.DrawImage(s.baseCourse.SubImage(trailRect).(*ebiten.Image), trailOpts)

	headingOpts := &ebiten.DrawImageOptions{}
	headingRect := image.Rect(int(s.x)-s.Bounds.Min.X, int(s.y)-s.Bounds.Min.Y, x1-s.Bounds.Min.X, y1-s.Bounds.Min.Y)
	headingOpts.GeoM.Translate(float64(s.Bounds.Min.X+headingRect.Min.X), float64(s.Bounds.Min.Y+headingRect.Min.Y))
	targetImage.DrawImage(s.course.SubImage(headingRect).(*ebiten.Image), headingOpts)
}

func (s *Ship) createShipImages() {
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

	if s.cargo > 0 {
		v, i := ui.Triangle(plumeW+2, 2, shipW-6, shipH-4, s.cargoColor)
		s.image.DrawTriangles(v, i, ui.Src, nil)
		s.plume.DrawTriangles(v, i, ui.Src, nil)
	}
}

func (s *Ship) createCourseImages() {
	var baseX, baseY int
	x0, y0 := s.origin.Planet().Center()
	x1, y1 := s.destination.Planet().Center()

	if x0 > x1 {
		baseX = x1
	} else if x1 > x0 {
		baseX = x0
	} else {
		baseX = x0
	}
	if y0 > y1 {
		baseY = y1
	} else if y1 > y0 {
		baseY = y0
	} else {
		baseY = y0
	}

	dc := gg.NewContext(s.Bounds.Dx(), s.Bounds.Dy())
	dc.SetRGB255(int(ui.NonFocusColor.R), int(ui.NonFocusColor.G), int(ui.NonFocusColor.B))
	dc.DrawLine(float64(x0)-float64(baseX), float64(y0)-float64(baseY), float64(x1)-float64(baseX), float64(y1)-float64(baseY))
	dc.SetLineWidth(ui.Border)
	dc.Stroke()
	s.baseCourse = ebiten.NewImageFromImage(dc.Image())

	dc = gg.NewContext(s.Bounds.Dx(), s.Bounds.Dy())
	dc.SetRGB255(int(ui.FocusedColor.R), int(ui.FocusedColor.G), int(ui.FocusedColor.B))
	dc.DrawLine(float64(x0)-float64(baseX), float64(y0)-float64(baseY), float64(x1)-float64(baseX), float64(y1)-float64(baseY))
	dc.SetLineWidth(ui.Border)
	dc.Stroke()
	s.course = ebiten.NewImageFromImage(dc.Image())

}
