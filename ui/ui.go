package ui

import (
	"embed"
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

func GameData(path string) ([]byte, error) {
	return gd.ReadFile(path)
}

//go:embed fonts audio images
var gd embed.FS

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

func Font(which string) font.Face {
	ttbytes, err := GameData(which)
	if err == nil {
		tt, err := opentype.Parse(ttbytes)
		if err == nil {
			fontface, err := opentype.NewFace(tt, &opentype.FaceOptions{
				Size:    fontSize,
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
