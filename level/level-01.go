package level

import (
	"github.com/jcgraybill/ship-shape/resource"
	"github.com/jcgraybill/ship-shape/structure"
)

var level01 = Level{
	Title:         "welcome to ship-shape",
	W:             800,
	H:             600,
	startingMoney: 200,
	allowedResources: []int{
		resource.Habitability,
		resource.Ice,
		resource.Water,
		resource.Population,
	},
	allowedStructures: []int{
		structure.Outpost,
		structure.Water,
	},
	Progress: level01Progress,
}

func level01Progress() bool {
	return false
}
