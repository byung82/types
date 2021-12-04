package types

import (
	"encoding/json"
)

type SecureCode string

func (s SecureCode) Secure() (string, error) {
	return "****", nil
}

func (s SecureCode) String() string {
	if v, err := s.Secure(); err != nil {
		return ""
	} else {
		return v
	}
}

func (s SecureCode) MarshalJSON() ([]byte, error) {

	v, err := s.Secure()

	if err != nil {
		return json.Marshal(nil)
	} else if len(v) == 0 {
		return json.Marshal(nil)
	} else {
		return json.Marshal(v)
	}
}

func (s SecureCode) Text() string {
	return string(s)
}
