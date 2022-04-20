package bar

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jcgraybill/ship-shape/ui"
)

type Bar struct {
	x, y int
	w, h int

	image   *ebiten.Image
	opts    *ebiten.DrawImageOptions
	bar     *ebiten.Image
	barOpts *ebiten.DrawImageOptions
}

func New(x, y, w int, value uint8, barColor color.RGBA) *Bar {
	b := Bar{
		x:       x,
		y:       y,
		w:       w,
		h:       ui.BarHeight,
		barOpts: &ebiten.DrawImageOptions{},
	}

	b.image = ebiten.NewImage(b.w, b.h)
	b.image.Fill(ui.NonFocusColor)

	barWidth := (int(value) * b.w) / 255
	b.bar = ebiten.NewImage(1, b.h)
	b.bar.Fill(barColor)
	if barWidth > 0 {
		b.barOpts.GeoM.Scale(float64(barWidth), 1)
		b.image.DrawImage(b.bar, b.barOpts)
	}
	b.opts = &ebiten.DrawImageOptions{}
	b.opts.GeoM.Translate(float64(b.x), float64(b.y))

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
	return b.h
}

func (b *Bar) UpdateValue(value uint8) {
	b.image.Fill(ui.NonFocusColor)
	barWidth := (int(value) * b.w) / 255

	if barWidth > 0 {
		b.barOpts.GeoM.Reset()
		b.barOpts.GeoM.Scale(float64(barWidth), 1)
		b.image.DrawImage(b.bar, b.barOpts)
	}

}

func (b *Bar) UpdateText(string) { return }
