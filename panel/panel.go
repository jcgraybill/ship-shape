package panel

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jcgraybill/ship-shape/panel/bar"
	"github.com/jcgraybill/ship-shape/panel/button"
	"github.com/jcgraybill/ship-shape/panel/label"
	"github.com/jcgraybill/ship-shape/ui"
)

type Panel struct {
	x, y           int
	w, h           int
	background     *ebiten.Image
	displayOptions *ebiten.DrawImageOptions

	interior               *ebiten.Image
	interiorDisplayOptions *ebiten.DrawImageOptions

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
	p.x = w - ui.PanelWidth - ui.PanelExternalPadding
	p.y = ui.PanelExternalPadding
	p.w = ui.PanelWidth
	p.h = h - ui.PanelExternalPadding*2
	p.background = ebiten.NewImage(p.w, p.h)
	p.background.Fill(color.White)
	p.displayOptions = &ebiten.DrawImageOptions{}
	p.interior = ebiten.NewImage(p.w-ui.Border*2, p.h-ui.Border*2)
	p.interiorDisplayOptions = &ebiten.DrawImageOptions{}
	p.interiorDisplayOptions.GeoM.Translate(float64(ui.Border), float64(ui.Border))
	p.displayOptions.GeoM.Translate(float64(p.x), float64(p.y))
	p.elements = make([]widget, 0)
	return &p
}

func (p *Panel) Draw(image *ebiten.Image) {
	p.interior.Fill(color.Black)
	for _, ui := range p.elements {
		p.interior.DrawImage(ui.Draw())
	}
	p.background.DrawImage(p.interior, p.interiorDisplayOptions)

	image.DrawImage(p.background, p.displayOptions)

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
	p.elements = append(p.elements, label.New(ui.Buffer, p.firstAvailableSpot(), p.w-ui.Buffer*2-ui.Border*2, p.h-p.firstAvailableSpot()-ui.Buffer*2, text))
}

func (p *Panel) AddButton(text string, callback func()) {
	p.elements = append(p.elements, button.New(ui.Buffer, p.firstAvailableSpot(), p.w-ui.Buffer*2-ui.Border*2, p.h-p.firstAvailableSpot()-ui.Buffer*2, text, callback))
}

func (p *Panel) AddBar(value uint8, color color.RGBA) {
	p.elements = append(p.elements, bar.New(ui.Buffer, p.firstAvailableSpot(), p.w-ui.Buffer*2-ui.Border*2, value, color))
}

func (p *Panel) firstAvailableSpot() int {
	i := ui.Buffer
	for _, element := range p.elements {
		i += element.Height()
		i += ui.Buffer
	}
	return i
}

func (p *Panel) Clear() {
	p.elements = nil
}
