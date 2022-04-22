package main

import "github.com/jcgraybill/ship-shape/structure"

func structuresGenerateIncome(g *Game) {
	for _, s := range g.structures {
		if s.Class() == structure.Residential {
			s.GenerateIncome()
		}
	}
}

func payWorkers(g *Game) {
	wages := 0
	for _, s := range g.structures {
		wages += s.LaborCost()
	}
	if wages <= g.money {
		g.money -= wages
	} else {
		g.money = 0
	}
}

func distributeWorkers(g *Game) {
	for _, s := range g.structures {
		s.AssignWorkers(0)
	}
	budget := g.money
	for workersToAssign := g.pop; workersToAssign > 0; {
		workersAssigned := false
		for _, s := range g.structures {
			if s.ActiveWorkers() < s.WorkerCapacity() && workersToAssign > 0 && s.CanProduce() && !s.IsPaused() {
				if budget >= s.WorkerCost() {
					budget -= s.WorkerCost()
					s.AssignWorkers(s.ActiveWorkers() + 1)
					workersToAssign -= 1
					workersAssigned = true
					if s.IsHighlighted() {
						g.panel.Clear()
						showStructurePanel(g, s)
					}
				}
			}
		}
		if !workersAssigned {
			workersToAssign = 0
		}
	}
}
