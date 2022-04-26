package ship

import (
	"github.com/fogleman/gg"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jcgraybill/ship-shape/ui"
)

func (s *Ship) Draw(image *ebiten.Image) {
	if s.plumeVisible {
		image.DrawImage(s.plume, s.opts)
	} else {
		image.DrawImage(s.image, s.opts)
	}
}

func (s *Ship) createBaseCourseLine() {
	x0, y0 := s.origin.Planet().Center()
	x1, y1 := s.destination.Planet().Center()
	dc := gg.NewContext(s.course.Bounds().Dx(), s.course.Bounds().Dy())
	dc.SetRGB255(int(ui.NonFocusColor.R), int(ui.NonFocusColor.G), int(ui.NonFocusColor.B))
	dc.DrawLine(float64(x0)-float64(s.baseX), float64(y0)-float64(s.baseY), float64(x1)-float64(s.baseX), float64(y1)-float64(s.baseY))
	dc.SetLineWidth(ui.Border)
	dc.Stroke()
	s.baseCourse = ebiten.NewImageFromImage(dc.Image())
}

func (s *Ship) updateCourseLine() {

	x1, y1 := s.destination.Planet().Center()

	s.course.DrawImage(s.baseCourse, nil)
	dc := gg.NewContext(s.course.Bounds().Dx(), s.course.Bounds().Dy())
	dc.SetRGB255(int(ui.FocusedColor.R), int(ui.FocusedColor.G), int(ui.FocusedColor.B))
	dc.DrawLine(float64(s.x)-float64(s.baseX), float64(s.y)-float64(s.baseY), float64(x1)-float64(s.baseX), float64(y1)-float64(s.baseY))
	dc.SetLineWidth(ui.Border)
	dc.Stroke()
	s.course.DrawImage(ebiten.NewImageFromImage(dc.Image()), nil)
}

func (s *Ship) DrawCourse(image *ebiten.Image) {
	s.updateCourseLine()
	image.DrawImage(s.course, s.courseOpts)
}
