package structure

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
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

		structureType: structureType,
	}

	s.createBounds()
	s.planet.ReplaceWithStructure()

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

	s.adjustPopulationCapacity()

	return &s
}

func (s *Structure) createBounds() {
	var w, h int
	ttf := ui.Font(ui.TtfRegular)
	textBounds := text.BoundString(*ttf, s.data.DisplayName)
	contentWidth := textBounds.Dx()
	if contentWidth < ui.PlanetSize {
		contentWidth = ui.PlanetSize
	}

	w = ui.Border + ui.Buffer + contentWidth + ui.Buffer + ui.Border
	h = ui.Border + ui.Buffer + textBounds.Dy() + ui.Buffer + ui.PlanetSize + ui.Buffer + ui.Border

	x, y := s.planet.Center()
	x = x - w/2
	y = y - ui.PlanetSize/2 - ui.Buffer - textBounds.Dy() - ui.Buffer - ui.Border

	s.Bounds = image.Rect(x, y, x+w, y+h)
}
