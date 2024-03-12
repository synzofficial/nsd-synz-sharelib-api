package jsonmodel

import (
	"encoding/json"
	"time"

	datetimeconst "github.com/synzofficial/nsd-synz-sharelib-api/pkg/constant/datetime-const"
)

func jsonToTime(b []byte, format datetimeconst.DateFormats, loc *time.Location) (time.Time, error) {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return time.Time{}, err
	}

	var lt time.Time
	var err error
	if loc == nil {
		lt, err = time.Parse(format, s)
		if err != nil {
			return time.Time{}, err
		}
	} else {
		lt, err = time.ParseInLocation(format, s, loc)
		if err != nil {
			return time.Time{}, err
		}
	}

	return lt, nil
}
