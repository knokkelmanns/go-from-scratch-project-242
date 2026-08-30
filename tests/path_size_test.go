package tests

import (
	"testing"

	"code"
)

func TestGetPathSize(t *testing.T) {
	tests := []struct {
		name string
		path string
		want string
	}{
		{
			name: "regular file",
			path: "testdata/test.txt",
			want: "6B",
		},
		{
			name: "directory",
			path: "testdata/flat",
			want: "8B",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := code.GetPathSize(tt.path)
			if err != nil {
				t.Fatalf("GetPathSize(%q) returned error: %v", tt.path, err)
			}
			if got != tt.want {
				t.Errorf("GetPathSize(%q) = %q, want %q", tt.path, got, tt.want)
			}
		})
	}
}

func TestGetPathSize_NotExist(t *testing.T) {
	_, err := code.GetPathSize("testdata/does-not-exist")
	if err == nil {
		t.Fatal("GetPathSize on a non-existent path should return an error")
	}
}
