package structure

import "encoding/json"

type StructureData struct {
	DisplayName string
	Cost        int
}

var data = []byte(`{ 
	"home": { "DisplayName": "home planet", "Cost": 32 },
	"outpost": { "DisplayName": "outpost","Cost": 8 }, 
	"water": { "DisplayName": "hydrologic extractor","Cost": 16 }
	}`)

func GetStructureData() (map[string]StructureData, error) {
	var sd map[string]StructureData
	err := json.Unmarshal(data, &sd)
	return sd, err
}
