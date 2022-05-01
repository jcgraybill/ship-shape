package label

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/jcgraybill/ship-shape/ui"
	"golang.org/x/image/font"
)

type Label struct {
	bounds   image.Rectangle
	ttf      *font.Face
	message  string
	source   func() string
	inverted bool

	image *ebiten.Image
	opts  *ebiten.DrawImageOptions
}

func New(x, y, w, h int, style string, inverted bool, source func() string) *Label {
	l := Label{
		ttf:      ui.Font(style),
		inverted: inverted,
		source:   source,
	}

	l.message = l.source()
	textBounds := text.BoundString(*(l.ttf), l.message)
	l.bounds = image.Rect(x, y, x+w, y+textBounds.Dy()+int((*l.ttf).Metrics().Descent/ui.DPI))

	return &l
}

func (l *Label) Draw() (*ebiten.Image, *ebiten.DrawImageOptions) {
	if l.image == nil {
		l.createImages()
	}
	if newMessage := l.source(); newMessage != l.message {
		l.message = newMessage
		l.updateText()
	}
	return l.image, l.opts
}

func (l *Label) Height() int {
	return l.bounds.Dy()
}

func (l *Label) LeftMouseButtonPress(x, y int) bool {
	return false
}

func (l *Label) LeftMouseButtonRelease(x, y int) bool {
	return false
}

func (l *Label) createImages() {
	l.image = ebiten.NewImage(l.bounds.Dx(), l.bounds.Dy())
	l.updateText()
	l.opts = &ebiten.DrawImageOptions{}
	l.opts.GeoM.Translate(float64(l.bounds.Min.X), float64(l.bounds.Min.Y))
}

func (l *Label) updateText() {
	if l.inverted {
		l.image.Fill(ui.FocusedColor)
		text.Draw(l.image, l.message, *l.ttf, 0, int((*l.ttf).Metrics().Ascent/ui.DPI), ui.BackgroundColor)
	} else {
		l.image.Fill(ui.BackgroundColor)
		text.Draw(l.image, l.message, *l.ttf, 0, int((*l.ttf).Metrics().Ascent/ui.DPI), ui.FocusedColor)
	}
}
