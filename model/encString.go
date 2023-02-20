package model

import (
	"encoding/json"
	"errors"
)

// MarshalJSON ...
func (e *EncString) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.String())
}

// UnmarshalJSON ...
func (e *EncString) UnmarshalJSON(b []byte) error {
	var v interface{}
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	switch value := v.(type) {
	case string:
		e.Message = string(value)
		return nil
	default:
		return errors.New("Invalid Data")
	}
}

// IsEmpty ...
func (e *EncString) IsEmpty() bool {
	return e.Message == ""
}
