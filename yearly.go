package main

import "github.com/jcgraybill/ship-shape/structure"

func structuresGenerateIncome(g *Game) {
	for _, s := range g.player.Structures() {
		if s.Class() == structure.Residential {
			s.GenerateIncome()
		}
	}
}

func payWorkers(g *Game) {

	for _, s := range g.player.Structures() {
		g.player.RemoveMoney(uint(s.LaborCost()))
	}
}

func distributeWorkers(g *Game) {
	for _, s := range g.player.Structures() {
		s.AssignWorkers(0)
	}
	budget := g.player.Money()
	for workersToAssign, _, _ := g.player.Population(); workersToAssign > 0; {
		workersAssigned := false
		for _, s := range g.player.Structures() {
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
