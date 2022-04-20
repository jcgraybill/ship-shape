package main

import (
	"github.com/jcgraybill/ship-shape/ship"
	"github.com/jcgraybill/ship-shape/structure"
)

type Bid struct {
	Structure *structure.Structure
	Resource  int
	Urgency   uint8
}

func bidForResources(g *Game) {
	bids := make([]*Bid, 0)
	for _, structure := range g.structures {

		if resource, urgency := structure.Bid(); urgency > 0 {
			bids = append(bids, &Bid{Structure: structure, Resource: resource, Urgency: urgency})
		}
	}

	for _, structure := range g.structures {
		if structure.HasShips() {
			var topBid int
			var topBidValue float64 = 0
			for i, bid := range bids {
				if bid.Resource == structure.Produces() && structure.Storage()[bid.Resource].Amount > 0 {

					shipAlreadyInLane := false
					for _, ship := range g.ships {
						_, origin, destination := ship.Manifest()
						if (origin == structure && destination == bid.Structure) ||
							(origin == bid.Structure && destination == structure) {
							shipAlreadyInLane = true
						}
					}
					if !shipAlreadyInLane {
						x1, y1 := structure.Planet().Center()
						x2, y2 := bid.Structure.Planet().Center()
						value := float64(bid.Urgency) / distance(float64(x1), float64(y1), float64(x2), float64(y2))
						if value > topBidValue {
							topBid = i
							topBidValue = value
						}
					}
				}
			}
			if topBidValue > 0 {

				ship := ship.New(structure, bids[topBid].Structure)
				ship.LoadCargo(bids[topBid].Resource)
				structure.LaunchShip(bids[topBid].Resource)
				g.ships[g.count] = ship
				if structure.IsHighlighted() {
					g.panel.Clear()
					showStructurePanel(g, structure)
				}
				break // prevents another structure from accepting the same bid
			}
		}
	}
}
