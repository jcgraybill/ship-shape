package structure

import (
	"embed"
	"encoding/json"
)

const (
	StructureDataLength = 3
	StructuresJSONFile  = "structures.json"
)

const (
	Home int = iota
	Outpost
	Water
)

type StructureData struct {
	DisplayName string
	Produces    Production
}

type Production struct {
	Resource int
	Rate     uint8
	Requires map[int]int
}

//go:embed structures.json
var structureJSON embed.FS

func GetStructureData() [StructureDataLength]StructureData {
	var sd [StructureDataLength]StructureData
	data, err := structureJSON.ReadFile(StructuresJSONFile)
	if err == nil {
		err := json.Unmarshal(data, &sd)
		if err == nil {
			return sd
		}
	}
	panic(err)
}
