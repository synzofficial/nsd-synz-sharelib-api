package datetimeconst

type DateFormats = string

const (
	DATE_FORMAT     DateFormats = "2006-01-02"
	DATETIME_FORMAT DateFormats = "2006-01-02T15:04:05"
	TIME_FORMAT     DateFormats = "15:04:05"

	DATETIME_WITH_MILLIS_TZ_FORMAT DateFormats = "2006-01-02T15:04:05.000-07"
	DATETIME_NO_T                  DateFormats = "2006-01-02 15:04:05"

	BO_DATETIME_FORMAT DateFormats = "02/01/2006 15:04:05"
	BO_DATE_FORMAT DateFormats = "02/01/2006"
)
