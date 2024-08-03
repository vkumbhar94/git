package main

import "testing"

func Test_xyz(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "as",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := xyz(); (err != nil) != tt.wantErr {
				t.Errorf("xyz() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
