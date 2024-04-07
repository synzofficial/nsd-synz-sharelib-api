package typeconvertutil

import (
	"bytes"
	"encoding/json"
)

func ToBodyReader(in any) (*bytes.Reader, error) {
	jsonBody, err := json.Marshal(in)
	if err != nil {
		return nil, err
	}
	bodyReader := bytes.NewReader(jsonBody)
	return bodyReader, nil
}
