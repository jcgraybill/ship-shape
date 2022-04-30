package main

import (
	"github.com/jcgraybill/ship-shape/ship"
	"github.com/jcgraybill/ship-shape/structure"
)

func (g *Game) collectIncome() {
	if exists, cap := g.player.Capitol(); exists {

		if cap.HasShips() {
			var topOfferStructure *structure.Structure
			var topOfferValue float64 = 0

			for _, s := range g.player.Structures() {
				if s.Income() > 0 {
					var offerValue float64
					x1, y1 := cap.Planet().Center()
					x2, y2 := s.Planet().Center()
					offerValue = float64(s.Income()) / distance(x1, y1, x2, y2)

					if offerValue > topOfferValue {
						shipAlreadyInLane := false
						for _, ship := range g.player.Ships() {
							_, origin, destination := ship.Manifest()
							if (origin == cap && destination == s) ||
								(origin == s && destination == cap) {
								shipAlreadyInLane = true
							}
						}
						if !shipAlreadyInLane {
							topOfferStructure = s
							topOfferValue = offerValue
						}
					}
				}
			}

			if topOfferValue > 0 {
				ship := ship.New(cap, topOfferStructure, ship.Income)
				cap.LaunchShip(0)
				g.player.Ships()[g.count] = ship
			}
		}
	}
}
