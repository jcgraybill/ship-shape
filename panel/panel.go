package panel

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jcgraybill/ship-shape/panel/bar"
	"github.com/jcgraybill/ship-shape/panel/button"
	"github.com/jcgraybill/ship-shape/panel/divider"
	"github.com/jcgraybill/ship-shape/panel/invertedLabel"
	"github.com/jcgraybill/ship-shape/panel/label"
	"github.com/jcgraybill/ship-shape/ui"
)

type Panel struct {
	x, y           int
	w, h           int
	locked         int
	background     *ebiten.Image
	displayOptions *ebiten.DrawImageOptions

	interior               *ebiten.Image
	interiorDisplayOptions *ebiten.DrawImageOptions
	elements               []widget
}

type widget interface {
	LeftMouseButtonPress(int, int) bool
	LeftMouseButtonRelease(int, int) bool
	Draw() (*ebiten.Image, *ebiten.DrawImageOptions)
	Height() int
	UpdateText(string)
	UpdateValue(uint8)
}

type updateableLabel interface {
}

func New(w, h int) *Panel {
	var p Panel
	p.x = w - ui.PanelWidth - ui.PanelExternalPadding
	p.y = ui.PanelExternalPadding
	p.w = ui.PanelWidth
	p.h = h - ui.PanelExternalPadding*2
	p.background = ebiten.NewImage(p.w, p.h)
	p.background.Fill(ui.FocusedColor)
	p.displayOptions = &ebiten.DrawImageOptions{}
	p.interior = ebiten.NewImage(p.w-ui.Border*2, p.h-ui.Border*2)
	p.interiorDisplayOptions = &ebiten.DrawImageOptions{}
	p.interiorDisplayOptions.GeoM.Translate(float64(ui.Border), float64(ui.Border))
	p.displayOptions.GeoM.Translate(float64(p.x), float64(p.y))
	p.elements = make([]widget, 0)
	return &p
}

func (p *Panel) Draw(image *ebiten.Image) {
	p.interior.Fill(ui.BackgroundColor)
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

func (p *Panel) AddInvertedLabel(text string, style string) {
	p.elements = append(p.elements, invertedLabel.New(0, p.firstAvailableSpot(), p.w, p.h-p.firstAvailableSpot()-ui.Buffer*2, text, style))
}

func (p *Panel) AddLabel(text string, style string) {
	p.elements = append(p.elements, label.New(ui.Buffer, p.firstAvailableSpot(), p.w-ui.Buffer*2-ui.Border*2, p.h-p.firstAvailableSpot()-ui.Buffer*2, text, style))
}

func (p *Panel) AddButton(text string, callback func()) {
	p.elements = append(p.elements, button.New(ui.Buffer, p.firstAvailableSpot(), p.w-ui.Buffer*2-ui.Border*2, p.h-p.firstAvailableSpot()-ui.Buffer*2, text, callback))
}

func (p *Panel) AddBar(value uint8, color color.RGBA) {
	p.elements = append(p.elements, bar.New(ui.Buffer, p.firstAvailableSpot(), p.w-ui.Buffer*2-ui.Border*2, value, color))
}

func (p *Panel) AddDivider() {
	p.elements = append(p.elements, divider.New(0, p.firstAvailableSpot(), p.w))
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
	p.elements = p.elements[:p.locked]
}

func (p *Panel) Resize(w, h int) {
	p.x = w - ui.PanelWidth - ui.PanelExternalPadding
	p.h = h - ui.PanelExternalPadding*2
	p.displayOptions.GeoM.Reset()
	p.displayOptions.GeoM.Translate(float64(p.x), float64(p.y))

	p.background = ebiten.NewImage(p.w, p.h)
	p.background.Fill(ui.FocusedColor)
	p.interior = ebiten.NewImage(p.w-ui.Border*2, p.h-ui.Border*2)
}

func (p *Panel) Lock(n int) {
	if len(p.elements) >= n {
		p.locked = n
	}
}

func (p *Panel) UpdateLabel(n int, newText string) {
	if n <= p.locked {
		p.elements[n].UpdateText(newText)
	}
}

func (p *Panel) UpdateBar(n int, newValue uint8) {
	if n <= p.locked {
		p.elements[n].UpdateValue(newValue)
	}
}
