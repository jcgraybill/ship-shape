package level

import (
	"github.com/jcgraybill/ship-shape/player"
	"github.com/jcgraybill/ship-shape/resource"
	"github.com/jcgraybill/ship-shape/structure"
)

var level01 = Level{
	Title:         "welcome to ship-shape",
	W:             800,
	H:             600,
	startingMoney: 200,
	startingYear:  2250,
	allowedResources: []int{
		resource.Habitability,
		resource.Ice,
		resource.Water,
		resource.Population,
	},
	allowedStructures: []int{
		structure.Outpost,
		structure.Water,
		structure.HQ,
	},
	progress: level01Progress,
	message:  "welcome to ship-shape\nthank you for playing",
}

func level01Progress(lvl *Level, p *player.Player) bool {
	for _, s := range p.Structures() {
		if s.StructureType() == structure.Outpost {
			lvl.message = "you built an outpost"
		}
	}
	return false
}
