package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "hexlet-path-size",
		Usage: "анализатор размера диска",
		Action: func(ctx context.Context, cmd *cli.Command) error {
			path := cmd.Args().First()
			result, err := GetPathSize(path)
			if err != nil {
				return cli.Exit(err.Error(), 1)
			}
			fmt.Println(result + "\t" + path)
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}

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
	} else {
		result := fmt.Sprintf("%dB", info.Size())
		return result, nil
	}
}
