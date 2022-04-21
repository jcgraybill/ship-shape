package main

import (
	"math"

	"github.com/jcgraybill/ship-shape/ship"
	"github.com/jcgraybill/ship-shape/structure"
	"github.com/jcgraybill/ship-shape/ui"
)

type Bid struct {
	Structure *structure.Structure
	Resource  int
	Urgency   uint8
}

func structuresBidForResources(g *Game) {
	if g.count%ui.BidFrequency == 0 {

		bids := make([]*Bid, 0)
		for _, structure := range g.structures {
			for resource, urgency := range structure.Bid() {
				bids = append(bids, &Bid{Structure: structure, Resource: resource, Urgency: urgency})
			}
		}

		for _, structure := range g.structures {
			if structure.HasShips() {
				var topBid *Bid
				var topBidValue float64 = 0
				for _, bid := range bids {
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
								topBid = bid
								topBidValue = value
							}
						}
					}
				}
				if topBidValue > 0 {

					ship := ship.New(structure, topBid.Structure)
					ship.LoadCargo(topBid.Resource, g.resourceData[topBid.Resource].Color)
					structure.LaunchShip(topBid.Resource)
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
}

func distance(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(math.Pow(math.Abs(x1-x2), 2) + math.Pow(math.Abs(y1-y2), 2))
}
