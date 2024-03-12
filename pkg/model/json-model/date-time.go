package jsonmodel

import (
	"fmt"
	"time"

	datetimeconst "github.com/synzofficial/nsd-synz-sharelib-api/pkg/constant/datetime-const"
)

type DateTimeWithMillis time.Time

func (t DateTimeWithMillis) String() string {
	return time.Time(t).Local().Format(datetimeconst.DATETIME_WITH_MILLIS_TZ_FORMAT)
}

func (t DateTimeWithMillis) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", time.Time(t).Local().Format(string(datetimeconst.DATETIME_WITH_MILLIS_TZ_FORMAT)))
	return []byte(stamp), nil
}

func (t *DateTimeWithMillis) UnmarshalJSON(b []byte) error {
	parsedTime, err := jsonToTime(b, datetimeconst.DATETIME_WITH_MILLIS_TZ_FORMAT, time.Local)
	if err != nil {
		return err
	}

	*t = DateTimeWithMillis(parsedTime)
	return nil
}
