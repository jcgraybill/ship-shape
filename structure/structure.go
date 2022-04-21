package structure

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/jcgraybill/ship-shape/planet"
	"github.com/jcgraybill/ship-shape/resource"
	"github.com/jcgraybill/ship-shape/ui"
)

type Structure struct {
	x, y, w, h       int
	highlighted      bool
	image            *ebiten.Image
	highlightedImage *ebiten.Image
	displayOpts      *ebiten.DrawImageOptions
	planet           *planet.Planet
	data             StructureData
	storage          map[int]Storage
	resourcesWanted  []int
	berths, ships    int
	workers          int
	income           int
	structureType    int
}

func New(structureType int, sd StructureData, p *planet.Planet) *Structure {
	s := Structure{
		data:          sd,
		planet:        p,
		highlighted:   false,
		workers:       0,
		displayOpts:   &ebiten.DrawImageOptions{},
		structureType: structureType,
	}

	s.planet.ReplaceWithStructure()
	px, py := s.planet.Center()
	s.image, s.x, s.y, s.w, s.h = s.generateImage(px, py, ui.NonFocusColor)
	s.highlightedImage, _, _, _, _ = s.generateImage(px, py, ui.FocusedColor)
	s.berths, s.ships = s.data.Berths, s.data.Berths
	s.storage = make(map[int]Storage)

	s.resourcesWanted = make([]int, 0)
	for _, st := range s.data.Storage {
		s.storage[st.Resource] = Storage{
			Resource: st.Resource,
			Capacity: st.Capacity,
			Amount:   st.Amount,
		}

		if st.Resource != s.data.Produces.Resource {
			s.resourcesWanted = append(s.resourcesWanted, st.Resource)
		}
	}

	s.displayOpts.GeoM.Translate(float64(s.x), float64(s.y))

	if s.structureType == Outpost && s.storage[resource.Population].Capacity > 0 {
		cap := float64(s.storage[resource.Population].Capacity) * (float64(s.planet.Resources()[resource.Habitability]) / 255)
		s.storage[resource.Population] = Storage{
			Resource: resource.Population,
			Amount:   s.storage[resource.Population].Amount,
			Capacity: uint8(math.Ceil(cap)),
		}
	}

	return &s
}

func (s *Structure) generateImage(planetCenterX, planetCenterY int, uiColor color.Color) (*ebiten.Image, int, int, int, int) {
	var x, y, w, h int
	ttf := ui.Font(ui.TtfRegular)
	textBounds := text.BoundString(ttf, s.data.DisplayName)
	contentWidth := textBounds.Dx()
	if contentWidth < ui.PlanetSize {
		contentWidth = ui.PlanetSize
	}
	image := ebiten.NewImage(ui.Border+ui.Buffer+contentWidth+ui.Buffer+ui.Border, ui.Border+ui.Buffer+textBounds.Dy()+ui.Buffer+ui.PlanetSize+ui.Buffer+ui.Border)
	image.Fill(uiColor)

	interior := ebiten.NewImage(ui.Buffer+contentWidth+ui.Buffer, ui.Buffer+textBounds.Dy()+ui.Buffer+ui.PlanetSize+ui.Buffer)
	interior.Fill(ui.BackgroundColor)
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(ui.Border, ui.Border)
	text.Draw(interior, s.data.DisplayName, ttf, ui.Buffer, int(ttf.Metrics().Ascent/ui.DPI)+ui.Buffer, uiColor)

	popts := &ebiten.DrawImageOptions{}
	popts.GeoM.Translate(float64(ui.Buffer), float64(ui.Buffer+textBounds.Dy()+ui.Buffer))
	interior.DrawImage(s.planet.Image(), popts)

	image.DrawImage(interior, opts)

	x = planetCenterX - ui.PlanetSize/2 - ui.Buffer - ui.Border
	y = planetCenterY - ui.PlanetSize/2 - ui.Buffer - textBounds.Dy() - ui.Buffer - ui.Border
	w = ui.Border + ui.Buffer + contentWidth + ui.Buffer + ui.Border
	h = ui.Border + ui.Buffer + textBounds.Dy() + ui.Buffer + ui.PlanetSize + ui.Buffer + ui.Border
	return image, x, y, w, h
}

func (s *Structure) MouseButton(x, y int) bool {
	if s.x < x && s.x+s.w > x {
		if s.y < y && s.y+s.h > y {
			return true
		}
	}
	return false
}

func (s *Structure) Draw(image *ebiten.Image) {
	if s.IsHighlighted() {
		image.DrawImage(s.highlightedImage, s.displayOpts)
	} else {
		image.DrawImage(s.image, s.displayOpts)
	}
}

func (s *Structure) Name() string {
	return s.data.DisplayName
}

func (s *Structure) Planet() *planet.Planet {
	return s.planet
}

func (s *Structure) Storage() map[int]Storage {
	return s.storage
}

func (s *Structure) Highlight() {
	s.highlighted = true
}

func (s *Structure) Unhighlight() {
	s.highlighted = false
}

func (s *Structure) IsHighlighted() bool {
	return s.highlighted
}

func (s *Structure) Produces() int {
	return s.data.Produces.Resource
}

func (s *Structure) HasShips() bool {
	if s.structureType == Capitol {
		if s.ships > 0 && s.workers > 0 && s.ships <= s.workers {
			return true
		}
	}

	if s.ships > 0 {
		return true
	}
	return false
}

func (s *Structure) Workers() int {
	return s.workers
}
func (s *Structure) WorkersNeeded() int {
	return s.data.Workers
}

func (s *Structure) AssignWorkers(workers int) {
	s.workers = workers
}

func (s *Structure) LaborCost() int {
	return s.workers * s.data.WorkerCost
}

func (s *Structure) WorkerCost() int {
	return s.data.WorkerCost
}

func (s *Structure) CanProduce() bool {
	if s.structureType == Capitol {
		return true
	}
	if s.Storage()[s.data.Produces.Resource].Amount < s.Storage()[s.data.Produces.Resource].Capacity {
		return true
	}
	return false
}

func (s *Structure) Income() int {
	return s.income
}

func (s *Structure) StructureType() int {
	return s.structureType
}

func (s *Structure) CollectIncome() int {
	income := s.income
	s.income = 0
	return income
}

func (s *Structure) Upgradeable() int {
	if s.data.Upgrade.Structure > 0 {
		upgradeable := true
		for _, r := range s.data.Upgrade.Required {
			if s.storage[r].Amount < s.storage[r].Capacity {
				upgradeable = false
			}
		}
		if upgradeable {
			return s.data.Upgrade.Structure
		}
	}
	return 0
}

func (s *Structure) Upgrade(st int, sd StructureData) {
	s.structureType = st
	s.data = sd
	s.image, _, _, s.w, s.h = s.generateImage(s.x, s.x, ui.NonFocusColor)
	s.highlightedImage, _, _, _, _ = s.generateImage(s.x, s.x, ui.FocusedColor)
	s.berths, s.ships = sd.Berths, sd.Berths

	storage := make(map[int]Storage)

	s.resourcesWanted = make([]int, 0)
	for _, st := range s.data.Storage {
		storage[st.Resource] = Storage{
			Resource: st.Resource,
			Capacity: st.Capacity,
			Amount:   s.storage[st.Resource].Amount,
		}

		if st.Resource != s.data.Produces.Resource {
			s.resourcesWanted = append(s.resourcesWanted, st.Resource)
		}
	}

	s.storage = storage
}
