package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/urfave/cli"
)

func main() {
	var in, out, dir string
	var append bool
	app := cli.NewApp()
	app.Usage = "jcbのCSVファイルをExcelに変換"
	app.UsageText = filepath.Base(os.Args[0]) + " {-f file | -d dir} -o out [-a]"

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
		cli.BoolFlag{
			Name:        "a",
			Usage:       "追加書き",
			Destination: &append,
		},
	}
	app.HideVersion = true

	app.Action = func(c *cli.Context) error {
		if err := argcheck(in, dir, out); err != nil {
			return cli.NewExitError(app.UsageText, 1)
		}
		return dotask(in, dir, out, append)
	}

	app.Run(os.Args)
}

func argcheck(in, dir, out string) error {
	if (in == "" && dir == "") || out == "" {
		return errors.New("Invalid Argment")
	}
	if in != "" && dir != "" {
		return errors.New("-fか-dのどちらかのみを指定")
	}
	return nil
}

func dotask(in, dir, out string, append bool) error {
	if dir != "" {
		files, err := ioutil.ReadDir(dir)
		if err != nil {
			return err
		}
		for _, file := range files {
			if filepath.Ext(file.Name()) != ".csv" {
				continue
			}
			fmt.Println(filepath.Join(dir, file.Name()))
		}
	}
	return nil
}
