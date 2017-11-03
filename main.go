package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Usage = "jcbのCSVファイルをExcelに変換"

	app.Action = func(c *cli.Context) error {
		fmt.Println("Hello World")
		return nil
	}
	app.Run(os.Args)

}
