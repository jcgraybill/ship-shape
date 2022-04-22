package ship

import (
	"math/rand"

	"github.com/jcgraybill/ship-shape/ui"
)

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
