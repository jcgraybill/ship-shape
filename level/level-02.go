package level

import (
	"github.com/jcgraybill/ship-shape/player"
	"github.com/jcgraybill/ship-shape/resource"
	"github.com/jcgraybill/ship-shape/structure"
)

var level02 = Level{
	title:         "level 2",
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
	update: func(lvl *Level, p *player.Player) bool {
		lvl.progress = 0
		for _, s := range p.Structures() {
			if s.StructureType() == structure.Outpost {
				lvl.progress += uint(s.Storage()[resource.Population].Amount)
			}
		}
		if lvl.progress >= lvl.goal {
			return true
		}
		return false
	},
	label:    "total population",
	goal:     12,
	progress: 0,
	message: `welcome to ship-shape
This is level 2!`,
}
