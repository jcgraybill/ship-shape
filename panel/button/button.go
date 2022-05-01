package button

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"

	"github.com/jcgraybill/ship-shape/ui"
)

type Button struct {
	bounds  image.Rectangle
	pressed bool
	active  bool
	source  func() bool
	message string

	image         *ebiten.Image
	inactiveImage *ebiten.Image
	pressedImage  *ebiten.Image
	opts          *ebiten.DrawImageOptions
	action        func()
}

func New(x, y, w, h int, message string, action func(), source func() bool) *Button {
	b := Button{
		action:  action,
		pressed: false,
		source:  source,
		message: message,
	}

	b.active = source()
	ttf := ui.Font(ui.TtfRegular)
	textBounds := text.BoundString(*ttf, message)
	b.bounds = image.Rect(x, y, x+w, y+textBounds.Dy()+ui.Buffer*2+ui.Border+2)

	return &b
}

func (b *Button) Draw() (*ebiten.Image, *ebiten.DrawImageOptions) {
	if b.image == nil || b.inactiveImage == nil || b.pressedImage == nil {
		b.createImages()
	}
	if newValue := b.source(); newValue != b.active {
		b.active = newValue
	}

	if b.pressed {
		return b.pressedImage, b.opts
	} else if b.active {
		return b.image, b.opts
	} else {
		return b.inactiveImage, b.opts
	}
}

func (b *Button) Height() int {
	return b.bounds.Dy()
}

func (b *Button) LeftMouseButtonPress(x, y int) bool {
	if !b.active {
		return false
	}
	if b.bounds.At(x, y) == color.Opaque {
		go ui.PlayButtonSound()
		b.pressed = true
		return true
	}
	return false
}

func (b *Button) LeftMouseButtonRelease(x, y int) bool {
	if !b.active {
		return false
	}

	if b.pressed {
		if b.bounds.At(x, y) == color.Opaque {
			b.action()
			return true
		}
		b.pressed = false
	}
	return false
}

func (b *Button) createImages() {
	ttf := ui.Font(ui.TtfRegular)
	b.image = b.createSingleButton(ttf, ui.FocusedColor, ui.BackgroundColor)
	b.inactiveImage = b.createSingleButton(ttf, ui.NonFocusColor, ui.BackgroundColor)
	b.pressedImage = b.createSingleButton(ttf, ui.BackgroundColor, ui.FocusedColor)
	b.opts = &ebiten.DrawImageOptions{}
	b.opts.GeoM.Translate(float64(b.bounds.Min.X), float64(b.bounds.Min.Y))

}

func (b *Button) createSingleButton(ttf *font.Face, foregroundColor color.RGBA, backgroundColor color.RGBA) *ebiten.Image {
	image := ebiten.NewImage(b.bounds.Dx(), b.bounds.Dy())
	image.Fill(foregroundColor)
	interior := ebiten.NewImage(b.bounds.Dx()-ui.Border*2, b.bounds.Dy()-ui.Border*2)
	interior.Fill(backgroundColor)

	textBounds := text.BoundString(*ttf, b.message)
	text.Draw(interior, b.message, *ttf, b.bounds.Dx()/2-textBounds.Dx()/2, int((*ttf).Metrics().Ascent/ui.DPI)+ui.Buffer, foregroundColor)
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(ui.Border, ui.Border)
	image.DrawImage(interior, opts)
	return image
}
