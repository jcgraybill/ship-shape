package planet

import (
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/jcgraybill/ship-shape/resource"
	"github.com/jcgraybill/ship-shape/ui"

	"github.com/fogleman/gg"
)

func (p *Planet) Draw(image *ebiten.Image) {
	image.DrawImage(p.blackImage, p.displayOpts)
	ui.ShaderOpts.Images[1] = p.Image()
	ui.ShaderOpts.GeoM.Reset()
	ui.ShaderOpts.GeoM.Translate(float64(p.Bounds.Min.X), float64(p.Bounds.Min.Y))
	image.DrawRectShader(ui.PlanetSize, ui.PlanetSize, ui.Shader, ui.ShaderOpts)

	if p.visible {
		cx, _ := p.Center()
		textBounds := text.BoundString(*(p.ttf), p.Name())
		if p.highlighted {
			text.Draw(image, p.Name(), *(p.ttf), cx-textBounds.Dx()/2, p.Bounds.Min.Y, ui.FocusedColor)
		} else {
			text.Draw(image, p.Name(), *(p.ttf), cx-textBounds.Dx()/2, p.Bounds.Min.Y, ui.NonFocusColor)
		}
	}
}

func (p *Planet) generatePlanetImages() (*ebiten.Image, *ebiten.Image, *ebiten.Image) {
	radius := float64(basePlanetRadius + rand.Intn(ui.PlanetSize/2-basePlanetRadius))

	var R, G, B float64
	rd := resource.GetResourceData()
	n := float64(len(p.resources) * 255)
	for resource, level := range p.resources {
		R = R + float64(level)*float64(rd[resource].Color.R)/n
		G = G + float64(level)*float64(rd[resource].Color.G)/n
		B = B + float64(level)*float64(rd[resource].Color.B)/n
	}

	dc := gg.NewContext(ui.PlanetSize, ui.PlanetSize)
	dc.DrawCircle(ui.PlanetSize/2, ui.PlanetSize/2, radius)
	dc.SetRGB255(0, 0, 0)
	dc.Fill()
	black := ebiten.NewImageFromImage(dc.Image())

	dc = gg.NewContext(ui.PlanetSize, ui.PlanetSize)
	dc.DrawCircle(ui.PlanetSize/2, ui.PlanetSize/2, radius+ui.Border)
	dc.SetRGB255(int(ui.NonFocusColor.R), int(ui.NonFocusColor.G), int(ui.NonFocusColor.B))
	dc.Fill()
	dc.DrawCircle(ui.PlanetSize/2, ui.PlanetSize/2, radius)
	dc.SetRGB255(int(R), int(G), int(B))
	dc.Fill()
	image := ebiten.NewImageFromImage(dc.Image())

	dc = gg.NewContext(ui.PlanetSize, ui.PlanetSize)
	dc.DrawCircle(ui.PlanetSize/2, ui.PlanetSize/2, radius+ui.Border)
	dc.SetRGB255(int(ui.FocusedColor.R), int(ui.FocusedColor.G), int(ui.FocusedColor.B))
	dc.Fill()
	dc.DrawCircle(ui.PlanetSize/2, ui.PlanetSize/2, radius)
	dc.SetRGB255(int(R), int(G), int(B))
	dc.Fill()
	highlighted := ebiten.NewImageFromImage(dc.Image())

	return image, highlighted, black
}

func (p *Planet) Image() *ebiten.Image {
	if p.highlighted {
		return p.highlightedImage
	} else {
		return p.image
	}
}

func (p *Planet) Highlight() {
	p.highlighted = true
}

func (p *Planet) Unhighlight() {
	p.highlighted = false
}

func (p *Planet) MouseButton(x, y int) bool {
	if p.visible && p.Bounds.At(x, y) == color.Opaque {
		return true
	}
	return false
}
