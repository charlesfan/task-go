package model

import (
	"encoding/json"
)

const (
	DefaultSep = ":"
)

type NullInt struct {
	Int   int
	Valid bool
}

func (ni *NullInt) UnmarshalJSON(data []byte) error {
	var temp int
	if err := json.Unmarshal(data, &temp); err == nil {
		ni.Int = temp
		ni.Valid = true
		return nil
	}

	ni.Valid = false
	return nil
}

func (ni *NullInt) MarshalJSON() ([]byte, error) {
	if ni.Valid {
		return json.Marshal(ni.Int)
	}
	return json.Marshal(nil)
}
