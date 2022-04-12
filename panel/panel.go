package panel

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jcgraybill/ship-shape/panel/label"
	"github.com/jcgraybill/ship-shape/planet"
	"github.com/jcgraybill/ship-shape/util"
	"golang.org/x/image/font"
)

const (
	panelWidth   = 160
	panelPadding = 10
	border       = 1
)

type Panel struct {
	x, y    int
	w, h    int
	image   *ebiten.Image
	display *ebiten.DrawImageOptions

	interior *ebiten.Image
	intOpts  *ebiten.DrawImageOptions

	ttf   font.Face
	label *label.Label
}

func New(w, h int) *Panel {
	var p Panel
	p.x = w - panelWidth - panelPadding
	p.y = panelPadding
	p.w = panelWidth
	p.h = h - panelPadding*2
	p.image = generateImage(p.w, p.h)
	p.display = &ebiten.DrawImageOptions{}
	p.interior = ebiten.NewImage(p.w-border*2, p.h-border*2)
	p.intOpts = &ebiten.DrawImageOptions{}
	p.intOpts.GeoM.Translate(float64(border), float64(border))
	p.ttf = util.Font()
	p.display.GeoM.Translate(float64(p.x), float64(p.y))
	return &p
}

func generateImage(w, h int) *ebiten.Image {
	img := ebiten.NewImage(w, h)
	img.Fill(color.White)
	return img
}

func (p *Panel) Image() *ebiten.Image {
	p.interior.Fill(color.Black)

	if p.label != nil {
		p.label.Draw(p.interior)

	}
	p.image.DrawImage(p.interior, p.intOpts)
	return p.image
}

func (p *Panel) Location() *ebiten.DrawImageOptions {
	return p.display
}

func (p *Panel) MouseButton(x, y int) bool {
	if p.x < x && p.x+p.w > x {
		if p.y < y && p.y+p.h > y {
			fmt.Println(fmt.Sprintf("panel %d %d", x-p.x, y-p.y))
			if p.label != nil {
				p.label.MouseButton(x-p.x, y-p.y)
			}
			return true
		}
	}
	return false
}

func (p *Panel) ShowPlanet(planet *planet.Planet) {
	p.label = label.New(4, 16, p.w-4, p.h-16, fmt.Sprintf("planet: %s\ngravity: %d\nwater: %d", planet.Name(), planet.Gravity, planet.Water))
}

func (p *Panel) Clear() {
	p.label = nil
}
