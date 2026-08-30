package code

import (
	"fmt"
	"os"
)

func GetPathSize(path string) (string, error) {
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
			entryInfo, err := entry.Info()
			if err != nil {
				return "", err
			}
			total += entryInfo.Size()
		}
		result := fmt.Sprintf("%dB", total)
		return result, nil
	}
	result := fmt.Sprintf("%dB", info.Size())
	return result, nil
}
