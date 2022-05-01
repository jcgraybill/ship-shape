package bar

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jcgraybill/ship-shape/ui"
)

type Bar struct {
	bounds image.Rectangle
	value  uint8
	source func() uint8

	barColor color.RGBA
	bar      *ebiten.Image
	fill     *ebiten.Image
	opts     *ebiten.DrawImageOptions
}

func New(x, y, w int, source func() uint8, barColor color.RGBA) *Bar {
	b := Bar{
		source:   source,
		barColor: barColor,
		bounds:   image.Rect(x, y, x+w, y+ui.BarHeight),
	}
	return &b
}

func (b *Bar) Draw() (*ebiten.Image, *ebiten.DrawImageOptions) {
	if b.bar == nil || b.fill == nil {
		b.createImages()
	}
	if newValue := b.source(); newValue != b.value {
		b.value = newValue
		b.updateImages()
	}
	return b.bar, b.opts
}

func (b *Bar) Height() int {
	return b.bounds.Dy()
}

func (b *Bar) LeftMouseButtonPress(int, int) bool {
	return false
}
func (b *Bar) LeftMouseButtonRelease(int, int) bool {
	return false
}

func (b *Bar) createImages() {
	b.fill = ebiten.NewImage(b.bounds.Dx(), b.bounds.Dy())
	b.bar = ebiten.NewImage(b.bounds.Dx(), b.bounds.Dy())

	b.bar.Fill(ui.NonFocusColor)
	b.fill.Fill(b.barColor)
	b.opts = &ebiten.DrawImageOptions{}
	b.opts.GeoM.Translate(float64(b.bounds.Min.X), float64(b.bounds.Min.Y))
}

func (b *Bar) updateImages() {
	b.bar.Fill(ui.NonFocusColor)
	fillWidth := (int(b.value) * b.bounds.Dx()) / 255
	if fillWidth > 0 {
		b.bar.DrawImage(b.fill.SubImage(image.Rect(0, 0, fillWidth, ui.BarHeight)).(*ebiten.Image), nil)
	}
}
