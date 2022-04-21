package main

import (
	"github.com/jcgraybill/ship-shape/ship"
	"github.com/jcgraybill/ship-shape/structure"
)

type income struct {
	s      *structure.Structure
	amount int
}

func collectIncome(g *Game) {
	if g.capitols > 0 {
		avail := make([]*income, 0)
		var cap *structure.Structure
		for _, s := range g.structures {
			if s.StructureType() == structure.Capitol {
				cap = s
			}
			if s.Income() > 0 {
				avail = append(avail, &income{s, s.Income()})
			}
		}

		if cap.HasShips() {
			var topOffer *income
			var topOfferValue float64 = 0
			for _, offer := range avail {

				shipAlreadyInLane := false
				for _, ship := range g.ships {
					_, origin, destination := ship.Manifest()
					if (origin == cap && destination == offer.s) ||
						(origin == offer.s && destination == cap) {
						shipAlreadyInLane = true
					}
				}

				if !shipAlreadyInLane {
					x1, y1 := cap.Planet().Center()
					x2, y2 := offer.s.Planet().Center()
					value := float64(offer.amount) / distance(float64(x1), float64(y1), float64(x2), float64(y2))
					if value > topOfferValue {
						topOffer = offer
						topOfferValue = value
					}
				}
			}

			if topOfferValue > 0 {
				ship := ship.New(cap, topOffer.s, ship.Income)
				cap.LaunchShip(0)
				g.ships[g.count] = ship
			}
		}
	}
}
