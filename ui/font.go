package ui

import (
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

const (
	TtfRegular = "fonts/OpenSans_SemiCondensed-Regular.ttf"
	TtfBold    = "fonts/OpenSans_SemiCondensed-Bold.ttf"
	fontSize   = 13
	DPI        = 72
)

var regular *font.Face
var bold *font.Face

func Font(which string) *font.Face {
	switch which {
	case TtfRegular:
		if regular == nil {
			regular = loadFont(TtfRegular)
		}
		return regular
	case TtfBold:
		if bold == nil {
			bold = loadFont(TtfBold)
		}
		return bold
	}
	return nil
}

func loadFont(which string) *font.Face {
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
				return &fontface
			}
			panic(err)
		}
		panic(err)
	}
	panic(err)
}
