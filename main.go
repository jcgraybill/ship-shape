package main

import (
	"image"
	"image/color"
	"math/rand"
	"os"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

const (
	w          = 1024
	h          = 768
	starriness = 3000
	ttf        = "Urbanist-Regular.ttf"
)

var (
	emptyImage = ebiten.NewImage(3, 3)
)

type Game struct {
	bg  *ebiten.Image
	ttf font.Face
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(g.bg, nil)
	src := emptyImage.SubImage(image.Rect(1, 1, 2, 2)).(*ebiten.Image)
	v, i := rectangle(50, 50, 120, 120, color.RGBA{0x00, 0x80, 0x00, 0xff})
	screen.DrawTriangles(v, i, src, nil)
	v, i = circle(120, 300, 60, color.RGBA{0x80, 0x00, 0x00, 0xff})
	screen.DrawTriangles(v, i, src, nil)
	v, i = line(400, 100, 600, 200, 2, color.RGBA{0x00, 0x00, 0xff, 0xff})
	screen.DrawTriangles(v, i, src, nil)
	v, i = triangle(400, 200, 100, 100, color.RGBA{0xff, 0x00, 0xff, 0xff})
	screen.DrawTriangles(v, i, src, nil)
	text.Draw(screen, "hello world", g.ttf, 20, 40, color.White)
}

func init() {
	rand.Seed(time.Now().UnixNano())
	emptyImage.Fill(color.White)
}

func main() {
	ebiten.SetWindowSize(w, h)
	ebiten.SetWindowTitle("ship shape")

	image := ebiten.NewImage(w, h)
	image.Fill(color.Black)
	starField(image)
	g := Game{
		bg:  image,
		ttf: loadFont(ttf),
	}

	if err := ebiten.RunGame(&g); err != nil {
		panic(err)
	}
}

func (g *Game) Layout(int, int) (int, int) {
	return w, h
}

func starField(image *ebiten.Image) {
	star := ebiten.NewImage(2, 2)

	for y := 0; y < image.Bounds().Dy(); y++ {
		for x := 0; x < image.Bounds().Dx(); x++ {
			if rand.Intn(starriness) == 0 {
				hue := uint8(rand.Intn(255))
				star.Fill(color.RGBA{hue, hue, hue, 255})
				opts := &ebiten.DrawImageOptions{}
				opts.GeoM.Translate(float64(x), float64(y))
				image.DrawImage(star, opts)
			}
		}
	}
}

func loadFont(path string) font.Face {
	ttbytes, err := os.ReadFile(path)
	if err == nil {
		tt, err := opentype.Parse(ttbytes)
		if err == nil {
			fontface, err := opentype.NewFace(tt, &opentype.FaceOptions{
				Size:    24,
				DPI:     72,
				Hinting: font.HintingFull,
			})
			if err == nil {
				return fontface
			}
			panic(err)
		}
		panic(err)
	}
	panic(err)
}
