package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

// MarshalJSON ...
func (x *EncString) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}

// UnmarshalJSON ...
func (x *EncString) UnmarshalJSON(b []byte) error {
	var v interface{}
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	switch value := v.(type) {
	case string:
		x.Message = string(value)
		return nil
	default:
		return errors.New("Invalid Data")
	}
}

// Scan ...
func (x *EncString) Scan(src interface{}) error {
	var data string
	switch value := src.(type) {
	case []byte:
		data = string(value)
	default:
		return errors.New("Invalid Message")
	}
	*x = EncString{Message: data}
	return nil
}

// Value ...
func (x *EncString) Value() (driver.Value, error) {
	if x.Message == "" {
		return nil, nil
	}
	encData := x.Message
	return encData, nil
}

// IsEmpty ...
func (x *EncString) IsEmpty() bool {
	return x.Message == ""
}
