package bar

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jcgraybill/ship-shape/ui"
)

type Bar struct {
	Bounds image.Rectangle
	value  uint8

	bar  *ebiten.Image
	fill *ebiten.Image
	opts *ebiten.DrawImageOptions
}

func New(x, y, w int, value uint8, barColor color.RGBA) *Bar {
	var b Bar
	b.value = value
	b.Bounds = image.Rect(x, y, x+w, y+ui.BarHeight)

	b.fill = ebiten.NewImage(b.Bounds.Dx(), b.Bounds.Dy())
	b.bar = ebiten.NewImage(b.Bounds.Dx(), b.Bounds.Dy())

	b.bar.Fill(ui.NonFocusColor)
	b.fill.Fill(barColor)

	barWidth := (int(value) * b.Bounds.Dx()) / 255

	if barWidth > 0 {
		b.bar.DrawImage(b.fill.SubImage(image.Rect(0, 0, barWidth, ui.BarHeight)).(*ebiten.Image), nil)
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
	return b.bar, b.opts
}

func (b *Bar) Height() int {
	return b.Bounds.Dy()
}

func (b *Bar) UpdateValue(value uint8) {
	if value != b.value {
		b.bar.Fill(ui.NonFocusColor)
		barWidth := (int(value) * b.Bounds.Dx()) / 255
		if barWidth > 0 {
			b.bar.DrawImage(b.fill.SubImage(image.Rect(0, 0, barWidth, ui.BarHeight)).(*ebiten.Image), nil)
		}
		b.value = value
	}
}

func (b *Bar) UpdateText(string) { return }
