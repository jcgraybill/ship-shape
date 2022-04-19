package button

import (
	"bytes"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
	"github.com/hajimehoshi/ebiten/v2/text"

	"github.com/jcgraybill/ship-shape/ui"
)

type Button struct {
	x, y    int
	w, h    int
	pressed bool

	image  *ebiten.Image
	opts   *ebiten.DrawImageOptions
	audio  *audio.Player
	action func()
}

func New(x, y, w, h int, message string, action func()) *Button {
	b := Button{
		x:       x,
		y:       y,
		w:       w,
		action:  action,
		pressed: false,
	}

	ttf := ui.Font()
	textBounds := text.BoundString(ttf, message)
	b.h = textBounds.Dy() + ui.Buffer*2 + ui.Border + 2

	image := ebiten.NewImage(b.w, b.h)
	image.Fill(ui.FocusedColor)

	interior := ebiten.NewImage(b.w-ui.Border*2, b.h-ui.Border*2)
	interior.Fill(ui.BackgroundColor)
	text.Draw(interior, message, ttf, b.w/2-textBounds.Dx()/2, int(ttf.Metrics().Ascent/ui.DPI)+ui.Buffer, ui.FocusedColor)
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(ui.Border, ui.Border)
	image.DrawImage(interior, opts)
	b.image = image

	b.opts = &ebiten.DrawImageOptions{}
	b.opts.GeoM.Translate(float64(b.x), float64(b.y))

	audioContext := audio.CurrentContext()
	audioBytes, err := ui.GameData("audio/button.wav")
	if err != nil {
		panic(err)
	}
	d, err := wav.Decode(audioContext, bytes.NewReader(audioBytes))
	if err != nil {
		panic(err)
	}
	b.audio, err = audioContext.NewPlayer(d)
	if err != nil {
		panic(err)
	}

	return &b
}

func (b *Button) LeftMouseButtonPress(x, y int) bool {
	if b.x < x && b.x+b.w > x {
		if b.y < y && b.y+b.h > y {
			b.pressed = true
			b.opts.ColorM.Scale(-1, -1, -1, 1)
			b.opts.ColorM.Translate(1, 1, 1, 0)
			go b.playSound()
			return true
		}
	}
	return false
}

func (b *Button) LeftMouseButtonRelease(x, y int) bool {
	if b.pressed {
		if b.x < x && b.x+b.w > x {
			if b.y < y && b.y+b.h > y {
				b.action()
				return true
			}
		}
		b.pressed = false
		b.opts.ColorM.Scale(-1, -1, -1, 1)
		b.opts.ColorM.Translate(1, 1, 1, 0)
	}
	return false
}

func (b *Button) Draw() (*ebiten.Image, *ebiten.DrawImageOptions) {
	return b.image, b.opts
}

func (b *Button) Height() int {
	return b.h
}

func (b *Button) playSound() {
	b.audio.Rewind()
	b.audio.Play()
}
