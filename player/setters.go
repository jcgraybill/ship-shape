package player

import "github.com/jcgraybill/ship-shape/structure"

func (p *Player) AddMoney(money uint) {
	p.money += int(money)
}

func (p *Player) RemoveMoney(money uint) {
	p.money -= int(money)
	if p.money < 0 {
		p.money = 0
	}
}

func (p *Player) SetPopulation(pop, maxPop, workersNeeded int) {
	p.population, p.maxPopulation, p.workersNeeded = pop, maxPop, workersNeeded
}

func (p *Player) AddStructure(s *structure.Structure) {
	p.structures = append(p.structures, s)
	if s.Class() == structure.Tax {
		p.capitol = s
	}
}
