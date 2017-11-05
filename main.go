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
	var isappend bool
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
			Destination: &isappend,
		},
	}
	app.HideVersion = true

	app.Action = func(c *cli.Context) error {
		if err := argcheck(in, dir, out); err != nil {
			return cli.NewExitError(app.UsageText, 1)
		}
		return dotask(in, dir, out, isappend)
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

func outfileopen(out string, isappend bool) (*os.File, error) {
	var mode int
	if isappend == true {
		mode = os.O_APPEND | os.O_WRONLY
	} else {
		mode = os.O_WRONLY | os.O_CREATE | os.O_TRUNC
	}
	ofile, err := os.OpenFile(out, mode, 0666)
	if err != nil {
		return nil, err
	}
	return ofile, nil
}

func inputlist(in, dir string) ([]string, error) {
	var paths []string
	if in != "" {
		return append(paths, in), nil
	}

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		if filepath.Ext(file.Name()) != ".csv" {
			continue
		}
		paths = append(paths, filepath.Join(dir, file.Name()))
	}
	return paths, nil
}

func dotask(in, dir, out string, isappend bool) error {
	paths, err := inputlist(in, dir)
	if err != nil {
		return err
	}

	ofile, err := outfileopen(out, isappend)
	if err != nil {
		return err
	}
	defer ofile.Close()

	for i := 0; i < len(paths); i++ {
		fmt.Printf("%d:%s\n", i, paths[0])
	}
	return nil
}
