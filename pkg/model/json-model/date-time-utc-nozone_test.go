package jsonmodel

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	datetimeconst "github.com/synzofficial/nsd-synz-sharelib-api/pkg/constant/datetime-const"
)

type structWithDateTimeUTCNoZone struct {
	Now DateTimeUTCNoZone `json:"now"`
}

func TestParseDateTimeUTCNoZone(t *testing.T) {
	testCases := []struct {
		input       []byte
		wantTime    DateTimeUTCNoZone
		wantMarshal []byte
		wantErr     error
	}{
		{
			input: []byte(`{"now":"2023-05-16 07:00:00"}`),
			wantTime: func() DateTimeUTCNoZone {
				parsedTime, _ := time.Parse(datetimeconst.DATETIME_WITH_MILLIS_TZ_FORMAT, "2023-05-16T14:00:00.000+07")
				return DateTimeUTCNoZone(parsedTime.UTC())
			}(),
			wantMarshal: []byte(`{"now":"2023-05-16 07:00:00"}`),
			wantErr:     nil,
		},
	}
	for _, testCase := range testCases {
		var parsedStruct structWithDateTimeUTCNoZone
		err := json.Unmarshal(testCase.input, &parsedStruct)
		if err != nil {
			assert.Equal(t, testCase.wantErr, err)
		} else {
			assert.Equal(t, testCase.wantTime, parsedStruct.Now)
		}

		jsonByte, err := json.Marshal(parsedStruct)
		if err != nil {
			assert.Equal(t, testCase.wantErr, err)
		} else {
			assert.Equal(t, testCase.wantMarshal, jsonByte)
		}
	}

}
