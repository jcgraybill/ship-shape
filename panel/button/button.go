package button

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/jcgraybill/ship-shape/util"
)

const (
	border = 1
)

type Button struct {
	x, y int
	w, h int

	image *ebiten.Image
	opts  *ebiten.DrawImageOptions
}

func New(x, y, w, h int, message string) *Button {
	l := Button{
		x: x,
		y: y,
		w: w,
		h: h,
	}
	ttf := util.Font()
	textBounds := text.BoundString(ttf, message)
	if textBounds.Dx() > w || textBounds.Dy() > h {
		//TODO: text is larger than bounding box
		fmt.Println("TODO: text is larger than bounding box")
	} else {
		l.w, l.h = w, textBounds.Dy()
	}

	image := ebiten.NewImage(w-border*2, textBounds.Dy()+border*4)
	image.Fill(color.White)

	interior := ebiten.NewImage(w-border*4, textBounds.Dy()+border*2)
	interior.Fill(color.Black)
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(border, border)
	text.Draw(interior, message, ttf, w/2-textBounds.Dx()/2, int(ttf.Metrics().Ascent/util.DPI)+border, color.White)
	image.DrawImage(interior, opts)
	l.image = image

	l.opts = &ebiten.DrawImageOptions{}
	l.opts.GeoM.Translate(float64(x), float64(y))

	return &l
}

func (l *Button) MouseButton(x, y int) bool {
	if l.x < x && l.x+l.w > x {
		if l.y < y && l.y+l.h > y {
			fmt.Println(fmt.Sprintf("button %d %d", x-l.x, y-l.y))
			return true
		}
	}
	return false
}

func (l *Button) Draw() (*ebiten.Image, *ebiten.DrawImageOptions) {
	return l.image, l.opts
}

func (l *Button) Height() int {
	return l.h
}
