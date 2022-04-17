package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/jcgraybill/ship-shape/planet"
	"github.com/jcgraybill/ship-shape/structure"
)

func (g *Game) Update() error {
	g.count++

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		handleLeftMouseButtonPress(g)
	} else if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		handleLeftMouseButtonRelease(g)
	}

	//TODO let every structure update its state
	for _, structure := range g.structures {
		if err := structure.Update(); err != nil {
			return err
		}
	}

	return nil
}

func handleLeftMouseButtonRelease(g *Game) {
	g.panel.LeftMouseButtonRelease(ebiten.CursorPosition())
}

func handleLeftMouseButtonPress(g *Game) {
	if !g.panel.LeftMouseButtonPress(ebiten.CursorPosition()) {
		g.panel.Clear()

		for _, planet := range g.planets {
			if planet.MouseButton(ebiten.CursorPosition()) {
				planet.Highlight()
				g.panel.AddLabel(planet.Describe())
				g.panel.AddButton("build "+g.structureData["water"].DisplayName, generateConstructionCallback(g, planet, "water"))
				g.panel.AddButton("build "+g.structureData["outpost"].DisplayName, generateConstructionCallback(g, planet, "outpost"))
			} else {
				planet.Unhighlight()
			}
		}

		for _, structure := range g.structures {
			if structure.MouseButton(ebiten.CursorPosition()) {
				g.panel.AddLabel(structure.Describe())
			}
		}

	}
}

func generateConstructionCallback(g *Game, p *planet.Planet, structureType string) func() {
	return func() {
		g.panel.Clear()
		structure := structure.New(g.structureData[structureType], p)
		g.panel.AddLabel(structure.Describe())
		g.structures = append(g.structures, structure)
	}
}
