package structure

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/jcgraybill/ship-shape/ui"
)

func (s *Structure) Draw(image *ebiten.Image) {
	if s.image == nil || s.highlightedImage == nil {
		s.image = s.generateImage(ui.NonFocusColor)
		s.highlightedImage = s.generateImage(ui.FocusedColor)
		s.displayOpts = &ebiten.DrawImageOptions{}
		s.displayOpts.GeoM.Translate(float64(s.Bounds.Min.X), float64(s.Bounds.Min.Y))
	}
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

func (s *Structure) generateImage(uiColor color.Color) *ebiten.Image {
	structureImage := ebiten.NewImage(s.Bounds.Dx(), s.Bounds.Dy())
	structureImage.Fill(uiColor)

	interior := ebiten.NewImage(s.Bounds.Dx()-ui.Border*2, s.Bounds.Dy()-ui.Border*2)
	interior.Fill(ui.BackgroundColor)
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(ui.Border, ui.Border)

	ttf := ui.Font(ui.TtfRegular)
	text.Draw(interior, s.data.DisplayName, *ttf, ui.Buffer, int((*ttf).Metrics().Ascent/ui.DPI)+ui.Buffer, uiColor)
	structureImage.DrawImage(interior, opts)

	return structureImage
}
