package ship

import (
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
