package ship

import (
	"image/color"
)

func (s *Ship) LoadCargo(resource int, cargoColor color.RGBA) {
	s.image, s.plume = nil, nil
	s.cargo = resource
	s.cargoColor = cargoColor
}
