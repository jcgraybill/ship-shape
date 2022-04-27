package divider

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jcgraybill/ship-shape/ui"
)

type Divider struct {
	Bounds image.Rectangle

	image *ebiten.Image
	opts  *ebiten.DrawImageOptions
}

func New(x, y, w int) *Divider {
	var d Divider
	d.Bounds = image.Rect(x, y, x+w, y+ui.Border)
	d.image = ebiten.NewImage(d.Bounds.Dx(), d.Bounds.Dy())
	d.image.Fill(ui.FocusedColor)

	d.opts = &ebiten.DrawImageOptions{}
	d.opts.GeoM.Translate(float64(d.Bounds.Min.X), float64(d.Bounds.Min.Y))

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
	return d.Bounds.Dy()
}

func (d *Divider) UpdateValue(uint8) { return }
func (d *Divider) UpdateText(string) { return }
