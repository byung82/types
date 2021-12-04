package types

import (
	"encoding/json"
	"errors"
	"github.com/mcnijman/go-emailaddress"
	"regexp"
)

type SecureEmail string

func (s SecureEmail) Secure() (string, error) {
	email, err := emailaddress.Parse(string(s))

	if err != nil {
		return "", errors.New("InvalidEmailFormat")
	}

	re := regexp.MustCompile(`^([a-z0-9._%+\-]+)@([a-z0-9.\-]+)\.([a-z]{2,4})`)
	return re.ReplaceAllString(email.String(), `***@***.$3`), nil
}

func (s SecureEmail) String() string {
	if v, err := s.Secure(); err != nil {
		return ""
	} else {
		return v
	}
}

func (s SecureEmail) MarshalJSON() ([]byte, error) {
	if v, err := s.Secure(); err != nil || len(v) == 0 {
		return json.Marshal(nil)
	} else {
		return json.Marshal(v)
	}
}

func (s SecureEmail) Text() string {
	return string(s)
}
