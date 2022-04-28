package ship

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

func (s *Ship) Draw(targetImage *ebiten.Image) {
	if s.plumeVisible {
		targetImage.DrawImage(s.plume, s.opts)
	} else {
		targetImage.DrawImage(s.image, s.opts)
	}
}

func (s *Ship) DrawCourse(targetImage *ebiten.Image) {

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
