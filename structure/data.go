package structure

import (
	"embed"
	"encoding/json"
)

const (
	StructureDataLength = 8
	StructuresJSONFile  = "structures.json"
)

const (
	Capitol int = iota
	Outpost
	Water
	Habitat
	Settlement
	Mine
	Smelter
	Factory
)

type StructureData struct {
	DisplayName string
	Produces    Production
	Storage     []Storage
	Consumes    []Consumption
	Workers     int
	WorkerCost  int
	Berths      int
	Cost        int
}

type Consumption struct {
	Resource int
	Rate     uint8
}

type Production struct {
	Resource int
	Rate     uint8
	Requires []Ingredient
}

type Ingredient struct {
	Resource int
	Quantity uint8
}

type Storage struct {
	Resource int
	Capacity uint8
	Amount   uint8
}

//go:embed structures.json
var structureJSON embed.FS

//.
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
