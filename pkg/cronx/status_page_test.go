package cronx

import (
	"testing"
)

func TestGetStatusPageTemplate(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "Success",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := GetStatusPageTemplate()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetStatusPageTemplate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
