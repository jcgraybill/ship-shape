package main

import (
	"fmt"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/jcgraybill/ship-shape/panel"
	"github.com/jcgraybill/ship-shape/planet"
	"github.com/jcgraybill/ship-shape/resource"
	"github.com/jcgraybill/ship-shape/structure"
	"github.com/jcgraybill/ship-shape/ui"
)

type Bid struct {
	Structure *structure.Structure
	Resource  int
	Urgency   uint8
}

func (g *Game) Update() error {
	g.count++

	handleInputEvents(g)

	if g.count%ui.BidFrequency == 0 {
		bidForResources(g)
	}

	return nil
}

func bidForResources(g *Game) {
	bids := make([]*Bid, 0)

	for _, structure := range g.structures {
		if structure.Produce(g.count) && structure.IsHighlighted() {
			g.panel.Clear()
			showStructurePanel(g, structure)
		}

		if resource, urgency := structure.Bid(); urgency > 0 {
			bids = append(bids, &Bid{Structure: structure, Resource: resource, Urgency: urgency})
			fmt.Println(fmt.Sprintf("%s %s bids %d for %s", structure.Name(), structure.Planet().Name(), urgency, g.resourceData[resource].DisplayName))
		}
	}

	for _, structure := range g.structures {
		var topBid int
		var topBidValue float64 = 0
		for i, bid := range bids {
			if bid.Resource == structure.Produces() && structure.Storage()[bid.Resource].Amount > 0 {
				x1, y1 := structure.Planet().Center()
				x2, y2 := bid.Structure.Planet().Center()
				value := float64(bid.Urgency) / distance(float64(x1), float64(y1), float64(x2), float64(y2))
				if value > topBidValue {
					topBid = i
					topBidValue = value
				}
			}
		}
		if topBidValue > 0 {
			fmt.Println(fmt.Sprintf("%s (%s) accepts %s (%s)'s bid for %s at %d / %f", structure.Name(), structure.Planet().Name(), bids[topBid].Structure.Name(), bids[topBid].Structure.Planet().Name(), g.resourceData[bids[topBid].Resource].DisplayName, bids[topBid].Urgency, topBidValue))
		}
	}
}

func distance(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(math.Pow(math.Abs(x1-x2), 2) + math.Pow(math.Abs(y1-y2), 2))
}

func handleInputEvents(g *Game) {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		if !g.panel.LeftMouseButtonPress(ebiten.CursorPosition()) {
			g.panel.Clear()

			for _, planet := range g.planets {
				if planet.MouseButton(ebiten.CursorPosition()) {
					planet.Highlight()
					showPlanet(g.panel, planet, g.resourceData)
					g.panel.AddButton("build "+g.structureData[structure.Water].DisplayName, generateConstructionCallback(g, planet, structure.Water))
					g.panel.AddButton("build "+g.structureData[structure.Outpost].DisplayName, generateConstructionCallback(g, planet, structure.Outpost))
				} else {
					planet.Unhighlight()
				}
			}

			for _, structure := range g.structures {
				if structure.MouseButton(ebiten.CursorPosition()) {
					structure.Highlight()
					showStructurePanel(g, structure)
				} else {
					structure.Unhighlight()
				}
			}

		}
	} else if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		g.panel.LeftMouseButtonRelease(ebiten.CursorPosition())
	}
}

func generateConstructionCallback(g *Game, p *planet.Planet, structureType int) func() {
	return func() {
		g.panel.Clear()
		structure := structure.New(g.structureData[structureType], p)
		showStructurePanel(g, structure)
		structure.Highlight()
		g.structures = append(g.structures, structure)
	}
}

func showPlanet(panel *panel.Panel, p *planet.Planet, rd [resource.ResourceDataLength]resource.ResourceData) {
	panel.AddLabel(fmt.Sprintf("planet: %s", p.Name()))
	for resource, level := range p.Resources() {
		panel.AddLabel(rd[resource].DisplayName)
		panel.AddBar(level, rd[resource].Color)
	}
}

func showStructure(panel *panel.Panel, s *structure.Structure, rd [resource.ResourceDataLength]resource.ResourceData) {
	panel.AddLabel(s.Name())

	if len(s.Storage()) > 0 {
		panel.AddDivider()
		panel.AddLabel("storage:")
		for _, st := range s.Storage() {
			panel.AddLabel(fmt.Sprintf("%s (%d/%d)", rd[st.Resource].DisplayName, st.Amount, st.Capacity))
			panel.AddBar(uint8((255*int(st.Amount))/int(st.Capacity)), rd[st.Resource].Color)
		}
	}
}

func showStructurePanel(g *Game, structure *structure.Structure) {
	showStructure(g.panel, structure, g.resourceData)
	g.panel.AddDivider()
	showPlanet(g.panel, structure.Planet(), g.resourceData)
	g.panel.AddDivider()
}
