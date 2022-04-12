package util

import (
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

func StarField(w, h, starriness int) *ebiten.Image {
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
