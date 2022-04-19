package planet

import (
	"fmt"
	"image/color"
	"math/rand"

	"golang.org/x/image/font"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/jcgraybill/ship-shape/resource"
	"github.com/jcgraybill/ship-shape/ui"
)

const (
	basePlanetRadius = 8
)

type Planet struct {
	x, y             int
	resources        map[int]uint8
	resourceData     [resource.ResourceDataLength]resource.ResourceData
	highlighted      bool
	image            *ebiten.Image
	highlightedImage *ebiten.Image
	name             string
	displayOpts      *ebiten.DrawImageOptions
	ttf              font.Face
	visible          bool
}

func New(x, y int, resources map[int]uint8, resourceData [resource.ResourceDataLength]resource.ResourceData) *Planet {
	var p Planet

	p.x, p.y = x, y
	p.visible = true

	p.resources = resources
	p.resourceData = resourceData

	p.name = p.generateName()

	p.image, p.highlightedImage = p.generatePlanetImages()

	p.highlighted = false

	p.displayOpts = &ebiten.DrawImageOptions{}
	p.displayOpts.GeoM.Translate(float64(p.x-ui.PlanetSize/2), float64(p.y-ui.PlanetSize/2))
	p.ttf = ui.Font(ui.TtfRegular)
	return &p
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

func (p *Planet) Center() (int, int) {
	return p.x, p.y
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

func (p *Planet) generateName() string {
	var n string
	letters := []string{"alpha", "beta", "gamma", "delta", "epsilon", "zêta", "êta", "thêta", "iota", "kappa", "lambda", "mu", "nu", "xi", "omikron", "pi", "rho", "sigma", "tau", "upsilon", "phi", "chi", "psi", "omega"}
	n = fmt.Sprintf("%s-%d", letters[rand.Intn(len(letters))], rand.Intn(1000))
	return n
}

func (p *Planet) Name() string {
	return p.name
}

func (p *Planet) Draw(image *ebiten.Image) {
	if p.visible {
		image.DrawImage(p.Image(), p.displayOpts)
		cx, cy := p.Center()
		textBounds := text.BoundString(p.ttf, p.Name())
		text.Draw(image, p.Name(), p.ttf, cx-textBounds.Dx()/2, cy-16, ui.FocusedColor)
	}
}

func (p *Planet) ReplaceWithStructure() {
	p.Unhighlight()
	p.visible = false
}

func (p *Planet) Resources() map[int]uint8 {
	return p.resources
}

func (p *Planet) In(x, y int) bool {
	if p.x-ui.PlanetSize/2 < x && p.x+ui.PlanetSize/2 > x {
		if p.y-ui.PlanetSize/2 < y && p.y+ui.PlanetSize/2 > y {
			return true
		}
	}
	return false
}
