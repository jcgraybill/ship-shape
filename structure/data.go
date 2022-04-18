package structure

import (
	"embed"
	"encoding/json"
	"fmt"
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
	Storage     Storage
}

type Production struct {
	Resource int
	Rate     uint8
	Requires map[int]int
}

type Storage struct {
	Resource int
	Storage  uint8
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
			fmt.Println(sd)
			return sd
		}
	}
	panic(err)
}
