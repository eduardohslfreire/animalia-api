package enum

import (
	"database/sql/driver"
	"fmt"
)

// RoleName ...
type RoleName string

const (
	// FirstMinister ...
	FirstMinister RoleName = "First Minister"
	// Treasurer ...
	Treasurer RoleName = "Treasurer"
	// General ...
	General RoleName = "General"
	// SecretaryOfState ...
	SecretaryOfState RoleName = "Secretary of State"
	// MinisterOfState ..
	MinisterOfState RoleName = "Minister of State"
	// Civil ...
	Civil RoleName = "Civil"
)

// Scan ...
func (r *RoleName) Scan(value interface{}) error {
	asString, ok := value.(string)
	if !ok {
		return fmt.Errorf("Value is not a string type, but %T", value)
	}
	*r = RoleName(asString)
	return nil
}

// Value ...
func (r RoleName) Value() (driver.Value, error) { return string(r), nil }
