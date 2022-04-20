package main

import (
	"github.com/jcgraybill/ship-shape/resource"
	"github.com/jcgraybill/ship-shape/ship"
)

func (g *Game) Update() error {
	g.count++

	handleMouseClicks(g)
	handleKeyPresses(g)

	structuresProduce(g)
	structuresBidForResources(g)
	shipsArrive(g)

	updatePopulation(g)

	return nil
}

func updatePopulation(g *Game) {
	g.pop, g.maxPop, g.workersNeeded = 0, 0, 0
	for _, structure := range g.structures {
		g.pop += int(structure.Storage()[resource.Population].Amount)
		g.maxPop += int(structure.Storage()[resource.Population].Capacity)
		g.workersNeeded += structure.WorkersNeeded()
		structure.AssignWorkers(0)
	}
	for workersToAssign := g.pop; workersToAssign > 0; {
		workersAssigned := false
		for _, structure := range g.structures {
			if structure.Workers() < structure.WorkersNeeded() && workersToAssign > 0 {
				structure.AssignWorkers(structure.Workers() + 1)
				workersToAssign -= 1
				workersAssigned = true
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
			showPopulationPanel(g.panel, g.pop, g.maxPop, g.workersNeeded)
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
					showPopulationPanel(g.panel, g.pop, g.maxPop, g.workersNeeded)
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
