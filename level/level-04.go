package level

import (
	"github.com/jcgraybill/ship-shape/player"
	"github.com/jcgraybill/ship-shape/resource"
	"github.com/jcgraybill/ship-shape/structure"
)

var level05 = Level{message: ":("} // placeholder

var level04 = Level{
	title:         "manufacturing",
	W:             1900,
	H:             1080,
	startingMoney: 5000,
	startingYear:  2300,
	label:         "metal",
	goal:          16,
	progress:      0,
	nextLevel:     New(&level05),

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
planets with high levels of IRON
to extract ore. Each SMELTER can
store 8 metal, so you'll need 
two smelters to meet this 
level's goal.
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
				mine = true
			}

		}
		if mine && !smelter {
			lvl.message = `mine !smelter

Balance
your economy carefully, to generate 
enough revenue to afford
everything.`
		} else if !mine && smelter {
			lvl.message = `!mine smelter

Balance
your economy carefully, to generate 
enough revenue to afford
everything.`
		} else if mine && smelter {
			lvl.message = `mine smelter`
		}

		if lvl.progress >= lvl.goal {
			lvl.message = `Bravo!

Click 'next'.`
			return true
		}
		return false
	},
}
