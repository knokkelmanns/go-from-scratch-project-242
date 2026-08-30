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
		all   bool
		want  string
	}{
		{
			name:  "regular file",
			path:  "testdata/test.txt",
			human: false,
			all:   false,
			want:  "6B",
		},
		{
			name:  "directory",
			path:  "testdata/flat",
			human: false,
			all:   false,
			want:  "8B",
		},
		{
			name:  "regular file human-readable bytes",
			path:  "testdata/test.txt",
			human: true,
			all:   false,
			want:  "6.0B",
		},
		{
			name:  "regular file human-readable kilobytes",
			path:  "testdata/big.txt",
			human: true,
			all:   false,
			want:  "2.0KB",
		},
		{
			name:  "directory without hidden files",
			path:  "testdata/withhidden",
			human: false,
			all:   false,
			want:  "5B",
		},
		{
			name:  "directory with hidden files",
			path:  "testdata/withhidden",
			human: false,
			all:   true,
			want:  "7B",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := code.GetPathSize(tt.path, tt.human, tt.all)
			if err != nil {
				t.Fatalf("GetPathSize(%q, %v, %v) returned error: %v", tt.path, tt.human, tt.all, err)
			}
			if got != tt.want {
				t.Errorf("GetPathSize(%q, %v, %v) = %q, want %q", tt.path, tt.human, tt.all, got, tt.want)
			}
		})
	}
}

func TestGetPathSize_NotExist(t *testing.T) {
	_, err := code.GetPathSize("testdata/does-not-exist", false, false)
	if err == nil {
		t.Fatal("GetPathSize on a non-existent path should return an error")
	}
}
