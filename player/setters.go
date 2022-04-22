package player

func (p *Player) AddMoney(money uint) {
	p.money += int(money)
}

func (p *Player) RemoveMoney(money uint) {
	p.money -= int(money)
}

func (p *Player) SetPopulation(pop, maxPop, workersNeeded int) {
	p.pop, p.maxPop, p.workersNeeded = pop, maxPop, workersNeeded
}
