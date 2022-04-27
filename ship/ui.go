package ship

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jcgraybill/ship-shape/ui"
)

func (s *Ship) Draw(targetImage *ebiten.Image) {
	if s.plumeVisible {
		targetImage.DrawImage(s.plume, s.opts)
	} else {
		targetImage.DrawImage(s.image, s.opts)
	}
}

func (s *Ship) DrawCourse(targetImage *ebiten.Image) {

	var ax, ay, bx, by, cx, cy, dx, dy int

	x0, y0 := s.origin.Planet().Center()
	x1, y1 := s.destination.Planet().Center()

	if x1 > x0 {
		ax = x0
		bx = int(s.x)
		cx = int(s.x)
		dx = x1
	} else if x1 < x0 {
		ax = int(s.x)
		bx = x0
		cx = x1
		dx = int(s.x)
	} else if x1 == x0 {
		ax = x0
		bx = x0 + ui.Border
		cx = x0
		dx = x0 + ui.Border
	}

	if y1 > y0 {
		ay = y0
		by = int(s.y)
		cy = int(s.y)
		dy = y1
	} else if y1 < y0 {
		ay = int(s.y)
		by = y0
		cy = y1
		dy = int(s.y)
	} else if y1 == y0 {
		ay = y0
		by = y0 + ui.Border
		cy = y0
		dy = y0 + ui.Border
	}

	trailOpts := &ebiten.DrawImageOptions{}
	trailOpts.GeoM.Translate(float64(ax), float64(ay))
	trailRect := image.Rect(ax-int(s.baseX), ay-int(s.baseY), bx-int(s.baseX), by-int(s.baseY))
	targetImage.DrawImage(s.baseCourse.SubImage(trailRect).(*ebiten.Image), trailOpts)

	headingOpts := &ebiten.DrawImageOptions{}
	headingOpts.GeoM.Translate(float64(cx), float64(cy))
	headingRect := image.Rect(cx-int(s.baseX), cy-int(s.baseY), dx-int(s.baseX), dy-int(s.baseY))
	targetImage.DrawImage(s.course.SubImage(headingRect).(*ebiten.Image), headingOpts)
}
