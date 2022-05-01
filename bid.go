package main

import (
	"math"

	"github.com/jcgraybill/ship-shape/ship"
	"github.com/jcgraybill/ship-shape/structure"
)

const PriorityBidValue = 255

func (g *Game) structuresBidForResources() {
	for _, s0 := range g.player.Structures() {
		if s0.Class() == structure.Extractor || s0.Class() == structure.Processor {
			if s0.Storage()[s0.Produces()].Amount > 0 && s0.HasShips() {
				var topBidStructure *structure.Structure
				var topBidValue float64 = 0
				for _, s1 := range g.player.Structures() {
					if !s1.IsPaused() {
						if s0.Produces() != s1.Produces() {
							if _, ok := s1.Storage()[s0.Produces()]; ok {
								if s1.Storage()[s0.Produces()].Amount+s1.Storage()[s0.Produces()].Incoming < s1.Storage()[s0.Produces()].Capacity {
									var bidValue float64
									if s1.IsPrioritized() {
										bidValue = PriorityBidValue
									} else {
										bidValue = float64(s1.Storage()[s0.Produces()].Capacity - s1.Storage()[s0.Produces()].Amount - s1.Storage()[s0.Produces()].Incoming)
										bidValue = bidValue * (255 / float64(s1.Storage()[s0.Produces()].Capacity))

										x1, y1 := s0.Planet().Center()
										x2, y2 := s1.Planet().Center()
										bidValue = bidValue / distance(x1, y1, x2, y2)
									}

									if bidValue > topBidValue {
										shipAlreadyInLane := false
										for _, sh := range g.player.Ships() {
											_, origin, destination := sh.Manifest()
											if (origin == s0 && destination == s1) ||
												(origin == s1 && destination == s0) {
												shipAlreadyInLane = true
											}
										}
										if !shipAlreadyInLane {
											topBidStructure = s1
											topBidValue = bidValue
										}

									}
								}
							}
						}
					}
				}
				if topBidValue > 0 {
					sh := ship.New(s0, topBidStructure, ship.Cargo)
					sh.LoadCargo(s0.Produces(), g.resourceData[s0.Produces()].Color)
					s0.LaunchShip(s0.Produces())
					topBidStructure.AwaitDelivery(s0.Produces())
					g.player.Ships()[g.count] = sh
				}
			}
		}
	}
}

func distance(x1, y1, x2, y2 int) float64 {
	return math.Sqrt(math.Pow(math.Abs(float64(x1-x2)), 2) + math.Pow(math.Abs(float64(y1-y2)), 2))
}
