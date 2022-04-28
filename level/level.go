package level

import (
	"math/rand"

	"github.com/jcgraybill/ship-shape/planet"
	"github.com/jcgraybill/ship-shape/player"
	"github.com/jcgraybill/ship-shape/resource"
	"github.com/jcgraybill/ship-shape/ui"
)

type Level struct {
	title             string
	W, H              int
	startingMoney     uint
	startingYear      uint
	allowedResources  []int
	allowedStructures []int
	update            func(*Level, *player.Player) bool
	planets           []*planet.Planet
	message, label    string
	progress, goal    uint
	goalMet           bool
	nextLevel         *Level
}

func StartingLevel() *Level {
	return &level01
}

func New(lvl *Level) *Level {
	lvl.planets = make([]*planet.Planet, 0)

	if lvl.W == 0 {
		lvl.W = 1
	}
	if lvl.H == 0 {
		lvl.H = 1
	}

	rd := resource.GetResourceData()

	cellsize := ui.PlanetSize * ui.PlanetDistance
	for i := 0; i < lvl.H/cellsize; i++ {
		//j < w/cellsize-1 prevents creating planets underneath the panel
		for j := 0; j < lvl.W/cellsize-1; j++ {
			x := j*cellsize + rand.Intn(cellsize-ui.PlanetSize*2) + ui.PlanetSize
			y := i*cellsize + rand.Intn(cellsize-ui.PlanetSize*2) + ui.PlanetSize

			planetResources := make(map[int]uint8)
			for _, r := range lvl.allowedResources {
				if rd[r].Source == resource.Planetary {
					planetResources[r] = uint8(rand.Intn(255))
				}
			}
			lvl.planets = append(lvl.planets, planet.New(x, y, planetResources))
		}
	}

	return lvl
}

func (lvl *Level) Update(p *player.Player) {
	if lvl.update(lvl, p) {
		lvl.goalMet = true
	}
}
