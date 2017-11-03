package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/urfave/cli"
)

func main() {
	var in, out, dir string
	app := cli.NewApp()
	app.Usage = "jcbのCSVファイルをExcelに変換"
	app.UsageText = filepath.Base(os.Args[0]) + " {-f file | -d dir} -o out [-m]"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "f",
			Usage:       "JCBの明細ファイル(CSV)",
			Destination: &in,
		},
		cli.StringFlag{
			Name:        "d",
			Usage:       "JCBの明細ファイル(CSV)の格納ディレクトリ",
			Destination: &dir,
		},
		cli.StringFlag{
			Name:        "o",
			Usage:       "出力ファイル名",
			Destination: &out,
		},
	}
	app.HideVersion = true

	app.Action = func(c *cli.Context) error {
		if err := argcheck(in, dir, out); err != nil {
			return cli.NewExitError(app.UsageText, 1)
		}
		fmt.Println("Hello World")
		return nil
	}

	app.Run(os.Args)
}

func argcheck(in, dir, out string) error {
	if (in == "" && dir == "") || out == "" {
		return errors.New("Invalid Argment")
	}
	if in != "" && dir != "" {
		return errors.New("-iか-dのどちらかのみを指定")
	}
	return nil
}
