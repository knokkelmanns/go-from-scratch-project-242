package tests

import (
	"testing"

	"code"
)

func TestGetPathSize(t *testing.T) {
	tests := []struct {
		name  string
		path  string
		human bool
		want  string
	}{
		{
			name:  "regular file",
			path:  "testdata/test.txt",
			human: false,
			want:  "6B",
		},
		{
			name:  "directory",
			path:  "testdata/flat",
			human: false,
			want:  "8B",
		},
		{
			name:  "regular file human-readable bytes",
			path:  "testdata/test.txt",
			human: true,
			want:  "6.0B",
		},
		{
			name:  "regular file human-readable kilobytes",
			path:  "testdata/big.txt",
			human: true,
			want:  "2.0KB",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := code.GetPathSize(tt.path, tt.human)
			if err != nil {
				t.Fatalf("GetPathSize(%q, %v) returned error: %v", tt.path, tt.human, err)
			}
			if got != tt.want {
				t.Errorf("GetPathSize(%q, %v) = %q, want %q", tt.path, tt.human, got, tt.want)
			}
		})
	}
}

func TestGetPathSize_NotExist(t *testing.T) {
	_, err := code.GetPathSize("testdata/does-not-exist", false)
	if err == nil {
		t.Fatal("GetPathSize on a non-existent path should return an error")
	}
}
