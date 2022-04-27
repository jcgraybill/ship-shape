package panel

import (
	"image"
	"image/color"

	"github.com/fogleman/gg"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jcgraybill/ship-shape/panel/bar"
	"github.com/jcgraybill/ship-shape/panel/button"
	"github.com/jcgraybill/ship-shape/panel/divider"
	"github.com/jcgraybill/ship-shape/panel/invertedLabel"
	"github.com/jcgraybill/ship-shape/panel/label"
	"github.com/jcgraybill/ship-shape/ui"
)

type Panel struct {
	Bounds         image.Rectangle
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
	p.Bounds = image.Rect(w-ui.PanelWidth-ui.PanelExternalPadding, ui.PanelExternalPadding, w-ui.PanelExternalPadding, h-ui.PanelExternalPadding)

	p.background = p.createBackgroundImage(p.Bounds.Dx(), p.Bounds.Dy())
	p.displayOptions = &ebiten.DrawImageOptions{}
	p.interior = ebiten.NewImage(p.Bounds.Dx()-ui.Border*2, p.Bounds.Dy()-ui.Border*2)
	p.interiorDisplayOptions = &ebiten.DrawImageOptions{}
	p.interiorDisplayOptions.GeoM.Translate(float64(p.Bounds.Min.X+ui.Border), float64(p.Bounds.Min.Y+ui.Border))
	p.displayOptions.GeoM.Translate(float64(p.Bounds.Min.X), float64(p.Bounds.Min.Y))
	p.elements = make([]widget, 0)
	return &p
}

func (p *Panel) createBackgroundImage(w, h int) *ebiten.Image {
	dc := gg.NewContext(w, h)
	dc.SetRGB255(int(ui.FocusedColor.R), int(ui.FocusedColor.G), int(ui.FocusedColor.B))
	dc.DrawRoundedRectangle(0, 0, float64(w), float64(h), 10)
	dc.Fill()
	dc.SetRGB255(int(ui.BackgroundColor.R), int(ui.BackgroundColor.G), int(ui.BackgroundColor.B))
	dc.DrawRoundedRectangle(ui.Border, ui.Border, float64(w-ui.Border*2), float64(h-ui.Border*2), 10)
	dc.Fill()

	return ebiten.NewImageFromImage(dc.Image())
}

func (p *Panel) Draw(image *ebiten.Image) {
	p.interior.Clear()
	for _, ui := range p.elements {
		p.interior.DrawImage(ui.Draw())
	}

	image.DrawImage(p.background, p.displayOptions)
	image.DrawImage(p.interior, p.interiorDisplayOptions)

}

func (p *Panel) LeftMouseButtonPress(x, y int) bool {
	if p.Bounds.At(x, y) == color.Opaque {
		for _, widget := range p.elements {
			widget.LeftMouseButtonPress(x-p.Bounds.Min.X, y-p.Bounds.Min.Y)
		}
		return true
	}
	return false
}

func (p *Panel) LeftMouseButtonRelease(x, y int) bool {
	if p.Bounds.At(x, y) == color.Opaque {
		for _, widget := range p.elements {
			widget.LeftMouseButtonRelease(x-p.Bounds.Min.X, y-p.Bounds.Min.Y)
		}
		return true
	}

	for _, widget := range p.elements {
		widget.LeftMouseButtonRelease(-1, -1)
	}
	return false
}

func (p *Panel) AddInvertedLabel(text string, style string) {
	p.elements = append(p.elements, invertedLabel.New(0, p.firstAvailableSpot(), p.Bounds.Dx(), p.Bounds.Dy()-p.firstAvailableSpot()-ui.Buffer*2, text, style))
}

func (p *Panel) AddLabel(text string, style string) {
	p.elements = append(p.elements, label.New(ui.Buffer, p.firstAvailableSpot(), p.Bounds.Dx()-ui.Buffer*2-ui.Border*2, p.Bounds.Dy()-p.firstAvailableSpot()-ui.Buffer*2, text, style))
}

func (p *Panel) AddButton(text string, callback func()) *button.Button {
	b := button.New(ui.Buffer, p.firstAvailableSpot(), p.Bounds.Dx()-ui.Buffer*2-ui.Border*2, p.Bounds.Dy()-p.firstAvailableSpot()-ui.Buffer*2, text, callback)
	p.elements = append(p.elements, b)
	return b
}

func (p *Panel) AddBar(value uint8, color color.RGBA) {
	p.elements = append(p.elements, bar.New(ui.Buffer, p.firstAvailableSpot(), p.Bounds.Dx()-ui.Buffer*2-ui.Border*2, value, color))
}

func (p *Panel) AddDivider() {
	p.elements = append(p.elements, divider.New(0, p.firstAvailableSpot(), p.Bounds.Dx()))
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
	p.Bounds = image.Rect(w-ui.PanelWidth-ui.PanelExternalPadding, ui.PanelExternalPadding, w-ui.PanelExternalPadding, h-ui.PanelExternalPadding)

	p.displayOptions.GeoM.Reset()
	p.displayOptions.GeoM.Translate(float64(p.Bounds.Min.X), float64(p.Bounds.Min.Y))
	p.interiorDisplayOptions.GeoM.Reset()
	p.interiorDisplayOptions.GeoM.Translate(float64(p.Bounds.Min.X+ui.Border), float64(p.Bounds.Min.Y+ui.Border))

	p.background = p.createBackgroundImage(p.Bounds.Dx(), p.Bounds.Dy())

	p.interior = ebiten.NewImage(p.Bounds.Dx()-ui.Border*2, p.Bounds.Dy()-ui.Border*2)
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
