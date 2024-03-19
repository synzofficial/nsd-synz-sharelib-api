package typeconvertutil

import (
	"time"

	jsonmodel "github.com/synzofficial/nsd-synz-sharelib-api/pkg/model/json-model"
)

func JsonDatetimePtrToTimePtr(j *jsonmodel.DateTimeLocal) *time.Time {
	if j == nil {
		return nil
	}
	return ToPtr(time.Time(*j))
}
