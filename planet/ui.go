package planet

import (
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/jcgraybill/ship-shape/ui"
)

func (p *Planet) Draw(image *ebiten.Image) {
	if p.visible {
		image.DrawImage(p.Image(), p.displayOpts)
		cx, cy := p.Center()
		textBounds := text.BoundString(p.ttf, p.Name())
		text.Draw(image, p.Name(), p.ttf, cx-textBounds.Dx()/2, cy-16, ui.FocusedColor)
	}
}

func (p *Planet) In(x, y int) bool {
	if p.x-ui.PlanetSize/2 < x && p.x+ui.PlanetSize/2 > x {
		if p.y-ui.PlanetSize/2 < y && p.y+ui.PlanetSize/2 > y {
			return true
		}
	}
	return false
}

func (p *Planet) generatePlanetImages() (*ebiten.Image, *ebiten.Image) {
	base := ebiten.NewImage(ui.PlanetSize, ui.PlanetSize)
	image := ebiten.NewImage(ui.PlanetSize, ui.PlanetSize)
	highlighted := ebiten.NewImage(ui.PlanetSize, ui.PlanetSize)

	radius := float32(basePlanetRadius + rand.Intn(ui.PlanetSize/2-basePlanetRadius))

	planetColor := color.RGBA{}
	var R, G, B int

	for resource, level := range p.resources {
		R = R + int(level)*int(p.resourceData[resource].Color.R)
		G = G + int(level)*int(p.resourceData[resource].Color.G)
		B = B + int(level)*int(p.resourceData[resource].Color.B)
	}

	n := len(p.resources) * 255

	planetColor.R = uint8(R / n)
	planetColor.G = uint8(G / n)
	planetColor.B = uint8(B / n)
	planetColor.A = 0xff

	v, i := ui.Circle(ui.PlanetSize/2, ui.PlanetSize/2, radius, planetColor)
	base.DrawTriangles(v, i, ui.Src, nil)

	radius = radius + ui.Border

	v, i = ui.Circle(ui.PlanetSize/2, ui.PlanetSize/2, radius, ui.NonFocusColor)
	image.DrawTriangles(v, i, ui.Src, nil)
	image.DrawImage(base, nil)

	v, i = ui.Circle(ui.PlanetSize/2, ui.PlanetSize/2, radius, ui.FocusedColor)
	highlighted.DrawTriangles(v, i, ui.Src, nil)
	highlighted.DrawImage(base, nil)

	return image, highlighted
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
	if p.visible && p.x-ui.PlanetSize/2 < x && p.x+ui.PlanetSize/2 > x {
		if p.y-ui.PlanetSize/2 < y && p.y+ui.PlanetSize/2 > y {
			return true
		}
	}
	return false
}
