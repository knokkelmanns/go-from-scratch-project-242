package code

import (
	"fmt"
	"os"
	"strings"
)

func GetPathSize(path string, human bool, all bool) (string, error) {
	info, err := os.Lstat(path)
	if err != nil {
		return "", err
	}
	if info.IsDir() {
		entries, err := os.ReadDir(path)
		if err != nil {
			return "", err
		}
		var total int64
		for _, entry := range entries {
			if !all && strings.HasPrefix(entry.Name(), ".") {
				continue
			}
			entryInfo, err := entry.Info()
			if err != nil {
				return "", err
			}
			total += entryInfo.Size()
		}
		result := formatSize(total, human)
		return result, nil
	}
	result := formatSize(info.Size(), human)
	return result, nil
}

func formatSize(size int64, human bool) string {
	if !human {
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
