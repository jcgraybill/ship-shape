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
		for _, s := range g.player.Structures() {
			for resource, urgency := range s.Bid() {
				bids = append(bids, &Bid{Structure: s, Resource: resource, Urgency: urgency})
			}
		}

		for _, s := range g.player.Structures() {
			if s.HasShips() {
				var topBid *Bid
				var topBidValue float64 = 0
				for _, bid := range bids {
					if bid.Resource == s.Produces() && s.Storage()[bid.Resource].Amount > 0 {

						shipAlreadyInLane := false
						for _, sh := range g.player.Ships() {
							_, origin, destination := sh.Manifest()
							if (origin == s && destination == bid.Structure) ||
								(origin == bid.Structure && destination == s) {
								shipAlreadyInLane = true
							}
						}
						if !shipAlreadyInLane {
							x1, y1 := s.Planet().Center()
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

					sh := ship.New(s, topBid.Structure, ship.Cargo)
					sh.LoadCargo(topBid.Resource, g.resourceData[topBid.Resource].Color)
					s.LaunchShip(topBid.Resource)
					g.player.Ships()[g.count] = sh
					if s.IsHighlighted() {
						g.panel.Clear()
						showStructurePanel(g, s)
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
