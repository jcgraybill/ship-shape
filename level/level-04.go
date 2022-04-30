package level

import (
	"github.com/jcgraybill/ship-shape/player"
	"github.com/jcgraybill/ship-shape/resource"
	"github.com/jcgraybill/ship-shape/structure"
)

var level04 = Level{
	title:         "manufacturing",
	W:             1900,
	H:             1080,
	startingMoney: 5000,
	label:         "metal",
	goal:          16,
	progress:      0,
	nextLevel:     &level05,

	allowedResources: []int{
		resource.Habitability,
		resource.Population,
		resource.Ice,
		resource.Water,
		resource.Iron,
		resource.Ore,
		resource.Metal,
	},
	allowedStructures: []int{
		structure.HQ,
		structure.Outpost,
		structure.Water,
		structure.Habitat,
		structure.Mine,
		structure.Smelter,
	},

	message: `Some structures can MANUFACTURE
goods, converting raw materials
into new kinds of resources.

The SMELTER uses ORE and water
to make METAL. Build MINES on
planets with high levels of IRON to
extract ore. Each SMELTER can store 
8 metal, so you'll need two smelters
to meet this level's goal.
`,

	update: func(lvl *Level, p *player.Player) bool {
		lvl.progress = 0
		mine, smelter := false, false
		for _, s := range p.Structures() {
			if s.StructureType() == structure.Smelter {
				lvl.progress += uint(s.Storage()[resource.Metal].Amount)
			}
			if s.StructureType() == structure.Mine {
				mine = true
			}
			if s.StructureType() == structure.Smelter {
				smelter = true
			}

		}
		if mine && !smelter {
			lvl.message = `MINES extract ORE. Once you have
enough money, buy one or more 
SMELTERS to convert the ore into 
METAL.

Balance your economy carefully, to 
generate enough revenue to afford
everything. If a structure is costing
too much to operate, you can 
PAUSE PRODUCTION for a while.`
		} else if !mine && smelter {
			lvl.message = `The SMELTER converts
ORE and WATER into METAL. With
no mines producing ore, your SMELTER
has nothing to do!

Until you have one or more mines, 
consider PAUSING PRODUCTION, so you
aren't paying to operate an idle
smelter.`
		} else if mine && smelter {
			lvl.message = `MINES and SMELTERS
work together to produce METAL.

Your SMELTERS, outposts, and habitats
all need water. Delivery ships try
to provide balanced levels of supplies
to everyone who needs them. For a 
small fee, you can ask them to 
PRIORITIZE DELIVERY to some 
structures.`
		}

		if lvl.progress >= lvl.goal {
			lvl.message = `Bravo!

You've mastered the skill of
PRODUCTION.

The next level will introduce
FACTORIES, which convert METAL
into MACHINERY. With MACHINERY
you can upgrade your HABITATS
to even larger SETTLEMENTS.

Click 'next' to continue.`
			return true
		}
		return false
	},
}
