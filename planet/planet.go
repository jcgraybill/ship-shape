package planet

import (
	"fmt"
	"image/color"
	"math/rand"

	"golang.org/x/image/font"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/jcgraybill/ship-shape/ui"
)

const (
	basePlanetRadius = 8
	basePlanetColor  = 0x7f
)

type Planet struct {
	x, y                int
	Water, Habitability uint8
	highlighted         bool
	image               *ebiten.Image
	highlightedImage    *ebiten.Image
	name                string
	displayOpts         *ebiten.DrawImageOptions
	ttf                 font.Face
	visible             bool
}

func New(x, y int, water, habitability uint8) *Planet {
	var p Planet

	p.x, p.y = x, y
	p.visible = true

	if water == 0 {
		p.Water = uint8(rand.Intn(255))
	} else {
		p.Water = water
	}

	if habitability == 0 {
		p.Habitability = uint8(rand.Intn(255))
	} else {
		p.Habitability = habitability
	}

	p.image = p.generatePlanetImage()
	p.highlightedImage = p.generateHighlightedImage()
	p.highlighted = false
	p.name = p.generateName()
	p.displayOpts = &ebiten.DrawImageOptions{}
	p.displayOpts.GeoM.Translate(float64(p.x-ui.PlanetSize/2), float64(p.y-ui.PlanetSize/2))
	p.ttf = ui.Font()
	return &p
}

func (p *Planet) generatePlanetImage() *ebiten.Image {
	image := ebiten.NewImage(ui.PlanetSize, ui.PlanetSize)

	radius := float32(basePlanetRadius + p.Habitability/32)

	waterColor := color.RGBA{R: 0x00, G: 0x00, B: 0xff, A: 0xff}

	planetColor := color.RGBA{}
	planetColor.R = basePlanetColor + uint8(int(waterColor.R)*int(p.Water)/(255*2))
	planetColor.G = basePlanetColor + uint8(int(waterColor.G)*int(p.Water)/(255*2))
	planetColor.B = basePlanetColor + uint8(int(waterColor.B)*int(p.Water)/(255*2))
	planetColor.A = 0xff

	v, i := ui.Circle(ui.PlanetSize/2, ui.PlanetSize/2, radius, planetColor)
	image.DrawTriangles(v, i, ui.Src, nil)
	return image
}

func (p *Planet) generateHighlightedImage() *ebiten.Image {
	image := ebiten.NewImage(ui.PlanetSize, ui.PlanetSize)
	radius := float32(basePlanetRadius+p.Habitability/32) + ui.Border

	v, i := ui.Circle(ui.PlanetSize/2, ui.PlanetSize/2, radius, color.RGBA{0xff, 0xff, 0xff, 0xff})
	image.DrawTriangles(v, i, ui.Src, nil)
	image.DrawImage(p.image, nil)
	return image
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
		text.Draw(image, p.Name(), p.ttf, cx-textBounds.Dx()/2, cy-16, color.White)
	}
}

func (p *Planet) ReplaceWithStructure() {
	p.Unhighlight()
	p.visible = false
}

func (p *Planet) Describe() string {
	return fmt.Sprintf("planet: %s\nhabitability: %d\nwater: %d", p.Name(), p.Habitability, p.Water)
}
