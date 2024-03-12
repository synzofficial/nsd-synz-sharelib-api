package jsonmodel

import (
	"fmt"
	"time"

	datetimeconst "github.com/synzofficial/nsd-synz-sharelib-api/pkg/constant/datetime-const"
)

type DateTimeLocal time.Time

func (t DateTimeLocal) String() string {
	return time.Time(t).Local().Format(datetimeconst.DATETIME_NO_T)
}

func (t DateTimeLocal) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", time.Time(t).Local().Format(string(datetimeconst.DATETIME_NO_T)))
	return []byte(stamp), nil
}

func (t *DateTimeLocal) UnmarshalJSON(b []byte) error {
	parsedTime, err := jsonToTime(b, datetimeconst.DATETIME_NO_T, time.Local)
	if err != nil {
		return err
	}

	*t = DateTimeLocal(parsedTime.Local())
	return nil
}
