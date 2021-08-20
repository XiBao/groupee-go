package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

// Uint64 marshal uint64 to string
type Uint64 uint64

// MarshalJSON implement json marshal interface
func (u Uint64) MarshalJSON() ([]byte, error) {
	if u == 0 {
		return nil, nil
	}
	return []byte(fmt.Sprintf(`"%s"`, strconv.FormatUint(uint64(u), 10))), nil
}

// UnmarshalJSON implement json marshal interface
func (u *Uint64) UnmarshalJSON(bs []byte) (err error) {
	var i uint64
	if err := json.Unmarshal(bs, &i); err == nil {
		*u = Uint64(i)
		return nil
	}
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return errors.New("expected a string or an integer")
	}
	if err := json.Unmarshal([]byte(s), &i); err != nil {
		return err
	}
	*u = Uint64(i)
	return nil
}

// Uint64 return uint64 value
func (u Uint64) Uint64() uint64 {
	return uint64(u)
}
