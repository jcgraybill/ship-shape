package level

import (
	"github.com/jcgraybill/ship-shape/player"
	"github.com/jcgraybill/ship-shape/resource"
	"github.com/jcgraybill/ship-shape/structure"
)

var level05 = Level{
	title:         "factories",
	W:             2048,
	H:             2048,
	startingMoney: 5000,
	startingYear:  2350,
	label:         "settlements",
	goal:          2,
	progress:      0,
	nextLevel:     New(&level06),

	allowedResources: []int{
		resource.Habitability,
		resource.Population,
		resource.Ice,
		resource.Water,
		resource.Iron,
		resource.Ore,
		resource.Metal,
		resource.Machinery,
	},
	allowedStructures: []int{
		structure.HQ,
		structure.Outpost,
		structure.Water,
		structure.Habitat,
		structure.Mine,
		structure.Smelter,
		structure.Factory,
		structure.Settlement,
	},

	message: `You now have access to two
new structures: 

FACTORIES turn METAL into
MACHINERY

MACHINERY allows you to upgrade
habitats to SETTLEMENTS, supporting
even larger populations!

Create two SETTLEMENTS.
`,

	update: func(lvl *Level, p *player.Player) bool {
		lvl.progress = 0
		for _, s := range p.Structures() {
			if s.StructureType() == structure.Settlement {
				lvl.progress += 1
			}
		}

		if lvl.progress >= lvl.goal {
			lvl.message = `Nicely done!

Multi-stage supply chains can get
tricky to keep balanced. Remember
that it's always better to place
structures that exchange materials
with each other close together.

Click 'next' to continue.`
			return true
		}
		return false
	},
}
