package structure

import (
	"embed"
	"encoding/json"
)

const (
	StructureDataLength = 13
	StructuresJSONFile  = "structures.json"
)

const (
	HQ int = iota
	Outpost
	Water
	Habitat
	Settlement
	Mine
	Smelter
	Factory
	Silica
	ChipFoundry
	Assembly
	Colony
	Capitol
)

const (
	Tax int = iota
	Residential
	Extractor
	Processor
)

type StructureData struct {
	DisplayName string
	Produces    Production
	Storage     []Storage
	Consumes    []Consumption
	Workers     int
	WorkerCost  int
	Berths      int
	MinShips    int
	Cost        int
	Prioritize  int
	Class       int
	Upgrade     Upgrade
	Downgrade   Upgrade
	Buildable   bool
}

type Upgrade struct {
	Structure int
	Required  []int
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
