package models

import (
	"testing"

	"github.com/alexadastra/hisho/hisho-core-service/internal/app/models"
)

func TestPage_ToLimitOffset(t *testing.T) {
	type fields struct {
		Size   int64
		Number int64
	}
	tests := []struct {
		name           string
		fields         fields
		expectedLimit  int64
		expectedOffset int64
	}{
		{
			name: "first page",
			fields: fields{
				Size:   10,
				Number: 1,
			},
			expectedLimit:  10,
			expectedOffset: 0,
		},
		{
			name: "second page",
			fields: fields{
				Size:   10,
				Number: 2,
			},
			expectedLimit:  10,
			expectedOffset: 10,
		},
		{
			name: "third page",
			fields: fields{
				Size:   10,
				Number: 3,
			},
			expectedLimit:  10,
			expectedOffset: 20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &models.Page{
				Size:   tt.fields.Size,
				Number: tt.fields.Number,
			}
			got, got1 := p.ToLimitOffset()
			if got != tt.expectedLimit {
				t.Errorf("Page.ToLimitOffset() got = %v, want %v", got, tt.expectedLimit)
			}
			if got1 != tt.expectedOffset {
				t.Errorf("Page.ToLimitOffset() got1 = %v, want %v", got1, tt.expectedOffset)
			}
		})
	}
}
