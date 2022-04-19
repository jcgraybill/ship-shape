package divider

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jcgraybill/ship-shape/ui"
)

type Divider struct {
	x, y int
	w, h int

	image *ebiten.Image
	opts  *ebiten.DrawImageOptions
}

func New(x, y, w int) *Divider {
	d := Divider{
		x: x,
		y: y,
		w: w,
		h: ui.Border,
	}

	d.image = ebiten.NewImage(d.w, d.h)
	d.image.Fill(color.White)

	d.opts = &ebiten.DrawImageOptions{}
	d.opts.GeoM.Translate(float64(d.x), float64(d.y))

	return &d
}

func (d *Divider) LeftMouseButtonPress(int, int) bool {
	return false
}
func (d *Divider) LeftMouseButtonRelease(int, int) bool {
	return false
}

func (d *Divider) Draw() (*ebiten.Image, *ebiten.DrawImageOptions) {
	return d.image, d.opts
}

func (d *Divider) Height() int {
	return d.h
}
