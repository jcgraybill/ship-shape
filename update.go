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

var ut [13]uint64
var um [13]uint64

func (g *Game) Update() error {
	g.count++

	ut[0], um[0] = g.measure(g.updatePlayerPanel)

	ut[1], um[1] = g.measure(g.handleMouseClicks)
	ut[2], um[2] = g.measure(g.handleKeyPresses)

	ut[3], um[3] = g.measure(g.structuresProduce)
	ut[4], um[4] = g.measure(g.structuresConsume)

	if g.count%ui.BroadcastFrequency == 0 {
		ut[5], um[5] = g.measure(g.structuresBidForResources)
		ut[6], um[6] = g.measure(g.collectIncome)
	} else {
		ut[5], um[5], ut[6], um[6] = 0, 0, 0, 0
	}

	ut[7], um[7] = g.measure(g.shipsArrive)

	ut[8], um[8] = g.measure(g.updatePopulation)
	ut[9], um[9] = g.measure(g.structuresGenerateIncome)

	if g.count%ui.YearLength == 0 {
		g.year += 1
		ut[10], um[10] = g.measure(g.payWorkers)
		ut[11], um[11] = g.measure(g.distributeWorkers)
	} else {
		ut[10], um[10], ut[11], um[11] = 0, 0, 0, 0
	}

	ut[12], ut[12] = g.measure(g.updateLevel)
	UpdateTimeLogger.Println(g.count, ",", g.csv(ut[:]))
	UpdateMemLogger.Println(g.count, ",", g.csv(um[:]))
	g.instrument()
	return nil
}

func (g *Game) updateLevel() {
	g.level.Update(g.player)
}

func (g *Game) updatePopulation() {
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

func (g *Game) structuresProduce() {
	for _, s := range g.player.Structures() {
		if s.Produce(g.count) && s.IsHighlighted() {
			g.panel.Clear()
			g.updatePopulation()
			showStructurePanel(g, s)
		}
	}
}

func (g *Game) structuresConsume() {
	for _, s := range g.player.Structures() {
		consumed, downgrade := s.Consume(g.count)
		if downgrade > 0 {
			s.Upgrade(downgrade, &g.structureData[downgrade])
		}

		if (consumed || downgrade > 0) && s.IsHighlighted() {
			g.panel.Clear()
			g.updatePopulation()
			showStructurePanel(g, s)
		}
	}
}

func (g *Game) shipsArrive() {
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
