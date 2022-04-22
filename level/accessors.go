package level

import (
	"github.com/jcgraybill/ship-shape/planet"
)

func (lvl *Level) Planets() []*planet.Planet {
	return lvl.planets
}

func (lvl *Level) AllowedStructures() []int {
	return lvl.allowedStructures
}

func (lvl *Level) StartingMoney() uint {
	return lvl.startingMoney
}
