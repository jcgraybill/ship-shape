package structure

import "encoding/json"

type StructureData struct {
	DisplayName string
	Cost        int
}

const StructureDataLength = 3

const (
	Home int = iota
	Outpost
	Water
)

var data = []byte(`[ 
	{ "DisplayName": "home planet", "Cost": 32 },
	{ "DisplayName": "outpost","Cost": 8 }, 
	{ "DisplayName": "hydrology plant","Cost": 16 }
	]`)

func GetStructureData() ([StructureDataLength]StructureData, error) {
	var sd [StructureDataLength]StructureData
	err := json.Unmarshal(data, &sd)
	return sd, err
}
