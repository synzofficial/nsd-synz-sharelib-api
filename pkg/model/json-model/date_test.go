package jsonmodel

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type testStruct struct {
	StartDate DateMinTime        `json:"start_date"`
	EndDate   DateMaxTime        `json:"end_date"`
	Now       DateTimeWithMillis `json:"now"`
}

func TestParseDateTime(t *testing.T) {
	input := `{"start_date":"2022-10-31","end_date":"2022-10-31","now":"2022-10-31T14:59:59.999+07"}`
	var inputStruct testStruct
	err := json.Unmarshal([]byte(input), &inputStruct)
	if err != nil {
		t.Errorf("error unmarshal struct: %v", err)
	}

	// check DateMinTime
	assert.Equal(t, time.Date(2022, 10, 31, 0, 0, 0, 0, time.Local), time.Time(inputStruct.StartDate))

	// check DateMaxTime
	assert.Equal(t, time.Date(2022, 11, 01, 0, 0, 0, 0, time.Local).Add(-time.Nanosecond), time.Time(inputStruct.EndDate))

	// check DateTimeWithMillis
	assert.Equal(t, time.Date(2022, 10, 31, 14, 59, 59, 999000000, time.Local), time.Time(inputStruct.Now))

	jsonByte, err := json.Marshal(&inputStruct)
	if err != nil {
		t.Errorf("error marshal struct: %v", err)
	}

	assert.Equal(t, input, string(jsonByte))

}
