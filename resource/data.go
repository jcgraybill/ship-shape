package resource

import (
	"encoding/json"
	"image/color"
)

type ResourceData struct {
	DisplayName string
	Color       color.RGBA
}

const ResourceDataLength = 3

const (
	Habitability int = iota
	Water
	Iron
)

var data = []byte(`[
	{ "DisplayName": "habitability", "Color": {"R": 0, "G": 0, "B": 255, "A": 255} },
	{ "DisplayName": "water",        "Color": {"R": 0, "G": 255, "B": 0, "A": 0} },
	{ "DisplayName": "iron",         "Color": {"R": 255, "G": 0, "B": 0, "A": 0} }
	]`)

func GetResourceData() ([ResourceDataLength]ResourceData, error) {
	var rd [ResourceDataLength]ResourceData
	err := json.Unmarshal(data, &rd)
	return rd, err
}
