package label

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/jcgraybill/ship-shape/ui"
)

// TODO text.BoundString underestimates this typeface's height by a few pixels,
// unless there's an "h" in the string. Work around for now, submit a github issue.
const (
	bottomBuffer = 4
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
	ttf := ui.Font()
	textBounds := text.BoundString(ttf, message)
	if textBounds.Dx() > w || textBounds.Dy() > h {
		//TODO: text is larger than bounding box
		fmt.Println("TODO: text is larger than bounding box")
	} else {
		l.w, l.h = w, textBounds.Dy()+bottomBuffer
	}

	image := ebiten.NewImage(l.w, l.h)
	image.Fill(color.Black)

	text.Draw(image, message, ttf, 0, int(ttf.Metrics().Ascent/ui.DPI), color.White)
	l.image = image

	l.opts = &ebiten.DrawImageOptions{}
	l.opts.GeoM.Translate(float64(x), float64(y))

	return &l
}

func (l *Label) LeftMouseButtonPress(x, y int) bool {
	if l.x < x && l.x+l.w > x {
		if l.y < y && l.y+l.h > y {
			return true
		}
	}
	return false
}

func (l *Label) LeftMouseButtonRelease(x, y int) bool {
	return false
}

func (l *Label) Draw() (*ebiten.Image, *ebiten.DrawImageOptions) {
	return l.image, l.opts
}

func (l *Label) Height() int {
	return l.h
}
