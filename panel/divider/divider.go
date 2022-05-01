package divider

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jcgraybill/ship-shape/ui"
)

type Divider struct {
	bounds image.Rectangle

	image *ebiten.Image
	opts  *ebiten.DrawImageOptions
}

func New(x, y, w int) *Divider {
	d := Divider{
		bounds: image.Rect(x, y, x+w, y+ui.Border),
	}

	return &d
}

func (d *Divider) Draw() (*ebiten.Image, *ebiten.DrawImageOptions) {
	if d.image == nil {
		d.createImages()
	}
	return d.image, d.opts
}

func (d *Divider) Height() int {
	return d.bounds.Dy()
}

func (d *Divider) LeftMouseButtonPress(int, int) bool {
	return false
}
func (d *Divider) LeftMouseButtonRelease(int, int) bool {
	return false
}

func (d *Divider) createImages() {
	d.image = ebiten.NewImage(d.bounds.Dx(), d.bounds.Dy())
	d.image.Fill(ui.FocusedColor)

	d.opts = &ebiten.DrawImageOptions{}
	d.opts.GeoM.Translate(float64(d.bounds.Min.X), float64(d.bounds.Min.Y))
}
