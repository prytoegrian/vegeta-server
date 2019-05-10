package models

import (
	"time"
)

// AttackInfo encapsulates the attack information for attacks
// submitted to the dispatcher
type AttackInfo struct {
	// id is a attack UUID generated for each attack submitted
	ID string `json:"id,omitempty"`
	// Status captures the attack status in the scheduler pipeline
	Status AttackStatus `json:"status,omitempty"`
	// Params captures the attack parameters
	Params    AttackParams `json:"params,omitempty"`
	CreatedAt string       `json:"created_at"`
	UpdatedAt string       `json:"updated_at"`
	Created2  JSONTime     `json:"created2"`
}

// AttackDetails captures the AttackInfo for COMPLETED attacks,
// along with the result as a byte array
type AttackDetails struct {
	AttackInfo
	Result []byte `json:"result,omitempty"`
}

type JSONTime time.Time

// MarshalJSON implements the json.Marshaler interface.
func (t JSONTime) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, time.RFC1123)
	b = append(b, '"')
	return b, nil
}

// FilterParams defines a map structure for the filter parameters received via
// query params in the request URL
type FilterParams map[string]interface{}

// Filter defines a type that must be implemented by
// an attack filter
type Filter func(AttackDetails) bool

// StatusFilter implements a attack status filter
// in the Filter function format
func StatusFilter(status string) Filter {
	return func(a AttackDetails) bool {
		if status == "" {
			return true
		}

		if a.Status == AttackStatus(status) {
			return true
		}

		return false
	}
}
