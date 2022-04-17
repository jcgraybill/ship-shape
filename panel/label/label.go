package label

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/jcgraybill/ship-shape/ui"
)

type Label struct {
	x, y int
	w, h int

	image *ebiten.Image
	opts  *ebiten.DrawImageOptions
}

func New(x, y, w, h int, message string) *Label {
	l := Label{
		x: x,
		y: y,
		w: w,
	}
	ttf := ui.Font()
	textBounds := text.BoundString(ttf, message)

	// FIXME text.BoundString underestimates this typeface's height by a few pixels
	l.h = textBounds.Dy() + int(ttf.Metrics().Descent/ui.DPI)

	image := ebiten.NewImage(l.w, l.h)
	image.Fill(color.Black)

	text.Draw(image, message, ttf, 0, int(ttf.Metrics().Ascent/ui.DPI), color.White)
	l.image = image

	l.opts = &ebiten.DrawImageOptions{}
	l.opts.GeoM.Translate(float64(l.x), float64(l.y))

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
	return l.h
}
