package util

import (
	"image/color"
	"math/rand"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

const (
	starriness = 3000
	ttf        = "fonts/OpenSans_SemiCondensed-Regular.ttf"
	DPI        = 72
)

func StarField(w, h int) *ebiten.Image {
	field := ebiten.NewImage(w, h)
	field.Fill(color.Black)

	star := ebiten.NewImage(2, 2)

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if rand.Intn(starriness) == 0 {
				hue := uint8(rand.Intn(255))
				star.Fill(color.RGBA{hue, hue, hue, 255})
				opts := &ebiten.DrawImageOptions{}
				opts.GeoM.Translate(float64(x), float64(y))
				field.DrawImage(star, opts)
			}
		}
	}
	return field
}

func Font() font.Face {
	ttbytes, err := os.ReadFile(ttf)
	if err == nil {
		tt, err := opentype.Parse(ttbytes)
		if err == nil {
			fontface, err := opentype.NewFace(tt, &opentype.FaceOptions{
				Size:    12,
				DPI:     DPI,
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
