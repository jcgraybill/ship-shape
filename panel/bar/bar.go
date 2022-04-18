package bar

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jcgraybill/ship-shape/ui"
)

var neutralColor = color.RGBA{R: 0x7d, G: 0x7d, B: 0x7d, A: 0xff}

type Bar struct {
	x, y  int
	w, h  int
	value uint8
	color color.RGBA

	image *ebiten.Image
	opts  *ebiten.DrawImageOptions
}

func New(x, y, w int, value uint8, barColor color.RGBA) *Bar {
	b := Bar{
		x:     x,
		y:     y,
		w:     w,
		h:     ui.BarHeight,
		value: value,
		color: barColor,
	}

	b.image = ebiten.NewImage(b.w, b.h)
	b.image.Fill(neutralColor)

	barWidth := (int(b.value) * b.w) / 255
	if barWidth > 0 {
		bar := ebiten.NewImage(barWidth, b.h)
		bar.Fill(b.color)
		b.image.DrawImage(bar, nil)
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
