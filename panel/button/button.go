package button

import (
	"bytes"
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
	"github.com/hajimehoshi/ebiten/v2/text"

	"github.com/jcgraybill/ship-shape/ui"
)

const (
	border = 1
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
		h:       h,
		action:  action,
		pressed: false,
	}
	ttf := ui.Font()
	textBounds := text.BoundString(ttf, message)
	if textBounds.Dx() > w || textBounds.Dy() > h {
		//TODO: text is larger than bounding box
		fmt.Println("TODO: text is larger than bounding box")
	} else {
		b.w, b.h = w, textBounds.Dy()
	}

	image := ebiten.NewImage(w-border*2, textBounds.Dy()+border*4)
	image.Fill(color.White)

	interior := ebiten.NewImage(w-border*4, textBounds.Dy()+border*2)
	interior.Fill(color.Black)
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(border, border)
	text.Draw(interior, message, ttf, w/2-textBounds.Dx()/2, int(ttf.Metrics().Ascent/ui.DPI)+border, color.White)
	image.DrawImage(interior, opts)
	b.image = image

	b.opts = &ebiten.DrawImageOptions{}
	b.opts.GeoM.Translate(float64(x), float64(y))

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
