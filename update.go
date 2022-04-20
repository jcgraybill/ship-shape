package main

import (
	"github.com/jcgraybill/ship-shape/resource"
	"github.com/jcgraybill/ship-shape/ship"
	"github.com/jcgraybill/ship-shape/ui"
)

func (g *Game) Update() error {
	g.count++

	handleMouseClicks(g)
	handleKeyPresses(g)

	structuresProduce(g)
	structuresBidForResources(g)
	shipsArrive(g)

	updatePopulation(g)
	payWorkers(g)

	return nil
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
		structure.AssignWorkers(0)
	}
	budget := g.money
	for workersToAssign := g.pop; workersToAssign > 0; {
		workersAssigned := false
		for _, structure := range g.structures {
			if structure.Workers() < structure.WorkersNeeded() && workersToAssign > 0 {
				if budget >= structure.WorkerCost() {
					budget -= structure.WorkerCost()
					structure.AssignWorkers(structure.Workers() + 1)
					workersToAssign -= 1
					workersAssigned = true
				}
			}
		}
		if !workersAssigned {
			workersToAssign = 0
		}
	}
}

func structuresProduce(g *Game) {
	for _, structure := range g.structures {
		if structure.Produce(g.count) && structure.IsHighlighted() {
			g.panel.Clear()
			updatePopulation(g)
			showPlayerPanel(g.panel, g.money, g.pop, g.maxPop, g.workersNeeded)
			showStructurePanel(g, structure)
		}
	}
}

func shipsArrive(g *Game) {
	for key, s := range g.ships {
		if s.Update(g.count) { //ship has arrived
			cargo, origin, destination := s.Manifest()
			if cargo > 0 {
				destination.ReceiveCargo(cargo)
				if destination.IsHighlighted() {
					g.panel.Clear()
					showPlayerPanel(g.panel, g.money, g.pop, g.maxPop, g.workersNeeded)
					showStructurePanel(g, destination)
				}
				returnShip := ship.New(destination, origin)
				g.ships[key] = returnShip
			} else {
				destination.ReturnShip()
				delete(g.ships, key)
			}
		}
	}
}
