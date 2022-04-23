package level

import (
	"github.com/jcgraybill/ship-shape/player"
	"github.com/jcgraybill/ship-shape/resource"
	"github.com/jcgraybill/ship-shape/structure"
)

var level04 = Level{} // placeholder

var level03 = Level{
	title:         "upgrades and resources",
	W:             1366,
	H:             768,
	startingMoney: 4000,
	startingYear:  2265,
	label:         "habitats",
	goal:          3,
	progress:      0,
	nextLevel:     New(&level04),

	allowedResources: []int{
		resource.Habitability,
		resource.Population,
		resource.Ice,
		resource.Water,
	},
	allowedStructures: []int{
		structure.Outpost,
		structure.Water,
		structure.Habitat,
	},

	message: `Resources and upgrades

You can increase your population
by upgrading your outposts to 
HABITATS.

Outposts need plenty of WATER to 
upgrade to habitats.

Fortunately, ICE deposits found
on many planets can be harvested
for water.

Find a planet with a lot of ice,
and build an ICE HARVESTER. Also
build an outpost or two.`,

	update: func(lvl *Level, p *player.Player) bool {
		lvl.progress = 0
		ice := false
		for _, s := range p.Structures() {
			if s.StructureType() == structure.Habitat {
				lvl.progress += 1
			}
			if s.StructureType() == structure.Water {
				ice = true
			}
		}
		if lvl.progress > 0 {
			lvl.message = `Good job.

HABITATS can support a much
larger population than outposts.

A larger population comes with
a cost, though. You'll need to
keep your habitat well-
supplied with water.

If a habitat ever runs out of
water, residents will become
unhappy and move away. Your 
structure will turn back into
an OUTPOST, and you'll need
to pay to upgrade it again.

Keep building HABITATS.
`
		} else if ice {
			lvl.message = `Structures such as the 
ice harvester need WORKERS!

Workers come from the 
population of your outposts and
habitats. Once per year, workers
are assigned to structures that
have available jobs. Better-staffed
structures are more productive.

Labor costs money. You can see what
wages the workers will be paid 
by clicking each structure.

Structures automatically deliver the
goods they produce to other
structures that need those goods. 

Once an outpost has a full supply of
water, upgrade it to a HABITAT!`
		}

		if p.Money() < 800 {
			p.AddMoney(400)
		}

		if lvl.progress >= lvl.goal {
			lvl.message = `Nice!

The same way planets with more
HABITABILITY yield more productive
outposts, planets with more ICE
improve the productivity of
ICE HARVESTERS.

Click 'next' to learn about earning
money.`
			return true
		}
		return false
	},
}
