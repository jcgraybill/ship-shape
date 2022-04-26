package ship

import (
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
	//s.baseCourse.SubImage(image.Rect(int(s.x)-int(s.baseX), int(s.y)-int(s.baseY), int(x1)-int(s.baseX), int(y1)-int(s.baseY))).(*ebiten.Image)
	targetImage.DrawImage(s.baseCourse, s.courseOpts)
}
