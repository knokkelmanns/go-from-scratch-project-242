package code

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func GetPathSize(path string, recursive, human, all bool) (string, error) {
	size, err := getSizeBytes(path, all, recursive)
	if err != nil {
		return "", err
	}
	return formatSize(size, human), nil
}

func getSizeBytes(path string, all, recursive bool) (int64, error) {
	info, err := os.Lstat(path)
	if err != nil {
		return 0, err
	}
	if info.IsDir() {
		entries, err := os.ReadDir(path)
		if err != nil {
			return 0, err
		}
		var total int64
		for _, entry := range entries {
			if !all && strings.HasPrefix(entry.Name(), ".") {
				continue
			}
			if entry.IsDir() {
				if recursive {
					subtotal, err := getSizeBytes(filepath.Join(path, entry.Name()), all, recursive)
					if err != nil {
						return 0, err
					}
					total += subtotal
				}
				continue
			}
			entryInfo, err := entry.Info()
			if err != nil {
				return 0, err
			}
			total += entryInfo.Size()
		}
		return total, nil
	}
	return info.Size(), nil
}

func formatSize(size int64, human bool) string {
	if !human || size < 1024 {
		return fmt.Sprintf("%dB", size)
	}
	units := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}
	unitIndex := 0
	floatSize := float64(size)
	for floatSize >= 1024 {
		floatSize /= 1024
		unitIndex += 1
	}
	return fmt.Sprintf("%.1f%s", floatSize, units[unitIndex])
}
