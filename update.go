package main

import (
	"image/color"

	"github.com/jcgraybill/ship-shape/resource"
	"github.com/jcgraybill/ship-shape/ship"
	"github.com/jcgraybill/ship-shape/structure"
	"github.com/jcgraybill/ship-shape/ui"
)

// FIXME - reactivate any buttons that were
// deactivated for lack of funds, when funds
// become available.
func (g *Game) Update() error {
	g.count++

	t := g.measure(updatePlayerPanel)
	UpdateLogger.Printf("%d", t.Milliseconds())

	handleMouseClicks(g)
	handleKeyPresses(g)

	structuresProduce(g)
	structuresConsume(g)

	if g.count%ui.BroadcastFrequency == 0 {
		structuresBidForResources(g)
		collectIncome(g)
	}

	shipsArrive(g)

	updatePopulation(g)
	structuresGenerateIncome(g)

	if g.count%ui.YearLength == 0 {
		g.year += 1
		payWorkers(g)
		distributeWorkers(g)
	}

	g.level.Update(g.player)
	g.instrument()
	return nil
}

func updatePopulation(g *Game) {
	pop, maxPop, workersNeeded := 0, 0, 0
	for _, s := range g.player.Structures() {
		if pr, ok := s.Storage()[resource.Population]; ok {
			pop += int(pr.Amount)
			maxPop += int(pr.Capacity)
		}
		workersNeeded += s.WorkerCapacity()
	}
	g.player.SetPopulation(pop, maxPop, workersNeeded)
}

func structuresProduce(g *Game) {
	for _, s := range g.player.Structures() {
		if s.Produce(g.count) && s.IsHighlighted() {
			g.panel.Clear()
			updatePopulation(g)
			showStructurePanel(g, s)
		}
	}
}

func structuresConsume(g *Game) {
	for _, s := range g.player.Structures() {
		consumed, downgrade := s.Consume(g.count)
		if downgrade > 0 {
			s.Upgrade(downgrade, &g.structureData[downgrade])
		}

		if (consumed || downgrade > 0) && s.IsHighlighted() {
			g.panel.Clear()
			updatePopulation(g)
			showStructurePanel(g, s)
		}
	}
}

func shipsArrive(g *Game) {
	for key, sh := range g.player.Ships() {
		if sh.Update(g.count) { //ship has arrived
			cargo, origin, destination := sh.Manifest()

			if sh.ShipType() == ship.Income && origin.Class() == structure.Tax {
				returnShip := ship.New(destination, origin, ship.Income)
				returnShip.LoadCargo(destination.CollectIncome(), color.RGBA{0xd4, 0xaf, 0x47, 0xff})
				g.player.Ships()[key] = returnShip
				continue
			}

			if sh.ShipType() == ship.Income && destination.Class() == structure.Tax {
				g.player.AddMoney(uint(cargo))
				destination.ReturnShip()
				delete(g.player.Ships(), key)
				continue
			}

			if cargo > 0 {
				destination.ReceiveCargo(cargo)
				if destination.IsHighlighted() {
					g.panel.Clear()
					showStructurePanel(g, destination)
				}
				returnShip := ship.New(destination, origin, ship.Cargo)
				g.player.Ships()[key] = returnShip
			} else {
				destination.ReturnShip()
				delete(g.player.Ships(), key)
			}
		}
	}
}
