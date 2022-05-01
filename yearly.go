package main

import "github.com/jcgraybill/ship-shape/structure"

func (g *Game) structuresGenerateIncome() {
	for _, s := range g.player.Structures() {
		if s.Class() == structure.Residential {
			s.GenerateIncome()
		}
	}
}

func (g *Game) payWorkers() {

	for _, s := range g.player.Structures() {
		g.player.RemoveMoney(uint(s.LaborCost()))
	}
}

func (g *Game) distributeWorkers() {
	for _, s := range g.player.Structures() {
		s.AssignWorkers(0)
	}
	budget := g.player.Money()
	for workersToAssign := g.player.Population(); workersToAssign > 0; {
		workersAssigned := false
		for _, s := range g.player.Structures() {
			if s.ActiveWorkers() < s.WorkerCapacity() && workersToAssign > 0 && s.CanProduce() && !s.IsPaused() {
				if budget >= s.WorkerCost() {
					budget -= s.WorkerCost()
					s.AssignWorkers(s.ActiveWorkers() + 1)
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
