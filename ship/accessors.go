package ship

import "github.com/jcgraybill/ship-shape/structure"

func (s *Ship) Manifest() (int, *structure.Structure, *structure.Structure) {
	return s.cargo, s.origin, s.destination
}

func (s *Ship) ShipType() int {
	return s.shipType
}
