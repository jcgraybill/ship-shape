package resource

import (
	"embed"
	"encoding/json"
	"image/color"
)

const (
	ResourceDataLength = 12
	ResourcesJSONFile  = "resources.json"
)

const (
	Habitability int = iota
	Ice
	Iron
	Water
	Population
	Ore
	Metal
	Machinery
	Sand
	Silicon
	IntegratedCircuits
	Computers
)

type ResourceData struct {
	DisplayName string
	Color       color.RGBA
}

//go:embed resources.json
var resourceJSON embed.FS

func GetResourceData() [ResourceDataLength]ResourceData {
	var rd [ResourceDataLength]ResourceData
	data, err := resourceJSON.ReadFile(ResourcesJSONFile)
	if err == nil {
		err := json.Unmarshal(data, &rd)
		if err == nil {
			return rd
		}
	}
	panic(err)
}
