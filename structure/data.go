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
	Storage     []Storage
	Consumes    int
	Berths      int
}

type Production struct {
	Resource int
	Rate     uint8
	Requires int
}

type Storage struct {
	Resource int
	Capacity uint8
	Amount   uint8
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
