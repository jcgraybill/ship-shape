package ui

/*
	v, i := ui.Rectangle(50, 50, 120, 120, color.RGBA{0x00, 0x80, 0x00, 0xff})
	image.DrawTriangles(v, i, ui.Src, nil)
	v, i := ui.Circle(120, 300, 60, color.RGBA{0x80, 0x00, 0x00, 0xff})
	image.DrawTriangles(v, i, ui.Src, nil)
	v, i := ui.Line(400, 100, 600, 200, 2, color.RGBA{0x00, 0x00, 0xff, 0xff})
	image.DrawTriangles(v, i, ui.Src, nil)
	v, i := ui.Triangle(400, 200, 100, 100, color.RGBA{0xff, 0x00, 0xff, 0xff})
	image.DrawTriangles(v, i, ui.Src, nil)

*/
import (
	"image"
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

var Src *ebiten.Image

func init() {
	emptyImage := ebiten.NewImage(3, 3)
	emptyImage.Fill(color.White)
	Src = emptyImage.SubImage(image.Rect(1, 1, 2, 2)).(*ebiten.Image)

}

func Line(x0, y0, x1, y1, width float32, clr color.RGBA) ([]ebiten.Vertex, []uint16) {

	theta := math.Atan2(float64(y1-y0), float64(x1-x0))
	theta += math.Pi / 2
	dx := float32(math.Cos(theta))
	dy := float32(math.Sin(theta))

	r := float32(clr.R) / 0xff
	g := float32(clr.G) / 0xff
	b := float32(clr.B) / 0xff
	a := float32(clr.A) / 0xff

	return []ebiten.Vertex{
		{
			DstX:   x0 - width*dx/2,
			DstY:   y0 - width*dy/2,
			SrcX:   1,
			SrcY:   1,
			ColorR: r,
			ColorG: g,
			ColorB: b,
			ColorA: a,
		},
		{
			DstX:   x0 + width*dx/2,
			DstY:   y0 + width*dy/2,
			SrcX:   1,
			SrcY:   1,
			ColorR: r,
			ColorG: g,
			ColorB: b,
			ColorA: a,
		},
		{
			DstX:   x1 - width*dx/2,
			DstY:   y1 - width*dy/2,
			SrcX:   1,
			SrcY:   1,
			ColorR: r,
			ColorG: g,
			ColorB: b,
			ColorA: a,
		},
		{
			DstX:   x1 + width*dx/2,
			DstY:   y1 + width*dy/2,
			SrcX:   1,
			SrcY:   1,
			ColorR: r,
			ColorG: g,
			ColorB: b,
			ColorA: a,
		},
	}, []uint16{0, 1, 2, 1, 2, 3}

}

func Triangle(x, y, w, h float32, clr color.RGBA) ([]ebiten.Vertex, []uint16) {
	r := float32(clr.R) / 0xff
	g := float32(clr.G) / 0xff
	b := float32(clr.B) / 0xff
	a := float32(clr.A) / 0xff
	x0 := x
	y0 := y
	x1 := x + w
	y1 := y + h/2
	y2 := y + h

	return []ebiten.Vertex{
		{
			DstX:   x0,
			DstY:   y0,
			SrcX:   1,
			SrcY:   1,
			ColorR: r,
			ColorG: g,
			ColorB: b,
			ColorA: a,
		},
		{
			DstX:   x1,
			DstY:   y1,
			SrcX:   1,
			SrcY:   1,
			ColorR: r,
			ColorG: g,
			ColorB: b,
			ColorA: a,
		},
		{
			DstX:   x0,
			DstY:   y2,
			SrcX:   1,
			SrcY:   1,
			ColorR: r,
			ColorG: g,
			ColorB: b,
			ColorA: a,
		},
	}, []uint16{0, 1, 2}
}
