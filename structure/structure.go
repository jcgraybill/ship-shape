package structure

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/jcgraybill/ship-shape/planet"
	"github.com/jcgraybill/ship-shape/util"
)

const (
	buffer = 2
)

type Structure struct {
	x, y, w, h int
	image      *ebiten.Image
	display    *ebiten.DrawImageOptions
	planet     *planet.Planet
	data       StructureData
}

func New(sd StructureData, p *planet.Planet) *Structure {
	var s Structure
	s.data = sd
	s.planet = p
	s.image, s.x, s.y, s.w, s.h = s.generateImage(p.Center())

	s.display = &ebiten.DrawImageOptions{}
	s.display.GeoM.Translate(float64(s.x), float64(s.y))
	p.ReplaceWithStructure()
	return &s
}

func (s *Structure) generateImage(planetCenterX, planetCenterY int) (*ebiten.Image, int, int, int, int) {
	var x, y, w, h int
	ttf := util.Font()
	textBounds := text.BoundString(ttf, s.data.DisplayName)
	image := ebiten.NewImage(buffer+textBounds.Dx()+buffer, buffer+textBounds.Dy()+buffer+util.PlanetSize+buffer)
	image.Fill(color.White)

	interior := ebiten.NewImage(buffer/2+textBounds.Dx()+buffer/2, buffer/2+textBounds.Dy()+buffer+util.PlanetSize+buffer/2)
	interior.Fill(color.Black)
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(buffer/2, buffer/2)
	text.Draw(interior, s.data.DisplayName, ttf, buffer/2, int(ttf.Metrics().Ascent/util.DPI)+buffer/2, color.White)

	popts := &ebiten.DrawImageOptions{}
	popts.GeoM.Translate(float64(buffer/2), float64(buffer/2+textBounds.Dy()+buffer))
	interior.DrawImage(s.planet.Image(), popts)

	image.DrawImage(interior, opts)

	x = planetCenterX - util.PlanetSize/2 - buffer
	y = planetCenterY - util.PlanetSize/2 - buffer - textBounds.Dy() - buffer
	w = buffer + textBounds.Dx() + buffer
	h = buffer + textBounds.Dy() + buffer + util.PlanetSize + buffer
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
	image.DrawImage(s.Image(), s.Location())
}

func (s *Structure) Location() *ebiten.DrawImageOptions {
	return s.display
}

func (s *Structure) Image() *ebiten.Image {
	return s.image
}

func (s *Structure) Name() string {
	return s.data.DisplayName
}

func (s *Structure) Planet() *planet.Planet {
	return s.planet
}

func (s *Structure) Describe() string {
	return fmt.Sprintf("%s\nplanet: %s\nhabitability: %d\nwater: %d", s.Name(), s.Planet().Name(), s.Planet().Habitability, s.Planet().Water)
}
