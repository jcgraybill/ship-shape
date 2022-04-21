package main

import (
	"image/color"

	"github.com/jcgraybill/ship-shape/resource"
	"github.com/jcgraybill/ship-shape/ship"
	"github.com/jcgraybill/ship-shape/structure"
	"github.com/jcgraybill/ship-shape/ui"
)

func (g *Game) Update() error {
	g.count++
	updatePlayerPanel(g)
	handleMouseClicks(g)
	handleKeyPresses(g)

	structuresGenerateIncome(g)
	structuresProduce(g)
	structuresConsume(g)
	structuresBidForResources(g)
	collectIncome(g)
	shipsArrive(g)

	updatePopulation(g)
	payWorkers(g)

	return nil
}

func structuresGenerateIncome(g *Game) {
	if g.count%ui.DayLength == 0 {
		for _, structure := range g.structures {
			structure.GenerateIncome()
		}
	}
}

func payWorkers(g *Game) {
	if g.count%ui.DayLength == 0 {
		wages := 0
		for _, structure := range g.structures {
			wages += structure.LaborCost()
		}
		if wages <= g.money {
			g.money -= wages
		} else {
			g.money = 0
		}

	}
}

func updatePopulation(g *Game) {
	g.pop, g.maxPop, g.workersNeeded = 0, 0, 0
	for _, structure := range g.structures {
		g.pop += int(structure.Storage()[resource.Population].Amount)
		g.maxPop += int(structure.Storage()[resource.Population].Capacity)
		g.workersNeeded += structure.WorkersNeeded()
	}
	if g.count%ui.DayLength == 0 {
		for _, structure := range g.structures {
			structure.AssignWorkers(0)
		}
		budget := g.money
		for workersToAssign := g.pop; workersToAssign > 0; {
			workersAssigned := false
			for _, structure := range g.structures {
				if structure.Workers() < structure.WorkersNeeded() && workersToAssign > 0 && structure.CanProduce() && !structure.IsPaused() {
					if budget >= structure.WorkerCost() {
						budget -= structure.WorkerCost()
						structure.AssignWorkers(structure.Workers() + 1)
						workersToAssign -= 1
						workersAssigned = true
						if structure.IsHighlighted() {
							g.panel.Clear()
							showStructurePanel(g, structure)
						}
					}
				}
			}
			if !workersAssigned {
				workersToAssign = 0
			}
		}
	}
}

func structuresProduce(g *Game) {
	for _, structure := range g.structures {
		if structure.Produce(g.count) && structure.IsHighlighted() {
			g.panel.Clear()
			updatePopulation(g)
			showStructurePanel(g, structure)
		}
	}
}

func structuresConsume(g *Game) {
	for _, structure := range g.structures {
		consumed, downgrade := structure.Consume(g.count)
		if downgrade > 0 {
			structure.Upgrade(downgrade, g.structureData[downgrade])
		}

		if (consumed || downgrade > 0) && structure.IsHighlighted() {
			g.panel.Clear()
			updatePopulation(g)
			showStructurePanel(g, structure)
		}
	}
}

func shipsArrive(g *Game) {
	for key, s := range g.ships {
		if s.Update(g.count) { //ship has arrived
			cargo, origin, destination := s.Manifest()

			if s.ShipType() == ship.Income && origin.StructureType() == structure.HQ {
				returnShip := ship.New(destination, origin, ship.Income)
				returnShip.LoadCargo(destination.CollectIncome(), color.RGBA{0xd4, 0xaf, 0x47, 0xff})
				g.ships[key] = returnShip
				continue
			}

			if s.ShipType() == ship.Income && destination.StructureType() == structure.HQ {
				g.money += cargo
				destination.ReturnShip()
				delete(g.ships, key)
				continue
			}

			if cargo > 0 {
				destination.ReceiveCargo(cargo)
				if destination.IsHighlighted() {
					g.panel.Clear()
					showStructurePanel(g, destination)
				}
				returnShip := ship.New(destination, origin, ship.Cargo)
				g.ships[key] = returnShip
			} else {
				destination.ReturnShip()
				delete(g.ships, key)
			}
		}
	}
}
