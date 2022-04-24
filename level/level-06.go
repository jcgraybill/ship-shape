package level

import (
	"github.com/jcgraybill/ship-shape/player"
	"github.com/jcgraybill/ship-shape/resource"
	"github.com/jcgraybill/ship-shape/structure"
)

var level06 = Level{
	title:         "sandbox",
	W:             4096,
	H:             4096,
	startingMoney: 10000,
	startingYear:  2500,
	label:         "colonies",
	goal:          5,
	progress:      0,
	nextLevel:     nil,

	allowedResources: []int{
		resource.Habitability,
		resource.Population,
		resource.Ice,
		resource.Water,
		resource.Iron,
		resource.Ore,
		resource.Metal,
		resource.Machinery,
		resource.Sand,
		resource.Silicon,
		resource.IntegratedCircuits,
		resource.Computers,
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
		structure.Silica,
		structure.ChipFoundry,
		structure.Assembly,
		structure.Colony,
		structure.Capitol,
	},

	message: `Now it's time to show off
your skills by creating COMPUTERS!

SILICA PURIFIERS create SILICON. 
Place them on planets with plenty 
of SAND. CHIP FOUNDRIES turn 
silicon into INTEGRATED CIRCUITS. 
Use integrated circuits to upgrade 
a factory to an ASSEMBLY PLANT. 
Assembly plants create COMPUTERS 
out of integrated circuits and 
machinery. Finally, use COMPUTERS 
to upgrade your headquarters to a 
CAPITOL, or your settlement to
a COLONY.

Create 4 COLONIES. Good luck!`,

	update: func(lvl *Level, p *player.Player) bool {
		lvl.progress = 0
		for _, s := range p.Structures() {
			if s.StructureType() == structure.Colony {
				lvl.progress += 1
			}
		}

		if lvl.progress >= lvl.goal {
			lvl.message = `Congratulations!

That was no small feat. Note the
year (above) - this shows how 
quickly you beat the final level
of this demo. 

Visit us on Github at
github.com/jcgraybill/ship-shape
to share how you did, as well as
any feedback. Come back soon for
more! Thank you for playing.

-Jules`
			return true
		}
		return false
	},
}
