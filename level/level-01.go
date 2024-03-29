package level

import (
	"github.com/jcgraybill/ship-shape/player"
	"github.com/jcgraybill/ship-shape/resource"
	"github.com/jcgraybill/ship-shape/structure"
)

var level01 = Level{
	title:         "welcome!",
	W:             1900,
	H:             1080,
	startingMoney: 3200,
	label:         "total population",
	goal:          120,
	progress:      0,
	nextLevel:     &level02,

	allowedResources: []int{
		resource.Environment,
		resource.Population,
	},
	allowedStructures: []int{
		structure.Outpost,
	},

	message: `Welcome to ship-shape!

Your mission is to build a thriving
civilization among the stars.

To do so, you build STRUCTURES
on PLANETS. Click on a planet to see
more about it.

Planets have better or worse 
ENVIRONMENT ratings. Planets with 
better environments can support 
larger populations. 

Find a planet with a good
ENVIRONMENT and try building an 
OUTPOST. Hint: look for a bar 
that's mostly green.`,

	update: func(lvl *Level, p *player.Player) bool {
		lvl.progress = 0
		outposts := 0

		for _, s := range p.Structures() {
			if s.StructureType() == structure.Outpost {
				lvl.progress += uint(s.Storage()[resource.Population].Amount)
				outposts++
			}
		}

		if outposts == 1 {
			lvl.message = `Well done!
			
You can watch the population of 
your OUTPOST starting to grow, up
to its maximum size.

Structures such as outposts cost
money to build. You can see how
much money is in your bank just
above here (the amount after 
"bank:").

Try building more outposts. Can 
you get to 120 POPULATION with 
the  money you have available?`
		}

		if outposts > 1 {
			if p.Money() < 800 {
				if p.MaxPopulation() < 120 {
					p.AddMoney(800)
					lvl.message = `Uh oh. Looks like you're
running low on funds. 

Here's a bit more money. Keep
building outposts. 

Remember, look for planets with
high ENVIRONMENT ratings to get
the largest population in your
outposts.

Try scrolling around: there are
many more planets to choose from
than the first few you see here.`
				} else {
					lvl.message = `You've spent all your money, and
built enough outposts to support
a population of 120. 

Now just sit back and watch your
population grow!`
				}
			}
		}

		if lvl.progress >= lvl.goal {
			lvl.message = `Brilliant!

You've mastered building outposts. 
Outposts are one of the most 
important structures in the game, 
as they provide the homes where 
your colonists will live.

Click "next" to learn about 
resources and upgrades.`
			return true
		}
		return false
	},
}
