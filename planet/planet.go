package planet

import (
	"fmt"
	"image"
	"math/rand"

	"golang.org/x/image/font"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jcgraybill/ship-shape/resource"
	"github.com/jcgraybill/ship-shape/ui"
)

const (
	basePlanetRadius = 8
)

type Planet struct {
	x, y             int
	resources        map[int]uint8
	resourceData     *[resource.ResourceDataLength]resource.ResourceData
	highlighted      bool
	image            *ebiten.Image
	highlightedImage *ebiten.Image
	blackImage       *ebiten.Image
	name             string
	displayOpts      *ebiten.DrawImageOptions
	ttf              *font.Face
	visible          bool
	Bounds           image.Rectangle
}

func New(x, y int, resources map[int]uint8, resourceData *[resource.ResourceDataLength]resource.ResourceData) *Planet {
	var p Planet

	p.x, p.y = x, y
	p.visible = true

	p.resources = resources
	p.resourceData = resourceData

	p.name = p.generateName()

	p.image, p.highlightedImage, p.blackImage = p.generatePlanetImages()

	p.highlighted = false

	p.displayOpts = &ebiten.DrawImageOptions{}
	p.displayOpts.GeoM.Translate(float64(p.x-ui.PlanetSize/2), float64(p.y-ui.PlanetSize/2))
	p.ttf = ui.Font(ui.TtfRegular)
	p.Bounds = image.Rect(p.x-ui.PlanetSize/2, p.y-ui.PlanetSize/2, p.x+ui.PlanetSize/2, p.y+ui.PlanetSize/2)
	return &p
}

func (p *Planet) generateName() string {
	var n string
	letters := []string{"alpha", "beta", "gamma", "delta", "epsilon", "zêta", "êta", "thêta", "iota", "kappa", "lambda", "mu", "nu", "xi", "omikron", "pi", "rho", "sigma", "tau", "upsilon", "phi", "chi", "psi", "omega"}
	n = fmt.Sprintf("%s-%d", letters[rand.Intn(len(letters))], rand.Intn(1000))
	return n
}

func (p *Planet) ReplaceWithStructure() {
	p.Unhighlight()
	p.visible = false
}
