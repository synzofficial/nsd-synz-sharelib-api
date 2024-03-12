package jsonmodel

import (
	"fmt"
	"time"

	datetimeconst "github.com/synzofficial/nsd-synz-sharelib-api/pkg/constant/datetime-const"
)

// use this type to parse yyyy-mm-dd to time.Time at midnight
type DateMinTime time.Time

func ToDateMinTime(t time.Time) DateMinTime {
	lt := t.Local()
	return DateMinTime(time.Date(lt.Year(), lt.Month(), lt.Day(), 0, 0, 0, 0, time.Local))
}

func ParseDateMinTime(t string, format datetimeconst.DateFormats) (DateMinTime, error) {
	parsedDate, err := time.ParseInLocation(format, t, time.Local)
	if err != nil {
		return DateMinTime{}, err
	}
	return ToDateMinTime(parsedDate), nil
}

func (t DateMinTime) String() string {
	return time.Time(t).Local().Format(datetimeconst.DATE_FORMAT)
}

func (t DateMinTime) MarshalJSON() ([]byte, error) {
	lt := time.Time(t).Local()
	ltMin := time.Date(lt.Year(), lt.Month(), lt.Day(), 0, 0, 0, 0, time.Local)
	stamp := fmt.Sprintf("\"%s\"", ltMin.Format(string(datetimeconst.DATE_FORMAT)))
	return []byte(stamp), nil
}

func (t *DateMinTime) UnmarshalJSON(b []byte) error {
	lt, err := jsonToTime(b, datetimeconst.DATE_FORMAT, time.Local)
	if err != nil {
		return err
	}

	ltMin := time.Date(lt.Year(), lt.Month(), lt.Day(), 0, 0, 0, 0, time.Local)
	*t = DateMinTime(ltMin)
	return nil
}

// use this type to parse yyyy-mm-dd to time.Time at 1 ns before midnight of the next day
type DateMaxTime time.Time

func ToDateMaxTime(t time.Time) DateMaxTime {
	lt := t.Local()
	return DateMaxTime(time.Date(lt.Year(), lt.Month(), lt.Day(), 0, 0, 0, 0, time.Local).Add(24 * time.Hour).Add(-time.Nanosecond))
}

func ParseDateMaxTime(t string, format datetimeconst.DateFormats) (DateMaxTime, error) {
	parsedDate, err := time.ParseInLocation(format, t, time.Local)
	if err != nil {
		return DateMaxTime{}, err
	}
	return ToDateMaxTime(parsedDate), nil
}

func (t DateMaxTime) String() string {
	return time.Time(t).Local().Format(datetimeconst.DATE_FORMAT)
}

func (t DateMaxTime) MarshalJSON() ([]byte, error) {
	lt := time.Time(t).Local()
	ltMin := time.Date(lt.Year(), lt.Month(), lt.Day(), 0, 0, 0, 0, time.Local)
	stamp := fmt.Sprintf("\"%s\"", ltMin.Format(string(datetimeconst.DATE_FORMAT)))
	return []byte(stamp), nil
}

func (t *DateMaxTime) UnmarshalJSON(b []byte) error {
	lt, err := jsonToTime(b, datetimeconst.DATE_FORMAT, time.Local)
	if err != nil {
		return err
	}

	ltMin := time.Date(lt.Year(), lt.Month(), lt.Day(), 0, 0, 0, 0, time.Local).Add(24 * time.Hour).Add(-time.Nanosecond)
	*t = DateMaxTime(ltMin)
	return nil
}

// use this to parse string to time.Time
func ToDateTime(s string, layout string) time.Time {
	dateResult, _ := time.ParseInLocation(layout, s, time.Local)
	return dateResult
}
