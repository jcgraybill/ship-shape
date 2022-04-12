package label

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/jcgraybill/ship-shape/util"
	"golang.org/x/image/font"
)

type Label struct {
	x, y int
	w, h int

	ttf   font.Face
	text  string
	color color.Gray16
}

func New(x, y, w, h int, message string) *Label {
	l := Label{
		x:     x,
		y:     y,
		w:     w,
		h:     h,
		text:  message,
		ttf:   util.Font(),
		color: color.White,
	}

	textBounds := text.BoundString(l.ttf, l.text)
	if textBounds.Dx() > w || textBounds.Dy() > h {
		//TODO: text is larger than bounding box
		fmt.Println("TODO: text is larger than bounding box")
	} else {
		l.w, l.h = textBounds.Dx(), textBounds.Dy()
	}
	return &l
}

//FIXME:x is bottom of first line of text, so misses clicks in top line
func (l *Label) MouseButton(x, y int) bool {
	if l.x < x && l.x+l.w > x {
		if l.y < y && l.y+l.h > y {
			fmt.Println(fmt.Sprintf("label %d %d", x-l.x, y-l.y))
			return true
		}
	}
	return false
}

func (l *Label) Draw(image *ebiten.Image) {
	text.Draw(image, l.text, l.ttf, l.x, l.y, l.color)
}
