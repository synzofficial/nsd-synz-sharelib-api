package jsonmodel

import (
	"fmt"
	"time"

	datetimeconst "github.com/synzofficial/nsd-synz-sharelib-api/pkg/constant/datetime-const"
)

type DateTimeUTCNoZone time.Time

func (t DateTimeUTCNoZone) String() string {
	return time.Time(t).UTC().Format(datetimeconst.DATETIME_NO_T)
}

func (t DateTimeUTCNoZone) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", time.Time(t).UTC().Format(string(datetimeconst.DATETIME_NO_T)))
	return []byte(stamp), nil
}

func (t *DateTimeUTCNoZone) UnmarshalJSON(b []byte) error {
	parsedTime, err := jsonToTime(b, datetimeconst.DATETIME_NO_T, time.UTC)
	if err != nil {
		return err
	}

	*t = DateTimeUTCNoZone(parsedTime.UTC())
	return nil
}
