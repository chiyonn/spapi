package types

import (
	"time"
	"strings"
	"encoding/json"
)

type NullableTime struct {
	Time *time.Time
}

func (nt *NullableTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)
	if s == "" || s == "null" {
		nt.Time = nil
		return nil
	}
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		return err
	}
	nt.Time = &t
	return nil
}

func (nt NullableTime) MarshalJSON() ([]byte, error) {
	if nt.Time == nil {
		return json.Marshal(nil)
	}
	return json.Marshal(nt.Time.Format(time.RFC3339))
}

