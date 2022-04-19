package ship

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jcgraybill/ship-shape/ui"
)

const shipSize = 16

type Ship struct {
	x, y  int
	image *ebiten.Image
	opts  *ebiten.DrawImageOptions
}

func New(x, y int) *Ship {
	s := Ship{
		x: x,
		y: y,
	}

	s.image = ebiten.NewImage(shipSize, shipSize)
	s.image.Fill(color.Black)
	v, i := ui.Triangle(0, 0, shipSize, shipSize, color.RGBA{0xcc, 0xcc, 0xcc, 0xff})
	s.image.DrawTriangles(v, i, ui.Src, nil)

	s.opts = &ebiten.DrawImageOptions{}
	s.opts.GeoM.Translate(float64(s.x), float64(s.y))

	return &s
}

func (s *Ship) Draw(image *ebiten.Image) {
	image.DrawImage(s.image, s.opts)

}

func (s *Ship) MouseButton(x, y int) bool {
	if s.x < x && s.x+shipSize > x {
		if s.y < y && s.y+shipSize > y {
			return true
		}
	}
	return false
}
