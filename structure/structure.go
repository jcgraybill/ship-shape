package structure

// TODO - don't create image when created,
// only on first update, to remove all display
// code from Update() loop.

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jcgraybill/ship-shape/planet"
	"github.com/jcgraybill/ship-shape/ui"
)

type Structure struct {
	highlighted             bool
	paused                  bool
	prioritized             bool
	image                   *ebiten.Image
	highlightedImage        *ebiten.Image
	displayOpts             *ebiten.DrawImageOptions
	planet                  *planet.Planet
	data                    *StructureData
	storage                 map[int]*Storage
	berths, ships, inFlight int
	workers                 int
	income                  float64
	structureType           int
	Bounds                  image.Rectangle
}

func New(structureType int, sd *StructureData, p *planet.Planet) *Structure {
	s := Structure{
		data:        sd,
		planet:      p,
		highlighted: false,
		paused:      false,
		prioritized: false,
		workers:     0,

		displayOpts:   &ebiten.DrawImageOptions{},
		structureType: structureType,
	}

	s.planet.ReplaceWithStructure()
	px, py := s.planet.Center()
	s.image, s.Bounds = s.generateImage(px, py, ui.NonFocusColor)
	s.highlightedImage, _ = s.generateImage(px, py, ui.FocusedColor)
	s.berths, s.ships, s.inFlight = s.data.Berths, s.data.Berths, 0
	if s.Class() == Tax {
		s.ships = s.workers
	}

	s.storage = make(map[int]*Storage)

	for _, st := range s.data.Storage {
		s.storage[st.Resource] = &Storage{
			Resource: st.Resource,
			Capacity: st.Capacity,
			Amount:   st.Amount,
		}
	}

	s.displayOpts.GeoM.Translate(float64(s.Bounds.Min.X), float64(s.Bounds.Min.Y))

	s.adjustPopulationCapacity()

	return &s
}
