package types

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strings"
)

type SecureCardNumber string

// Value implements the driver Valuer interface.
func (d SecureCardNumber) Value() (driver.Value, error) {
	if len(string(d)) == 0 {
		return nil, nil
	}

	return driver.Value(d.Secure()), nil
}

// MarshalJSON implements the json.Marshaler interface.
func (d SecureCardNumber) MarshalJSON() ([]byte, error) {
	str := d.Secure()

	return json.Marshal(str)
}

func (d SecureCardNumber) Secure() string {
	var str string
	r := strings.NewReplacer("-", "")
	str = r.Replace(string(d))

	if strings.Contains(string(d), "*") {
		str = string(d)
	} else {
		str = fmt.Sprintf("%s-%s**-****-%s", str[:4], str[4:6], str[12:])
	}

	return str
}

func (d SecureCardNumber) CardTxt() string {
	size := len(d)

	if size == 15 {
		return string(d[size-3:])
	} else {
		return string(d[size-4:])
	}
}

func (d SecureCardNumber) BinCode() string {
	r := strings.NewReplacer("-", "")

	str := r.Replace(string(d))

	return str[:6]
}

func (d SecureCardNumber) String() string {
	return d.Secure()
}

func (d SecureCardNumber) UnSecure() string {
	return string(d)
}
