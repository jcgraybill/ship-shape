package ship

import (
	"image/color"
	"math/rand"

	"github.com/jcgraybill/ship-shape/ui"
)

func (s *Ship) Update(count uint) bool {
	if count%plumeCycleTime == 0 {
		if rand.Intn(plumeFrequency) == 0 {
			s.plumeVisible = false
		} else {
			s.plumeVisible = true
		}
	}

	if s.destination.Planet().Bounds.At(int(s.x), int(s.y)) == color.Opaque {
		return true
	} else {
		if count%ui.ShipSpeed == 0 {
			s.x += s.dx
			s.y += s.dy
			s.opts.GeoM.Translate(s.dx, s.dy)
		}
	}
	return false
}
