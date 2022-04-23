package level

import (
	"github.com/jcgraybill/ship-shape/player"
	"github.com/jcgraybill/ship-shape/resource"
	"github.com/jcgraybill/ship-shape/structure"
)

var level03 = Level{
	title:         "earning money",
	W:             1366,
	H:             768,
	startingMoney: 4000,
	startingYear:  2280,
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
		structure.HQ,
		structure.Outpost,
		structure.Water,
		structure.Habitat,
	},

	message: `Earning money

Structures and labor cost MONEY.
How can you earn money to pay for
more of them?

The POPULATION in your outposts
and habitats is constantly 
generating REVENUE. To collect
that revenue, you need a 
HEADQUARTERS.

Build three HABITATS again. This
time you have less starting
money, so you will need a
HEADQUARTERS to afford them.
	`,

	update: func(lvl *Level, p *player.Player) bool {
		lvl.progress = 0
		hq := false
		for _, s := range p.Structures() {
			if s.StructureType() == structure.Habitat {
				lvl.progress += 1
			}
			if s.StructureType() == structure.HQ {
				hq = true
			}
		}
		if hq {
			lvl.message = `The headquarters collects
revenue by sending ships to 
visit outposts and habitats,
just like the cargo ships that
deliver goods.

Any planet can support a 
headquarters equally well. It
doesn't consume or depend on any
particular resources.

Hint: place structures that 
interact with each other close
together. That gives their ships
shorter journeys, which makes it
easier to keep supplies flowing.`
		}

		if lvl.progress >= lvl.goal {
			lvl.message = `Great!

You built three settlements, and
balanced your economy.

You can upgrade habitats to
SETTLEMENTS, which support even
larger populations (and earn
even more money) by providing
them MACHINERY.

To create MACHINERY, you'll need
to learn about MANUFACTURING.

Click 'next'.`
			return true
		}
		return false
	},
}
