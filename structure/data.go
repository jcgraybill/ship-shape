package structure

import "encoding/json"

type StructureData struct {
	DisplayName string
	Cost        int
}

var data = []byte(`{ 
	"habitat": { "DisplayName": "habitat","Cost": 8 }, 
	"water": { "DisplayName": "water purification plant","Cost": 16 },
	"admin": { "DisplayName": "administrative center", "Cost": 32 }
	}`)

func GetStructureData() (map[string]StructureData, error) {
	var sd map[string]StructureData
	err := json.Unmarshal(data, &sd)
	return sd, err
}
