package label

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/jcgraybill/ship-shape/util"
)

type Label struct {
	x, y int
	w, h int

	image *ebiten.Image
	opts  *ebiten.DrawImageOptions
}

func New(x, y, w, h int, message string) *Label {
	l := Label{
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

	image := ebiten.NewImage(w, textBounds.Dy())
	image.Fill(color.Black)

	text.Draw(image, message, ttf, 0, int(ttf.Metrics().Ascent/util.DPI), color.White)
	l.image = image

	l.opts = &ebiten.DrawImageOptions{}
	l.opts.GeoM.Translate(float64(x), float64(y))

	return &l
}

func (l *Label) MouseButton(x, y int) bool {
	if l.x < x && l.x+l.w > x {
		if l.y < y && l.y+l.h > y {
			return true
		}
	}
	return false
}

func (l *Label) Draw() (*ebiten.Image, *ebiten.DrawImageOptions) {
	return l.image, l.opts
}

func (l *Label) Height() int {
	return l.h
}
