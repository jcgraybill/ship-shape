package ui

import "image/color"

const (
	starriness = 3000
	PlanetSize = 32

	W = 1000
	H = 1000

	WindowW = 800
	WindowH = 480

	Buffer    = 4
	Border    = 1
	BarHeight = 4

	PanelWidth           = 160
	PanelExternalPadding = 10

	TtfRegular = "fonts/OpenSans_SemiCondensed-Regular.ttf"
	TtfBold    = "fonts/OpenSans_SemiCondensed-Bold.ttf"
	fontSize   = 13
	DPI        = 72

	BaseProductionRate = 4  // smaller numbers are faster
	BidFrequency       = 60 // smaller numbers are more frequent
	ShipSpeed          = 2  // smaller numbers are faster
)

var FocusedColor = color.RGBA{0xff, 0xff, 0xff, 0xff}
var NonFocusColor = color.RGBA{0x7f, 0x7f, 0x7f, 0xff}
var BackgroundColor = color.RGBA{0x00, 0x00, 0x00, 0xff}
