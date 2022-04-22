package ship

import (
	"image/color"

	"github.com/jcgraybill/ship-shape/ui"
)

func (s *Ship) LoadCargo(resource int, cargoColor color.RGBA) {
	v, i := ui.Triangle(plumeW+2, 2, shipW-6, shipH-4, cargoColor)
	s.image.DrawTriangles(v, i, ui.Src, nil)
	s.plume.DrawTriangles(v, i, ui.Src, nil)
	s.cargo = resource
}
