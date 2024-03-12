package jsonmodel

import (
	"encoding/json"
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestMoney_MarshalJSON(t *testing.T) {
	type arg struct {
		A Money `json:"a"`
	}

	tests := []struct {
		name    string
		input   arg
		want    []byte
		wantErr error
	}{
		{
			name:    "test marshal no rounding",
			input:   arg{A: Money(decimal.NewFromFloat(100.45))},
			want:    []byte(`{"a":100.45}`),
			wantErr: nil,
		},
		{
			name:    "test marshal with rounding",
			input:   arg{A: Money(decimal.NewFromFloat(100.555))},
			want:    []byte(`{"a":100.56}`),
			wantErr: nil,
		},
		{
			name:    "test marshal zero preserved",
			input:   arg{A: Money(decimal.NewFromFloat(100))},
			want:    []byte(`{"a":100.00}`),
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := json.Marshal(tt.input)
			if err != nil {
				assert.Equal(t, tt.wantErr, err)
			} else {
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestMoney_UnmarshalJSON(t *testing.T) {
	type output struct {
		A Money `json:"a"`
	}

	tests := []struct {
		name    string
		input   []byte
		want    Money
		wantErr bool
	}{
		{
			name:    "unmarshal number success",
			input:   []byte(`{"a":100.555}`),
			want:    Money(decimal.NewFromFloat(100.555)),
			wantErr: false,
		},
		{
			name:    "unmarshal number str success",
			input:   []byte(`{"a":"100.555"}`),
			want:    Money(decimal.NewFromFloat(100.555)),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got output
			err := json.Unmarshal(tt.input, &got)
			if err != nil && !tt.wantErr {
				t.Errorf("error occurred: %v", err)
			} else {
				assert.Equal(t, tt.want, got.A)
			}
		})
	}
}
