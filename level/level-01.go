package level

import (
	"github.com/jcgraybill/ship-shape/player"
	"github.com/jcgraybill/ship-shape/resource"
	"github.com/jcgraybill/ship-shape/structure"
)

var level01 = Level{
	title:         "welcome to ship-shape",
	W:             1920,
	H:             1080,
	startingMoney: 200,
	startingYear:  2250,
	nextLevel:     New(&level02),
	allowedResources: []int{
		resource.Habitability,
		resource.Population,
	},
	allowedStructures: []int{
		structure.Outpost,
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
	message: `Welcome to ship-shape!

Your mission is to build a thriving
civilization among the stars.

To do so, you build STRUCTURES
on PLANETS. Click on a planet to see
more about it.
	`,
}
