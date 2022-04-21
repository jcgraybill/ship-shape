package ui

import "image/color"

const (
	starriness     = 3000
	PlanetSize     = 32
	PlanetDistance = 6 // lower numbers are denser
	StartingMoney  = 5000

	W = 2048
	H = 2048

	WindowW           = 800
	WindowH           = 600
	ArrowKeyMoveSpeed = 4 // larger numbers are faster

	Buffer    = 4
	Border    = 1
	BarHeight = 4

	PanelWidth           = 160
	PanelExternalPadding = 10

	TtfRegular = "fonts/OpenSans_SemiCondensed-Regular.ttf"
	TtfBold    = "fonts/OpenSans_SemiCondensed-Bold.ttf"
	fontSize   = 13
	DPI        = 72

	IncomeRate         = 0.5
	MaxCapitols        = 1
	DayLength          = 3600 // smaller numbers are faster
	BaseProductionRate = 3600 // smaller numbers are faster
	BidFrequency       = 60   // smaller numbers are more frequent
	ShipSpeed          = 2    // smaller numbers are faster
)

var FocusedColor = color.RGBA{0xff, 0xff, 0xff, 0xff}
var NonFocusColor = color.RGBA{0x7f, 0x7f, 0x7f, 0xff}
var BackgroundColor = color.RGBA{0x00, 0x00, 0x00, 0xff}
