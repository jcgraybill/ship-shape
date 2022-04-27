package structure

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/jcgraybill/ship-shape/ui"
)

func (s *Structure) Draw(image *ebiten.Image) {
	if s.IsHighlighted() {
		image.DrawImage(s.highlightedImage, s.displayOpts)
	} else {
		image.DrawImage(s.image, s.displayOpts)
	}
}

func (s *Structure) MouseButton(x, y int) bool {
	if s.Bounds.At(x, y) == color.Opaque {
		return true
	}
	return false
}

func (s *Structure) Highlight() {
	s.highlighted = true
	s.planet.Highlight()
}

func (s *Structure) Unhighlight() {
	s.highlighted = false
	s.planet.Unhighlight()

}

func (s *Structure) IsHighlighted() bool {
	return s.highlighted
}

func (s *Structure) generateImage(planetCenterX, planetCenterY int, uiColor color.Color) (*ebiten.Image, image.Rectangle) {
	var x, y, w, h int
	ttf := ui.Font(ui.TtfRegular)
	textBounds := text.BoundString(*ttf, s.data.DisplayName)
	contentWidth := textBounds.Dx()
	if contentWidth < ui.PlanetSize {
		contentWidth = ui.PlanetSize
	}

	w = ui.Border + ui.Buffer + contentWidth + ui.Buffer + ui.Border
	h = ui.Border + ui.Buffer + textBounds.Dy() + ui.Buffer + ui.PlanetSize + ui.Buffer + ui.Border

	structureImage := ebiten.NewImage(w, h)
	structureImage.Fill(uiColor)

	interior := ebiten.NewImage(w-ui.Border*2, h-ui.Border*2)
	interior.Fill(ui.BackgroundColor)
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(ui.Border, ui.Border)
	text.Draw(interior, s.data.DisplayName, *ttf, ui.Buffer, int((*ttf).Metrics().Ascent/ui.DPI)+ui.Buffer, uiColor)
	structureImage.DrawImage(interior, opts)

	x = planetCenterX - w/2
	y = planetCenterY - ui.PlanetSize/2 - ui.Buffer - textBounds.Dy() - ui.Buffer - ui.Border
	return structureImage, image.Rect(x, y, x+w, y+h)
}
