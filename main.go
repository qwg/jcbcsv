package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	var in, out, dir string
	app := cli.NewApp()
	app.Usage = "jcbのCSVファイルをExcelに変換"
	app.UsageText = os.Args[0] + " {-f file | -d dir} -o out [-m]"

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

	app.Action = func(c *cli.Context) error {
		if (in == "" && dir == "") || out == "" {
			fmt.Println(app.UsageText)
			return nil
		}
		fmt.Println("Hello World")
		return nil
	}
	app.Run(os.Args)

}
