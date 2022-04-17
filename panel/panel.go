package panel

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jcgraybill/ship-shape/panel/button"
	"github.com/jcgraybill/ship-shape/panel/label"
	"github.com/jcgraybill/ship-shape/ui"
)

const (
	panelWidth   = 160
	panelPadding = 10
)

type Panel struct {
	x, y    int
	w, h    int
	image   *ebiten.Image
	display *ebiten.DrawImageOptions

	interior *ebiten.Image
	intOpts  *ebiten.DrawImageOptions

	elements []widget
}

type widget interface {
	LeftMouseButtonPress(int, int) bool
	LeftMouseButtonRelease(int, int) bool
	Draw() (*ebiten.Image, *ebiten.DrawImageOptions)
	Height() int
}

func New(w, h int) *Panel {
	var p Panel
	p.x = w - panelWidth - panelPadding
	p.y = panelPadding
	p.w = panelWidth
	p.h = h - panelPadding*2
	p.image = generateImage(p.w, p.h)
	p.display = &ebiten.DrawImageOptions{}
	p.interior = ebiten.NewImage(p.w-ui.Border*2, p.h-ui.Border*2)
	p.intOpts = &ebiten.DrawImageOptions{}
	p.intOpts.GeoM.Translate(float64(ui.Border), float64(ui.Border))
	p.display.GeoM.Translate(float64(p.x), float64(p.y))
	p.elements = make([]widget, 0)
	return &p
}

func generateImage(w, h int) *ebiten.Image {
	img := ebiten.NewImage(w, h)
	img.Fill(color.White)
	return img
}

func (p *Panel) Image() *ebiten.Image {
	p.interior.Fill(color.Black)

	for _, ui := range p.elements {
		p.interior.DrawImage(ui.Draw())
	}
	p.image.DrawImage(p.interior, p.intOpts)
	return p.image
}

func (p *Panel) Location() *ebiten.DrawImageOptions {
	return p.display
}

func (p *Panel) LeftMouseButtonPress(x, y int) bool {
	if p.x < x && p.x+p.w > x {
		if p.y < y && p.y+p.h > y {
			for _, widget := range p.elements {
				widget.LeftMouseButtonPress(x-p.x, y-p.y)
			}
			return true
		}
	}
	return false
}

func (p *Panel) LeftMouseButtonRelease(x, y int) bool {
	if p.x < x && p.x+p.w > x {
		if p.y < y && p.y+p.h > y {
			for _, widget := range p.elements {
				widget.LeftMouseButtonRelease(x-p.x, y-p.y)
			}
			return true
		}
	}
	for _, widget := range p.elements {
		widget.LeftMouseButtonRelease(-1, -1)
	}
	return false
}

func (p *Panel) AddLabel(text string) {
	p.elements = append(p.elements, label.New(ui.Buffer, p.firstAvailableSpot(), p.w-ui.Buffer*2, p.h-ui.Buffer*2, text))
}

func (p *Panel) AddButton(text string, callback func()) {
	p.elements = append(p.elements, button.New(ui.Buffer, p.firstAvailableSpot(), p.w-ui.Buffer*2, p.h-ui.Buffer*2, text, callback))
}

func (p *Panel) firstAvailableSpot() int {
	i := ui.Buffer
	for _, element := range p.elements {
		i += element.Height()
		i += ui.Buffer * 4
	}
	return i
}

func (p *Panel) Clear() {
	p.elements = nil
}
