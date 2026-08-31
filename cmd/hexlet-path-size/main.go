package main

import (
	"code"
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
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "all",
				Aliases: []string{"a"},
				Usage:   "include hidden files and directories",
			},
			&cli.BoolFlag{
				Name:    "human",
				Aliases: []string{"H"},
				Usage:   "human-readable sizes (auto-select unit)",
			},
			&cli.BoolFlag{
				Name:    "recursive",
				Aliases: []string{"r"},
				Usage:   "recursive size of directories",
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			path := cmd.Args().First()
			human := cmd.Bool("human")
			recursive := cmd.Bool("recursive")
			all := cmd.Bool("all")
			result, err := code.GetPathSize(path, human, all, recursive)
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
