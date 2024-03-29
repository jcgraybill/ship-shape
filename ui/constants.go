package ui

import "image/color"

const (
	NameofGame     = "ship shape"
	starriness     = 3000
	PlanetSize     = 32
	PlanetDistance = 6 // lower numbers are denser

	WindowW           = 1024
	WindowH           = 768
	ArrowKeyMoveSpeed = 4 // larger numbers are faster

	Buffer    = 4
	Border    = 1
	BarHeight = 4

	PanelWidth           = 200
	PanelExternalPadding = 10
	IncomeRate           = 3

	YearLength         = 1200 // smaller numbers are faster
	BaseProductionRate = 3600 // smaller numbers are faster
	ShipSpeed          = 1    // smaller numbers are faster
)

var FocusedColor = color.RGBA{0x78, 0xcc, 0xe2, 0xff}
var NonFocusColor = color.RGBA{0x7f, 0x7f, 0x7f, 0xff}
var LevelProgressColor = color.RGBA{0x80, 0x00, 0x20, 0xff}
var BackgroundColor = color.RGBA{0x00, 0x00, 0x00, 0xff}
