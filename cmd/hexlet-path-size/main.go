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
		Action: func(ctx context.Context, cmd *cli.Command) error {
			path := cmd.Args().First()
			result, err := code.GetPathSize(path)
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
