package portabledec

import (
	"database/sql/driver"
	"errors"
	"github.com/shopspring/decimal"
)

type JSONDec struct {
	decimal.Decimal
}

func (t JSONDec) MarshalJSON() ([]byte, error) {
	return []byte("\"" + t.StringFixed(8) + "\""), nil
}

func (t *JSONDec) UnmarshalJSON(b []byte) error {
	s := string(b)
	if len(s) > 0 && s[0] == '"' {
		s = s[1:]
	}
	if len(s) > 0 && s[len(s)-1] == '"' {
		s = s[:len(s)-1]
	}
	dec, err := decimal.NewFromString(s)
	if err != nil {
		return errors.New("Incompatible type for JSONDec")
	}
	*t = JSONDec{dec}
	return nil
}

func (t JSONDec) Value() (driver.Value, error) {
	return []byte(t.String()), nil
}

func (s *JSONDec) Scan(src interface{}) error {
	switch src.(type) {
	case []byte:
		newDec, err := decimal.NewFromString(string(src.([]byte)))
		if err != nil {
			return err
		}
		*s = JSONDec{newDec}
	default:
		return errors.New("Incompatible type for SDecimal")
	}
	return nil
}
