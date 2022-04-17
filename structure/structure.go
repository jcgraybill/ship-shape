package structure

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/jcgraybill/ship-shape/planet"
	"github.com/jcgraybill/ship-shape/ui"
)

type Structure struct {
	x, y, w, h  int
	image       *ebiten.Image
	displayOpts *ebiten.DrawImageOptions
	planet      *planet.Planet
	data        StructureData
}

func New(sd StructureData, p *planet.Planet) *Structure {
	var s Structure
	s.data = sd
	s.planet = p
	p.ReplaceWithStructure()
	s.image, s.x, s.y, s.w, s.h = s.generateImage(p.Center())

	s.displayOpts = &ebiten.DrawImageOptions{}
	s.displayOpts.GeoM.Translate(float64(s.x), float64(s.y))

	return &s
}

func (s *Structure) generateImage(planetCenterX, planetCenterY int) (*ebiten.Image, int, int, int, int) {
	var x, y, w, h int
	ttf := ui.Font()
	textBounds := text.BoundString(ttf, s.data.DisplayName)
	contentWidth := textBounds.Dx()
	if contentWidth < ui.PlanetSize {
		contentWidth = ui.PlanetSize
	}
	image := ebiten.NewImage(ui.Border+ui.Buffer+contentWidth+ui.Buffer+ui.Border, ui.Border+ui.Buffer+textBounds.Dy()+ui.Buffer+ui.PlanetSize+ui.Buffer+ui.Border)
	image.Fill(color.White)

	interior := ebiten.NewImage(ui.Buffer+contentWidth+ui.Buffer, ui.Buffer+textBounds.Dy()+ui.Buffer+ui.PlanetSize+ui.Buffer)
	interior.Fill(color.Black)
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(ui.Border, ui.Border)
	text.Draw(interior, s.data.DisplayName, ttf, ui.Buffer, int(ttf.Metrics().Ascent/ui.DPI)+ui.Buffer, color.White)

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
	image.DrawImage(s.image, s.displayOpts)
}

func (s *Structure) Describe() string {
	return fmt.Sprintf("%s\nplanet: %s\nhabitability: %d\nwater: %d", s.data.DisplayName, s.planet.Name(), s.planet.Habitability, s.planet.Water)
}