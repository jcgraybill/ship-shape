package label

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/jcgraybill/ship-shape/ui"
	"golang.org/x/image/font"
)

type Label struct {
	Bounds image.Rectangle
	ttf    *font.Face

	image *ebiten.Image
	opts  *ebiten.DrawImageOptions
}

func New(x, y, w, h int, message string, style string) *Label {
	var l Label
	l.ttf = ui.Font(style)

	// FIXME text.BoundString underestimates this typeface's height by a few pixels
	// or I'm using the wrong metric below to locate the dot position
	textBounds := text.BoundString(*(l.ttf), message)
	l.Bounds = image.Rect(x, y, x+w, y+textBounds.Dy()+int((*l.ttf).Metrics().Descent/ui.DPI))

	image := ebiten.NewImage(l.Bounds.Dx(), l.Bounds.Dy())
	image.Fill(ui.BackgroundColor)

	text.Draw(image, message, *l.ttf, 0, int((*l.ttf).Metrics().Ascent/ui.DPI), ui.FocusedColor)
	l.image = image

	l.opts = &ebiten.DrawImageOptions{}
	l.opts.GeoM.Translate(float64(l.Bounds.Min.X), float64(l.Bounds.Min.Y))

	return &l
}

func (l *Label) LeftMouseButtonPress(x, y int) bool {
	return false
}

func (l *Label) LeftMouseButtonRelease(x, y int) bool {
	return false
}

func (l *Label) Draw() (*ebiten.Image, *ebiten.DrawImageOptions) {
	return l.image, l.opts
}

func (l *Label) Height() int {
	return l.Bounds.Dy()
}

func (l *Label) UpdateValue(uint8) { return }
func (l *Label) UpdateText(newText string) {
	l.image.Fill(ui.BackgroundColor)
	text.Draw(l.image, newText, *l.ttf, 0, int((*l.ttf).Metrics().Ascent/ui.DPI), ui.FocusedColor)
}
