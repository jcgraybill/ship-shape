package button

import (
	"bytes"
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
	"github.com/hajimehoshi/ebiten/v2/text"

	"github.com/jcgraybill/ship-shape/ui"
)

type Button struct {
	Bounds  image.Rectangle
	pressed bool
	active  bool

	image         *ebiten.Image
	inactiveImage *ebiten.Image
	opts          *ebiten.DrawImageOptions
	audio         *audio.Player
	action        func()
}

func New(x, y, w, h int, message string, action func()) *Button {
	b := Button{
		action:  action,
		pressed: false,
		active:  true,
	}

	ttf := ui.Font(ui.TtfRegular)
	textBounds := text.BoundString(*ttf, message)
	b.Bounds = image.Rect(x, y, x+w, y+textBounds.Dy()+ui.Buffer*2+ui.Border+2)

	image := ebiten.NewImage(b.Bounds.Dx(), b.Bounds.Dy())
	image.Fill(ui.FocusedColor)

	interior := ebiten.NewImage(b.Bounds.Dx()-ui.Border*2, b.Bounds.Dy()-ui.Border*2)
	interior.Fill(ui.BackgroundColor)
	text.Draw(interior, message, *ttf, b.Bounds.Dx()/2-textBounds.Dx()/2, int((*ttf).Metrics().Ascent/ui.DPI)+ui.Buffer, ui.FocusedColor)
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(ui.Border, ui.Border)
	image.DrawImage(interior, opts)
	b.image = image

	dimage := ebiten.NewImage(b.Bounds.Dx(), b.Bounds.Dy())
	dimage.Fill(ui.NonFocusColor)

	dinterior := ebiten.NewImage(b.Bounds.Dx()-ui.Border*2, b.Bounds.Dy()-ui.Border*2)
	dinterior.Fill(ui.BackgroundColor)
	text.Draw(dinterior, message, *ttf, b.Bounds.Dx()/2-textBounds.Dx()/2, int((*ttf).Metrics().Ascent/ui.DPI)+ui.Buffer, ui.NonFocusColor)
	dimage.DrawImage(dinterior, opts)
	b.inactiveImage = dimage

	b.opts = &ebiten.DrawImageOptions{}
	b.opts.GeoM.Translate(float64(b.Bounds.Min.X), float64(b.Bounds.Min.Y))

	audioContext := audio.CurrentContext()
	audioBytes, err := ui.GameData("audio/button.ogg")
	if err == nil {
		d, err := vorbis.Decode(audioContext, bytes.NewReader(audioBytes))
		if err == nil {
			b.audio, err = audioContext.NewPlayer(d)
			if err != nil {
				panic(err)
			}
		} else {
			panic(err)
		}
	} else {
		panic(err)
	}

	return &b
}

func (b *Button) LeftMouseButtonPress(x, y int) bool {
	if !b.active {
		return false
	}
	if b.Bounds.At(x, y) == color.Opaque {
		b.pressed = true
		b.opts.ColorM.Scale(-1, -1, -1, 1)
		b.opts.ColorM.Translate(1, 1, 1, 0)
		go b.playSound()
		return true
	}
	return false
}

func (b *Button) LeftMouseButtonRelease(x, y int) bool {
	if !b.active {
		return false
	}

	if b.pressed {
		if b.Bounds.At(x, y) == color.Opaque {
			b.action()
			return true
		}
		b.pressed = false
		b.opts.ColorM.Scale(-1, -1, -1, 1)
		b.opts.ColorM.Translate(1, 1, 1, 0)
	}
	return false
}

func (b *Button) Draw() (*ebiten.Image, *ebiten.DrawImageOptions) {
	if b.active {
		return b.image, b.opts
	} else {
		return b.inactiveImage, b.opts
	}
}

func (b *Button) Height() int {
	return b.Bounds.Dy()
}

func (b *Button) playSound() {
	b.audio.Rewind()
	b.audio.Play()
}

func (b *Button) UpdateValue(uint8) { return }
func (b *Button) UpdateText(string) { return }

func (b *Button) Activate()   { b.active = true }
func (b *Button) DeActivate() { b.active = false }
