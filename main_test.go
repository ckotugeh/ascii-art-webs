package main

import (
	
	"testing"
)

func TestGenerateASCIIArt(t *testing.T) {
	// Table-driven test cases
	tests := []struct {
		name    string
		text    string
		banner  string
		want    string // Simplified check for this example
		wantErr bool
	}{
		{
			name:    "Basic character check",
			text:    "A",
			banner:  "test_banner", // Ensure this file exists in banners/
			wantErr: false,
		},
		{
			name:    "Multi-line input",
			text:    "A\nA",
			banner:  "test_banner",
			wantErr: false,
		},
		{
			name:    "Missing banner file",
			text:    "Hello",
			banner:  "non_existent",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := generateASCIIArt(tt.text, tt.banner)
			
			// Check if error matches expectation
			if (err != nil) != tt.wantErr {
				t.Errorf("generateASCIIArt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// If no error, ensure output isn't empty
			if err == nil && got == "" {
				t.Errorf("generateASCIIArt() returned empty string for valid input")
			}
		})
	}
}
