package bar

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jcgraybill/ship-shape/ui"
)

type Bar struct {
	Bounds image.Rectangle

	image   *ebiten.Image
	opts    *ebiten.DrawImageOptions
	bar     *ebiten.Image
	barOpts *ebiten.DrawImageOptions
}

func New(x, y, w int, value uint8, barColor color.RGBA) *Bar {
	var b Bar
	b.Bounds = image.Rect(x, y, x+w, y+ui.BarHeight)

	b.image = ebiten.NewImage(b.Bounds.Dx(), b.Bounds.Dy())
	b.image.Fill(ui.NonFocusColor)

	barWidth := (int(value) * b.Bounds.Dx()) / 255
	b.bar = ebiten.NewImage(1, b.Bounds.Dy())
	b.bar.Fill(barColor)
	b.barOpts = &ebiten.DrawImageOptions{}
	if barWidth > 0 {
		b.barOpts.GeoM.Scale(float64(barWidth), 1)
		b.image.DrawImage(b.bar, b.barOpts)
	}
	b.opts = &ebiten.DrawImageOptions{}
	b.opts.GeoM.Translate(float64(b.Bounds.Min.X), float64(b.Bounds.Min.Y))

	return &b
}

func (b *Bar) LeftMouseButtonPress(int, int) bool {
	return false
}
func (b *Bar) LeftMouseButtonRelease(int, int) bool {
	return false
}

func (b *Bar) Draw() (*ebiten.Image, *ebiten.DrawImageOptions) {
	return b.image, b.opts
}

func (b *Bar) Height() int {
	return b.Bounds.Dy()
}

func (b *Bar) UpdateValue(value uint8) {
	b.image.Fill(ui.NonFocusColor)
	barWidth := (int(value) * b.Bounds.Dx()) / 255

	if barWidth > 0 {
		b.barOpts.GeoM.Reset()
		b.barOpts.GeoM.Scale(float64(barWidth), 1)
		b.image.DrawImage(b.bar, b.barOpts)
	}

}

func (b *Bar) UpdateText(string) { return }
