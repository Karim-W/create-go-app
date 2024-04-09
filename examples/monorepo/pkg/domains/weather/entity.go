// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    entity, err := UnmarshalEntity(bytes)
//    bytes, err = entity.Marshal()

package weather

import "encoding/json"

func UnmarshalEntity(data []byte) (Entity, error) {
	var r Entity
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Entity) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Entity struct {
	Date         string `json:"date"`
	TemperatureC int64  `json:"temperatureC"`
	TemperatureF int64  `json:"temperatureF"`
	Summary      string `json:"summary"`
}
