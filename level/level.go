package level

import (
	"math/rand"

	"github.com/jcgraybill/ship-shape/planet"
	"github.com/jcgraybill/ship-shape/player"
	"github.com/jcgraybill/ship-shape/resource"
	"github.com/jcgraybill/ship-shape/ui"
)

type Level struct {
	Title             string
	W, H              int
	startingMoney     uint
	startingYear      uint
	allowedResources  []int
	allowedStructures []int
	progress          func(*Level, *player.Player) bool
	planets           []*planet.Planet
	message           string
}

func New(which uint) *Level {
	lvl := &level01
	lvl.planets = make([]*planet.Planet, 0)
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
			lvl.planets = append(lvl.planets, planet.New(x, y, planetResources, rd))
		}
	}

	return lvl
}

func (lvl *Level) Progress(p *player.Player) bool {
	return lvl.progress(lvl, p)
}
